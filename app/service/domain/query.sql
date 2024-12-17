-- name: GetDomainDetail :many
select  id,
        domain,
        status,
        created
from Domains where ($1::varchar = '' or domain ilike "%" || $1::varchar || "%")
                or 1 = 1 limit $2 offset $3;


-- name: GetDomainCount :one
select count(*)
from Domains where ($1::varchar = '' or domain ilike "%" || $1::varchar || "%")
                or 1 = 1;

-- name: CreateDomain :one
INSERT INTO Domains (domain, status, created)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteDomain :exec
delete
from Domains
where id = $1;

-- name: UpdateDomain :exec
UPDATE Domains
set domain=$2,
    status=$3,
    created=$4
WHERE id = $1;
