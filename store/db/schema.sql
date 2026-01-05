CREATE TABLE IF NOT EXISTS members (
    membership_number INTEGER PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE,
    phone VARCHAR(20),
    is_bowling_member BOOLEAN DEFAULT FALSE,
    is_life_member BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS competitions (
    id SERIAL PRIMARY KEY,
    start_year INTEGER NOT NULL,
    division VARCHAR(50) NOT NULL,
    bowlslink_id VARCHAR(50),
    bowlers_per_team INTEGER NOT NULL,
    teams_per_side INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS rounds (
    id SERIAL PRIMARY KEY,
    competition_id INTEGER REFERENCES competitions(id) ON DELETE CASCADE,
    round_number INTEGER NOT NULL,
    duty_selector INTEGER NOT NULL REFERENCES members(membership_number),
    umpire INTEGER REFERENCES members(membership_number),
    played_at TIMESTAMPTZ NOT NULL,
    venue VARCHAR(100) NOT NULL,
    opponent VARCHAR(100) NOT NULL,
    is_played_on_surface BOOLEAN DEFAULT FALSE,
    surface VARCHAR(10) NOT NULL,
    low_rink_number INTEGER,
    high_rink_number INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS teams (
    id SERIAL PRIMARY KEY,
    round_id INTEGER REFERENCES rounds(id) ON DELETE CASCADE,
    team_number INTEGER NOT NULL,
    duty VARCHAR(10),
    role_lead INTEGER NOT NULL REFERENCES members(membership_number),
    role_second INTEGER REFERENCES members(membership_number),
    role_third INTEGER REFERENCES members(membership_number),
    role_skip INTEGER NOT NULL REFERENCES members(membership_number),
    shots_for INTEGER DEFAULT 0,
    shots_against INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sides (
    id SERIAL PRIMARY KEY,
    competition_id INTEGER REFERENCES competitions(id) ON DELETE CASCADE,
    team_1_id INTEGER REFERENCES teams(id) ON DELETE CASCADE,
    team_2_id INTEGER REFERENCES teams(id) ON DELETE CASCADE,
    team_3_id INTEGER REFERENCES teams(id) ON DELETE CASCADE,
    team_4_id INTEGER REFERENCES teams(id) ON DELETE CASCADE,
    points INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);