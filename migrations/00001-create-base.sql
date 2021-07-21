-- +migrate Up

-- author -----------------------------------
CREATE TABLE author (
  id serial PRIMARY KEY,
  name varchar(64) NOT NULL,
  description varchar(1024),
  boss boolean NOT NULL DEFAULT FALSE,
  created_at timestamptz NOT NULL DEFAULT now(),
  company_read boolean NOT NULL DEFAULT FALSE,
  employee_create boolean NOT NULL DEFAULT FALSE,
  employee_read boolean NOT NULL DEFAULT FALSE,
  employee_update boolean NOT NULL DEFAULT FALSE,
  employee_delete boolean NOT NULL DEFAULT FALSE
);

CREATE UNIQUE INDEX author_name_idx ON author (name);


-- company -------------------------------
CREATE TABLE company (
  id serial PRIMARY KEY,
  name varchar(64) NOT NULL,
  -- 公司规模: 0-10000
  people int NOT NULL,
  -- 1:免费模式, 2:普通模式, 3:旗舰版
  model smallint NOT NULL DEFAULT 1,
  -- 0:saas 1:私有化部署
  deploy_model smallint NOT NULL DEFAULT 1,
  created_at timestamptz NOT NULL DEFAULT now(),
  update_at timestamptz NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX company_name_x ON company (name);

-- account -------------------------------
CREATE TABLE account (
  id serial PRIMARY KEY,
  name varchar(64) NOT NULL,
  phone varchar(64) NOT NULL,
  email varchar(64),
  password varchar(64) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  update_at timestamptz NOT NULL DEFAULT now()
);


CREATE INDEX account_name_x ON account (name);
CREATE UNIQUE INDEX account_phone_x ON account (phone);


-- company -------------------------------
CREATE TABLE employee (
  id serial PRIMARY KEY,
  account_id serial NOT NULL,
  company_id serial NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  update_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX employee_account_id_x ON employee (account_id);
CREATE INDEX employee_company_id_x ON employee (company_id);


-- employee_author ----------------------------
CREATE TABLE employee_author (
  id serial PRIMARY KEY,
  employee_id serial NOT NULL,
  author_id serial NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX employee_author_employee_id_idx ON employee_author (employee_id);
CREATE INDEX employee_author_author_id_idx ON employee_author (author_id);


-- +migrate Down
DROP TABLE employee_author;
DROP TABLE author;
DROP TABLE employee;
DROP TABLE company;
DROP TABLE account;
