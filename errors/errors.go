package errors

import "fmt"

type Error struct {
  ErrorCode int,
  Node string,
  Value string,
  Message string
}

type ErrorCode int

const (
  ExpiredConfig ErrorCode = iota
  PreNBFConfig
  InvalidIssuer
  InvalidType
  InvalidAlg
  InvalidSubject
  InvalidAudience
  InvalidServiceID
  InvalidZone
  InvalidServiceName
  InvalidConfigID
)

func (e Error) String() string {
  return fmt.Sprintf("%q", e.Message) 
}
