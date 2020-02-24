package resository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/syarifme/simpleshop/models"
	repo "github.com/syarifme/simpleshop/repositories"
)

type userRepository struct {
	db *gorm.DB
}

func UserRepo(db *gorm.DB) repo.Repository {
	return &userRepository{
		db: db,
	}
}

func (userRepo *userRepository) GetAll() models.Response {
	var response models.Response
	var users []models.User
	userRepo.db.Find(&users)

	fmt.Printf("%+v \n", users)

	response = models.Response{
		Status:  200,
		Message: "Message",
	}
	return response
}
