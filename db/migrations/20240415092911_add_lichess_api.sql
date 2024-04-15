-- Modify "matches" table
ALTER TABLE "matches" ADD COLUMN "round_name" character varying(50) NOT NULL, ADD COLUMN "lichess_round_id" character varying(50) NOT NULL;
-- Modify "players" table
ALTER TABLE "players" ADD COLUMN "image_url" character varying(100) NULL;
-- Modify "tournaments" table
ALTER TABLE "tournaments" ADD COLUMN "lichess_tournament_id" character varying(50) NOT NULL;
