-- Postgres
DROP TYPE IF EXISTS legal_id_type;
CREATE TYPE legal_id_type AS ENUM ('SSN', 'EIN');

CREATE TABLE users (
  id            UUID,
  -- Based on firebase User uid. Not well documented, but if you dig into this:
  -- https://firebase.google.com/docs/auth/admin/manage-users
  -- we find that an assigned ID may be up to 128 characters, though generated
  -- uids are 28 characters at time of writing (2019-03-08), though that's not
  -- guaranteed
  auth_id       VARCHAR(128),
  legal_id      VARCHAR(128),
  legal_id_type legal_id_type,
  active        BOOLEAN DEFAULT TRUE NOT NULL,

  CONSTRAINT users_key PRIMARY KEY ( id ),
  CONSTRAINT users_auth_id_unique UNIQUE (auth_id),
  CONSTRAINT users_ref_subjects FOREIGN KEY ( id ) REFERENCES subjects ( id )
);

CREATE VIEW users_join_entity AS
  SELECT e.*, u.auth_id, u.legal_id, u.legal_id_type, u.active
    FROM users u JOIN entities e ON u.id=e.id;

CREATE TABLE apps_users (
  app_id          UUID NOT NULL,
  user_id         UUID NOT NULL,
  default_context UUID,
  CONSTRAINT apps_users_key PRIMARY KEY ( app_id, user_id ),
  CONSTRAINT apps_users_default_context_ref_entities FOREIGN KEY ( default_context ) REFERENCES entities ( id )
);
