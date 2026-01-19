-- name: AddTickerCategory :exec
INSERT INTO ticker_categories (title)
VALUES ($1)
RETURNING *;

-- name: FindTickerCategory :many
SELECT * FROM ticker_categories
WHERE title ILIKE '%' || $1 || '%'
ORDER BY title;

-- name: AddTickerMessage :exec
INSERT INTO ticker_messages (start_at,
                             end_at,
                             category_id,
                             info)
VALUES ($1, $2, $3, $4);

-- name: FindActiveTickers :many
SELECT m.start_at AS start_at,
       m.end_at AS end_at,
       c.title AS category,
       m.info AS info
FROM ticker_messages m
INNER JOIN ticker_categories c ON m.category_id = c.id
WHERE m.start_at <= now() AND m.end_at > now()
ORDER BY m.start_at;

-- name: FindMembers :many
SELECT * FROM members
ORDER BY last_name, first_name;

-- name: FindActiveFinancialMembers :many
SELECT * FROM members
WHERE is_active = TRUE AND is_financial = TRUE
ORDER BY last_name, first_name;

-- name: FindMemberById :one
SELECT * FROM members
WHERE membership_number = $1 AND is_active = TRUE;

-- name: FindActiveFinancialBowlingMembers :many
SELECT * FROM members
WHERE is_bowling_member = TRUE AND is_active = TRUE AND is_financial = TRUE
ORDER BY last_name, first_name;

-- name: FindInactiveMembers :many
SELECT * FROM members
WHERE is_active = FALSE
ORDER BY last_name, first_name;

-- name: FindMembersByName :many
SELECT * FROM members
WHERE (first_name ILIKE '%' || $1 || '%' OR last_name ILIKE '%' || $1 || '%')
ORDER BY last_name, first_name;

-- name: FindActiveMembersByName :many
SELECT * FROM members
WHERE (first_name ILIKE '%' || $1 || '%' OR last_name ILIKE '%' || $1 || '%')
AND is_active = TRUE
ORDER BY last_name, first_name;

-- name: FindActiveFinancialMembersByName :many
SELECT * FROM members
WHERE (first_name ILIKE '%' || $1 || '%' OR last_name ILIKE '%' || $1 || '%')
AND is_active = TRUE AND is_financial = TRUE
ORDER BY last_name, first_name;

-- name: FindLifeMembers :many
SELECT * FROM members
WHERE is_life_member = TRUE AND is_active = TRUE
ORDER BY last_name, first_name;

-- name: DeactivateMember :exec
UPDATE members
SET is_active = FALSE
WHERE membership_number = $1;

-- name: ReactivateMember :exec
UPDATE members
SET is_active = TRUE
WHERE membership_number = $1;

-- name: UpdateMemberDetails :exec
UPDATE members
SET first_name = $2,
    last_name = $3,
    email = $4,
    phone = $5,
    is_bowling_member = $6,
    is_life_member = $7,
    is_financial = $8,
    is_active = $9,
    updated_at = CURRENT_TIMESTAMP
WHERE membership_number = $1;

-- name: CreateMember :one
INSERT INTO members (membership_number,
                     first_name, 
                     last_name,
                     email,
                     phone,
                     is_bowling_member,
                     is_life_member,
                     is_financial,
                     is_active)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;