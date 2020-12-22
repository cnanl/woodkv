package types

type Method uint

const (
	Mget Method = iota
	Mput
	Mdelete
	MprefixScan
)

type Byte8 [8]byte
type Byte256 [256]byte

func (this Byte8) Less(item interface{}) bool {
	other := item.(Byte8)
	for i := 0; i < 8; i++ {
		if this[i] != other[i] {
			return this[i] < other[i]
		}
	}
	return false //return false if equal
}
func (by256 Byte256) V() {}

type KVRequest struct {
	Method Method
	Key    Byte8
	Value  Byte256
}

type KVReply struct {
	Value Byte256
}

type PrefixScanRequest struct {
	N int
}

type PrefixScanReply struct {
	Keys []Byte8
}

func (m Method) String() string {
	switch m {
	case Mget:
		return "Get"
	case Mput:
		return "Put"
	case Mdelete:
		return "Delete"
	case MprefixScan:
		return "PrefixScan"
	}
	return "error method"
}
