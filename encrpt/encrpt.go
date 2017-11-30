package encrpt

import (
	"crypto/aes"
	"errors"
)

const BLOCK_SIZE = 1024
const CHUNK_SIZE = 16

func EncryptBlock(data, key, blockID []byte) ([]byte, error) {
	if len(data) != BLOCK_SIZE {
		return nil, errors.New("Incorrect data size. Must be 2^10")
	}

	cip, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	block := make([]byte, BLOCK_SIZE)
	for i := 0; i < BLOCK_SIZE/CHUNK_SIZE; i++ {
		blockLicationKey := sliceTo(simpleSha256(sliceByteXOR(blockID, byte(i))), CHUNK_SIZE)
		dest := xorBytes(blockLicationKey, data[i*CHUNK_SIZE:])
		cip.Encrypt(block[i*CHUNK_SIZE:], dest)
	}
	return block, nil
}

func DecryptBlock(block, key, blockID []byte) ([]byte, error) {
	if len(block) != BLOCK_SIZE {
		return nil, errors.New("Incorrect data size. Must be 2^10")
	}

	cip, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	data := make([]byte, BLOCK_SIZE)
	for i := 0; i < BLOCK_SIZE/CHUNK_SIZE; i++ {
		blockLicationKey := sliceTo(simpleSha256(sliceByteXOR(blockID, byte(i))), CHUNK_SIZE)
		dest := make([]byte, CHUNK_SIZE)
		cip.Decrypt(dest, block[i*CHUNK_SIZE:])
		for j, val := range xorBytes(dest, blockLicationKey) {
			data[16*i+j] = val
		}
	}
	return data, nil
}
