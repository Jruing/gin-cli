-- name: GetRoleDetail :many
select id,
       rolename,
       status,
       created
from Roles where ($1::varchar = '' or rolename ilike "%" || $1::varchar || "%")
              or 1 = 1 limit $2 offset $3;



-- name: GetRoleCount :one
select count(*)
from Roles where ($1::varchar = '' or rolename ilike "%" || $1::varchar || "%")
              or 1 = 1;

-- name: CreateRole :one
INSERT INTO Roles (rolename, status, created)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteRole :exec
delete
from Roles
where id = $1;

-- name: UpdateRole :exec
UPDATE Roles
set rolename=$2,
    status=$3,
    created=$4
WHERE id = $1;