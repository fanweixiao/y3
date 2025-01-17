package y3

import (
	"github.com/yomorun/y3/utils"
)

// Tag represents the Tag of TLV,
// MSB used to represent the packet type, 0x80 means a node packet, otherwise is a primitive packet.
// Low 7 bits represent Sequence ID, like `key` in JSON format
type Tag struct {
	raw byte
}

// IsNode returns true is MSB is 1.
func (t *Tag) IsNode() bool {
	return t.raw&utils.MSB == utils.MSB
}

// IsSlice determine if the current node is a Slice
func (t *Tag) IsSlice() bool {
	return t.raw&utils.SliceFlag == utils.SliceFlag
}

// SeqID get the sequence ID, as key in JSON format
func (t *Tag) SeqID() byte {
	//return t.raw & packetutils.DropMSB
	return t.raw & utils.DropMSBArrayFlag
}

// NewTag create a NodePacket Tag field
func NewTag(b byte) *Tag {
	return &Tag{raw: b}
}

// Raw return the original byte
func (t *Tag) Raw() byte {
	return t.raw
}
