CREATE TABLE users (
    user_id SERIAL PRIMARY KEY NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    points INTEGER NOT NULL CHECK (points >= 0),
    oauth_id VARCHAR(50) UNIQUE,
    email_address VARCHAR(100) UNIQUE NOT NULL
);


CREATE TABLE players (
    player_id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    rating INTEGER NOT NULL,
    image_url VARCHAR(100)
);

CREATE TABLE tournaments (
    tournament_id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    lichess_tournament_id VARCHAR(50) NOT NULL UNIQUE,
    CONSTRAINT fk_matches_tournaments FOREIGN KEY (tournament_id) REFERENCES tournaments(tournament_id) ON DELETE CASCADE
);

CREATE TABLE matches (
    match_id SERIAL PRIMARY KEY NOT NULL,
    tournament_id INTEGER REFERENCES tournaments(tournament_id) NOT NULL,
    player1_id INTEGER REFERENCES players(player_id) NOT NULL,
    player2_id INTEGER REFERENCES players(player_id) NOT NULL,
    match_date TIMESTAMP NOT NULL,
    round_name VARCHAR(50) NOT NULL,
    lichess_round_id VARCHAR(50) NOT NULL,
    lichess_game_id VARCHAR(50),
    match_result INTEGER CHECK (match_result >= 0 AND match_result < 3) -- 0 for white, 1 for draw, 2 for black
);

CREATE TABLE bets (
    bet_id SERIAL PRIMARY KEY NOT NULL,
    user_id INTEGER REFERENCES users(user_id),
    match_id INTEGER REFERENCES matches(match_id),
    bet_points INTEGER NOT NULL,
    bet_result BOOLEAN, -- True if the user's bet was correct, False otherwise
    bet_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    bet_value INTEGER NOT NULL CHECK (bet_value >= 0 AND bet_value < 3), -- 0 for white, 1 for draw, 2 for black
    CONSTRAINT one_bet_per_user_per_match UNIQUE (user_id, match_id)
);

CREATE TABLE user_points_history (
    transaction_id SERIAL PRIMARY KEY NOT NULL,
    user_id INTEGER REFERENCES users(user_id),
    transaction_date TIMESTAMP NOT NULL,
    transaction_amount INTEGER NOT NULL, -- Positive for gaining points, negative for losing points
    remaining_points INTEGER NOT NULL
);
