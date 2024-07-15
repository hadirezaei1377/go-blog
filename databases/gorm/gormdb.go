package databases

import (
	"fmt"
	"go-blog/models"
	"go-blog/models/permissions"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormdb struct {
	db *gorm.DB
}

// AddComment implements databases.Database.
func (gdb *gormdb) AddComment(comment *models.Comment) error {
	panic("unimplemented")
}

// CheckCategoryExists implements databases.Database.
func (gdb *gormdb) CheckCategoryExists(name string) bool {
	panic("unimplemented")
}

// CreateCategory implements databases.Database.
func (gdb *gormdb) CreateCategory(catg *models.Category) (uint, error) {
	panic("unimplemented")
}

// GetCategories implements databases.Database.
func (gdb *gormdb) GetCategories() ([]models.Category, error) {
	panic("unimplemented")
}

// GetCategory implements databases.Database.
func (gdb *gormdb) GetCategory(name string) (*models.Category, error) {
	panic("unimplemented")
}

// GetComment implements databases.Database.
func (gdb *gormdb) GetComment(id uint) (*models.Comment, error) {
	panic("unimplemented")
}

func Connect(dsn string) (*gormdb, error) {
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("Database connection successfully opened")

	// Auto migration
	err = DB.AutoMigrate(models.User{}, models.Post{}, models.Category{}, models.Role{}, models.Comment{})
	if err != nil {
		return nil, err
	}
	gdb := &gormdb{db: DB}
	gdb.AddBasicRoles()
	fmt.Println("Database Migrated")
	return gdb, nil
}

// Add some basic roles manually
func (gdb *gormdb) AddBasicRoles() {
	gdb.CreateRole(&models.Role{Name: "superadmin", Permissions: permissions.Compress([]permissions.Permission{permissions.FullAccess})})
	gdb.CreateRole(&models.Role{Name: "moderator", Permissions: permissions.Compress([]permissions.Permission{permissions.FullContents})})
	gdb.CreateRole(&models.Role{Name: "author", Permissions: permissions.Compress([]permissions.Permission{permissions.CreatePost, permissions.FullContents})})
}
