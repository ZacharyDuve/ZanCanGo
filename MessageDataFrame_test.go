package zancango

import (
	"bytes"
	"testing"
)

func TestMessageDataFrameReturnsErrorWhenIdentifierIsLargerThan28Bits(t *testing.T) {
	_, err := NewMessageDataFrame(0x1FFF_FFFF, MessageDataFromByteSlice([]byte{0x01}))
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestMessageDataFrameReturnsNoErrorWhenIdentifierIs28BitsOrLess(t *testing.T) {
	_, err := NewMessageDataFrame(0x0FFF_FFFF, MessageDataFromByteSlice([]byte{0x01}))
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestMessageDataToBytesCorrectlyEncodes1ByteIdentifierAnd1ByteData(t *testing.T) {
	var identifier uint32 = 0x32
	var dataByte byte = 0xAA
	data := MessageDataFromByteSlice([]byte{dataByte})
	mdf, err := NewMessageDataFrame(identifier, data)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	expected := []byte{byte(identifier), dataByte}
	if !bytes.Equal(mdf.ToBytes(), expected) {
		t.Errorf("Expected %v, got %v", expected, mdf.ToBytes())
	}
}

func TestMessageDataToBytesCorrectlyEncodes1ByteIdentifierAnd2ByteData(t *testing.T) {
	var identifier uint32 = 0x32
	messageData := MessageDataFromUInt16(0xACDC)
	mdf, err := NewMessageDataFrame(identifier, messageData)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	expected := []byte{0b0100_0000 | byte(identifier), 0xAC, 0xDC}
	if !bytes.Equal(mdf.ToBytes(), expected) {
		t.Errorf("Expected %v, got %v", expected, mdf.ToBytes())
	}
}

func TestMessageDataToBytesCorrectlyEncodes1ByteIdentifierAnd3ByteData(t *testing.T) {
	var identifier uint32 = 0x32
	messageData := MessageDataFromUInt32(0xF4ACDC)
	mdf, err := NewMessageDataFrame(identifier, messageData)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	// Message Frame should transmit minimum number of bytes so u32 should be encoded as u24 since leading byte is 0
	expected := []byte{0b1000_0000 | byte(identifier), 0xF4, 0xAC, 0xDC}
	if !bytes.Equal(mdf.ToBytes(), expected) {
		t.Errorf("Expected %v, got %v", expected, mdf.ToBytes())
	}
}
