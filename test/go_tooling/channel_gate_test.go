package main

import "testing"

func generatedResultError(result *go___Result) string {
	if result == nil {
		return "<nil-result>"
	}
	message := result.error()
	if message == nil {
		return ""
	}
	return *message
}

func requireNativeChannelPanic(t *testing.T, operation string, run func()) {
	t.Helper()
	defer func() {
		if recovered := recover(); recovered == nil {
			t.Fatalf("%s did not preserve Go's native panic", operation)
		}
	}()
	run()
}

func TestTypedGeneratedChannelCloseContract(t *testing.T) {
	channel := go__concurrency_makeChan__int_95e97e5e(1)
	go__concurrency_send__int_95e97e5e(channel, 7)
	go__concurrency_close__int_95e97e5e(channel)

	buffered := go__concurrency_tryRecv__int_95e97e5e(channel)
	if !buffered.isOk() || buffered.unwrap() != 7 {
		t.Fatalf("buffered value before close drain = (%t, %v), want (true, 7)", buffered.isOk(), buffered.value)
	}
	closed := go__concurrency_tryRecv__int_95e97e5e(channel)
	if !closed.isErr() || generatedResultError(closed) != "closed" {
		t.Fatalf("tryRecv after close = (%t, %q), want (true, closed)", closed.isErr(), generatedResultError(closed))
	}
	if got := go__concurrency_recvOr__int_95e97e5e(channel, -1); got != -1 {
		t.Fatalf("recvOr after close = %d, want fallback -1", got)
	}
	if got := go__concurrency_recv__int_95e97e5e(channel); got != 0 {
		t.Fatalf("recv after close = %d, want Go zero value 0", got)
	}

	requireNativeChannelPanic(t, "send after close", func() {
		go__concurrency_send__int_95e97e5e(channel, 8)
	})
	requireNativeChannelPanic(t, "trySend after close", func() {
		go__concurrency_trySend__int_95e97e5e(channel, 8)
	})
	requireNativeChannelPanic(t, "double close", func() {
		go__concurrency_close__int_95e97e5e(channel)
	})
}

func TestGenericGeneratedChannelCloseContract(t *testing.T) {
	channel := make(chan any, 1)
	go__concurrency_send(channel, 7)
	go__concurrency_close(channel)

	buffered := go__concurrency_tryRecv(channel)
	if !buffered.isOk() || buffered.unwrap() != 7 {
		t.Fatalf("generic buffered receive = (%t, %v), want (true, 7)", buffered.isOk(), buffered.value)
	}
	closed := go__concurrency_tryRecv(channel)
	if !closed.isErr() || generatedResultError(closed) != "closed" {
		t.Fatalf("generic tryRecv after close = (%t, %q), want (true, closed)", closed.isErr(), generatedResultError(closed))
	}
	if got := go__concurrency_recvOr(channel, -1); got != -1 {
		t.Fatalf("generic recvOr after close = %v, want fallback -1", got)
	}

	requireNativeChannelPanic(t, "generic send after close", func() {
		go__concurrency_send(channel, 8)
	})
	requireNativeChannelPanic(t, "generic trySend after close", func() {
		go__concurrency_trySend(channel, 8)
	})
	requireNativeChannelPanic(t, "generic double close", func() {
		go__concurrency_close(channel)
	})
}

func TestGeneratedNilChannelNonBlockingContract(t *testing.T) {
	var channel chan int
	if go__concurrency_trySend__int_95e97e5e(channel, 1) {
		t.Fatal("trySend on a nil channel reported success")
	}
	empty := go__concurrency_tryRecv__int_95e97e5e(channel)
	if !empty.isErr() || generatedResultError(empty) != "empty" {
		t.Fatalf("tryRecv on nil channel = (%t, %q), want (true, empty)", empty.isErr(), generatedResultError(empty))
	}
	if got := go__concurrency_recvOr__int_95e97e5e(channel, -1); got != -1 {
		t.Fatalf("recvOr on nil channel = %d, want fallback -1", got)
	}
	requireNativeChannelPanic(t, "close nil channel", func() {
		go__concurrency_close__int_95e97e5e(channel)
	})
}
