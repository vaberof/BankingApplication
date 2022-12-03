package userpg

import "gorm.io/gorm"

type PostgresUserStorage struct {
	db *gorm.DB
}

func NewPostgresUserStorage(db *gorm.DB) *PostgresUserStorage {
	return &PostgresUserStorage{
		db: db,
	}
}

func (s *PostgresUserStorage) CreateUser(username string, password string) error {
	return s.createUserImpl(username, password)
}

func (s *PostgresUserStorage) GetUserById(userId uint) (*User, error) {
	return s.getUserByIdImpl(userId)
}

func (s *PostgresUserStorage) GetUserByUsername(username string) (*User, error) {
	return s.getUserByUsernameImpl(username)
}

func (s *PostgresUserStorage) createUserImpl(username string, password string) error {
	var user User

	user.Username = username
	user.Password = password

	err := s.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresUserStorage) getUserByIdImpl(userId uint) (*User, error) {
	var user User

	err := s.db.Table("users").Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *PostgresUserStorage) getUserByUsernameImpl(username string) (*User, error) {
	var user User

	err := s.db.Table("users").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
