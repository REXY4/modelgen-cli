package modelgen

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FieldInfo struct {
	Name string
	Type string
	Tag  string
}

func CreateEntityFile(filename, modelName string, fields []FieldInfo, force bool) error {
	// kalau file sudah ada dan tidak force, jangan overwrite
	if !force {
		if _, err := os.Stat(filename); err == nil {
			return fmt.Errorf("file %s already exists", filename)
		}
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("package %s\n\n", strings.ToLower(modelName)))

	// cek kalau ada field time.Time (default atau custom), import "time"
	needTime := true // karena CreatedAt & UpdatedAt pakai time.Time
	if needTime {
		sb.WriteString("import \"time\"\n\n")
	}

	sb.WriteString(fmt.Sprintf("type %s struct {\n", modelName))

	// ===== Tambahkan default fields =====
	sb.WriteString("    ID        uint      `json:\"id\"`\n")
	caser := cases.Title(language.English)
	// ===== Field custom dari user =====
	for _, f := range fields {
		mapped := MapJSONType(strings.ToLower(f.Name),f.Type)

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
	
	sb.WriteString("    CreatedAt time.Time `json:\"created_at\"`\n")
	sb.WriteString("    UpdatedAt time.Time `json:\"updated_at\"`\n")
	sb.WriteString("}\n")

	// buat folder kalau belum ada
	if err := os.MkdirAll(getDir(filename), os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(filename, []byte(sb.String()), 0644)
}


