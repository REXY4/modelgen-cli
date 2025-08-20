# modelgen-cli

CLI tool to generate GORM-ready models & entities in Go projects.

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

## Usage

### 1. Generate a model

```bash
go run main.go --name User --attributes name:string,email:string,age:int
```

- Creates:

  - `internal/models/user.go` → GORM-ready model
  - `internal/entity/user.go` → plain entity struct

---

### 2. Generate another model

```bash
go run main.go --name Product --attributes name:string,price:float
```

- Creates:

  - `internal/models/product.go`
  - `internal/entity/product.go`

---

### 3. Add relations (optional)

```bash
go run main.go --name User --attributes name:string,email:string,age:int \
  --relations Products:Product:one2many,Tags:Tag:many2many
```

- Updates `internal/models/user.go` to include:

```go
Products []Product `gorm:"foreignKey:UserID"`
Tags     []Tag     `gorm:"many2many:user_tag"`
```

---

### 4. Build CLI binary (optional)

```bash
go build -o modelgen
./modelgen --name User --attributes name:string,email:string
```

- Places a standalone executable (`modelgen` or `modelgen.exe`) that can be used anywhere.

---

### 5. Install via `go install`

```bash
go install github.com/REXY4/modelgen-cli@latest
modelgen-cli --name User --attributes name:string,email:string
```

- Installs CLI globally in `$GOPATH/bin` or `$HOME/go/bin`.

---

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
