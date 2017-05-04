package errors

import "fmt"

type Error struct {
  ErrorCode int,
  Message string
}

type ErrorCode int

const (
  someError ErrorCode = iota
  anotherError
)

func (e Error) String() string {
  return fmt.Sprintf("%q", e.Message) 
}
