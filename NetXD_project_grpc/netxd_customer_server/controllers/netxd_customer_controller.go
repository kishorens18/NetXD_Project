package controllers

import (
	"context"

	pro "github.com/kishorens18/NetXD_Project/NetXD_Project/grpc/customer"
	"github.com/kishorens18/NetXD_Project/netxd_customer_dal/interfaces"
	"github.com/kishorens18/NetXD_Project/netxd_customer_dal/models"
)

type RPCserver struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.ICustomer
)

func (s *RPCserver) CreateCustomer(ctx context.Context, req *pro.Customer) (*pro.CustomerResponse, error) {
	dbCustomer := &models.Customer{
		CustomerID: req.CustomerID}
	result, err := CustomerService.CreateCustomer(dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			CustomerID: result.CustomerID,
		}
		return responseCustomer, nil
	}
}
