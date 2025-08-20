# modelgen-cli

CLI tool to generate GORM-ready models, entities, and migrations in Go projects.

**Author:** M.RIZKI ISWANTO <REXY4>
**GitHub:** [github.com/REXY4/modelgen-cli](https://github.com/REXY4/modelgen-cli)

---

## Dependencies

- **`gorm.io/gorm`**

  - **Purpose:** Object-Relational Mapping (ORM) for Go.
  - **In CLI:** Generated models are ready to use with GORM, including `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, and GORM tags like `gorm:"column:..."` and `gorm:"primaryKey"`.
  - **Website:** [https://gorm.io](https://gorm.io)

- Standard Go libraries: `os`, `strings`, `flag`, `fmt`, `log`, `time`.

---

## CLI Options

```bash
--name string        # Model name, e.g.: User
--attributes string  # Model attributes, e.g.: name:string,email:string
--folder string      # Base folder of the project (default: ".")
--relations string   # Relations, e.g.: Products:Product:one2many,Tags:Tag:many2many
--init               # Generate configs/db.go for supported databases
--db string          # Database type: postgres, mysql, sqlite, sqlserver (default: postgres)
```

---

## Usage

### 1. Generate a model

```bash
go run main.go --name User --attributes name:string,email:string,age:int
```

- Creates:
  - `configs/db.go`
  - `internal/models/<Model Name>.go`
  - `internal/entity/<Entity Name>.go`
  - `internal/migrations/<timestamp>_create_<Migration Name>.go`

---

### 1. Install via `go install`

```bash
go install github.com/REXY4/modelgen-cli@latest

```

### 2. Initialize database config

```bash
# PostgreSQL
go run main.go --init --db postgres

# MySQL
go run main.go --init --db mysql

# SQLite
go run main.go --init --db sqlite

# SQL Server
go run main.go --init --db sqlserver

```

- Creates `configs/db.go` for the chosen database.
- Supported databases: **postgres, mysql, sqlite, sqlserver**.
- Automatically configures GORM connection using DSN.

### 2. Generate a model

```bash
modelgen --name User --attributes name:string,email:string

```

### 3. Add relations (optional)

```bash
modelgen --name User --attributes name:string,email:string,age:int \
  --relations Products:Product:one2many,Tags:Tag:many2many
```

- Updates `internal/models/user.go` to include:

```go
Products []Product `gorm:"foreignKey:UserID"`
Tags     []Tag     `gorm:"many2many:user_tag"`
```

## Supported Data Types

### 1. String Types

| Name       | Go Type  | GORM Tag Example     |
| ---------- | -------- | -------------------- |
| string     | `string` | `gorm:"column:name"` |
| text       | `string` | `gorm:"type:text"`   |
| varchar(n) | `string` | `gorm:"size:n"`      |

> Note: `text` and `varchar` differ only in the database; in Go they are both `string`.

---

### 2. Numeric Types

| Name    | Go Type   | GORM Tag Example              |
| ------- | --------- | ----------------------------- |
| int     | `int`     | `gorm:"column:age"`           |
| int8    | `int8`    | `gorm:"column:small_number"`  |
| int16   | `int16`   | `gorm:"column:small_number"`  |
| int32   | `int32`   | `gorm:"column:number"`        |
| int64   | `int64`   | `gorm:"column:number"`        |
| uint    | `uint`    | `gorm:"column:id;primaryKey"` |
| float   | `float64` | `gorm:"column:price"`         |
| float32 | `float32` | `gorm:"column:price"`         |
| decimal | `float64` | `gorm:"type:decimal(10,2)"`   |

---

### 3. Boolean Type

| Name | Go Type | GORM Tag Example          |
| ---- | ------- | ------------------------- |
| bool | `bool`  | `gorm:"column:is_active"` |

---

### 4. Time / Date Types

| Name      | Go Type     | GORM Tag Example           |
| --------- | ----------- | -------------------------- |
| datetime  | `time.Time` | `gorm:"column:created_at"` |
| timestamp | `time.Time` | `gorm:"column:updated_at"` |
| date      | `time.Time` | `gorm:"type:date"`         |

> GORM usually auto-generates `CreatedAt`, `UpdatedAt`, `DeletedAt` timestamps.

---

### 5. JSON / Array Types

| Name      | Go Type          | GORM Tag Example     |
| --------- | ---------------- | -------------------- |
| json      | `datatypes.JSON` | `gorm:"type:json"`   |
| string\[] | `pq.StringArray` | `gorm:"type:text[]"` |

> Requires imports:

```go
import "gorm.io/datatypes"
import "github.com/lib/pq"
```

---

### 6. Relations (special attributes)

| Relation Type | Go Type Example | GORM Tag Example            |
| ------------- | --------------- | --------------------------- |
| One-to-One    | `Profile`       | `gorm:"foreignKey:UserID"`  |
| One-to-Many   | `[]Product`     | `gorm:"foreignKey:UserID"`  |
| Many-to-Many  | `[]Tag`         | `gorm:"many2many:user_tag"` |

---

### 7. Migration Usage

After generating models with migrations:

```go
import "yourapp/internal/migrations"

func main() {
    migrations.MigrateAll()    // Run all migrations
    // migrations.RollbackAll() // Rollback all tables if needed
}
```

- `MigrateAll()` runs `Up<Model>()` for each generated model.
- `RollbackAll()` runs `Down<Model>()` for each generated model.

---
