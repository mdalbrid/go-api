package object

import (
	model "github.com/mdalbrid/models"
	"github.com/mdalbrid/models/ws_object"
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
	Column    string `json:"column" validate:"required,oneof=uuid explorationUUID groupUUID name image country comment sortWeight dateAdd dateEdit"`
	Direction string `json:"type" validate:"required,oneof=ASC DESC"`
}

func OrderDtoToObject(dto []OrderDto) []ws_object.Order {
	var orders []ws_object.Order
	for _, b := range dto {
		orders = append(orders, ws_object.Order{
			Column:    b.Column,
			Direction: b.Direction,
		})
	}
	return orders
}

func (dto *ListRequestDto) ToFilter() ws_object.FilterObject {
	return ws_object.FilterObject{
		Offset: dto.Offset,
		Limit:  dto.Limit,
		Orders: OrderDtoToObject(dto.Orders),
		Filter: ws_object.Filter{
			ExplorationUUID: dto.Filter.ExplorationUUID,
		},
	}
}

type CreateRequestDto struct {
	ExplorationUUID model.UUID `json:"explorationUUID"`
	GroupUUID       model.UUID `json:"groupUUID"`
	Name            string     `json:"name"`
	Image           string     `json:"image"`
	Country         string     `json:"country"`
	Comment         string     `json:"comment"`
}

type EditRequestDto struct {
	Uuid            model.UUID `json:"uuid"`
	ExplorationUUID model.UUID `json:"explorationUUID"`
	GroupUUID       model.UUID `json:"groupUUID"`
	Name            string     `json:"name"`
	Image           string     `json:"image"`
	Country         string     `json:"country"`
	Comment         string     `json:"comment"`
	SortWeight      int        `json:"sortWeight"`
}

type DeleteRequestDto struct {
	Uuid model.UUID `json:"uuid"`
}

type DetailsRequestDto struct {
	Uuid model.UUID `json:"uuid"`
}
