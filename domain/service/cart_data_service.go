package service

import (
	"github.com/2559065/cart/domain/model"
	"github.com/2559065/cart/domain/repository"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64 , error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int64) ([]model.Cart, error)

	CleanCart(int64) error
	DecrNum(int64, int64) error
	IncrNum(int64, int64) error
}


//创建
func NewCartDataService(categoryRepository repository.ICartRepository) ICartDataService {
	return &CategoryDataService{ categoryRepository }
}

type CategoryDataService struct {
	CartRepository repository.ICartRepository
}


//插入
func (u *CategoryDataService) AddCart(category *model.Cart) (int64 ,error) {
	 return u.CartRepository.CreateCart(category)
}

//删除
func (u *CategoryDataService) DeleteCart(categoryID int64) error {
	return u.CartRepository.DeleteCartByID(categoryID)
}

//更新
func (u *CategoryDataService) UpdateCart(category *model.Cart) error {
	return u.CartRepository.UpdateCart(category)
}

//查找
func (u *CategoryDataService) FindCartByID(categoryID int64) (*model.Cart, error) {
	return u.CartRepository.FindCartByID(categoryID)
}

//查找
func (u *CategoryDataService) FindAllCart(userID int64) ([]model.Cart, error) {
	return u.CartRepository.FindAll(userID)
}

func (u *CategoryDataService) CleanCart(userID int64) error {
	return u.CartRepository.CleanCart(userID)
}

func (u *CategoryDataService) DecrNum(cartID int64, num int64) error {
	return u.CartRepository.DecrNum(cartID, num)
}

func (u *CategoryDataService) IncrNum(cartID int64, num int64) error {
	return u.CartRepository.IncrNum(cartID, num)
}
