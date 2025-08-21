package modelgen

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func CreateMigrationFile(modelName, migrationsFolder, moduleName string) error {
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("%s/%s_create_%s.go", migrationsFolder, timestamp, strings.ToLower(modelName))

	content := fmt.Sprintf(`package migrations

import (
	"%s/configs"
	"%s/database/models"
)

// Up migrates table %s
func Up%s() {
	configs.DB.AutoMigrate(&models.%s{})
}

// Down rolls back table %s
func Down%s() {
	configs.DB.Migrator().DropTable(&models.%s{})
}
`, moduleName, moduleName, modelName, modelName, modelName, modelName, modelName, modelName)

	return os.WriteFile(filename, []byte(content), 0644)
}

func UpdateMasterMigration(migrationsFolder, modelName string) error {
	masterFile := fmt.Sprintf("%s/migrate.go", migrationsFolder)

	// Jika belum ada, buat file baru
	if _, err := os.Stat(masterFile); os.IsNotExist(err) {
		content := `package migrations

import "fmt"

func MigrateAll() {
	fmt.Println("Running migrations...")
	Up` + modelName + `()
	// Add other migrations here
	fmt.Println("Migrations completed!")
}

func RollbackAll() {
	fmt.Println("Rolling back migrations...")
	Down` + modelName + `()
	// Add other rollbacks here
	fmt.Println("Rollback completed!")
}
`
		return os.WriteFile(masterFile, []byte(content), 0644)
	}

	// Jika sudah ada, append import Up/Down baru jika belum ada
	data, err := os.ReadFile(masterFile)
	if err != nil {
		return err
	}

	text := string(data)
	if !strings.Contains(text, "Up"+modelName+"()") {
		text = strings.Replace(text, "// Add other migrations here", "Up"+modelName+"()\n\t// Add other migrations here", 1)
	}
	if !strings.Contains(text, "Down"+modelName+"()") {
		text = strings.Replace(text, "// Add other rollbacks here", "Down"+modelName+"()\n\t// Add other rollbacks here", 1)
	}

	return os.WriteFile(masterFile, []byte(text), 0644)
}
