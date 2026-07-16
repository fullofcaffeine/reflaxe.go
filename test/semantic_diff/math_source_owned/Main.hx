class Main {
	static function close(actual:Float, expected:Float, epsilon:Float = 1e-12):Bool {
		return Math.abs(actual - expected) <= epsilon;
	}

	static function zeroSign(value:Float):String {
		if (value != 0.0)
			return "nonzero";
		return 1.0 / value == Math.NEGATIVE_INFINITY ? "negative" : "positive";
	}

	static function main() {
		Sys.println("constants=" + close(Math.PI, 3.141592653589793) + ":" + Math.isFinite(Math.PI) + ":" + Math.isFinite(Math.POSITIVE_INFINITY) + ":"
			+ Math.isFinite(Math.NEGATIVE_INFINITY) + ":" + Math.isFinite(Math.NaN) + ":" + Math.isNaN(Math.NaN));

		Sys.println("rounding=" + Math.floor(3.8) + ":" + Math.floor(-3.2) + ":" + Math.ceil(3.2) + ":" + Math.ceil(-3.8) + ":" + Math.round(1.5) + ":"
			+ Math.round(0.5) + ":" + Math.round(-0.5) + ":" + Math.round(-1.5));
		Sys.println("floatRounding="
			+ (Math.ffloor(3.8) == 3.0)
			+ ":"
			+ (Math.fceil(-3.8) == -3.0)
			+ ":"
			+ (Math.fround(-1.5) == -1.0)
			+ ":"
			+ (Math.ffloor(Math.POSITIVE_INFINITY) == Math.POSITIVE_INFINITY)
			+ ":"
			+ Math.isNaN(Math.fround(Math.NaN)));

		Sys.println("trig=" + close(Math.sin(Math.PI / 6.0), 0.5) + ":" + close(Math.cos(Math.PI / 3.0), 0.5) + ":" + close(Math.tan(Math.PI / 4.0), 1.0)
			+ ":" + close(Math.asin(0.5), Math.PI / 6.0) + ":" + close(Math.acos(0.5), Math.PI / 3.0) + ":" + close(Math.atan(1.0), Math.PI / 4.0) + ":"
			+ close(Math.atan2(1.0, -1.0), Math.PI * 0.75));
		Sys.println("powers=" + close(Math.exp(1.0), 2.718281828459045) + ":" + close(Math.log(2.718281828459045), 1.0) + ":"
			+ close(Math.pow(2.0, 10.0), 1024.0) + ":" + close(Math.sqrt(2.0) * Math.sqrt(2.0), 2.0));
		Sys.println("domains=" + Math.isNaN(Math.sqrt(-1.0)) + ":" + Math.isNaN(Math.log(-1.0)) + ":" + Math.isNaN(Math.asin(2.0)) + ":"
			+ Math.isNaN(Math.pow(-2.0, 0.5)) + ":" + (Math.log(0.0) == Math.NEGATIVE_INFINITY));
		Sys.println("infinite="
			+ Math.isNaN(Math.sin(Math.POSITIVE_INFINITY))
			+ ":"
			+ Math.isNaN(Math.cos(Math.NEGATIVE_INFINITY))
			+ ":"
			+ Math.isNaN(Math.tan(Math.POSITIVE_INFINITY))
			+ ":"
			+ close(Math.atan(Math.POSITIVE_INFINITY), Math.PI / 2.0)
			+ ":"
			+ close(Math.atan2(Math.POSITIVE_INFINITY, 1.0), Math.PI / 2.0)
			+ ":"
			+ (Math.exp(Math.POSITIVE_INFINITY) == Math.POSITIVE_INFINITY)
			+ ":"
			+ (Math.exp(Math.NEGATIVE_INFINITY) == 0.0)
			+ ":"
			+ (Math.log(Math.POSITIVE_INFINITY) == Math.POSITIVE_INFINITY)
			+ ":"
			+ (Math.sqrt(Math.POSITIVE_INFINITY) == Math.POSITIVE_INFINITY));

		var negativeZero = -1.0 / Math.POSITIVE_INFINITY;
		var positiveZero = 0.0;
		Sys.println("signedZero=" + zeroSign(negativeZero) + ":" + zeroSign(Math.abs(negativeZero)) + ":" + zeroSign(Math.min(positiveZero, negativeZero))
			+ ":" + zeroSign(Math.min(negativeZero, positiveZero)) + ":" + zeroSign(Math.max(positiveZero, negativeZero)) + ":"
			+ zeroSign(Math.max(negativeZero, positiveZero)) + ":" + zeroSign(Math.sqrt(negativeZero)));
		Sys.println("nanMinMax="
			+ Math.isNaN(Math.min(Math.NaN, 1.0))
			+ ":"
			+ Math.isNaN(Math.max(1.0, Math.NaN))
			+ ":"
			+ (Math.abs(Math.NEGATIVE_INFINITY) == Math.POSITIVE_INFINITY));

		var randomInRange = true;
		for (_ in 0...32) {
			var value = Math.random();
			if (value < 0.0 || value >= 1.0)
				randomInRange = false;
		}
		Sys.println("random=" + randomInRange);
	}
}
