class CollectionFeatureKey {
	public function new() {}
}

class Main {
	static function main() {
		var key = new CollectionFeatureKey();
		var values = new haxe.ds.ObjectMap<CollectionFeatureKey, String>();
		values.set(key, "one");
		values.exists(key);
	}
}
