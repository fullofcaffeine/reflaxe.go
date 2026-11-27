/**
	Modules whose file names meet: `a_b.Util` and `a.b.Util` share the base
	`module_a_b_util`, so one of them needs a numbered name, and that number
	must not take the name of `a.b.Util_1`. The `m` group repeats the shape
	with the `_1` module sorting after both shared bases, and the `tools`
	group repeats it after the `_hx` build-constraint escape. Every module
	must get its own file, or one file overwrites another and Go reports
	"undefined" names.
**/
class Main {
	static function main() {
		Sys.println(a_b.Util.name());
		Sys.println(a.b.Util.name());
		Sys.println(a.b.Util_1.name());
		Sys.println(m.n.Util.name());
		Sys.println(m_n.Util.name());
		Sys.println(m_n.Util_1.name());
		Sys.println(tools.Linux.name());
		Sys.println(tools.Linux_hx.name());
		Sys.println(tools.Linux_hx_1.name());
	}
}
