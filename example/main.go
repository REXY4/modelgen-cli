package main

import (
	"fmt"
	"log"

	"github.com/REXY4/modelgen-cli/modelgen"
)

func main() {
	baseFolder := "." // project root

	// --- Contoh 1: Generate model User ---
	fieldsUser := []modelgen.FieldInfo{
		{Name: "name", Type: "string"},
		{Name: "email", Type: "string"},
		{Name: "age", Type: "int"},
	}

	err := modelgen.GenerateModelEntity("User", fieldsUser, baseFolder)
	if err != nil {
		log.Fatal(err)
	}

	// --- Contoh 2: Generate model Product ---
	fieldsProduct := []modelgen.FieldInfo{
		{Name: "name", Type: "string"},
		{Name: "price", Type: "float"},
	}

	err = modelgen.GenerateModelEntity("Product", fieldsProduct, baseFolder)
	if err != nil {
		log.Fatal(err)
	}

	// --- Contoh 3: Tambahkan relasi One2Many User → Product ---
	relations := []modelgen.RelationInfo{
		{FieldName: "Products", Target: "Product", Type: modelgen.One2Many},
	}

	modelFile := fmt.Sprintf("%s/internal/models/%s.go", baseFolder, "user")
	err = modelgen.AddRelation(modelFile, relations, "User")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Example: Model & Entity berhasil dibuat dengan relasi!")
}
