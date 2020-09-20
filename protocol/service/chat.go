package service

import (
	"context"
	"log"
	"test/chat"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *chat.Message) (*chat.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &chat.Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) GetOrder(ctx context.Context, req *chat.GetOrderReq) (*chat.GetOrderRes, error) {
	return GetOder(req), nil
}
