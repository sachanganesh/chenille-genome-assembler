package debruijn

import (
	"strings"
)

// ===================================
// Kmer
// ===================================

type Kmer struct {
	data		uint64
	length		uint8
}

func NewKmer(kmer string) Kmer {
	return Kmer{
		data: ConvertStringToUInt64(PreprocessString(kmer)),
		length: uint8(len(kmer)),
	}
}

// ===================================
// Kmer Functions
// ===================================

func (kmer *Kmer) RawLen() uint8 {
	return kmer.length
}

func (kmer *Kmer) Len() int {
	return int(kmer.length)
}

func (kmer *Kmer) GetValue() uint64 {
	return kmer.data
}

func (kmer *Kmer) GetLastNucleotide() int {
	return int(kmer.GetValue() & 0x03)
}

func (kmer *Kmer) String() string {
	alphabet := [4]string{"A", "C", "G", "T"}
	var rep strings.Builder

	i := 0
	for i < kmer.Len() {
		var tmp uint64 = (kmer.data >> uint64(i * 2)) & 0x03
		rep.WriteString(alphabet[tmp])
		i++
	}

	return rep.String()
}

func (kmer *Kmer) Equals(other Kmer) bool {
	return kmer.data == other.GetValue()
}

func (kmer *Kmer) GeneratePredecessor(nt byte) Kmer {
	alphabet := [4]byte{'A', 'C', 'G', 'T'}

	var i uint64
	for k := range alphabet {
		if alphabet[k] == nt {
			i = uint64(k)
			break
		}
	}

	tmp := GenerateNOneBits(kmer.Len() - 2)
	tmp = kmer.data & tmp

	pred := (i << kmer.RawLen() - 2) | tmp
	return Kmer{pred, kmer.RawLen()}
}

func (kmer *Kmer) GenerateSuccessor(nt byte) Kmer {
	alphabet := [4]byte{'A', 'C', 'G', 'T'}

	var i uint64
	for k := range alphabet {
		if alphabet[k] == nt {
			i = uint64(k)
			break
		}
	}

	tmp := GenerateNOneBits(kmer.Len() - 2) << 2
	tmp = kmer.data & tmp

	succ := i | tmp
	return Kmer{succ, kmer.RawLen()}
}

// ===================================
// Kmer Utilities
// ===================================

func GenerateNOneBits(n int) uint64 {
	var v uint64 = 0

	var i uint64 = 0
	for i < uint64(n) {
		v = (v << 1) | 0x01
		i++
	}

	return v
}

func ConvertStringToUInt64(kmer string) uint64 {
	var rep uint64

	for i := range kmer {
		var tmp uint64

		switch ch := kmer[i]; ch {
		case 'A':
			tmp = 0
		case 'C':
			tmp = 1
		case 'G':
			tmp = 2
		case 'T':
			tmp = 3
		}

		rep = (rep << 2) | tmp
	}

	return rep
}

func PreprocessString(kmer string) string {
	var tmp strings.Builder

	for _, nt := range kmer {
		if nt != 'A' && nt != 'C' && nt != 'G' && nt != 'T' {
			tmp.WriteString("A")
		} else {
			tmp.WriteString(string(nt))
		}
	}

	return tmp.String()
}
