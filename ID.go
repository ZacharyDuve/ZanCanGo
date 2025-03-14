package zancango

import "fmt"

const (
	// id_mask is 11 bits long
	id_mask ID = 0x0000_07FF
)

type ID uint32

func IDFromTypeAndAddress(t MessageType, a Address) (ID, error) {
	isValidMessageType := ValidateFrameType(t)

	if !isValidMessageType {
		return 0, fmt.Errorf("unable to create new ID due to invalid MessageType: %v", t)
	}

	return id_mask & (ID(t)<<8 | ID(a)), nil
}

func (this ID) MessageType() (MessageType, error) {

	mt := MessageType(this >> 8)

	if !ValidateFrameType(mt) {
		return 0, fmt.Errorf("invalid MessageType: %v", mt)
	}
	return mt, nil
}

func (this ID) Address() Address {
	// Since address is only 8 bits long, we can just cast it to an Address
	return Address(this)
}
