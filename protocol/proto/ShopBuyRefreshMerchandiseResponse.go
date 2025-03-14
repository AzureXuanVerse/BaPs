// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type ShopBuyRefreshMerchandiseResponse struct{
	message ProtoMessage
	ResponsePacket

    AccountCurrencyDB *AccountCurrencyDB
    ConsumeResultDB *ConsumeResultDB
    ParcelResultDB *ParcelResultDB
    ShopProductDB []*ShopProductDB
    MailDB *MailDB
}

func (x *ShopBuyRefreshMerchandiseResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ShopBuyRefreshMerchandiseResponse) ProtoReflect() Message {
	return x
}

func (x *ShopBuyRefreshMerchandiseResponse) GetProtocol() int32 {
	return mx.Protocol_Shop_BuyRefreshMerchandise
}

func (x *ShopBuyRefreshMerchandiseResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

