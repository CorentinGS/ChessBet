-- Modify "matches" table
ALTER TABLE "matches" ADD CONSTRAINT "matches_match_result_check" CHECK ((match_result >= 0) AND (match_result < 3)), ADD COLUMN "match_result" integer NULL;
