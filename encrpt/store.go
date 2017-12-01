package encrpt

import (
	"crypto/rand"
	//"errors"
)

type StoreUnit interface {
	GetKey() []byte
	GetBlock(index int) ([]byte, [][]byte, error)
	SetBlock(index int, data []byte) error
	GetMT() MerkleTree
	GetTopMTHash() []byte
}

type Store interface {
	NewUnit(size uint, key []byte) (*string, error)
	GetUnit(string) StoreUnit
}

type RamStore struct {
	units map[string]RamStoreUnit
}

func NewRamStore() *RamStore {
	rs := new(RamStore)
	rs.units = make(map[string]RamStoreUnit)
	return rs
}

func (rs *RamStore) NewUnit(size uint, key []byte) (*string, error) {
	id := make([]byte, 12)

	for {
		rand.Read(id)
		for i, val := range id {
			id[i] = 'a' + val%('z'-'a'+1)
		}
		_, ok := rs.units[string(id)]
		if !ok {
			break
		}
	}
	strID := string(id)
	rsu, err := NewRamStoreUnit(size, key)
	if err != nil {
		return nil, err
	}
	rs.units[strID] = *rsu
	return &strID, nil
}

func (rs *RamStore) GetUnit(strID string) StoreUnit {
	val, ok := rs.units[strID]
	if ok {
		ret := StoreUnit(&val)
		return ret
	}
	return nil
}
