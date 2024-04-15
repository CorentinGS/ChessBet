CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    points INTEGER NOT NULL CHECK (points >= 0),
    oauth_id VARCHAR(50) UNIQUE,
    email_address VARCHAR(100) UNIQUE NOT NULL
);


CREATE TABLE players (
    player_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    rating INTEGER NOT NULL
);

CREATE TABLE tournaments (
    tournament_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL
);

CREATE TABLE matches (
    match_id SERIAL PRIMARY KEY,
    tournament_id INTEGER REFERENCES tournaments(tournament_id),
    player1_id INTEGER REFERENCES players(player_id) NOT NULL,
    player2_id INTEGER REFERENCES players(player_id) NOT NULL,
    match_date TIMESTAMP NOT NULL
);

CREATE TABLE bets (
    bet_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id),
    match_id INTEGER REFERENCES matches(match_id),
    bet_points INTEGER NOT NULL,
    bet_result BOOLEAN, -- True if the user's bet was correct, False otherwise
    bet_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    bet_value INTEGER NOT NULL CHECK (bet_value >= 0 AND bet_value < 3), -- 0 for white, 1 for draw, 2 for black
    CONSTRAINT one_bet_per_user_per_match UNIQUE (user_id, match_id)
);

CREATE TABLE user_points_history (
    transaction_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id),
    transaction_date TIMESTAMP NOT NULL,
    transaction_amount INTEGER NOT NULL, -- Positive for gaining points, negative for losing points
    remaining_points INTEGER NOT NULL
);
