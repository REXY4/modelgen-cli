package modelgen

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)



func CreateModel(filename, modelName string, fields []FieldInfo, force bool) error {
	// kalau file sudah ada dan tidak force, jangan overwrite
	if !force {
		if _, err := os.Stat(filename); err == nil {
			return fmt.Errorf("file %s already exists", filename)
		}
	}

	var sb strings.Builder
	sb.WriteString("package model\n\n")


	// cek kalau ada field time.Time (default atau custom), import "time"
	needTime := true // karena CreatedAt & UpdatedAt pakai time.Time
	if needTime {
		sb.WriteString("import \"time\"\n\n")
	}

	sb.WriteString(fmt.Sprintf("type %s struct {\n", modelName))

	// ===== Tambahkan default fields =====
	sb.WriteString("    ID        uint      `gorm:\"primaryKey\"`\n")
	// ===== Field custom dari user =====
	caser := cases.Title(language.English)
	for _, f := range fields {
		mapped := MapType(f.Type)

		// gabung tags: gorm + custom
		tags := []string{}
		if mapped.GormTag != "" {
			tags = append(tags, mapped.GormTag)
		}
		if f.Tag != "" {
			tags = append(tags, f.Tag)
		}
		if len(tags) > 0 {
			sb.WriteString(fmt.Sprintf("    %s %s `%s`\n", caser.String(f.Name), mapped.GoType, strings.Join(tags, " ")))
		} else {
			sb.WriteString(fmt.Sprintf("    %s %s\n", caser.String(f.Name), mapped.GoType))
		}
	}	
	sb.WriteString("    CreatedAt time.Time `gorm:\"autoCreateTime\"`\n")
	sb.WriteString("    UpdatedAt time.Time `gorm:\"autoUpdateTime\"`\n")
	sb.WriteString("}\n")

	// buat folder kalau belum ada
	if err := os.MkdirAll(getDir(filename), os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(filename, []byte(sb.String()), 0644)
}



