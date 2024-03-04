package builder

import (
	"os"
	"path/filepath"

	"github.com/tinygo-org/tinygo/goenv"
)

var WasmBuiltins = Library{
	name: "wasmbuiltins",
	makeHeaders: func(target, includeDir string) error {
		os.Mkdir(includeDir+"/bits", 0o777)
		f, err := os.Create(includeDir + "/bits/alltypes.h")
		if err != nil {
			return err
		}
		if _, err := f.Write([]byte(wasmAllTypes)); err != nil {
			return err
		}
		return f.Close()
	},
	cflags: func(target, headerPath string) []string {
		libcDir := filepath.Join(goenv.Get("TINYGOROOT"), "lib/wasi-libc")
		return []string{
			"-Werror",
			"-Wall",
			"-std=gnu11",
			"-nostdlibinc",
			"-isystem", libcDir + "/libc-top-half/musl/arch/wasm32",
			"-isystem", libcDir + "/libc-top-half/musl/arch/generic",
			"-isystem", libcDir + "/libc-top-half/musl/src/internal",
			"-isystem", libcDir + "/libc-top-half/musl/src/include",
			"-isystem", libcDir + "/libc-top-half/musl/include",
			"-I" + headerPath,
		}
	},
	sourceDir: func() string { return filepath.Join(goenv.Get("TINYGOROOT"), "lib/wasi-libc") },
	librarySources: func(target string) ([]string, error) {
		return []string{
			"libc-top-half/musl/src/math/__math_divzero.c",
			"libc-top-half/musl/src/math/__math_invalid.c",
			"libc-top-half/musl/src/math/__math_oflow.c",
			"libc-top-half/musl/src/math/__math_uflow.c",
			"libc-top-half/musl/src/math/__math_xflow.c",
			"libc-top-half/musl/src/math/exp.c",
			"libc-top-half/musl/src/math/exp_data.c",
			"libc-top-half/musl/src/math/exp2.c",
			"libc-top-half/musl/src/math/log.c",
			"libc-top-half/musl/src/math/log_data.c",
		}, nil
	},
}

const wasmAllTypes = `
typedef __INT8_TYPE__   int8_t;
typedef __INT16_TYPE__  int16_t;
typedef __INT32_TYPE__  int32_t;
typedef __INT64_TYPE__  int64_t;
typedef __UINT8_TYPE__  uint8_t;
typedef __UINT16_TYPE__ uint16_t;
typedef __UINT32_TYPE__ uint32_t;
typedef __UINT64_TYPE__ uint64_t;

typedef double double_t;
`
