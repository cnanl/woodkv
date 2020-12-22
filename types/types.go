package types

//Byte8 implement store.Key interface
type Byte8 [8]byte

func (this Byte8) Less(item interface{}) bool {
	other := item.(Byte8)
	for i := 0; i < 8; i++ {
		if this[i] != other[i] {
			return this[i] < other[i]
		}
	}
	return false //return false if equal
}

//Byte256 implement store.Value interface
type Byte256 [256]byte

func (by256 Byte256) V() {}

type Method uint

const (
	Mget Method = iota
	Mput
	Mdelete
	MprefixScan
)

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

//These two structs are prepared for Get,Put,Delete RPC
//In fact, we can prepare different structs for each method. It will be better.
type KVRequest struct {
	Method Method
	Key    Byte8
	Value  Byte256
}

type KVReply struct {
	Value Byte256
}

//for PrefixScan RPC
type PrefixScanRequest struct {
	N int
}

type PrefixScanReply struct {
	Keys []Byte8
}
