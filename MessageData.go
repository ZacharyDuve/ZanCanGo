package zancango

import "fmt"

const (
	data_identifier_mask uint32 = 0x0FFF_FFFF
)

type messageData uint32

type MessageDataFrame struct {
	dataIdentifierLenBytes uint8
	dataIdentifier         uint32
	dataLenBytes           uint8
	data                   messageData
}

func NewMessageDataFrame(dataIdentifier uint32, data messageData) (*MessageDataFrame, error) {
	// First need to check that the dataIdentifier is valid

	if dataIdentifier&data_identifier_mask != 0 {
		return nil, fmt.Errorf("invalid dataIdentifier: %x, due to it exceeding mask %x", dataIdentifier, data_identifier_mask)
	}

	messageDataFrame := &MessageDataFrame{dataIdentifier: dataIdentifier, data: data}

	// Need to computer dataIdentifierLen

	for dataIdentCopy := dataIdentifier; dataIdentCopy != 0; dataIdentCopy >>= 8 {
		messageDataFrame.dataIdentifierLenBytes++
	}

	// Need to compute dataLen

	for dataCopy := uint32(data); dataCopy != 0; dataCopy >>= 8 {
		messageDataFrame.dataLenBytes++
	}

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

}
