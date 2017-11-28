package encrpt_test

import (
	"bytes"
	"crypto/rand"
	"github.com/drpytho/encrptCloud/encrpt"
	"testing"
)

func TestEncrypringRandomData(t *testing.T) {
	for i := 0; i < 1000; i++ {
		dat := make([]byte, 1024)
		key := make([]byte, 32)
		blockID := []byte("BlockID")
		_, err := rand.Read(dat)
		if err != nil {
			t.Skip("Failed to generate random data")
		}
		_, err1 := rand.Read(key)
		if err1 != nil {
			t.Skip("Failed to generate random key")
		}
		block, err := encrpt.EncryptBlock(dat, key, blockID)
		if err != nil {
			t.Error(err)
		}
		data, err := encrpt.DecryptBlock(block, key, blockID)
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal(dat, data) {
			t.Error("Data did not match")
		}
	}
}
