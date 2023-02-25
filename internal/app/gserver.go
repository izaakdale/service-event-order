package app

import (
	"context"
	"log"

	"github.com/izaakdale/service-event-order/internal/datastore"
	"github.com/izaakdale/service-event-order/pkg/proto/order"
)

type (
	GServer struct {
		order.OrderServiceServer
	}
)

func (gs *GServer) GetOrder(ctx context.Context, req *order.OrderRequest) (*order.OrderResponse, error) {
	log.Printf("hit get order\n")
	o, err := datastore.FetchOrderTickets(req.OrderId)
	if err != nil {
		return nil, err
	}

	tickets := make([]*order.Ticket, len(o.Tickets))
	for _, v := range o.Tickets {
		tickets = append(tickets, &order.Ticket{
			FirstName:  v.FirstName,
			Surname:    v.Surname,
			QrPath:     v.QRPath,
			TicketType: v.Type,
		})
	}

	return &order.OrderResponse{
		Email:   o.Email,
		Tickets: tickets,
	}, nil
}