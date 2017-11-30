package encrpt

import (
	"errors"
)

type StoreUnit interface {
	getKey() []byte
	getBlock(index int) ([]byte, [][]byte, error)
	setBlock(index int, data []byte) error
	getMT() MerkleTree
	getTopMTHash() []byte
}

type RamStoreUnit struct {
	data []byte
	key  []byte
	mt   MerkleTree
}

func NewRamStoreUnit(size uint, key []byte) (*RamStoreUnit, error) {
	rsu := new(RamStoreUnit)
	rsu.data = make([]byte, 1<<size)
	rsu.key = key
	mt, err := CreateMerkleTree(rsu.data, BLOCK_SIZE)
	if err != nil {
		return nil, errors.New("Could not build MerkleTree")
	}
	rsu.mt = mt
	return rsu, nil
}

func (rsu *RamStoreUnit) getKey() []byte {
	return rsu.key
}

func (rsu *RamStoreUnit) getBlock(index int) ([]byte, [][]byte, error) {
	block := rsu.data[index*BLOCK_SIZE : (index+1)*BLOCK_SIZE]
	data, err := DecryptBlock(block, rsu.key, []byte{byte(index)})
	if err != nil {
		return nil, nil, err
	}
	mtPath, err := rsu.mt.GetPathFromLeaf(index)
	if err != nil {
		return nil, nil, err
	}
	return data, mtPath, nil
}

func (rsu *RamStoreUnit) setBlock(index int, data []byte) error {
	block, err := EncryptBlock(data, rsu.key, []byte{byte(index)})
	if err != nil {
		return err
	}
	for i, val := range block {
		rsu.data[index*BLOCK_SIZE+i] = val
	}

	err = rsu.mt.ChangeLeaf(index, simpleSha256(block))
	if err != nil {
		return err
	}
	return nil
}

func (rsu *RamStoreUnit) getMT() MerkleTree {
	return rsu.mt
}

func (rsu *RamStoreUnit) getTopMTHash() []byte {
	return rsu.mt[0]
}
