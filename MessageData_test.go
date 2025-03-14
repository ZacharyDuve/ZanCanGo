package zancango

import (
	"bytes"
	"testing"
)

func TestMessageDataFromByteSlice(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03}
	md := MessageDataFromByteSlice(data)
	if !bytes.Equal(md.ToByteSlice(), data) {
		t.Errorf("Expected %v, got %v", data, md.ToByteSlice())
	}
}

func TestMessageDataFromByte(t *testing.T) {
	data := byte(0x01)
	md := MessageDataFromByte(data)
	if md.ToByte() != data {
		t.Errorf("Expected %v, got %v", data, md.ToByte())
	}
}

func TestMessageDataFromInt16(t *testing.T) {
	data := int16(0x0102)
	md := MessageDataFromInt16(data)
	if md.ToInt16() != data {
		t.Errorf("Expected %v, got %v", data, md.ToInt16())
	}
}

func TestMessageDataFromInt32(t *testing.T) {
	data := int32(0x01020304)
	md := MessageDataFromInt32(data)
	if md.ToInt32() != data {
		t.Errorf("Expected %v, got %v", data, md.ToInt32())
	}
}

func TestMessageDataFromUInt16(t *testing.T) {
	data := uint16(0x0102)
	md := MessageDataFromUInt16(data)
	if md.ToUInt16() != data {
		t.Errorf("Expected %v, got %v", data, md.ToUInt16())
	}
}

func TestMessageDataFromUInt32(t *testing.T) {
	data := uint32(0x01020304)
	md := MessageDataFromUInt32(data)
	if md.ToUInt32() != data {
		t.Errorf("Expected %v, got %v", data, md.ToUInt32())
	}
}

func TestMessageDataLenBytes(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03, 0x04}
	md := MessageDataFromByteSlice(data)
	if md.LenBytes() != len(data) {
		t.Errorf("Expected %v, got %v", len(data), md.LenBytes())
	}
}
