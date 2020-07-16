package analyzer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseAnalyzer(t *testing.T) {
	t.Helper()
	ba := &BaseAnalyzer{}

	assert.Equal(t, "1.2.3.4", ba.PickIP("foo 1.2.3.4 bar"))
	assert.Equal(t, false, ba.IsPublicIP("127.0.0.1"))
}

func TestStaticAnalyzer(t *testing.T) {
	t.Helper()

	sa := &StaticAnalyzer{}
	sa.Target = "testdata/helloworld"

	assert.Equal(t, false, sa.IsUpxPacked())

	info, _ := sa.GetFileInfo()
	assert.Equal(t, "ELF 64-bit LSB pie executable", strings.Split(info, ",")[0])

	exifInfo, _ := sa.GetExifInfo()
	assert.Equal(t, "ELF shared library", exifInfo["File Type"])

	ascii, _ := sa.GetAsciiStr()
	assert.Equal(t, "/lib64/ld-linux-x86-64.so.2", ascii[0])

	deps, _ := sa.GetLibraryDepends()
	assert.Equal(t, "libc.so.6", deps[1].name)
	assert.Equal(t, "/usr/lib/libc.so.6", deps[1].path)

	elfHdr, _ := sa.GetElfHeader()
	assert.Equal(t, "ELF64", elfHdr["Class"])
}
