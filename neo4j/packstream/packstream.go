package packstream

type Packer interface {
	Nil()
	Bool(b bool)
	Int(i int)
	Int8(i int8)
	Int16(i int16)
	Int32(i int32)
	Int64(i int64)
	Uint(i uint)
	Uint8(i uint8)
	Uint16(i uint16)
	Uint32(i uint32)
	Uint64(i uint64)
	Float32(f float32)
	Float64(f float64)
	String(s string)
	Bytes(b []byte)
	MapHeader(l int)
	ArrayHeader(l int)
	StructHeader(tag byte, num int)
}
