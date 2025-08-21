package modelgen

import (
	"fmt"
	"os"
	"strings"
)

func CreateServiceFile(filename, modelName, moduleName string) error {
	content := fmt.Sprintf(`package %s

import (
	models "%s/database/models"
	
)

type %sService interface {
	GetAll() ([]%s, error)
	GetByID(id uint) (%s, error)
	Create(data *%s) error
	Update(data *%s) error
	Delete(id uint) error
}

type %sServiceImpl struct {
	repo %sRepository
}

func New%sService(repo %sRepository) %sService {
	return &%sServiceImpl{repo}
}

func (s *%sServiceImpl) GetAll() ([]%s, error) {
	return s.repo.FindAll()
}

func (s *%sServiceImpl) GetByID(id uint) (%s, error) {
	return s.repo.FindByID(id)
}

func (s *%sServiceImpl) Create(data *%s) error {
	return s.repo.Create(data)
}

func (s *%sServiceImpl) Update(data *%s) error {
	return s.repo.Update(data)
}

func (s *%sServiceImpl) Delete(id uint) error {
	return s.repo.Delete(id)
}
`, strings.ToLower(modelName),moduleName,modelName,
		modelName, modelName, modelName, modelName, modelName,
		modelName, modelName,
		modelName, modelName, modelName, modelName,
		modelName, modelName,
		modelName, modelName,
		modelName, modelName,
		modelName, modelName)

	return os.WriteFile(filename, []byte(content), 0644)
}
