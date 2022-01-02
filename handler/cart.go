package handler

import (
	"context"
	"github.com/2559065/cart/domain/model"
	"github.com/2559065/cart/domain/service"
	"github.com/2559065/cart/proto/cart"
	"github.com/2559065/common"
)

type Cart struct{
	CartDataService service.ICartDataService
}

// 添加购物车
func (h *Cart)AddCart(ctx context.Context, req *cart.CartInfo, res *cart.ResponseAdd) (err error){
	cart := &model.Cart{}
	common.SwapTo(req, cart)
	res.CartId, err = h.CartDataService.AddCart(cart)
	return err
}
func (h *Cart) ClearCart(ctx context.Context, req *cart.Clean, res *cart.Response) error {
	if err := h.CartDataService.CleanCart(req.UserId); err != nil {
		return err
	}
	res.Msg = "购物车清空成功"
	return nil
}
func (h *Cart)Incr(ctx context.Context, req *cart.Item, res *cart.Response) error {
	if err := h.CartDataService.IncrNum(req.Id, req.ChangeNum); err != nil {
		return err
	}
	res.Msg = "购物车添加成功"
	return nil
}
func (h *Cart)Decr(ctx context.Context, req *cart.Item, res *cart.Response) error {
	if err := h.CartDataService.DecrNum(req.Id, req.ChangeNum); err != nil {
		return err
	}
	res.Msg = "购物车减少成功"
	return nil
}
func (h *Cart)DeleteItemByID(ctx context.Context, req *cart.CartID, res *cart.Response) error {
	if err := h.CartDataService.DeleteCart(req.Id); err != nil {
		return err
	}
	res.Msg = "购物车删除成功"
	return nil
}
func (h *Cart)GetAll(ctx context.Context, req *cart.CartFindAll, res *cart.CartAll) error {
	cartAll, err := h.CartDataService.FindAllCart(req.UserId)
	if err != nil {
		return err
	}
	for _, v := range cartAll {
		cart := &cart.CartInfo{}
		if err := common.SwapTo(v, cart); err != nil {
			return err
		}
		res.CartInfo = append(res.CartInfo, cart)
	}
	return nil
}
