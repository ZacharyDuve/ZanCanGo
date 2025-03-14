package zancango

type messageData []byte

func MessageDataFromByteSlice(data []byte) messageData {
	return messageData(data)
}

func (this messageData) ToByteSlice() []byte {
	return []byte(this)
}

func MessageDataFromByte(data byte) messageData {
	return messageData([]byte{data})
}

func (this messageData) ToByte() byte {
	return this[0]
}

func MessageDataFromInt16(data int16) messageData {
	return messageData([]byte{byte(data >> 8), byte(data)})
}

func (this messageData) ToInt16() int16 {
	return int16(this[0])<<8 | int16(this[1])
}

func MessageDataFromInt32(data int32) messageData {
	return messageData([]byte{byte(data >> 24), byte(data >> 16), byte(data >> 8), byte(data)})
}

func (this messageData) ToInt32() int32 {
	return int32(this[0])<<24 | int32(this[1])<<16 | int32(this[2])<<8 | int32(this[3])
}

func MessageDataFromUInt16(data uint16) messageData {
	return messageData([]byte{byte(data >> 8), byte(data)})
}

func (this messageData) ToUInt16() uint16 {
	return uint16(this[0])<<8 | uint16(this[1])
}

func MessageDataFromUInt32(data uint32) messageData {
	return messageData([]byte{byte(data >> 24), byte(data >> 16), byte(data >> 8), byte(data)})
}

func (this messageData) ToUInt32() uint32 {
	return uint32(this[0])<<24 | uint32(this[1])<<16 | uint32(this[2])<<8 | uint32(this[3])
}

// Length of the messageData in bytes
func (this messageData) LenBytes() int {
	return len(this)
}
