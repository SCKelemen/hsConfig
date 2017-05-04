package config

import (
  "fmt"
  "time"
  "github.com/SCKelemen/hsConfig/errors"
)

type Properties map[string]interface{}

func (p Properties) Validate() error {
  return "NOT IMPLEMENTED"
}

func (p Properties) Get(key string)interface{} {
  if p == nil {
    return nil
  }
  return p[key]
}

func (p Properties) Set(key string, val interface{}) {
  p[key] = val
}

func (p Properties) Rem(key string) {
  delete(p, key)
}

