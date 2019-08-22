package users

import (
  "time"

  . "github.com/Liquid-Labs/lc-entities-model/go/entities"
)

type Subject struct {
  Entity
}

type User struct {
  tableName   struct{} `sql:"select:users_join_entity"`
  Subject
  AuthID      string `json:"authId"`
  LegalID     string `json:"legalId"`
  LegalIDType string `json:"legalIdType"`
  Active      bool   `json:"active" sql:",notnull"`
  deletedAt   time.Time
}

func NewUser(
    resourceName ResourceName,
    name string,
    description string,
    authID string,
    legalID string,
    legalIDType string,
    active bool) *User {
  return &User{
      Subject: Subject{*NewEntity(resourceName, name, description, ``, false)},
      AuthID: authID,
      LegalID: legalID,
      LegalIDType: legalIDType,
      Active: active,
    }
}

func (u *User) Clone() *User {
  return &User{
    struct{}{},
    Subject{*u.Entity.Clone()},
    u.AuthID,
    u.LegalID,
    u.LegalIDType,
    u.Active,
    time.Time{},
  }
}

func (u *User) CloneNew() *User {
  newU := u.Clone()
  newU.Entity = *u.Entity.CloneNew()
  return newU
}

func (u *User) IsConcrete() bool { return false }

func (u *User) GetEntity() *Entity { return &u.Subject.Entity }

func (u *User) GetAuthID() string { return u.AuthID }
func (u *User) SetAuthID(id string) { u.AuthID = id }

func (u *User) GetLegalID() string { return u.LegalID }
func (u *User) SetLegalID(id string) { u.LegalID = id }

func (u *User) GetLegalIDType() string { return u.LegalIDType }
func (u *User) SetLegalIDType(t string) { u.LegalIDType = t }

func (u *User) IsActive() bool { return u.Active }
func (u *User) SetActive(a bool) { u.Active = a }
