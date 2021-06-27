-- +migrate Up

-- account -------------------------------
CREATE TABLE account (
  id serial PRIMARY KEY,
  name text NOT NULL,
  phone text NOT NULL,
  password text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  update_at timestamptz NOT NULL DEFAULT now()
);


CREATE INDEX account_name_x ON account (name);
CREATE UNIQUE INDEX account_phone_x ON account (phone);
COMMENT ON COLUMN account.id is '@gorm:primaryKey @validate:required,min=3,max=32';
COMMENT ON COLUMN account.phone is '@gorm:primaryKey @validate:required,min=3,max=32';

-- company -------------------------------
CREATE TABLE company (
  id serial PRIMARY KEY,
  account_id serial NOT NULL,
  name text NOT NULL,
  people text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  update_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX company_account_id_x ON company (account_id);
COMMENT ON COLUMN company.name is '@gorm:primaryKey @validate:required,min=3,max=32';

-- company -------------------------------
CREATE TABLE employ (
  id serial PRIMARY KEY,
  account_id serial NOT NULL,
  company_id serial NOT NULL,
  name text NOT NULL,
  is_boss text NOT NULL DEFAULT FALSE,
  created_at timestamptz NOT NULL DEFAULT now(),
  update_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX employ_account_id_x ON employ (account_id);
CREATE INDEX employ_company_id_x ON employ (company_id);


-- +migrate Down
DROP TABLE account;
DROP TABLE company;
DROP TABLE employ;