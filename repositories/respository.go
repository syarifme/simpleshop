package repositories

import (
	"github.com/syarifme/simpleshop/models"
)

type Repository interface {
	GetAll() models.Response
}
