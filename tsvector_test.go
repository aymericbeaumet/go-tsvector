package pgtypes_test

import (
	"testing"

	"github.com/aymericbeaumet/go-pgtypes"
	"github.com/go-test/deep"
	_ "github.com/lib/pq"
)

type tsvectorTestModel struct {
	ID      uint             `gorm:"primaryKey"`
	Text    string           `gorm:"not null"`
	TextTSV pgtypes.TSVector `gorm:"not null"`
}

func TestTSVectorSQLCast(t *testing.T) {
	var tsvector pgtypes.TSVector
	err := sqlDB.
		QueryRow("SELECT $1::tsvector", "The quick brown fox jumps over the lazy dog").
		Scan(&tsvector)
	if err != nil {
		t.Error(err)
	}

	expected := map[string][]int{
		"The":   nil,
		"brown": nil,
		"dog":   nil,
		"fox":   nil,
		"jumps": nil,
		"lazy":  nil,
		"over":  nil,
		"quick": nil,
		"the":   nil,
	}
	if diff := deep.Equal(tsvector.Lexemes(), expected); diff != nil {
		t.Error(diff)
	}
}

func TestTSVectorSQLScan(t *testing.T) {
	var tsvector pgtypes.TSVector
	err := sqlDB.
		QueryRow("SELECT to_tsvector($1)", "I am a test: the quick brown fox jumps over the lazy fox!").
		Scan(&tsvector)
	if err != nil {
		t.Error(err)
	}

	expected := map[string][]int{
		"brown": {7},
		"fox":   {8, 13},
		"jump":  {9},
		"lazi":  {12},
		"quick": {6},
		"test":  {4},
	}
	if diff := deep.Equal(tsvector.Lexemes(), expected); diff != nil {
		t.Error(diff)
	}
}

func TestTSVectorGORMCreateFind(t *testing.T) {
	text := "I am a test: the quick brown fox jumps over the lazy fox!"

	in := tsvectorTestModel{
		Text:    text,
		TextTSV: pgtypes.ToTSVector(text),
	}
	res := gormDB.Create(&in)
	if res.Error != nil {
		t.Error(res.Error)
	}

	var out tsvectorTestModel
	if res := gormDB.First(&out, in.ID); res.Error != nil {
		t.Error(res.Error)
	}

	expected := map[string][]int{
		"brown": {7},
		"fox":   {8, 13},
		"jump":  {9},
		"lazi":  {12},
		"quick": {6},
		"test":  {4},
	}
	if diff := deep.Equal(out.TextTSV.Lexemes(), expected); diff != nil {
		t.Error(diff)
	}
}
