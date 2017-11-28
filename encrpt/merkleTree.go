package encrpt

import (
	"bytes"
	"errors"
)

type MerkleTree [][]byte

func isPowerOfTwo(n int) bool {
	return (n != 0 && ((n & (n - 1)) == 0))
}

func buildMerkleTree(mt MerkleTree, i int) []byte {
	if len(mt[i]) != 0 {
		return mt[i]
	}
	mt[i] = simpleSha256(xorBytes(buildMerkleTree(mt, 2*i), buildMerkleTree(mt, 2*i+1)))
	return mt[i]
}

func CreateMerkleTree(data []byte, blockSize int) (MerkleTree, error) {
	if len(data)%blockSize != 0 {
		return nil, errors.New("The size of the data must be a multiple of the blockSize")
	}
	N := len(data) / blockSize
	// N Should be a power of two.
	if !isPowerOfTwo(N) {
		return nil, errors.New("The number of blocks must be a power of two, it was " + string(N))
	}

	mt := make(MerkleTree, 2*N-1)
	for i := 0; i < N; i++ {
		mt[N-2+i] = simpleSha256(data[i*blockSize : (i+1)*blockSize])
	}
	buildMerkleTree(mt, 0)

}
