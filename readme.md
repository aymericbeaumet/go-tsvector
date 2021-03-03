# go-tsvector

[![test status](https://img.shields.io/github/workflow/status/aymericbeaumet/go-tsvector/Continuous%20Integration?style=flat-square&logo=github)](https://github.com/aymericbeaumet/go-tsvector/actions)
[![github](https://img.shields.io/github/issues/aymericbeaumet/go-tsvector?style=flat-square&logo=github)](https://github.com/aymericbeaumet/go-tsvector/issues)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?style=flat-square&logo=go&logoColor=white)](https://pkg.go.dev/github.com/aymericbeaumet/go-tsvector)

- Postgres
  [`tsvector`](https://www.postgresql.org/docs/current/datatype-textsearch.html#DATATYPE-TSVECTOR)
  type definition for Go
- Tested to work with [`database/sql`](https://golang.org/pkg/database/sql/) and
  [GORM](https://gorm.io/)
- Compatible with any driver respecting the
  [`Scanner`](https://golang.org/pkg/database/sql/#Scanner) and
  [`Valuer`](https://golang.org/pkg/database/sql/driver/#Valuer) interfaces.

## Examples

### SQL scan with casting

```go
var out tsvector.TSVector
_ = sqlDB.QueryRow("SELECT $1::tsvector", "The quick brown fox jumps over the lazy dog").Scan(&out)
```

### SQL scan while calling the `to_tsvector` function

```go
var out tsvector.TSVector
_ = sqlDB.QueryRow("SELECT to_tsvector($1)", "The quick brown fox jumps over the lazy dog").Scan(&out)
```

### Gorm support

```go
type MyModel struct {
	ID  int64             `gorm:"primaryKey"`
	TSV tsvector.TSVector `gorm:"not null"`
}

in := &MyModel{
  TextTSV: tsvector.ToTSVector("The quick brown fox jumps over the lazy dog"),
}
_ = gormDB.Create(in)

var out MyModel
_ = gormDB.First(&out, in.ID)
```
