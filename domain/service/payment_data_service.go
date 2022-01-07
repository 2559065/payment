package service

import (
	"github.com/2559065/payment/domain/model"
	"github.com/2559065/payment/domain/repository"
)

type IPaymentDataService interface {
	AddPayment(*model.Payment) (int64 , error)
	DeletePayment(int64) error
	UpdatePayment(*model.Payment) error
	FindPaymentByID(int64) (*model.Payment, error)
	FindAllPayment() ([]model.Payment, error)
}


//创建
func NewPaymentDataService(PaymentRepository repository.IPaymentRepository) IPaymentDataService {
	return &PaymentDataService{PaymentRepository}
}

type PaymentDataService struct {
	PaymentRepository repository.IPaymentRepository
}

//插入
func (u *PaymentDataService) AddPayment(payment *model.Payment) (int64 ,error) {
	 return u.PaymentRepository.CreateOrder(payment)
}

//删除
func (u *PaymentDataService) DeletePayment(paymentID int64) error {
	return u.PaymentRepository.DeleteOrderByID(paymentID)
}

//更新
func (u *PaymentDataService) UpdatePayment(payment *model.Payment) error {
	return u.PaymentRepository.UpdateOrder(payment)
}

//查找
func (u *PaymentDataService) FindPaymentByID(paymentID int64) (*model.Payment, error) {
	return u.PaymentRepository.FindOrderByID(paymentID)
}

//查找
func (u *PaymentDataService) FindAllPayment() ([]model.Payment, error) {
	return u.PaymentRepository.FindAll()
}
