package handler

import (
	"context"
	"github.com/2559065/common"
	"github.com/2559065/payment/domain/model"
	"github.com/2559065/payment/domain/service"
	"github.com/2559065/payment/proto/payment"
)

type Payment struct{
	PaymentDataService service.IPaymentDataService
}



func (e *Payment)AddPayment(ctx context.Context, req *payment.PaymentInfo, res *payment.PaymentID) error {
	payment := &model.Payment{}
	if err := common.SwapTo(req, payment); err != nil {
		return err
	}
	paymentID, err := e.PaymentDataService.AddPayment(payment)
	if err != nil {
		return err
	}
	res.PaymentId = paymentID
	return nil
}
func (e *Payment)UpdatePayment(ctx context.Context, req *payment.PaymentInfo, res *payment.Response) error {
	payment := &model.Payment{}
	if err := common.SwapTo(req, payment); err != nil {
		return err
	}
	return e.PaymentDataService.UpdatePayment(payment)
}
func (e *Payment)DeletePaymentByID(ctx context.Context, req *payment.PaymentID, res *payment.Response) error {
	return e.PaymentDataService.DeletePayment(req.PaymentId)
}
func (e *Payment)FindPaymentByID(ctx context.Context, req *payment.PaymentID, res *payment.PaymentInfo) error {
	payment, err := e.PaymentDataService.FindPaymentByID(req.PaymentId)
	if err != nil {
		return err
	}
	return common.SwapTo(payment, res)
}
func (e *Payment)FindAllPayment(ctx context.Context, req *payment.All, res *payment.PaymentAll) error {
	allPayment, err := e.PaymentDataService.FindAllPayment()
	if err != nil {
		return err
	}
	for _, v := range allPayment {
		payment := &payment.PaymentInfo{}
		if err := common.SwapTo(v, payment); err != nil {
			return err
		}
		res.PaymentInfo = append(res.PaymentInfo, payment)
	}
	return nil
}
