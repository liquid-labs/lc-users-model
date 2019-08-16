package users_test

import (
  "reflect"
  "strings"
  "testing"
  "time"

  . "github.com/Liquid-Labs/lc-entities-model/go/entities"
  "github.com/stretchr/testify/assert"

  // the package we're testing
  . "github.com/Liquid-Labs/lc-users-model/go/users"
)

type TestUser struct {
  User
}
func (tu *TestUser) GetResourceName() ResourceName {
  return ResourceName(`testusers`)
}

func TestUsersClone(t *testing.T) {
  now := time.Now()
  orig := NewUser(&TestUser{}, `john`, `cool`, `azn-1`, `555`, `SSN`, true)
  orig.ID = EID(`abc`)
  orig.OwnerID = EID(`owner-A`)
  orig.CreatedAt = now
  orig.LastUpdated = now.Add(100)
  orig.DeletedAt = now.Add(200)
  clone := orig.Clone()

  assert.Equal(t, orig, clone, "Clone does not match.")

  clone.ID = EID(`hij`)
  clone.Name = `sally`
  clone.Description = `awesome`
  clone.OwnerID = EID(`owner-B`)
  clone.PubliclyReadable = false
  clone.CreatedAt = orig.CreatedAt.Add(20)
  clone.LastUpdated = orig.LastUpdated.Add(20)
  clone.DeletedAt = orig.DeletedAt.Add(20)
  clone.AuthID = `azn-3`
  clone.LegalID = `666`
  clone.LegalIDType = `EIN`
  clone.Active = false

  // TODO: abstract this
  oReflection := reflect.ValueOf(orig).Elem()
  cReflection := reflect.ValueOf(clone).Elem()
  for i := 0; i < oReflection.NumField(); i++ {
    name := oReflection.Type().FieldByIndex([]int{i}).Name
    if name[:1] == strings.ToUpper(name[:1]) {
      assert.NotEqualf(
        t,
        oReflection.Field(i).Interface(),
        cReflection.Field(i).Interface(),
        `Fields '%s' unexpectedly match.`,
        oReflection.Type().Field(i),
      )
    }
  }
}

func TestUsersCloneNew(t *testing.T) {
  now := time.Now()
  orig := NewUser(&TestUser{}, `john`, `cool`, `azn-1`, `555`, `SSN`, true)
  orig.ID = EID(`abc`)
  orig.OwnerID = EID(`owner-A`)
  orig.CreatedAt = now
  orig.LastUpdated = now.Add(100)
  orig.DeletedAt = now.Add(200)
  clone := orig.CloneNew()

  assert.Equal(t, EID(``), clone.ID)
  assert.Equal(t, orig.GetOwnerID(), clone.GetOwnerID())
  assert.Equal(t, time.Time{}, clone.GetCreatedAt())
  assert.Equal(t, time.Time{}, clone.GetLastUpdated())
  assert.Equal(t, time.Time{}, clone.GetDeletedAt())

  clone.ID = EID(`hij`)
  clone.Name = `sally`
  clone.Description = `awesome`
  clone.OwnerID = EID(`owner-B`)
  clone.PubliclyReadable = false
  clone.CreatedAt = orig.CreatedAt.Add(20)
  clone.LastUpdated = orig.LastUpdated.Add(20)
  clone.DeletedAt = orig.DeletedAt.Add(20)
  clone.AuthID = `azn-3`
  clone.LegalID = `666`
  clone.LegalIDType = `EIN`
  clone.Active = false

  // TODO: abstract this
  oReflection := reflect.ValueOf(orig).Elem()
  cReflection := reflect.ValueOf(clone).Elem()
  for i := 0; i < oReflection.NumField(); i++ {
    name := oReflection.Type().FieldByIndex([]int{i}).Name
    if name[:1] == strings.ToUpper(name[:1]) {
      assert.NotEqualf(
        t,
        oReflection.Field(i).Interface(),
        cReflection.Field(i).Interface(),
        `Fields '%s' unexpectedly match.`,
        oReflection.Type().Field(i),
      )
    }
  }
}
