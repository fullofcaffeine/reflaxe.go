import go.Chan;
import go.Go;
import go.Select;

class Main {
	static function main() {
		var gate:Chan<Int> = Go.newChan(1);
		var sendFirst = switch (Select.send(gate, 7)) {
			case Sent:
				"sent";
			case Defaulted:
				"default";
		};
		var sendSecond = switch (Select.send(gate, 8)) {
			case Sent:
				"sent";
			case Defaulted:
				"default";
		};
		var recvFirst = switch (Select.recv(gate)) {
			case Received(value):
				"recv:" + value;
			case Defaulted:
				"empty";
		};
		var recvSecond = switch (Select.recv(gate)) {
			case Received(value):
				"recv:" + value;
			case Defaulted:
				"empty";
		};

		var left:Chan<String> = Go.newChan(1);
		var right:Chan<String> = Go.newChan(1);
		right.send("beta");
		var recvTwo = switch (Select.recv2(left, right)) {
			case First(value):
				"first:" + value;
			case Second(value):
				"second:" + value;
			case Defaulted:
				"none";
		};

		var sendTwoA:Chan<Int> = Go.newChan(1);
		var sendTwoB:Chan<Int> = Go.newChan(1);
		var sendTwo = switch (Select.send2(sendTwoA, 11, sendTwoB, 22)) {
			case FirstSent:
				"first";
			case SecondSent:
				"second";
			case Defaulted:
				"none";
		};
		var sendTwoValues = sendTwoA.recvOr(-1) + "," + sendTwoB.recvOr(-1);

		Sys.println("select.send=" + sendFirst + "," + sendSecond);
		Sys.println("select.recv=" + recvFirst + "," + recvSecond);
		Sys.println("select.recv2=" + recvTwo);
		Sys.println("select.send2=" + sendTwo + " values=" + sendTwoValues);
	}
}
