// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package datatype

import (
	"database/sql"
	"time"
)

type DtCharacter struct {
	A sql.NullString
	B sql.NullString
	C sql.NullString
	D sql.NullString
	E sql.NullString
	F sql.NullString
	G sql.NullString
	H sql.NullString
}

type DtCharacterNotNull struct {
	A string
	B string
	C string
	D string
	E string
	F string
	G string
	H string
}

type DtDatetime struct {
	A sql.NullTime
	B sql.NullTime
	C sql.NullTime
}

type DtDatetimeNotNull struct {
	A time.Time
	B time.Time
	C time.Time
}

type DtNumeric struct {
	A sql.NullInt64
	B sql.NullInt64
	C sql.NullInt64
	D sql.NullInt64
	E sql.NullInt64
	F sql.NullInt64
	G sql.NullInt64
	H sql.NullInt64
	I sql.NullInt64
	J sql.NullFloat64
	K sql.NullFloat64
	L sql.NullFloat64
	M sql.NullFloat64
	N sql.NullFloat64
	O sql.NullFloat64
}

type DtNumericNotNull struct {
	A int64
	B int64
	C int64
	D int64
	E int64
	F int64
	G int64
	H int64
	I int64
	J float64
	K float64
	L float64
	M float64
	N float64
	O float64
}
