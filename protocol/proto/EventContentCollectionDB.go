// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type EventContentCollectionDB struct {
	EventContentId int64
	GroupId        int64
	UniqueId       int64
	ReceiveDate    mx.MxTime
}

func (x *EventContentCollectionDB) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}
