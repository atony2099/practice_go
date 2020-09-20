package service

import (
	"fmt"
	"test/chat"
)

func GetOder(req *chat.GetOrderReq) *chat.GetOrderRes {
	fmt.Println("current order id", req.Id)
	return &chat.GetOrderRes{OrderName: "tang", Price: 123}
}
