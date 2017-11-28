package encrpt

import (
	"crypto/aes"
	"errors"
)

func EncryptBlock(data, key, blockID []byte) ([]byte, error) {
	if len(data) != 1<<10 {
		return nil, errors.New("Incorrect data size. Must be 2^10")
	}

	cip, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	block := make([]byte, 1<<10)
	for i := 0; i < (1<<10)/(1<<4); i++ {
		blockLicationKey := sliceTo(simpleSha256(sliceByteXOR(blockID, byte(i))), 16)
		dest := xorBytes(blockLicationKey, data[i*16:])
		cip.Encrypt(block[i*16:], dest)
	}
	return block, nil
}

func DecryptBlock(block, key, blockID []byte) ([]byte, error) {
	if len(block) != 1<<10 {
		return nil, errors.New("Incorrect data size. Must be 2^10")
	}

	cip, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	data := make([]byte, 1<<10)
	for i := 0; i < (1<<10)/(1<<4); i++ {
		blockLicationKey := sliceTo(simpleSha256(sliceByteXOR(blockID, byte(i))), 16)
		dest := make([]byte, 16)
		cip.Decrypt(dest, block[i*16:])
		for j, val := range xorBytes(dest, blockLicationKey) {
			data[16*i+j] = val
		}
	}
	return data, nil
}
