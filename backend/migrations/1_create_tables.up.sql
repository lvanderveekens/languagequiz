BEGIN;

CREATE OR REPLACE FUNCTION trigger_set_updated_at() RETURNS TRIGGER AS $$
    BEGIN
        NEW.updated_at = NOW();
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS quiz(
    id UUID NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    language_tag TEXT NOT NULL,
    "name" TEXT NOT NULL
);

CREATE TRIGGER set_updated_at
    BEFORE UPDATE
    ON quiz
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

CREATE TABLE IF NOT EXISTS quiz_section(
    id UUID NOT NULL PRIMARY KEY,
    quiz_id UUID NOT NULL REFERENCES quiz (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "name" TEXT NOT NULL
);

CREATE TRIGGER set_updated_at
    BEFORE UPDATE
    ON quiz_section
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

CREATE TABLE IF NOT EXISTS exercise(
    id UUID NOT NULL PRIMARY KEY,
    quiz_section_id UUID NOT NULL REFERENCES quiz_section (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    feedback TEXT,
    "type" TEXT NOT NULL,
    question TEXT,
    choices TEXT[],
    answer TEXT,
    sentence TEXT,
    corrected_sentence TEXT
);

CREATE TRIGGER set_updated_at
    BEFORE UPDATE
    ON exercise
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

COMMIT;