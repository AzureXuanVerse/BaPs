// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type ShopBuyMerchandiseResponse struct{
	message ProtoMessage
	ResponsePacket

    AccountCurrencyDB *AccountCurrencyDB
    ConsumeResultDB *ConsumeResultDB
    ParcelResultDB *ParcelResultDB
    MailDB *MailDB
    ShopProductDB *ShopProductDB
}

func (x *ShopBuyMerchandiseResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ShopBuyMerchandiseResponse) ProtoReflect() Message {
	return x
}

func (x *ShopBuyMerchandiseResponse) GetProtocol() int32 {
	return mx.Protocol_Shop_BuyMerchandise
}

func (x *ShopBuyMerchandiseResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

