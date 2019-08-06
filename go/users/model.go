package users

import (
  "github.com/Liquid-Labs/lc-entities-model/go/entities"
)

type User struct {
  entities.Entity
  AuthID      string `json:"authId"`
  LegalID     string `json:"legalId"`
  LegalIDType string `json:"legalIdType"`
  Active      bool   `json:"active"`
}

func NewUser(name string,
    description string,
    authID string,
    legalID string,
    legalIDType string,
    active bool) *User {
  return &User{
      *entities.NewEntity(name, description, ``, false),
      authID,
      legalID,
      legalIDType,
      active,
    }
}

func (u *User) Clone() *User {
  return &User{
    *u.Entity.Clone(),
    u.AuthID,
    u.LegalID,
    u.LegalIDType,
    u.Active,
  }
}

func (u *User) CloneNew() *User {
  newU := u.Clone()
  newU.Entity = *u.Entity.CloneNew()
  return newU
}

func (u *User) GetAuthID() string { return u.AuthID }
func (u *User) SetAuthID(id string) { u.AuthID = id }

func (u *User) GetLegalID() string { return u.LegalID }
func (u *User) SetLegalID(id string) { u.LegalID = id }

func (u *User) GetLegalIDType() string { return u.LegalIDType }
func (u *User) SetLegalIDType(t string) { u.LegalIDType = t }

func (u *User) IsActive() bool { return u.Active }
func (u *User) SetActive(a bool) { u.Active = a }
