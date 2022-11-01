package migrations

import (
	"github.com/antoniopenizollo/crud-books/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB)  {
	db.AutoMigrate(models.Book{})
}
