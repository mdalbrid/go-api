package object_group

import (
	model "github.com/mdalbrid/models"
	"github.com/mdalbrid/models/ws_object_group"
)

type ListFilterDto struct {
	ExplorationUUID model.UUID `json:"explorationUUID"`
}

type ListRequestDto struct {
	Offset int           `json:"offset"`
	Limit  int           `json:"limit"`
	Filter ListFilterDto `json:"filter"`
	Orders []OrderDto    `json:"orders"`
}

type OrderDto struct {
	Column    string `json:"column" validate:"required,oneof=uuid name icon sortWeight dateAdd dateEdit"`
	Direction string `json:"type" validate:"required,oneof=ASC DESC"`
}

func OrderDtoToObject(dto []OrderDto) []ws_object_group.Order {
	var orders []ws_object_group.Order
	for _, b := range dto {
		orders = append(orders, ws_object_group.Order{
			Column:    b.Column,
			Direction: b.Direction,
		})
	}
	return orders
}

func (dto *ListRequestDto) ToFilter() ws_object_group.FilterObject {
	return ws_object_group.FilterObject{
		Offset: dto.Offset,
		Limit:  dto.Limit,
		Orders: OrderDtoToObject(dto.Orders),
		Filter: ws_object_group.Filter{
			ExplorationUUID: dto.Filter.ExplorationUUID,
		},
	}
}

type CreateRequestDto struct {
	ExplorationUUID model.UUID `json:"explorationUUID"`
	Name            string     `json:"name"`
	Icon            string     `json:"icon"`
}

type EditRequestDto struct {
	Uuid            model.UUID `json:"uuid"`
	ExplorationUUID model.UUID `json:"explorationUUID"`
	Name            string     `json:"name"`
	Icon            string     `json:"icon"`
	SortWeight      int        `json:"sortWeight"`
}

type DeleteRequestDto struct {
	Uuid model.UUID `json:"uuid"`
}

type DetailsRequestDto struct {
	Uuid model.UUID `json:"uuid"`
}
