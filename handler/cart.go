package handler

import (
	"cart/domain/model"
	"cart/domain/service"
	"cart/proto/cart"
	"context"
)

type Cart struct{
	CartDataService service.ICartDataService
}

func (h *Cart)AddCart(ctx context.Context, req *cart.CartInfo, res *cart.ResponseAdd) error {
	cart := &model.Cart{}

}
//ClearCart(context.Context, *Clean, *Response) error
//Incr(context.Context, *Item, *Response) error
//Decr(context.Context, *Item, *Response) error
//DeleteItemByID(context.Context, *CartID, *Response) error
//GetAll(context.Context, *CartFindAll, *CartAll) error
