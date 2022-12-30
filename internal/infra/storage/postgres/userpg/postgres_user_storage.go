package userpg

import (
	"github.com/sirupsen/logrus"
	domain "github.com/vaberof/MockBankingApplication/internal/domain/user"
	"gorm.io/gorm"
)

type PostgresUserStorage struct {
	db *gorm.DB
}

func NewPostgresUserStorage(db *gorm.DB) *PostgresUserStorage {
	return &PostgresUserStorage{
		db: db,
	}
}

func (s *PostgresUserStorage) CreateUser(username string, password string) (*domain.User, error) {
	return s.createUserImpl(username, password)
}

func (s *PostgresUserStorage) GetUserById(userId uint) (*domain.User, error) {
	return s.getUserByIdImpl(userId)
}

func (s *PostgresUserStorage) GetUserByUsername(username string) (*domain.User, error) {
	return s.getUserByUsernameImpl(username)
}

func (s *PostgresUserStorage) createUserImpl(username string, password string) (*domain.User, error) {
	var postgresUser PostgresUser

	postgresUser.Username = username
	postgresUser.Password = password

	err := s.db.Table("users").Create(&postgresUser).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "userpg",
			"func":    "createUserImpl",
		}).Error(err)

		return nil, err
	}

	domainUser := BuildDomainUser(&postgresUser)

	return domainUser, nil
}

func (s *PostgresUserStorage) getUserByIdImpl(userId uint) (*domain.User, error) {
	var postgresUser PostgresUser

	err := s.db.Table("users").Where("id = ?", userId).First(&postgresUser).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "userpg",
			"func":    "getUserByIdImpl",
		}).Error(err)

		return nil, err
	}

	domainUser := BuildDomainUser(&postgresUser)

	return domainUser, nil
}

func (s *PostgresUserStorage) getUserByUsernameImpl(username string) (*domain.User, error) {
	var postgresUser PostgresUser

	err := s.db.Table("users").Where("username = ?", username).First(&postgresUser).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "userpg",
			"func":    "getUserByUsernameImpl",
		}).Error(err)

		return nil, err
	}

	domainUser := BuildDomainUser(&postgresUser)

	return domainUser, nil
}
