package service

import "fmt"

type OrderService struct{}
type OderInfo interface {
	GetOrder(id string)
}

func (oderS *OrderService) GetOrder(s transport.GetOder) transport.OderInfo {
	fmt.Println("current id:", s.Id)

	return

}
