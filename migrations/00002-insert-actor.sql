-- +migrate Up

-- 初始化角色
INSERT INTO author (
  id,
  name,
  description,
  created_at,
  boss,
  company_read,
  employee_create,
  employee_read,
  employee_update,
  employee_delete
) VALUES (1, 'Manager', 'Have all permission', now(), true, true, true, true, true, true ) RETURNING *;

INSERT INTO author (
  id,
  name,
  description,
  created_at,
  boss,
  company_read,
  employee_create,
  employee_read,
  employee_update,
  employee_delete
) VALUES (2, 'Visitor', 'Only can login', now(), false, true, false, false, false, false) RETURNING *;

-- +migrate Down

