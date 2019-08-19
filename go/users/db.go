package users

import (
  "github.com/go-pg/pg/orm"

  . "github.com/Liquid-Labs/terror/go/terror"
  . "github.com/Liquid-Labs/lc-entities-model/go/entities"
)

var UserFields = append(EntityFields,
  `auth_id`,
  `legal_id`,
  `legal_id_type`,
  `active`,
)

// Create creates (or inserts) a new User record into the DB. As Users are logically abstract, one would typically only call this as part of another items create sequence.
func (u *User) CreateRaw(db orm.DB) Terror {
  if err := CreateEntityRaw(u, db); err != nil {
    return err
  } else {
    qs := db.Model((&u.Subject)).ExcludeColumn(EntityFields...)
    if _, err := qs.Insert(); err != nil {
      return ServerError(`There was a problem creating the subject record.`, err)
    } else {
      qu := db.Model(u).ExcludeColumn(EntityFields...)
      if _, err := qu.Insert(); err != nil {
        return ServerError(`There was a problem creating the user record.`, err)
      } else {
        return nil
      }
    }
  }
}

var updateExcludes = make([]string, len(EntityFields))
func init() {
  copy(updateExcludes, EntityFields)
  updateExcludes = append(updateExcludes, "id")
}

func (u *User) UpdateRawQueries(db orm.DB) []*orm.Query {
  qs := (&u.Entity).UpdateRawQueries(db)
  // No need for a Subjects query, there's nothing there. And there's a trivial bug (?) in go-pg (v8.0.5) where if you exclude all the fields (as we would) then it's the same as excluding none of the fields.
  qu := db.Model(u).
    ExcludeColumn(updateExcludes...).
    Where(`"user".id=?id`)
  qu.GetModel().Table().SoftDeleteField = nil

  return append(qs, qu)
}

// Updates a User record in the DB. As Users are logically abstract, one would typically only call this as part of another items update sequence.
func (u *User) UpdateRaw(db orm.DB) Terror {
  return DoRawUpdate(u.UpdateRawQueries(db), db)
}

// Archive updates a User record in the DB. As Users are logically abstract, one would typically only call this as part of another items archive sequence.
func (u *User) ArchiveRaw(db orm.DB) Terror {
  return (&u.Entity).ArchiveRaw(db)
}
