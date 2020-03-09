// Code generated by sqlc. DO NOT EDIT.

package booktest

import (
	"fmt"
	"time"
)

type BookTypeType string

const (
	FICTION    BookTypeType = "FICTION"
	NONFICTION BookTypeType = "NONFICTION"
)

func (e *BookTypeType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BookTypeType(s)
	case string:
		*e = BookTypeType(s)
	default:
		return fmt.Errorf("unsupported scan type for BookTypeType: %T", src)
	}
	return nil
}

type Author struct {
	AuthorID int
	Name     string
}

type Book struct {
	BookID    int
	AuthorID  int
	Isbn      string
	BookType  BookTypeType
	Title     string
	Yr        int
	Available time.Time
	Tags      string
}
