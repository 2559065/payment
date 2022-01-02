package repository

import (
	"cart/domain/model"
	"errors"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

// 创建CategoryRespository
func NewCategoryRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{db}
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

func (u *CartRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Cart{}).Error
}

//根据ID查找Category信息
func (u *CartRepository) FindCartByID(categoryID int64) (category *model.Cart, err error) {
	category = &model.Cart{}
	return category, u.mysqlDb.First(category, categoryID).Error
}

//创建Category信息
func (u *CartRepository) CreateCart(cart *model.Cart) (int64, error) {
	db := u.mysqlDb.FirstOrCreate(cart, model.Cart{ProductID: cart.ProductID,
		SizeID:  cart.SizeID,
		User_ID: cart.User_ID})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("购物车插入失败")
	}
	return cart.ID, nil
}

//根据ID删除Category信息
func (u *CartRepository) DeleteCartByID(categoryID int64) error {
	return u.mysqlDb.Where("id = ?", categoryID).Delete(&model.Cart{}).Error
}

//更新Category信息
func (u *CartRepository) UpdateCart(category *model.Cart) error {
	return u.mysqlDb.Model(category).Update(category).Error
}

//获取结果集
func (u *CartRepository) FindAll(UserId int64) (categoryAll []model.Cart, err error) {
	return categoryAll, u.mysqlDb.Where("user_id = ?", UserId).Find(&categoryAll).Error
}

// 根据用户ID清空购物车
func (u *CartRepository) CleanCart(userID int64) error {
	return u.mysqlDb.Where("user_id = ?", userID).Delete(&model.Cart{}).Error
}

// 添加商品数量
func (u *CartRepository) IncrNum(cartID int64, num int64) error {
	cart := &model.Cart{
		ID: cartID}
	return u.mysqlDb.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

func (u *CartRepository) DecrNum(cartID int64, num int64) error {
	cart := &model.Cart{ID: cartID}
	db := u.mysqlDb.Model(cart).Where("num >= ?", num).UpdateColumn("num", gorm.Expr("num - ?", num))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("减少失败")
	}
	return nil
}
