package gguf_info

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type (
	GGUF struct {
		Header Header
	}
	Header struct {
		Magic           uint32
		Version         uint32
		TensorCount     uint64
		MetadataKvCount uint64
		MetadataKV      []MetadataKV
	}
	MetadataKV struct {
		Key               String
		MetadataValueType ValueType
		Value             Value
	}
	TensorInfo struct {
		Name        string   // it must be at most 64 bytes long
		NDimensions uint32   // currently at most 4
		Dimensions  []uint64 // e.g. [4096, 32000]
		Type        Type
		Offset      uint64
	}
	String struct {
		Len  uint64
		Char []byte
	}
	Array struct {
		Type   ValueType
		Len    uint64
		Values []any
	}
	Type      uint32
	ValueType uint32
	Value     any
)

// New creates a gguf file format Reader
// Currently expects a Little Endian encoding.
func New(path string) (*GGUF, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	g := &GGUF{}
	if err := g.readHeader(f); err != nil {
		return nil, err
	}
	return g, nil
}

func (g *GGUF) readFloat64(f *os.File) (float64, error) {
	v := float64(0)
	return v, binary.Read(f, binary.LittleEndian, &v)
}

func (g *GGUF) readFloat32(f *os.File) (float32, error) {
	v := float32(0)
	return v, binary.Read(f, binary.LittleEndian, &v)
}

func (g *GGUF) readInt64(f *os.File) (int64, error) {
	v, err := g.readUint64(f)
	return int64(v), err
}

func (g *GGUF) readUint64(f *os.File) (uint64, error) {
	v := make([]byte, 8)
	_, err := f.Read(v)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(v), err
}

func (g *GGUF) readInt32(f *os.File) (int32, error) {
	v, err := g.readUint32(f)
	return int32(v), err
}

func (g *GGUF) readUint32(f *os.File) (uint32, error) {
	v := make([]byte, 4)
	_, err := f.Read(v)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(v), err
}

func (g *GGUF) readInt16(f *os.File) (int16, error) {
	v, err := g.readUint16(f)
	return int16(v), err
}

func (g *GGUF) readUint16(f *os.File) (uint16, error) {
	v := make([]byte, 2)
	_, err := f.Read(v)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(v), err
}

func (g *GGUF) readInt8(f *os.File) (int8, error) {
	v, err := g.readUint8(f)
	return int8(v), err
}

func (g *GGUF) readUint8(f *os.File) (uint8, error) {
	v := make([]byte, 1)
	_, err := f.Read(v)
	if err != nil {
		return 0, err
	}
	return v[0], err
}

func (g *GGUF) readBool(f *os.File) (bool, error) {
	v, err := g.readUint8(f)
	return v != 0, err
}

func (g *GGUF) readString(f *os.File) (String, error) {
	var err error
	var s String
	s.Len, err = g.readUint64(f)
	if err != nil {
		return s, err
	}
	s.Char = make([]byte, s.Len)
	if _, err := f.Read(s.Char); err != nil {
		return s, err
	}
	return s, nil
}

func (g *GGUF) readArray(f *os.File) (Array, error) {
	var a Array
	t, err := g.readUint32(f)
	if err != nil {
		return a, err
	}
	a.Type = ValueType(t)
	a.Len, err = g.readUint64(f)
	a.Values = make([]any, a.Len)
	for i := 0; i < int(a.Len); i++ {
		a.Values[i], err = g.readValueType(f, a.Type)
	}
	return a, err
}

