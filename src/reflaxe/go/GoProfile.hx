package reflaxe.go;

enum abstract GoProfile(String) from String to String {
  var Portable = "portable";
  var Idiomatic = "idiomatic";
  var Gopher = "gopher";
  var Metal = "metal";
}
