package repositories

import (
	"errors"
	"jewete/database"
	"jewete/entities"

	"gorm.io/gorm"
)

type Repository interface {
	Create(model interface{}) error
	Get(model interface{}) (*entities.User, error)
	GetById(id *string) (*entities.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuth() Repository {
	return &authRepository{
		db: database.GetInstance(),
	}
}

func (r *authRepository) Create(model interface{}) error {
	user := model.(*entities.User)

	username := r.db.Where("username = ?", user.Username).Or("email = ?", user.Email).First(user).Error
	if !errors.Is(username, gorm.ErrRecordNotFound) {
		return errors.New("username already exist")
	}

	email := r.db.Where("email = ?", user.Email).First(user).Error
	if !errors.Is(email, gorm.ErrRecordNotFound) {
		return errors.New("email already exist")
	}

	record := r.db.Create(user)
	if record.Error != nil {
		return record.Error
	}

	return nil
}

func (r *authRepository) Get(model interface{}) (*entities.User, error) {
	user := model.(*entities.User)

	res, err := r.db.Model(&entities.User{}).Where("username = ?", user.Username).Rows()
	if err != nil {
		return nil, err
	}

	var record entities.User
	for res.Next() {
		r.db.ScanRows(res, &record)
	}

	return &record, nil
}

func (r *authRepository) GetById(id *string) (*entities.User, error) {
	res, err := r.db.Model(&entities.User{}).Where("id = ?", id).Rows()
	if err != nil {
		return nil, err
	}

	var record entities.User
	for res.Next() {
		r.db.ScanRows(res, &record)
	}

	return &record, nil
}
