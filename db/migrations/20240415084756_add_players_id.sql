-- Modify "bets" table
ALTER TABLE "bets" ADD CONSTRAINT "bets_bet_value_check" CHECK ((bet_value >= 0) AND (bet_value < 3)), ADD COLUMN "bet_date" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, ADD COLUMN "bet_value" integer NOT NULL;
-- Create "players" table
CREATE TABLE "players" ("player_id" serial NOT NULL, "name" character varying(50) NOT NULL, "rating" integer NOT NULL, PRIMARY KEY ("player_id"));
-- Modify "matches" table
ALTER TABLE "matches" DROP COLUMN "player1", DROP COLUMN "player2", DROP COLUMN "win_probability_player1", DROP COLUMN "draw_probability", DROP COLUMN "win_probability_player2", ADD COLUMN "player1_id" integer NOT NULL, ADD COLUMN "player2_id" integer NOT NULL, ADD CONSTRAINT "matches_player1_id_fkey" FOREIGN KEY ("player1_id") REFERENCES "players" ("player_id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "matches_player2_id_fkey" FOREIGN KEY ("player2_id") REFERENCES "players" ("player_id") ON UPDATE NO ACTION ON DELETE NO ACTION;
