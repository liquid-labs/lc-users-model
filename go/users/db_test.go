package users_test

import (
  "os"
  "testing"
  "time"

  "github.com/go-pg/pg"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/require"
  "github.com/stretchr/testify/suite"

  "github.com/Liquid-Labs/lc-rdb-service/go/rdb"
  "github.com/Liquid-Labs/strkit/go/strkit"
  "github.com/Liquid-Labs/terror/go/terror"
  . "github.com/Liquid-Labs/lc-entities-model/go/entities"
  . "github.com/Liquid-Labs/lc-users-model/go/users"
)

func init() {
  terror.EchoErrorLog()
}

func retrieveUser(id EID) (*User, terror.Terror) {
  u := &User{Subject:Subject{Entity:Entity{ID: id}}}
  q := rdb.Connect().Model(u).Where(`"user".id=?id`)
  if err := q.Select(); err != nil && err != pg.ErrNoRows {
    return nil, terror.ServerError(`Problem retrieving entity.`, err)
  } else if err == pg.ErrNoRows {
    return nil, nil
  } else {
    return u, nil
  }
}

const (
  name = `John Doe`
  desc = `desc`
  legalID = `555-55-5555`
  legalIDType = `SSN`
  active = true
)

type UserIntegrationSuite struct {
  suite.Suite
  U *User
  AuthID string
}
func (s *UserIntegrationSuite) SetupTest() {
  s.AuthID = strkit.RandString(strkit.LettersAndNumbers, 16)
  s.U = NewUser(`users`, name, desc, s.AuthID, legalID, legalIDType, active)
}
func TestUserIntegrationSuite(t *testing.T) {
  if os.Getenv(`SKIP_INTEGRATION`) == `true` {
    t.Skip()
  } else {
    suite.Run(t, new(UserIntegrationSuite))
  }
}

func (s *UserIntegrationSuite) TestUserCreate() {
  require.NoError(s.T(), s.U.CreateRaw(rdb.Connect()), `Unexpected error creating test user`)
  // require.NoError(s.T(), rdb.Connect().Insert(s.U), `Unexpected error creating test entity`)
  // require.NoError(s.T(), err, `creating test entity`)
  assert.Equal(s.T(), name, s.U.GetName())
  assert.Equal(s.T(), desc, s.U.GetDescription())
  assert.Equal(s.T(), s.AuthID, s.U.GetAuthID())
  assert.Equal(s.T(), legalID, s.U.GetLegalID())
  assert.Equal(s.T(), legalIDType, s.U.GetLegalIDType())
  assert.Equal(s.T(), active, s.U.IsActive())
  // the default stuff
  assert.NotEqual(s.T(), EID(``), s.U.GetID(), `ID should have been set on insert.`)
  assert.NotEqual(s.T(), time.Time{}, s.U.GetCreatedAt(), `'Created at' should have been set on insert.`)
  assert.NotEqual(s.T(), time.Time{}, s.U.GetLastUpdated(), `'Last updated' should have been set on insert.`)
  assert.Equal(s.T(), false, s.U.IsPubliclyReadable())
  assert.Equal(s.T(), s.U.GetID(), s.U.GetOwnerID())
}

func (s *UserIntegrationSuite) TestUserRetrieve() {
  require.NoError(s.T(), s.U.CreateRaw(rdb.Connect()), `Unexpected error creating test user`)
  // require.NoError(s.T(), rdb.Connect().Insert(s.U), `Unexpected error creating test entity`)
  uCopy, err := retrieveUser(s.U.GetID())
  require.NoError(s.T(), err)
  assert.Equal(s.T(), s.U, uCopy)
}

func (s *UserIntegrationSuite) TestUsersUpdate() {
  require.NoError(s.T(), s.U.CreateRaw(rdb.Connect()), `Unexpected error creating test user`)
  s.U.SetName(`foo`)
  s.U.SetDescription(`bar`)
  s.U.SetPubliclyReadable(true)
  newAuthID := strkit.RandString(strkit.LettersAndNumbers, 16)
  s.U.SetAuthID(newAuthID)
  s.U.SetLegalID(`4444-44444`)
  s.U.SetLegalIDType(`EIN`)
  s.U.SetActive(false)
  require.NoError(s.T(), s.U.UpdateRaw(rdb.Connect()))
  assert.Equal(s.T(), `foo`, s.U.GetName())
  assert.Equal(s.T(), `bar`, s.U.GetDescription())
  assert.Equal(s.T(), true, s.U.IsPubliclyReadable())
  assert.Equal(s.T(), newAuthID, s.U.GetAuthID())
  assert.Equal(s.T(), `4444-44444`, s.U.GetLegalID())
  assert.Equal(s.T(), `EIN`, s.U.GetLegalIDType())
  assert.Equal(s.T(), false, s.U.IsActive())
  uCopy, err := retrieveUser(s.U.GetID())
  require.NoError(s.T(), err)
  assert.Equal(s.T(), s.U, uCopy)
}

func (s *UserIntegrationSuite) TestUserArchive() {
  require.NoError(s.T(), s.U.CreateRaw(rdb.Connect()), `Unexpected error creating test user`)
  require.NoError(s.T(), s.U.ArchiveRaw(rdb.Connect()))

  eCopy, err := retrieveUser(s.U.GetID())
  require.NoError(s.T(), err)
  assert.Nil(s.T(), eCopy)

  archived := &User{}
  assert.NoError(s.T(), rdb.Connect().Model(archived).Where(`"user".id=?`, s.U.GetID()).Deleted().Select())
  assert.Equal(s.T(), s.U, archived)
}
