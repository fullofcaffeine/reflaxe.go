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
		}
		var sendSecond = switch (Select.send(gate, 8)) {
			case Sent:
				"sent";
			case Defaulted:
				"default";
		}
		var recvFirst = switch (Select.recv(gate)) {
			case Received(value):
				"recv:" + value;
			case Defaulted:
				"empty";
		}
		var recvSecond = switch (Select.recv(gate)) {
			case Received(value):
				"recv:" + value;
			case Defaulted:
				"empty";
		}
		Sys.println(sendFirst);
		Sys.println(sendSecond);
		Sys.println(recvFirst);
		Sys.println(recvSecond);

		var left:Chan<String> = Go.newChan(1);
		var right:Chan<String> = Go.newChan(1);
		right.send("beta");
		var recvTwo = switch (Select.recv2(left, right)) {
			case First(value):
				"left:" + value;
			case Second(value):
				"right:" + value;
			case Defaulted:
				"none";
		}
		Sys.println(recvTwo);

		var sendTwoA:Chan<Int> = Go.newChan(1);
		var sendTwoB:Chan<Int> = Go.newChan(1);
		var sendTwo = switch (Select.send2(sendTwoA, 11, sendTwoB, 22)) {
			case FirstSent:
				"a";
			case SecondSent:
				"b";
			case Defaulted:
				"none";
		}
		var sendTwoValues = sendTwoA.recvOr(-1) + "," + sendTwoB.recvOr(-1);
		Sys.println(sendTwo);
		Sys.println(sendTwoValues);
	}
}
