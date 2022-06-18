CREATE SEQUENCE user_ids_sequence START 1;

CREATE TABLE IF NOT EXISTS users (
  id integer NOT NULL DEFAULT nextval('user_ids_sequence'),
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  status BOOLEAN NOT NULL DEFAULT false,
  created_at TIME DEFAULT CURRENT_TIMESTAMP,
  updated_at TIME DEFAULT CURRENT_TIMESTAMP
);
