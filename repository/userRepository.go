package repository

import (
	"gorm.io/gorm"
	"messagingapp/commons"
	"messagingapp/models"
	"errors"
)

type UserRepositoryDB struct {
	db *gorm.DB
}

type UserService interface {
	CreateUser(*models.User) error
	DeleteUser(*models.User) error
	GetUser(*models.User) error
	UpdateUser(*models.User) error
}

func NewUserRepositoryDB() *UserRepositoryDB {
	return &UserRepositoryDB{
		db: commons.InitGorm(),
	}
}

func (r *UserRepositoryDB) CreateUser(user *models.User) error {
	err := r.db.Create(user).Error
	return err
}

func (r *UserRepositoryDB) DeleteUser(user *models.User) error {
	err := r.db.Delete(user).Error
	return err
}

func (r *UserRepositoryDB) GetUser(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("user doesn't exists: " + err.Error())
	}
	return &user, err
}

func (r *UserRepositoryDB) UpdateUser(newUser *models.User) error { // new implementation on the new fiels if they're not written, risk of errors
	var oldUser models.User
	err := r.db.Where("email = ?", newUser.Email).First(&oldUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("user doesn't exists: " + err.Error())
		}
		return err
	}
	oldUser = *newUser
	r.db.Save(&oldUser)
	return nil
}