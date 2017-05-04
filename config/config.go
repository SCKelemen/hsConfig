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

// accessors for standard props
// Issuer
func (p Properties) SetIssuer(val string) {
  p.Set("iss", val)
}
func (p Properties) Issuer() (string, bool) {
  v, ok := p.Get("iss").(string)
  return v, ok
}
func (p Properties) RemoveIssuer() {
  p.Rem("iss")
}

// Subject
func (p Properties) SetSubject(val string) {
  p.Set("sub", val)
}
func (p Properties) Subject() (string, bool) {
  v, ok := p.Get("sub").(string)
  return v, ok
}
func (p Properties) RemoveSubject() {
  p.Rem("sub")
}

// AUDIENCE - needs to be updates

// Expiry
func (p Properties) SetExpiry(val string) {
  p.Set("exp", val)
}
func (p Properties) Expiry() (string, bool) {
  v, ok := p.Get("exp").(string)
  return v, ok
}
func (p Properties) RemoveExpiry() {
  p.Rem("exp")
}


// Not Before
func (p Properties) SetNotBefore(val string) {
  p.Set("nbf", val)
}
func (p Properties) NotBefore() (string, bool) {
  v, ok := p.Get("nbf").(string)
  return v, ok
}
func (p Properties) RemoveNotBefore() {
  p.Rem("nbf")
}


// Issued At
func (p Properties) SetIssuedAt(val string) {
  p.Set("iat", val)
}
func (p Properties) IssuedAt() (string, bool) {
  v, ok := p.Get("iat").(string)
  return v, ok
}
func (p Properties) RemoveIssusedAt() {
  p.Rem("iat")
}


// Configuration ID
func (p Properties) SetConfigID(val string) {
  p.Set("cid", val)
}
func (p Properties) ConfigID() (string, bool) {
  v, ok := p.Get("cid").(string)
  return v, ok
}
func (p Properties) RemoveConfigID() {
  p.Rem("cid")
}


// Service Name
func (p Properties) SetServiceName(val string) {
  p.Set("nam", val)
}
func (p Properties) ServiceName() (string, bool) {
  v, ok := p.Get("nam").(string)
  return v, ok
}
func (p Properties) RemoveServiceName() {
  p.Rem("nam")
}


// Geo Zone
func (p Properties) SetGeoZone(val string) {
  p.Set("geo", val)
}
func (p Properties) GeoZone() (string, bool) {
  v, ok := p.Get("geo").(string)
  return v, ok
}
func (p Properties) RemoveGeoZone() {
  p.Rem("geo")
}


// Service ID
func (p Properties) SetServiceID(val string) {
  p.Set("sid", val)
}
func (p Properties) ServiceID() (string, bool) {
  v, ok := p.Get("sid").(string)
  return v, ok
}
func (p Properties) RemoveServiceID() {
  p.Rem("sid")
}
