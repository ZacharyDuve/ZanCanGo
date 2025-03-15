package zancango

import "fmt"

const (
	data_identifier_mask         uint32 = 0x0FFF_FFFF
	data_identifier_inverse_mask uint32 = 0xFFFF_FFFF ^ data_identifier_mask
)

type MessageDataFrame struct {
	dataIdentifierLenBytes int
	dataIdentifier         uint32
	dataLenBytes           int
	data                   messageData
}

func NewMessageDataFrame(dataIdentifier uint32, data messageData) (*MessageDataFrame, error) {
	// First need to check that the dataIdentifier is valid

	if dataIdentifier&data_identifier_inverse_mask != 0 {
		return nil, fmt.Errorf("invalid dataIdentifier: %x, due to it exceeding mask %x", dataIdentifier, data_identifier_mask)
	}

	messageDataFrame := &MessageDataFrame{dataIdentifier: dataIdentifier, data: data}

	// Need to computer dataIdentifierLen

	for dataIdentCopy := dataIdentifier; dataIdentCopy != 0; dataIdentCopy >>= 8 {
		messageDataFrame.dataIdentifierLenBytes++
	}

	// Need to compute dataLen

	messageDataFrame.dataLenBytes = len(data)

	return messageDataFrame, nil
}

func (this *MessageDataFrame) DataIdentifier() uint32 {
	return this.dataIdentifier
}

func (this *MessageDataFrame) Data() messageData {
	return this.data
}

func (this *MessageDataFrame) ToBytes() []byte {
	bytes := make([]byte, this.dataIdentifierLenBytes+this.dataLenBytes)
	// Need to make a copy so that we don't modify the original dataIdentifier
	dataIdentifierCopy := this.dataIdentifier
	for i := this.dataIdentifierLenBytes - 1; i >= 0; i-- {
		bytes[i] = byte(dataIdentifierCopy)
		dataIdentifierCopy >>= 8

		if i == 0 {
			// Need to set the first byte to the dataIdentifierLenBytes
			dataLenBits := byte(this.dataLenBytes-1) << 6
			bytes[0] = bytes[0] | dataLenBits
		}
	}

	for i := 0; i < this.dataLenBytes; i++ {
		bytes[this.dataIdentifierLenBytes+i] = this.data[i]
	}

	return bytes
}

func MessageDataFrameFromBytes(data []byte) (*MessageDataFrame, error) {
	// First need to check that the dataIdentifier is valid

	// We need a minimum of 2 bytes to have a valid data frame
	if len(data) < 2 {
		return nil, fmt.Errorf("invalid data: %v, due to it being empty", data)
	}

	dataLength := int(data[0]&0xC0>>6) + 1

	dataIdentifierLen := len(data) - dataLength

	dataIdentifier := uint32(0)

	for i := 0; i < dataIdentifierLen; i++ {
		dataIdentifier |= uint32(data[i]) << uint(8*i)
	}

	return NewMessageDataFrame(dataIdentifier, MessageDataFromByteSlice(data[dataIdentifierLen:]))
}
