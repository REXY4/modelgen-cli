package main

import (
	"flag"
	"log"
	"strings"

	"github.com/REXY4/modelgen-cli/modelgen"
)

func main() {
	modelName := flag.String("name", "", "Model name, e.g., User")
	attributes := flag.String("attributes", "", "Model attributes, e.g., name:string,email:string")
	baseFolder := flag.String("path", ".", "Project base folder (default: current folder)")
	flag.Parse()

	if *modelName == "" || *attributes == "" {
		log.Fatal("Usage: go run main.go --name User --attributes name:string,email:string")
	}

	// --- Parse fields ---
	fields := []modelgen.FieldInfo{}
	for _, attr := range strings.Split(*attributes, ",") {
		parts := strings.Split(attr, ":")
		if len(parts) != 2 {
			log.Fatalf("Invalid attribute format: %s", attr)
		}
		fields = append(fields, modelgen.FieldInfo{
			Name: parts[0],
			Type: parts[1],
		})
	}

	// --- Generate Model, Entity, Migration ---
	err := modelgen.GenerateModelEntity(*modelName, fields, *baseFolder)
	if err != nil {
		log.Fatalf("Failed to generate model: %v", err)
	}

	log.Println("✅ Model, Entity, and Migration successfully created!")
}
