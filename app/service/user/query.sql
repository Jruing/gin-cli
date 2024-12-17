-- name: GetUserDetail :many
select id,
       nickname,
       username,
       password,
       sex,
       email,
       status,
       created
from Users
where ($1::varchar = '' or username ilike "%" || $1::varchar || "%")
    and ($2::varchar = '' or nickname ilike "%" || $2::varchar || "%")
   or 1 = 1
limit $3 offset $4;

-- name: GetUserCount :one
select count(*)
from Users
where ($1::varchar = '' or username ilike "%" || $1::varchar || "%")
    and ($2::varchar = '' or nickname ilike "%" || $2::varchar || "%")
   or 1 = 1;

-- name: CheckUser :one
select count(1)
from Users
where username = $1
  and password = $2;


-- name: CreateUser :exec
INSERT INTO Users (nickname, username, password, sex, email, status, created)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: DeleteUser :exec
delete
from Users
where id = $1;

-- name: UpdateUser :exec
UPDATE Users
set nickname=$2,
    username=$3,
    password=$4,
    sex=$5,
    email=$6,
    status=$7
WHERE id = $1;