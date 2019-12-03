package matchers

/*
var (
	TypeWasm = newType("wasm", "application/wasm")
	TypeExe  = newType("exe", "application/x-msdownload")
	TypeSwf  = newType("swf", "application/x-shockwave-flash")
	TypeEot  = newType("eot", "application/octet-stream")
)
*/

var Application = Map{
	//TypeWasm: Wasm,
	//TypeExe:  Exe,
	//TypeSwf:  Swf,
	//TypeEot:  Eot,
}

/*
// Wasm detects a Web Assembly 1.0 filetype.
func Wasm(buf []byte) bool {
	// WASM has starts with `\0asm`, followed by the version.
	// http://webassembly.github.io/spec/core/binary/modules.html#binary-magic
	return len(buf) >= 8 &&
		buf[0] == 0x00 && buf[1] == 0x61 &&
		buf[2] == 0x73 && buf[3] == 0x6D &&
		buf[4] == 0x01 && buf[5] == 0x00 &&
		buf[6] == 0x00 && buf[7] == 0x00
}

func Exe(buf []byte) bool {
	return len(buf) > 1 &&
		buf[0] == 0x4D && buf[1] == 0x5A
}

func Swf(buf []byte) bool {
	return len(buf) > 2 &&
		(buf[0] == 0x43 || buf[0] == 0x46) &&
		buf[1] == 0x57 && buf[2] == 0x53
}

func Eot(buf []byte) bool {
	return len(buf) > 35 &&
		buf[34] == 0x4C && buf[35] == 0x50 &&
		((buf[8] == 0x02 && buf[9] == 0x00 &&
			buf[10] == 0x01) || (buf[8] == 0x01 &&
			buf[9] == 0x00 && buf[10] == 0x00) ||
			(buf[8] == 0x02 && buf[9] == 0x00 &&
				buf[10] == 0x02))
}

*/
