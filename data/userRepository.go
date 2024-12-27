package data

import (
	"gorm.io/gorm"
	m "messagingapp/models"
)

type UserRep struct {
	db *gorm.DB
}

func (r *UserRep) Create(user *m.User) error {
	err := r.db.Create(user).Error
	return err
}

func (r *UserRep) Update(user *m.User) error {
	err := r.db.Save(user).Error
	return err
}

func (r *UserRep) Delete(user *m.User) error {
	err := r.db.Delete(user).Error
	return err
}

func (r *UserRep) Find(email string) (*m.User, error) {
	user := new(m.User)
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}