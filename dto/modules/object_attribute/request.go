package object_attribute

import (
	model "github.com/mdalbrid/models"
	"github.com/mdalbrid/models/ws_object_attribute"
)

type ListFilterDto struct {
	ObjectUUID model.UUID `json:"objectUUID"`
}

type ListRequestDto struct {
	Offset int           `json:"offset"`
	Limit  int           `json:"limit"`
	Filter ListFilterDto `json:"filter"`
	Orders []OrderDto    `json:"orders"`
}

type OrderDto struct {
	Column    string `json:"column" validate:"required,oneof=uuid name value sortWeight dateAdd dateEdit"`
	Direction string `json:"type" validate:"required,oneof=ASC DESC"`
}

func OrderDtoToObject(dto []OrderDto) []ws_object_attribute.Order {
	var orders []ws_object_attribute.Order
	for _, b := range dto {
		orders = append(orders, ws_object_attribute.Order{
			Column:    b.Column,
			Direction: b.Direction,
		})
	}
	return orders
}

func (dto *ListRequestDto) ToFilter() ws_object_attribute.FilterObject {
	return ws_object_attribute.FilterObject{
		Offset: dto.Offset,
		Limit:  dto.Limit,
		Orders: OrderDtoToObject(dto.Orders),
		Filter: ws_object_attribute.Filter{
			ObjectUUID: dto.Filter.ObjectUUID,
		},
	}
}

type CreateRequestDto struct {
	ObjectUUID model.UUID `json:"objectUUID"`
	Name       string     `json:"name"`
	Value      string     `json:"value"`
}

type EditRequestDto struct {
	Uuid       model.UUID `json:"uuid"`
	ObjectUUID model.UUID `json:"objectUUID"`
	Name       string     `json:"name"`
	Value      string     `json:"value"`
	SortWeight int        `json:"sortWeight"`
}

type DeleteRequestDto struct {
	Uuid model.UUID `json:"uuid"`
}

type DetailsRequestDto struct {
	Uuid model.UUID `json:"uuid"`
}
