package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/REXY4/modelgen-cli/modelgen"
)

func main() {
	fmt.Println("modelgen-cli v1.0.0")
	fmt.Println("Author: M.RIZKI ISWANTO")
	fmt.Println("GitHub: github.com/REXY4/modelgen-cli")

	// --- Flags ---
	modelName := flag.String("name", "", "Nama model, contoh: User")
	attributes := flag.String("attributes", "", "Atribut model, contoh: name:string,email:string")
	baseFolder := flag.String("folder", ".", "Base folder project")
	relations := flag.String("relations", "", "Relasi, contoh: Products:Product:one2many,Tags:Tag:many2many")
	initDB := flag.Bool("init", false, "Generate configs/db.go for supported databases")
	dbType := flag.String("db", "postgres", "Database type: postgres, mysql, sqlite, sqlserver")

	flag.Parse()

	// --- Jika init DB ---
	if *initDB {
		if err := modelgen.GenerateDBConfig(*baseFolder, *dbType); err != nil {
			log.Fatal("Failed to create db.go:", err)
		}
		fmt.Println("configs/db.go created successfully for", *dbType)
		return
	}

	// --- Pastikan modelName & attributes ada ---
	if *modelName == "" || *attributes == "" {
		fmt.Println("Use : go run main.go --name User --attributes name:string,email:string")
		return
	}

	// --- Parse fields ---
	fields := parseFields(*attributes)

	// --- Generate model & entity & migration ---
	if err := modelgen.GenerateModelEntity(*modelName, fields, *baseFolder); err != nil {
		log.Fatal(err)
	}

	// --- Parse & add relations ---
	if *relations != "" {
		rels := parseRelations(*relations)
		modelFile := fmt.Sprintf("%s/internal/models/%s.go", *baseFolder, strings.ToLower(*modelName))
		if err := modelgen.AddRelation(modelFile, rels, *modelName); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Done!")
}

// --- Helpers ---
func parseFields(attr string) []modelgen.FieldInfo {
	var fields []modelgen.FieldInfo
	for _, a := range strings.Split(attr, ",") {
		parts := strings.Split(a, ":")
		if len(parts) != 2 {
			log.Fatalf("Attribute format is invalid: %s", a)
		}
		fields = append(fields, modelgen.FieldInfo{Name: parts[0], Type: parts[1]})
	}
	return fields
}

func parseRelations(rel string) []modelgen.RelationInfo {
	var rels []modelgen.RelationInfo
	for _, r := range strings.Split(rel, ",") {
		parts := strings.Split(r, ":")
		if len(parts) != 3 {
			log.Fatalf("Relation format is invalid: %s", r)
		}
		rt := modelgen.RelationType(parts[2])
		rels = append(rels, modelgen.RelationInfo{FieldName: parts[0], Target: parts[1], Type: rt})
	}
	return rels
}
