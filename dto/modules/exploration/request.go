package exploration

import (
	model "github.com/mdalbrid/models"
	"github.com/mdalbrid/models/ws_exploration"
)

type ListFilterDto struct {
	FilterName string `json:"name"`
}

type ListRequestDto struct {
	Offset int           `json:"offset"`
	Limit  int           `json:"limit"`
	Orders []OrderDto    `json:"orders"`
	Filter ListFilterDto `json:"filter"`
}

type OrderDto struct {
	Column    string `json:"column" validate:"required,oneof=uuid name tags dateAdd dateEdit"`
	Direction string `json:"type" validate:"required,oneof=ASC DESC"`
}

func OrderDtoToObject(dto []OrderDto) []ws_exploration.Order {
	var orders []ws_exploration.Order
	for _, b := range dto {
		orders = append(orders, ws_exploration.Order{
			Column:    b.Column,
			Direction: b.Direction,
		})
	}
	return orders
}

func (dto *ListRequestDto) ToFilter() ws_exploration.FilterObject {
	return ws_exploration.FilterObject{
		Offset: dto.Offset,
		Limit:  dto.Limit,
		Orders: OrderDtoToObject(dto.Orders),
		Filter: ws_exploration.Filter{
			FilterName: dto.Filter.FilterName,
		},
	}
}

type CreateRequestDto struct {
	Name       string   `json:"name"`
	Comment    string   `json:"comment"`
	AccessType string   `json:"accessType"`
	Tags       []string `json:"tags"`
}

type EditRequestDto struct {
	Uuid       model.UUID `json:"uuid"`
	Name       string     `json:"name"`
	Comment    string     `json:"comment"`
	AccessType string     `json:"accessType"`
	Tags       []string   `json:"tags"`
}

type DeleteRequestDto struct {
	Uuid model.UUID `json:"uuid"`
}

type DetailsRequestDto struct {
	Uuid model.UUID `json:"uuid"`
}
