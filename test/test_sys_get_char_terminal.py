#!/usr/bin/env python3

from __future__ import annotations

import errno
import os
from pathlib import Path
import pty
import select
import shutil
import subprocess
import tempfile
import termios
import time


ROOT = Path(__file__).resolve().parent.parent
CASE = ROOT / "test" / "snapshot" / "sys" / "sys_get_char_terminal"
OUT = CASE / "out"
CHARACTER_FLAGS = termios.ICANON | termios.ECHO | getattr(termios, "ECHONL", 0)


def run_checked(
    command: list[str],
    *,
    cwd: Path,
    env: dict[str, str] | None = None,
    timeout: int = 120,
) -> None:
    merged_env = os.environ.copy()
    if env:
        merged_env.update(env)
    process = subprocess.run(
        command,
        cwd=cwd,
        env=merged_env,
        capture_output=True,
        text=True,
        timeout=timeout,
    )
    if process.returncode == 0:
        return
    output = "\n".join(part for part in (process.stdout, process.stderr) if part)
    raise AssertionError(f"{' '.join(command)} failed with {process.returncode}\n{output}")


def read_available(master_fd: int, deadline: float) -> bytes:
    output = bytearray()
    while time.monotonic() < deadline:
        readable, _, _ = select.select([master_fd], [], [], 0.05)
        if not readable:
            continue
        try:
            chunk = os.read(master_fd, 4096)
        except OSError as error:
            if error.errno == errno.EIO:
                break
            raise
        if not chunk:
            break
        output.extend(chunk)
    return bytes(output)


def read_until(master_fd: int, marker: bytes, deadline: float) -> bytes:
    output = bytearray()
    while marker not in output and time.monotonic() < deadline:
        readable, _, _ = select.select([master_fd], [], [], 0.05)
        if not readable:
            continue
        try:
            chunk = os.read(master_fd, 4096)
        except OSError as error:
            if error.errno == errno.EIO:
                break
            raise
        if not chunk:
            break
        output.extend(chunk)
    if marker not in output:
        raise AssertionError(f"PTY output never reached {marker!r}: {bytes(output)!r}")
    return bytes(output)


def wait_for_character_mode(slave_fd: int, process: subprocess.Popen[bytes], deadline: float) -> None:
    while time.monotonic() < deadline:
        if process.poll() is not None:
            raise AssertionError(f"program exited before entering character mode: {process.returncode}")
        local_flags = termios.tcgetattr(slave_fd)[3]
        if local_flags & CHARACTER_FLAGS == 0:
            return
        time.sleep(0.01)
    local_flags = termios.tcgetattr(slave_fd)[3]
    raise AssertionError(
        "Sys.getChar did not disable canonical input and terminal echo before reading "
        f"(lflag={local_flags:#x})"
    )


def run_terminal_case(binary: Path, echo: bool) -> None:
    master_fd, slave_fd = pty.openpty()
    initial_attributes = termios.tcgetattr(slave_fd)
    process: subprocess.Popen[bytes] | None = None
    try:
        argument = "echo" if echo else "no-echo"
        process = subprocess.Popen(
            [str(binary), argument],
            cwd=ROOT,
            stdin=slave_fd,
            stdout=slave_fd,
            stderr=slave_fd,
            close_fds=True,
        )
        deadline = time.monotonic() + 10
        output = bytearray(read_until(master_fd, b"ready|", deadline))
        wait_for_character_mode(slave_fd, process, deadline)

        os.write(master_fd, b"Z")
        process.wait(timeout=max(0.1, deadline - time.monotonic()))
        output.extend(read_available(master_fd, time.monotonic() + 0.5))
        if process.returncode != 0:
            raise AssertionError(f"terminal child exited with {process.returncode}: {bytes(output)!r}")

        restored_attributes = termios.tcgetattr(slave_fd)
        restored_flags = restored_attributes[3] & CHARACTER_FLAGS
        initial_flags = initial_attributes[3] & CHARACTER_FLAGS
        restored_minimum = restored_attributes[6][termios.VMIN]
        initial_minimum = initial_attributes[6][termios.VMIN]
        restored_timeout = restored_attributes[6][termios.VTIME]
        initial_timeout = initial_attributes[6][termios.VTIME]
        if (
            restored_flags != initial_flags
            or restored_minimum != initial_minimum
            or restored_timeout != initial_timeout
        ):
            raise AssertionError(
                "Sys.getChar did not restore terminal character settings: "
                f"flags before={initial_flags:#x}, after={restored_flags:#x}; "
                f"VMIN before={initial_minimum!r}, after={restored_minimum!r}; "
                f"VTIME before={initial_timeout!r}, after={restored_timeout!r}"
            )

        normalized = bytes(output).replace(b"\r\n", b"\n")
        expected = b"ready|Z|90|\n" if echo else b"ready||90|\n"
        if normalized != expected:
            raise AssertionError(
                f"echo={echo} PTY output mismatch: got {normalized!r}, want {expected!r}"
            )
    finally:
        if process is not None and process.poll() is None:
            process.kill()
            process.wait(timeout=5)
        os.close(master_fd)
        os.close(slave_fd)


