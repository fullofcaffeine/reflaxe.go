enum CollectionFeatureKey {
	Entry(value:Int);
}

class Main {
	static function main() {
		var values = new haxe.ds.EnumValueMap<CollectionFeatureKey, String>();
		values.set(CollectionFeatureKey.Entry(1), "one");
		values.exists(CollectionFeatureKey.Entry(1));
	}
}
