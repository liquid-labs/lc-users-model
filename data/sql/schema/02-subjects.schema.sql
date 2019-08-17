CREATE TABLE subjects (
  id         UUID,

  CONSTRAINT subjects_key PRIMARY KEY ( id ),
  CONSTRAINT subjects_refs_entities FOREIGN KEY ( id ) REFERENCES entities ( id )
);
