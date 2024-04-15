-- Modify "tournaments" table
ALTER TABLE "tournaments" ADD CONSTRAINT "fk_matches_tournaments" FOREIGN KEY ("tournament_id") REFERENCES "tournaments" ("tournament_id") ON UPDATE NO ACTION ON DELETE CASCADE;