func (g *GGUF) readValueType(f *os.File, t ValueType) (any, error) {
	switch t {
	case GGUF_METADATA_VALUE_TYPE_UINT8:
		return g.readUint8(f)
	case GGUF_METADATA_VALUE_TYPE_INT8:
		return g.readInt8(f)
	case GGUF_METADATA_VALUE_TYPE_UINT16:
		return g.readUint16(f)
	case GGUF_METADATA_VALUE_TYPE_INT16:
		return g.readInt16(f)
	case GGUF_METADATA_VALUE_TYPE_UINT32:
		return g.readUint32(f)
	case GGUF_METADATA_VALUE_TYPE_INT32:
		return g.readInt32(f)
	case GGUF_METADATA_VALUE_TYPE_FLOAT32:
		return g.readFloat32(f)
	case GGUF_METADATA_VALUE_TYPE_BOOL:
		return g.readBool(f)
	case GGUF_METADATA_VALUE_TYPE_STRING:
		return g.readString(f)
	case GGUF_METADATA_VALUE_TYPE_ARRAY:
		return g.readArray(f)
	case GGUF_METADATA_VALUE_TYPE_UINT64:
		return g.readUint64(f)
	case GGUF_METADATA_VALUE_TYPE_INT64:
		return g.readInt64(f)
	case GGUF_METADATA_VALUE_TYPE_FLOAT64:
		return g.readFloat64(f)
	default:
		return nil, fmt.Errorf("unknown metadata value type: %v", t)
	}
}

func (g *GGUF) readHeader(f *os.File) error {
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}
	headerPart := make([]byte, 4+4+8+8)
	if _, err := f.Read(headerPart); err != nil {
		return err
	}
	g.Header.Magic = binary.LittleEndian.Uint32(headerPart[0:4])
	g.Header.Version = binary.LittleEndian.Uint32(headerPart[4:8])
	g.Header.TensorCount = binary.LittleEndian.Uint64(headerPart[8:16])
	g.Header.MetadataKvCount = binary.LittleEndian.Uint64(headerPart[16:24])
	for i := 0; i < int(g.Header.MetadataKvCount); i++ {
		var err error
		mkv := MetadataKV{}
		mkv.Key, err = g.readString(f)
		if err != nil {
			return err
		}
		vt, err := g.readUint32(f)
		if err != nil {
			return err
		}
		mkv.Value, err = g.readValueType(f, ValueType(vt))
		if err != nil {
			return err
		}
		g.Header.MetadataKV = append(g.Header.MetadataKV, mkv)
	}

	return nil
}

func (g *GGUF) MetadataValue(key string, w io.Writer) error {
	for _, v := range g.Header.MetadataKV {
		if string(v.Key.Char) == key {
			return writeValue(v.Value, -1, w)
		}
	}
	return nil
}

func (g *GGUF) Out(w io.Writer) error {
	var err error
	magic := make([]byte, 4)
	binary.LittleEndian.PutUint32(magic, g.Header.Magic)

	_, err = fmt.Fprintf(w, "Header:\n")
	_, err = fmt.Fprintf(w, "\tMagic:            %s\n", string(magic))
	_, err = fmt.Fprintf(w, "\tVersion:          %d\n", g.Header.Version)
	_, err = fmt.Fprintf(w, "\tTensor Count:     %d\n", g.Header.TensorCount)
	_, err = fmt.Fprintf(w, "\tMetadata Entries: %d\n", g.Header.MetadataKvCount)
	if g.Header.MetadataKvCount > 0 {
		_, err = fmt.Fprintf(w, "\tMetadata:\n")
	}
	for _, m := range g.Header.MetadataKV {
		_, err = fmt.Fprintf(w, "\t\t* %s=", string(m.Key.Char))
		err = writeValue(m.Value, 50, w)
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("\n"))
	}

	return err
}

func writeValue(t any, maxLen int, w io.Writer) error {
	var err error
	switch v := t.(type) {
	case int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		_, err = fmt.Fprintf(w, "%d", v)
	case float32, float64:
		_, err = fmt.Fprintf(w, "%f", v)
	case bool:
		_, err = fmt.Fprintf(w, "%t", v)
	case String:
		_, err = fmt.Fprintf(w, "%s", v.Char)
	case Array:
		if maxLen < 0 || int(v.Len) < maxLen {
			return json.NewEncoder(w).Encode(v.Values)
		} else {
			_, err = w.Write([]byte("...truncated..."))
		}
	default:
		_, err = fmt.Fprintf(w, "%s", "-")
	}
	return err
}

func (a Array) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Values)
}

func (s String) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s.Char))
}