def run_redirected_case(binary: Path, echo: bool, input_bytes: bytes, expected: bytes) -> None:
    argument = "echo" if echo else "no-echo"
    process = subprocess.run(
        [str(binary), argument],
        cwd=ROOT,
        input=input_bytes,
        capture_output=True,
        timeout=10,
    )
    if process.returncode != 0:
        raise AssertionError(
            f"redirected child exited with {process.returncode}: "
            f"stdout={process.stdout!r}, stderr={process.stderr!r}"
        )
    if process.stdout != expected or process.stderr:
        raise AssertionError(
            f"redirected echo={echo} input={input_bytes!r} mismatch: "
            f"stdout={process.stdout!r}, stderr={process.stderr!r}, want={expected!r}"
        )


def main() -> int:
    if os.name != "posix":
        raise SystemExit("the Sys.getChar PTY contract requires a POSIX admitted-runtime host")

    shutil.rmtree(OUT, ignore_errors=True)
    with tempfile.TemporaryDirectory(prefix="haxe-go-sys-get-char-") as raw_temp:
        temp = Path(raw_temp)
        try:
            run_checked(
                ["haxe", "compile.hxml", "-D", "go_no_build"],
                cwd=CASE,
                env={"HAXE_NO_SERVER": "1"},
            )
            go_files = sorted(str(path) for path in OUT.rglob("*.go"))
            run_checked(["gofmt", "-w", *go_files], cwd=OUT)
            run_checked(["go", "test", "./..."], cwd=OUT)

            binary = temp / "sys-get-char"
            run_checked(
                [
                    "go",
                    "build",
                    "-gcflags=all=-d=checkptr=2",
                    "-o",
                    str(binary),
                    ".",
                ],
                cwd=OUT,
            )
            run_terminal_case(binary, echo=False)
            run_terminal_case(binary, echo=True)
            run_redirected_case(binary, echo=False, input_bytes=b"Q", expected=b"ready||81|\n")
            run_redirected_case(binary, echo=True, input_bytes=b"Q", expected=b"ready|Q|81|\n")
            run_redirected_case(binary, echo=False, input_bytes=b"", expected=b"ready|eof|\n")

            for goos, goarch in (
                ("linux", "amd64"),
                ("darwin", "arm64"),
                ("windows", "amd64"),
                ("freebsd", "amd64"),
            ):
                output = temp / f"sys-get-char-{goos}-{goarch}"
                run_checked(
                    ["go", "build", "-o", str(output), "."],
                    cwd=OUT,
                    env={"CGO_ENABLED": "0", "GOOS": goos, "GOARCH": goarch},
                )
        finally:
            shutil.rmtree(OUT, ignore_errors=True)

    print(
        "[sys-get-char-terminal] checkptr PTY, redirected echo/EOF, restoration, "
        "and cross-build contracts passed"
    )
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
