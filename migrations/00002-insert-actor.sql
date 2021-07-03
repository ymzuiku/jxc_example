-- +migrate Up

-- 初始化角色
INSERT INTO actor (
  id,
  name,
  description
) VALUES (1, 'Manager', 'Have all permission') RETURNING *;

INSERT INTO actor (
  id,
  name,
  description
) VALUES (2, 'Visitor', 'Only can login') RETURNING *;

-- 初始化权限
INSERT INTO actor_permission (
  name,
  actor_id,
  company_read,
  employ_create,
  employ_read,
  employ_update,
  employ_delete
) VALUES ('All', 1, 'y', 'y', 'y', 'y', 'y') RETURNING *;

INSERT INTO actor_permission (
  name,
  actor_id,
  company_read,
  employ_create,
  employ_read,
  employ_update,
  employ_delete
) VALUES ('Only SignIn', 2, 'y', 'n', 'n', 'n', 'n') RETURNING *;

-- +migrate Down
DELETE FROM actor;
