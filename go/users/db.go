package users

import (
  "github.com/go-pg/pg/orm"

  . "github.com/Liquid-Labs/lc-entities-model/go/entities"
)

var UserFields = append(EntityFields,
  `auth_id`,
  `legal_id`,
  `legal_id_type`,
  `active`,
)

var updateExcludes = make([]string, len(EntityFields))
func init() {
  copy(updateExcludes, EntityFields)
  updateExcludes = append(updateExcludes, "id")
}

func (u *User) CreateQueries(db orm.DB) []*orm.Query {
  return append(
    (&u.Entity).CreateQueries(db),
    db.Model((&u.Subject)).ExcludeColumn(EntityFields...),
    db.Model(u).ExcludeColumn(EntityFields...))
}

func (u *User) UpdateQueries(db orm.DB) []*orm.Query {
  qes := (&u.Entity).UpdateQueries(db)
  qu := db.Model(u).
    ExcludeColumn(updateExcludes...).
    Where(`"user".id=?id`)
  qu.GetModel().Table().SoftDeleteField = nil

  return append(qes, qu)
}

func (u *User) ArchiveQueries(db orm.DB) []*orm.Query {
  return (&u.Entity).ArchiveQueries(db)
}

func (u *User) DeleteQueries(db orm.DB) []*orm.Query {
  qs := []*orm.Query{db.Model(u).Where(`"user".id=?id`)}
  return append(qs, (&u.Entity).DeleteQueries(db)...)
}
