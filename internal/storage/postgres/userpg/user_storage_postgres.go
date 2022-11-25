package userpg

import "gorm.io/gorm"

type UserStoragePostgres struct {
	db *gorm.DB
}

func NewUserStoragePostgres(db *gorm.DB) *UserStoragePostgres {
	return &UserStoragePostgres{
		db: db,
	}
}

func (s *UserStoragePostgres) CreateUser(username string, password string) error {
	var user User

	user.Username = username
	user.Password = password

	err := s.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}