// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type RuntimeFieldHandle struct {
	*ValueType

	Value *IntPtr
}

func (x *RuntimeFieldHandle) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *RuntimeFieldHandle) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ValueType = packet.(*ValueType)
}
