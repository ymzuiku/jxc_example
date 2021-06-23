-- +migrate Up
CREATE TABLE users (
  id bigserial PRIMARY KEY,
  name text NOT NULL,
  phone text NOT NULL,
  password text NOT NULL,
  created_at timestamptz DEFAULT now(),
  update_at timestamptz DEFAULT now()
);

CREATE INDEX name_idx ON users (name);
CREATE UNIQUE INDEX phone_idx ON users (phone);


-- +migrate Down
DROP TABLE users;