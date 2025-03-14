package zancango

import (
	"github.com/ZacharyDuve/godatacollections"
	"github.com/ZacharyDuve/godatacollections/tree"
)

type MessageType uint8

const (
	FrameTypeEmergency MessageType = iota
	//Skip 0x02 for now
	_
	FrameTypeError
	FrameTypeTime
	FrameTypeSendData
	FrameTypeRequestData
	FrameTypeSetData
)

var validMessageTypes godatacollections.Set[MessageType, MessageType]

func compareMessageTypes(a, b MessageType) int {
	return int(a) - int(b)
}

func init() {
	var err error
	validMessageTypes, err = tree.NewBST[MessageType, MessageType](compareMessageTypes, func(k MessageType) MessageType { return k }, FrameTypeEmergency)

	if err != nil {
		panic(err)
	}
	// Manually inserting the Message type to utilize BST to be a tree instead of a list
	// Yes this is where something line a RB Tree or an AVL tree would be better

	// Putting error first to ensure that it gets fastest
	validMessageTypes.Insert(FrameTypeEmergency)
	validMessageTypes.Insert(FrameTypeSendData)
	validMessageTypes.Insert(FrameTypeError)
	validMessageTypes.Insert(FrameTypeTime)
	validMessageTypes.Insert(FrameTypeRequestData)
	validMessageTypes.Insert(FrameTypeSetData)
}

func ValidateFrameType(mT MessageType) bool {
	return validMessageTypes.Contains(mT)
}
