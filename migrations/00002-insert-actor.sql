-- +migrate Up

-- 初始化角色
INSERT INTO author (
  id,
  name,
  description,
  company_read,
  employ_create,
  employ_read,
  employ_update,
  employ_delete
) VALUES (1, 'Manager', 'Have all permission', 'y', 'y', 'y', 'y', 'y') RETURNING *;

INSERT INTO author (
  id,
  name,
  description,
  company_read,
  employ_create,
  employ_read,
  employ_update,
  employ_delete
) VALUES (2, 'Visitor', 'Only can login', 'y', 'n', 'n', 'n', 'n') RETURNING *;

-- +migrate Down

