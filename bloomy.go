package bloomy

import (
	"math"

	"github.com/bits-and-blooms/bitset"
	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	m      uint64
	k      uint32
	bitset *bitset.BitSet
}

func New(item uint32, falsePositiveRate float64) *BloomFilter {
	ln22 := math.Ln2 * math.Ln2
	m := uint64(math.Ceil((float64(-1) * float64(item) * math.Log(falsePositiveRate)) / ln22))
	k := uint32(math.Ceil((float64(-1) * math.Log(falsePositiveRate)) / math.Ln2))

	return &BloomFilter{
		m:      m,
		k:      k,
		bitset: bitset.New(uint(m)),
	}
}

func (bf *BloomFilter) getIndex(a, b, i uint64) uint64 {
	return (a + (b * i)) % bf.m
}

func (bf *BloomFilter) Insert(value []byte) {
	a, b := murmur3.Sum128(value)

	for k := uint32(0); k < bf.k; k += 1 {
		index := bf.getIndex(a, b, uint64(k))
		bf.bitset.Set(uint(index))
	}
}

func (bf *BloomFilter) Contains(value []byte) bool {
    a, b:= murmur3.Sum128(value)
}

func (bf *BloomFilter) Contains(value []byte) bool {
	a, b := murmur3.Sum128(value)
	for k := uint32(0); k < bf.k; k += 1 {
		idx := bf.getIndex(a, b, uint64(k))
		if !bf.bitset.Test(uint(idx)) {
			return false
		}
	}
	return true
}
