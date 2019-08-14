CREATE TABLE subjects (
  id BIGINT,
  CONSTRAINT subjects_key PRIMARY KEY ( id ),
  CONSTRAINT subjects_refs_entities FOREIGN KEY ( id ) REFERENCES entities ( id )
);
