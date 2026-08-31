@:goBuildConstraint("cgo")
@:keep
class Platform {}

@:goBuildConstraint("!cgo")
@:keep
class PlatformFallback {}
