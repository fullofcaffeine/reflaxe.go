class Base {
	public function new() {}

	public function who():Int {
		return 1;
	}

	public function callWho():Int {
		return who();
	}
}
