-- Modify "tournaments" table
ALTER TABLE "tournaments" ALTER COLUMN "start_date" TYPE timestamp, ALTER COLUMN "end_date" TYPE timestamp, ADD CONSTRAINT "tournaments_lichess_tournament_id_key" UNIQUE ("lichess_tournament_id");
