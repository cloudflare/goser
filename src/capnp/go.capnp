@0xd12a1c51fedd6c88;
annotation package(file) :Text;
annotation import(file) :Text;
annotation tag(enumerant) : Text;
annotation notag(enumerant) : Void;
annotation jsonfmt(field) : Jsonfmts;
struct Jsonfmts {
	pkg @0 : Text;
	fmt @1 : Text;
}
$package("capn");
