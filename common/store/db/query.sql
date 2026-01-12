-- name: AddTickerCategory :exec
INSERT INTO ticker_categories (title)
VALUES ($1)
RETURNING *;

-- name: AddTickerMessage :exec
INSERT INTO ticker_messages (start_at,
                             end_at,
                             category_id,
                             info)
VALUES ($1, $2, $3, $4);

-- name: GetActiveTickers :many
SELECT m.start_at AS start_at,
       m.end_at AS end_at,
       c.title AS category,
       m.info AS info
FROM ticker_messages m
INNER JOIN ticker_categories c ON m.category_id = c.id
WHERE m.start_at <= now() AND m.end_at > now()
ORDER BY m.start_at;

-- name: GetMembers :many
SELECT * FROM members
WHERE is_active = TRUE
ORDER BY last_name, first_name;

-- name: GetMemberById :one
SELECT * FROM members
WHERE membership_number = $1 AND is_active = TRUE;

-- name: GetBowlingMembers :many
SELECT * FROM members
WHERE is_bowling_member = TRUE AND is_active = TRUE
ORDER BY last_name, first_name;

-- name: GetInactiveMembers :many
SELECT * FROM members
WHERE is_active = FALSE
ORDER BY last_name, first_name;

-- name: FindMembersByName :many
SELECT * FROM members
WHERE (first_name ILIKE '%' || $1 || '%' OR last_name ILIKE '%' || $1 || '%')
AND is_active = TRUE
ORDER BY last_name, first_name;

-- name: DeactivateMember :exec
UPDATE members
SET is_active = FALSE
WHERE membership_number = $1;

-- name: ReactivateMember :exec
UPDATE members
SET is_active = TRUE
WHERE membership_number = $1;

-- name: UpdateMemberEmail :exec
UPDATE members
SET email = $2, updated_at = CURRENT_TIMESTAMP
WHERE membership_number = $1;

-- name: UpdateMemberPhone :exec
UPDATE members
SET phone = $2, updated_at = CURRENT_TIMESTAMP
WHERE membership_number = $1;

-- name: GetLifeMembers :many
SELECT * FROM members
WHERE is_life_member = TRUE AND is_active = TRUE
ORDER BY last_name, first_name;

-- name: CreateMember :one
INSERT INTO members (membership_number,
                     first_name, 
                     last_name,
                     email,
                     phone,
                     is_bowling_member,
                     is_life_member)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateMembershipType :exec
UPDATE members
SET is_bowling_member = $2, is_life_member = $3, updated_at = CURRENT_TIMESTAMP
WHERE membership_number = $1;