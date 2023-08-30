package interfaces

import "github.com/kishorens18/NetXD_Project/netxd_customer_dal/models"

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.DBResponse, error)
}
