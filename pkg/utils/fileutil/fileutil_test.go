package fileutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashFunc(t *testing.T) {
	t.Helper()

	target := "testdata/helloworld"

	filesize, _ := GetFileSize(target)
	assert.EqualValues(t, filesize, 16536)

	digest, _ := GetFileMd5(target)
	assert.Equal(t, digest, "3fb856dc07a7de73fab89c34e53b7c5c")

	digest, _ = GetFileSha1(target)
	assert.Equal(t, digest, "280939497881ed7864380115a4238c6ab988e17a")

	digest, _ = GetFileSha256(target)
	assert.Equal(t, digest, "4cf37a3e0ab6cf02907e0eeec54139ef72add0c9b3f67d99e655f6c96290bb86")

	digest, _ = GetFileSSDeep(target)
	assert.Equal(t, digest, "96:RpYiTJB+BawDSXXz0DaHZCEw7/LuB3EBqoAKK4rq:R1dw4wAX4aHoeuso")
}
