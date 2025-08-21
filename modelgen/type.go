package modelgen

import (
	"fmt"
	"regexp"
	"strings"
)

type GoField struct {
	GoType  string
	GormTag string
}

// ambil directory dari path
func getDir(path string) string {
	idx := strings.LastIndex(path, "/")
	if idx == -1 {
		return "."
	}
	return path[:idx]
}

func MapType(t string) GoField {
	t = strings.ToLower(strings.TrimSpace(t)) // normalize

	switch {
	// varchar(n)
	case strings.HasPrefix(t, "varchar"):
		re := regexp.MustCompile(`varchar\((\d+)\)`)
		match := re.FindStringSubmatch(t)
		if len(match) == 2 {
			size := match[1]
			return GoField{
				GoType:  "string",
				GormTag: fmt.Sprintf("gorm:\"type:varchar(%s);size:%s\"", size, size),
			}
		}
		return GoField{
			GoType:  "string",
			GormTag: "gorm:\"type:varchar\"",
		}

	// decimal(p,s)
	case strings.HasPrefix(t, "decimal"):
		return GoField{
			GoType:  "float64",
			GormTag: fmt.Sprintf("gorm:\"type:%s\"", t),
		}

	// text & string
	case t == "string":
		return GoField{
			GoType:  "string",
			GormTag: "gorm:\"type:varchar(255);size:255\"",
		}
	case t == "text":
		return GoField{
			GoType:  "string",
			GormTag: "gorm:\"type:text\"",
		}

	// number
	case t == "int":
		return GoField{
			GoType:  "int",
			GormTag: "gorm:\"type:int\"",
		}
	case t == "int64":
		return GoField{
			GoType:  "int64",
			GormTag: "gorm:\"type:bigint\"",
		}
	case t == "float":
		return GoField{
			GoType:  "float64",
			GormTag: "gorm:\"type:float\"",
		}
	case t == "float32":
		return GoField{
			GoType:  "float32",
			GormTag: "gorm:\"type:float4\"",
		}

	// boolean
	case t == "bool", t == "boolean":
		return GoField{
			GoType:  "bool",
			GormTag: "gorm:\"type:boolean\"",
		}

	// datetime
	case t == "datetime", t == "timestamp":
		return GoField{
			GoType:  "time.Time",
			GormTag: "gorm:\"type:timestamp\"",
		}
	case t == "date":
		return GoField{
			GoType:  "time.Time",
			GormTag: "gorm:\"type:date\"",
		}

	default:
		return GoField{
			GoType:  "string",
			GormTag: fmt.Sprintf("gorm:\"type:%s\"", t),
		}
	}
}


func MapJSONType(fieldName, t string) GoField {
	t = strings.ToLower(strings.TrimSpace(t)) // normalize

	switch {
	// varchar(n) -> string
	case strings.HasPrefix(t, "varchar"):
		return GoField{
			GoType:  "string",
			GormTag: fmt.Sprintf("json:\"%s\"", fieldName),
		}

	// decimal(p,s) -> float64
	case strings.HasPrefix(t, "decimal"):
		return GoField{
			GoType:  "float64",
			GormTag: fmt.Sprintf("json:\"%s\"", fieldName),
		}

	// text & string
	case t == "string", t == "text":
		return GoField{
			GoType:  "string",
			GormTag: fmt.Sprintf("json:\"%s\"", fieldName),
		}

	// number
	case t == "int":
		return GoField{
			GoType:  "int",
			GormTag: fmt.Sprintf("json:\"%s\"", fieldName),
		}
	case t == "int64":
		return GoField{
			GoType:  "int64",
			GormTag: fmt.Sprintf("json:\"%s\"", fieldName),
		}
	case t == "float", t == "float64":
		return GoField{
			GoType:  "float64",
			GormTag: fmt.Sprintf("json:\"%s\"", fieldName),
		}
	case t == "float32":
		return GoField{
			GoType:  "float32",
			GormTag: fmt.Sprintf("json:\"%s\"", fieldName),
		}

	// boolean
	case t == "bool", t == "boolean":
		return GoField{
			GoType:  "bool",
			GormTag: fmt.Sprintf("json:\"%s\"", fieldName),
		}

	// datetime
	case t == "datetime", t == "timestamp", t == "date":
		return GoField{
			GoType:  "time.Time",
			GormTag: fmt.Sprintf("json:\"%s\"", fieldName),
		}

	// default fallback
	default:
		return GoField{
			GoType:  "string",
			GormTag: fmt.Sprintf("json:\"%s\"", fieldName),
		}
	}
}





