package go;

enum SelectRecv<T> {
	Received(value:T);
	Defaulted;
}

enum SelectRecv2<A, B> {
	First(value:A);
	Second(value:B);
	Defaulted;
}

enum SelectSend {
	Sent;
	Defaulted;
}

enum SelectSend2 {
	FirstSent;
	SecondSent;
	Defaulted;
}

class Select {
	/**
		`@:generic` forces per-type specialization so `go.Chan<T>` operations stay typed at call sites.
		This helps metal output keep typed channel shim paths instead of collapsing into broader `any` fallback lanes.
	**/
	@:generic
	public static function recv<T>(channel:Chan<T>):SelectRecv<T> {
		var received = channel.tryRecv();
		if (received.isOk()) {
			return Received(received.unwrap());
		}
		return Defaulted;
	}

	@:generic
	public static function recv2<A, B>(first:Chan<A>, second:Chan<B>):SelectRecv2<A, B> {
		var firstRecv = first.tryRecv();
		if (firstRecv.isOk()) {
			return First(firstRecv.unwrap());
		}

		var secondRecv = second.tryRecv();
		if (secondRecv.isOk()) {
			return Second(secondRecv.unwrap());
		}

		return Defaulted;
	}

	@:generic
	public static function send<T>(channel:Chan<T>, value:T):SelectSend {
		if (channel.trySend(value)) {
			return Sent;
		}
		return Defaulted;
	}

	@:generic
	public static function send2<A, B>(first:Chan<A>, firstValue:A, second:Chan<B>, secondValue:B):SelectSend2 {
		if (first.trySend(firstValue)) {
			return FirstSent;
		}
		if (second.trySend(secondValue)) {
			return SecondSent;
		}
		return Defaulted;
	}
}
