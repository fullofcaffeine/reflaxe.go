package profile;

class StorySignalMetrics {
	public final cardCount:Int;
	public final highValue:Int;
	public final openHighValue:Int;

	public function new(cardCount:Int, highValue:Int, openHighValue:Int) {
		this.cardCount = cardCount;
		this.highValue = highValue;
		this.openHighValue = openHighValue;
	}
}
