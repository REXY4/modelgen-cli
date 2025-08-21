package modelgen

import (
	"fmt"
	"os"
	"strings"
)

// GenerateModelEntity membuat GORM-ready model, entity, dan file migrasi
func GenerateModelEntity(modelName string, fields []FieldInfo, baseFolder string) error {
	if modelName == "" || len(fields) == 0 {
		return fmt.Errorf("modelName dan fields harus diisi")
	}
	modelsFolder := fmt.Sprintf("%s/database/models", baseFolder)
	internalFolder := fmt.Sprintf("%s/internal/%s", baseFolder,strings.ToLower(modelName))
	migrationsFolder := fmt.Sprintf("%s/database/migrations", baseFolder)

	if err := os.MkdirAll(modelsFolder, os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(internalFolder, os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(migrationsFolder, os.ModePerm); err != nil {
		return err
	}

	

	moduleName, err := getModuleName(baseFolder)
	if err != nil {
		return err
	}

	// create entity file
	// entityFile := fmt.Sprintf("%s/entity.go", internalFolder)
	// if err := CreateEntityFile(entityFile, modelName, fields, false); err != nil {
	// 	return err
	// }

	// create model file
	modelFile := fmt.Sprintf("%s/%s.go", modelsFolder, strings.ToLower(modelName))
	if err :=CreateModel(modelFile, modelName, fields, true); err != nil {
		return err
	}

	// create entity file
	entityFile := fmt.Sprintf("%s/entity.go", internalFolder)
	if err := CreateEntityFile(entityFile, modelName, fields, false); err != nil {
		return err
	}
	
	repoFile := fmt.Sprintf("%s/repository.go", internalFolder)
	if err:= CreateRepositoryFile(repoFile, modelName, moduleName);err != nil{
		return err
	}

	serviceFile := fmt.Sprintf("%s/service.go", internalFolder)
	if err:= CreateServiceFile(serviceFile, modelName, moduleName);err != nil{
		return err
	}

	if err := CreateMigrationFile(modelName, migrationsFolder, moduleName); err != nil {
		return err
	}

	if err:= UpdateMasterMigration(migrationsFolder, modelName);err !=nil{
		return err
	}

	return nil
}




func capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func joinLines(lines []string) string {
	if len(lines) == 0 {
		return ""
	}
	return "\n" + strings.Join(lines, "\n")
}
