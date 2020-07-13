package analyzer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ShouldBeEqual = "They should be equal"
)

func TestBaseAnalyzer(t *testing.T) {
	t.Helper()
	ba := &BaseAnalyzer{}

	assert.Equal(t, ba.PickIP("foo 1.2.3.4 bar"), "1.2.3.4", ShouldBeEqual)
	assert.Equal(t, ba.IsPublicIP("127.0.0.1"), false, ShouldBeEqual)
}

func TestStaticAnalyzer(t *testing.T) {
	t.Helper()

	sa := &StaticAnalyzer{}
	sa.Target = "testdata/helloworld"

	assert.Equal(t, sa.IsUpxPacked(), false, ShouldBeEqual)
	info, _ := sa.GetFileInfo()
	assert.Equal(t, strings.Split(info, ",")[0], "ELF 64-bit LSB pie executable")
}
