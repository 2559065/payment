package repository

import (
	"errors"
	"github.com/2559065/payment/domain/model"
	"github.com/jinzhu/gorm"
)

type IPaymentRepository interface {
	InitTable() error
	FindOrderByID(int64) (*model.Payment, error)
	CreateOrder(*model.Payment) (int64, error)
	DeleteOrderByID(int64) error
	UpdateOrder(*model.Payment) error
	FindAll() ([]model.Payment, error)

	UpdateShipStatus(int64, int32) error
	UpdatePayStatus(int64,int32) error
}

// 创建CategoryRespository
func NewPaymentRepository(db *gorm.DB) IPaymentRepository {
	return &PaymentRepository{db}
}

type PaymentRepository struct {
	mysqlDb *gorm.DB
}

func (u *PaymentRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Payment{}).Error
}

//根据ID查找Category信息
func (u *PaymentRepository) FindOrderByID(orderID int64) (order *model.Payment, err error) {
	order = &model.Payment{}
	return order, u.mysqlDb.Preload("OrderDetail").First(order, orderID).Error
}

//创建Category信息
func (u *PaymentRepository) CreateOrder(order *model.Payment) (int64, error) {
	return order.ID, u.mysqlDb.Create(order).Error
}

//根据ID删除Category信息
func (u *PaymentRepository) DeleteOrderByID(orderID int64) error {
	// 彻底删除Order信息
	if err := u.mysqlDb.Where("id = ?", orderID).Delete(&model.Payment{}).Error; err != nil {
		return err
	}
	return nil
}

//更新Category信息
func (u *PaymentRepository) UpdateOrder(order *model.Payment) error {
	return u.mysqlDb.Model(order).Update(order).Error
}

//获取结果集
func (u *PaymentRepository) FindAll() (orderAll []model.Payment, err error) {
	return orderAll, u.mysqlDb.Preload("OrderDetail").Find(&orderAll).Error
}

// 更新订单的发货状态
func (u *PaymentRepository) UpdateShipStatus(orderID int64, shipStatus int32) error {
	db := u.mysqlDb.Model(&model.Payment{}).Where("id = ?", orderID).UpdateColumn("ship_status", shipStatus)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return nil
}

// 更新订单的支付状态
func (u *PaymentRepository) UpdatePayStatus(orderID int64,payStatus int32) error {
	db := u.mysqlDb.Model(&model.Payment{}).Where("order_id = ?", orderID).UpdateColumn("pay_status", payStatus)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return nil
}
