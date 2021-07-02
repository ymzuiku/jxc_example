-- +migrate Up

-- 初始化角色
INSERT INTO actor (
  name,
  description
) VALUES ('Manager', 'Have all permission') RETURNING *;

INSERT INTO actor (
  name,
  description
) VALUES ('Visitor', 'Only can login') RETURNING *;

-- 初始化权限
INSERT INTO actor_permission (
  name,
  company_read,
  employ_create,
  employ_read,
  employ_update,
  employ_delete
) VALUES ('All', 'y', 'y', 'y', 'y', 'y') RETURNING *;

INSERT INTO actor_permission (
  name,
  company_read,
  employ_create,
  employ_read,
  employ_update,
  employ_delete
) VALUES ('Only SignIn', 'y', 'n', 'n', 'n', 'n') RETURNING *;

-- +migrate Down
DELETE FROM actor;
