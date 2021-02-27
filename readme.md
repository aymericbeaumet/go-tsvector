# pgtypes

[![test status](https://img.shields.io/github/workflow/status/aymericbeaumet/go-pgtypes/Continuous%20Integration?style=flat-square&logo=github)](https://github.com/aymericbeaumet/go-pgtypes/actions)
[![github](https://img.shields.io/github/issues/aymericbeaumet/go-pgtypes?style=flat-square&logo=github)](https://github.com/aymericbeaumet/go-pgtypes/issues)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?style=flat-square&logo=go&logoColor=white)](https://pkg.go.dev/github.com/aymericbeaumet/go-pgtypes)

Additional Postgres types definitions. Tested to work with
[`database/sql`](https://golang.org/pkg/database/sql/) and
[GORM](https://gorm.io/). Should be compatible with any driver respecting the
[`Scanner`](https://golang.org/pkg/database/sql/#Scanner) and
[`Valuer`](https://golang.org/pkg/database/sql/driver/#Valuer) interfaces.

Currently supports
[`tsvector`](https://www.postgresql.org/docs/current/datatype-textsearch.html#DATATYPE-TSVECTOR),
[`tsquery`](https://www.postgresql.org/docs/current/datatype-textsearch.html#DATATYPE-TSQUERY)
and some of their related functions.
