BEGIN;

ALTER TABLE exercise ADD COLUMN "type" TEXT NOT NULL;

ALTER TABLE exercise ADD COLUMN question TEXT;
ALTER TABLE exercise ADD COLUMN choices TEXT[];
ALTER TABLE exercise ADD COLUMN answer TEXT;

ALTER TABLE exercise ADD COLUMN sentence TEXT;
ALTER TABLE exercise ADD COLUMN corrected_sentence TEXT;

COMMIT;