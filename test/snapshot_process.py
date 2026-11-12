"""Own snapshot launcher descendants without changing command timeout budgets.

POSIX commands get a new session/process group. Windows uses a small supervisor
which joins a kill-on-close Job Object before spawning the command. Assigning an
already-running launcher to a job would race its first child creation. The
Windows job does not permit breakaway. Its completed supervisor also reclaims
remaining children on normal exit. POSIX ownership covers descendants which
retain the inherited process group; deliberately detached sessions are outside it.
"""

from __future__ import annotations

import contextlib
import json
import os
from pathlib import Path
import signal
import subprocess
import sys
import tempfile


def run_command(cmd: list[str], cwd: Path, timeout_s: int, env: dict[str, str] | None = None) -> subprocess.CompletedProcess[str]:
    merged_env = os.environ.copy()
    if env:
        merged_env.update(env)
    # Keep the error transport private and separate from compiler stdout/stderr.
    # The Windows helper reports native spawn errors so callers retain OSError,
    # rather than receiving a successful helper launch with a synthetic exit.
    with (tempfile.TemporaryDirectory(prefix="snapshot-process-") if os.name == "nt" else contextlib.nullcontext()) as raw:
        error_path = Path(raw) / "spawn-error.json" if raw is not None else None
        launch = [sys.executable, str(Path(__file__).resolve()), str(error_path), *cmd] if error_path is not None else cmd
        with subprocess.Popen(
            launch, cwd=cwd, stdout=subprocess.PIPE, stderr=subprocess.PIPE,
            text=True, env=merged_env, start_new_session=os.name == "posix",
        ) as process:
            try:
                stdout, stderr = process.communicate(timeout=timeout_s)
            except subprocess.TimeoutExpired as error:
                _kill_owned_tree(process)
                if os.name == "nt":
                    # Match subprocess.run's Windows output collection contract.
                    error.stdout, error.stderr = process.communicate()
                else:
                    process.wait()
                error.cmd = cmd
                raise
            except BaseException:
                _kill_owned_tree(process)
                process.wait()
                raise
            if error_path is not None and error_path.exists():
                detail = json.loads(error_path.read_text(encoding="utf-8"))
                raise OSError(detail["errno"], detail["strerror"], detail["filename"], detail["winerror"], detail["filename2"])
            return subprocess.CompletedProcess(cmd, process.returncode, stdout, stderr)


def _kill_owned_tree(process: subprocess.Popen[str]) -> None:
    if os.name == "posix":
        try:
            # The group remains identifiable even if its launcher has already
            # exited while a descendant keeps the captured output pipes open.
            os.killpg(process.pid, signal.SIGKILL)
        except ProcessLookupError:
            pass
    else:
        # The Windows supervisor's non-inherited job handle closes at death.
        # JOB_OBJECT_LIMIT_KILL_ON_JOB_CLOSE then terminates all job members.
        process.kill()


def _windows_job() -> tuple[object, object]:
    """Enter a private, non-breakaway Windows job before starting any child."""
    import ctypes
    from ctypes import wintypes

    class BasicLimits(ctypes.Structure):
        _fields_ = [
            ("PerProcessUserTimeLimit", ctypes.c_longlong),
            ("PerJobUserTimeLimit", ctypes.c_longlong),
            ("LimitFlags", wintypes.DWORD),
            ("MinimumWorkingSetSize", ctypes.c_size_t),
            ("MaximumWorkingSetSize", ctypes.c_size_t),
            ("ActiveProcessLimit", wintypes.DWORD),
            ("Affinity", ctypes.c_size_t),
            ("PriorityClass", wintypes.DWORD),
            ("SchedulingClass", wintypes.DWORD),
        ]

    class IoCounters(ctypes.Structure):
        _fields_ = [(name, ctypes.c_ulonglong) for name in (
            "ReadOperationCount", "WriteOperationCount", "OtherOperationCount",
            "ReadTransferCount", "WriteTransferCount", "OtherTransferCount",
        )]

    class ExtendedLimits(ctypes.Structure):
        _fields_ = [
            ("BasicLimitInformation", BasicLimits), ("IoInfo", IoCounters),
            ("ProcessMemoryLimit", ctypes.c_size_t), ("JobMemoryLimit", ctypes.c_size_t),
            ("PeakProcessMemoryUsed", ctypes.c_size_t), ("PeakJobMemoryUsed", ctypes.c_size_t),
        ]

    kernel = ctypes.WinDLL("kernel32", use_last_error=True)
    kernel.CreateJobObjectW.argtypes = [ctypes.c_void_p, wintypes.LPCWSTR]
    kernel.CreateJobObjectW.restype = wintypes.HANDLE
    kernel.SetInformationJobObject.argtypes = [wintypes.HANDLE, ctypes.c_int, ctypes.c_void_p, wintypes.DWORD]
    kernel.SetInformationJobObject.restype = wintypes.BOOL
    kernel.GetCurrentProcess.argtypes = []
    kernel.GetCurrentProcess.restype = wintypes.HANDLE
    kernel.AssignProcessToJobObject.argtypes = [wintypes.HANDLE, wintypes.HANDLE]
    kernel.AssignProcessToJobObject.restype = wintypes.BOOL
    kernel.CloseHandle.argtypes = [wintypes.HANDLE]
    kernel.CloseHandle.restype = wintypes.BOOL
    job = kernel.CreateJobObjectW(None, None)
    if not job:
        raise ctypes.WinError(ctypes.get_last_error())
    try:
        limits = ExtendedLimits()
        limits.BasicLimitInformation.LimitFlags = 0x2000  # KILL_ON_JOB_CLOSE; no breakaway.
        if not kernel.SetInformationJobObject(job, 9, ctypes.byref(limits), ctypes.sizeof(limits)):
            raise ctypes.WinError(ctypes.get_last_error())
        if not kernel.AssignProcessToJobObject(job, kernel.GetCurrentProcess()):
            raise ctypes.WinError(ctypes.get_last_error())
    except BaseException:
        kernel.CloseHandle(job)
        raise
    return kernel, job


def _windows_main(error_path: Path, cmd: list[str]) -> int:
    try:
        kernel, job = _windows_job()
        # Keep the only job handle alive until this supervisor exits. It is not
        # inherited by the command, so killing the supervisor closes ownership.
        result = subprocess.call(cmd)
        # Closing our own kill-on-close job also terminates this supervisor;
        # retain the handle until OS process teardown preserves command status.
        return result
    except OSError as error:
        error_path.write_text(json.dumps({
            "errno": error.errno, "strerror": error.strerror,
            "filename": error.filename, "winerror": getattr(error, "winerror", None),
            "filename2": error.filename2,
        }), encoding="utf-8")
        return 127


if __name__ == "__main__":
    import ctypes
    from ctypes import wintypes

    # Use the native DWORD boundary for high Windows statuses, independent of
    # Python-version exit-code narrowing (including historical 32-bit long overflow).
    exit_process = ctypes.WinDLL("kernel32", use_last_error=True).ExitProcess
    exit_process.argtypes = [wintypes.UINT]
    exit_process.restype = None
    exit_process(_windows_main(Path(sys.argv[1]), sys.argv[2:]))
