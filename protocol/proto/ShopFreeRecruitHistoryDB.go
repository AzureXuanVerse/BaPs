// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ShopFreeRecruitHistoryDB struct {
	UniqueId       int64
	RecruitCount   int32
	LastUpdateDate mx.MxTime
}

func (x *ShopFreeRecruitHistoryDB) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}
