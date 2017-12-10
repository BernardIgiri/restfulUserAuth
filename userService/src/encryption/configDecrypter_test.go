package encryption_test

import (
	"testing"

	"application/encryption"
	"github.com/stretchr/testify/assert"
)

func TestConfigDecrypt(t *testing.T) {
	d := encryption.NewConfigDecrypter([]byte("1234567890123456"))
	const data = "nDK0gO1ExTPa60VtFjUXjv40VOMBEBtlHtAVqSvLeb9/p92Amla3C8s+6sVMd+Cw"
	const expected = "Hello bob, how are you?"

	actual, err := d.Decrypt(data)

	assert.Nil(t, err)
	assert.Equal(t, expected, string(actual))
}
