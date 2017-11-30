package encrpt

import (
	//"bytes"
	"errors"
	"math"
)

type MerkleTree [][]byte

func isPowerOfTwo(n int) bool {
	return (n != 0 && ((n & (n - 1)) == 0))
}

func buildMerkleTree(mt MerkleTree, i int) []byte {
	if len(mt[i-1]) != 0 {
		return mt[i-1]
	}
	mt[i-1] = simpleSha256(xorBytes(buildMerkleTree(mt, 2*i), buildMerkleTree(mt, 2*i+1)))
	return mt[i-1]
}

func (mt *MerkleTree) ChangeNode(i int, to []byte) error {
	if len(*mt) <= i || i < 0 {
		return errors.New("Access outside of bound in Merkletree")
	}

	(*mt)[i] = to

	if i == 0 {
		return nil
	}

	if i%2 == 1 {
		return mt.ChangeNode(i/2, simpleSha256(xorBytes((*mt)[i-1], to)))
	} else {
		return mt.ChangeNode(i/2, simpleSha256(xorBytes((*mt)[i+1], to)))
	}
}

func (mt *MerkleTree) ChangeLeaf(i int, to []byte) error {
	return mt.ChangeNode((len(*mt)-3)/2+i, to)
}

func (mt MerkleTree) GetPathFrom(index int) ([][]byte, error) {
	if len(mt) <= index || index < 0 {
		return nil, errors.New("Access outside of bound in Merkletree")
	}

	l := uint(math.Floor(math.Log2(float64(index)) + 1))
	path := make([][]byte, l)
	for i := uint(0); i < l; i++ {
		newI := ((index) / (1 << i))
		if newI%2 == 1 {
			path[i] = mt[newI-1]
		} else {
			path[i] = mt[newI+1]
		}
	}
	return path, nil
}

func (mt MerkleTree) GetPathFromLeaf(index int) ([][]byte, error) {
	return mt.GetPathFrom((len(mt)-3)/2 + index)
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
	buildMerkleTree(mt, 1)

	return mt, nil
}
