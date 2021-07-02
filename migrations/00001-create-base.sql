-- +migrate Up

-- account -------------------------------
CREATE TYPE ok AS ENUM ('n', 'y');

CREATE TABLE account (
  id serial PRIMARY KEY,
  name varchar(64) NOT NULL,
  phone varchar(64) NOT NULL,
  email varchar(64) NOT NULL,
  password varchar(64) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  update_at timestamptz NOT NULL DEFAULT now()
);


CREATE INDEX account_name_x ON account (name);
CREATE UNIQUE INDEX account_phone_x ON account (phone);

-- company -------------------------------
-- 0:免费模式, 1:普通模式, 2:旗舰版
CREATE TYPE company_model AS ENUM ('free', 'normal', 'prod');
-- 0:saas 1:私有化部署
CREATE TYPE company_deploy_model AS ENUM ('saas', 'private');

CREATE TABLE company (
  id serial PRIMARY KEY,
  account_id serial NOT NULL,
  name varchar(64) NOT NULL,
  people varchar(64) NOT NULL,
  model company_model NOT NULL DEFAULT 'free',
  deploy_model company_deploy_model NOT NULL DEFAULT 'saas',
  created_at timestamptz NOT NULL DEFAULT now(),
  update_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX company_account_id_x ON company (account_id);

-- company -------------------------------
CREATE TABLE employ (
  id serial PRIMARY KEY,
  account_id serial NOT NULL,
  company_id serial NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  update_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX employ_account_id_x ON employ (account_id);
CREATE INDEX employ_company_id_x ON employ (company_id);

-- employ_actor ----------------------------
CREATE TABLE employ_actor (
  id serial PRIMARY KEY,
  employ_id serial NOT NULL,
  actor_id serial NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX employ_actor_employ_id_idx ON employ_actor (employ_id);
CREATE INDEX employ_actor_actor_id_idx ON employ_actor (actor_id);

-- actor -----------------------------------
CREATE TABLE actor (
  id serial PRIMARY KEY,
  name varchar(64) NOT NULL,
  description varchar(1024) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX actor_name_idx ON actor (name);

-- actor_permission -------------------------
CREATE TABLE actor_permission (
  id serial PRIMARY KEY,
  name varchar(64) NOT NULL,
  company_read ok NOT NULL DEFAULT 'n',
  employ_create ok NOT NULL DEFAULT 'n',
  employ_read ok NOT NULL DEFAULT 'n',
  employ_update ok NOT NULL DEFAULT 'n',
  employ_delete ok NOT NULL DEFAULT 'n'
);


-- +migrate Down
DROP TABLE account;
DROP TABLE company;
DROP TABLE employ;
DROP TABLE actor;
DROP TABLE employ_actor;
DROP TABLE actor_permission;
-- type 需要在table干掉之后才能drop
DROP TYPE ok;
DROP TYPE company_model;
DROP TYPE company_deploy_model;