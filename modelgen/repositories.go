package modelgen

import (
	"fmt"
	"os"
	"strings"
)

func CreateRepositoryFile(filename, modelName, moduleName string) error {
	content := fmt.Sprintf(`package %s

import (
	models "%s/database/models"
	"gorm.io/gorm"
)

type %sRepository interface {
	FindAll() ([]%s, error)
	FindByID(id uint) (%s, error)
	Create(data *%s) error
	Update(data *%s) error
	Delete(id uint) error
}

type %sRepositoryImpl struct {
	db *gorm.DB
}

func New%sRepository(db *gorm.DB) %sRepository {
	return &%sRepositoryImpl{db}
}

func (r *%sRepositoryImpl) FindAll() ([]%s, error) {
	var items []%s
	err := r.db.Find(&items).Error
	return items, err
}

func (r *%sRepositoryImpl) FindByID(id uint) (%s, error) {
	var item %s
	err := r.db.First(&item, id).Error
	return item, err
}

func (r *%sRepositoryImpl) Create(data *%s) error {
	return r.db.Create(data).Error
}

func (r *%sRepositoryImpl) Update(data *%s) error {
	return r.db.Save(data).Error
}

func (r *%sRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&%s{}, id).Error
}
`, strings.ToLower(modelName),moduleName, modelName,modelName, modelName, modelName, modelName,
		modelName, modelName, modelName, modelName,
		modelName, modelName, modelName,
		modelName, modelName, modelName,
		modelName, modelName,
		modelName, modelName,
		modelName, modelName)

	return os.WriteFile(filename, []byte(content), 0644)
}