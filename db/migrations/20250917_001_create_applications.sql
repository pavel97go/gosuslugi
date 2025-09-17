-- +goose Up
DO $do$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'application_type') THEN
    CREATE TYPE application_type AS ENUM ('passport', 'certificate');
  END IF;
END
$do$;

DO $do$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'application_status') THEN
    CREATE TYPE application_status AS ENUM ('draft', 'submitted', 'approved', 'rejected');
  END IF;
END
$do$;

CREATE TABLE IF NOT EXISTS applications (
    id            BIGSERIAL PRIMARY KEY,
    citizen_name  TEXT                     NOT NULL,
    document_type application_type         NOT NULL,
    data          JSONB                    NOT NULL DEFAULT '{}',
    status        application_status       NOT NULL DEFAULT 'draft',
    created_at    TIMESTAMPTZ              NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ              NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_applications_status         ON applications (status);
CREATE INDEX IF NOT EXISTS idx_applications_document_type  ON applications (document_type);
CREATE INDEX IF NOT EXISTS idx_applications_created_at     ON applications (created_at DESC);

DROP FUNCTION IF EXISTS set_updated_at();

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $set_updated_at$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$set_updated_at$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_applications_updated_at ON applications;

CREATE TRIGGER trg_applications_updated_at
BEFORE UPDATE ON applications
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

-- +goose Down
DROP TRIGGER IF EXISTS trg_applications_updated_at ON applications;
DROP FUNCTION IF EXISTS set_updated_at();
DROP INDEX IF EXISTS idx_applications_created_at;
DROP INDEX IF EXISTS idx_applications_document_type;
DROP INDEX IF EXISTS idx_applications_status;
DROP TABLE IF EXISTS applications;

DO $do$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type t
                 JOIN pg_depend d ON d.refobjid = t.oid
                 WHERE t.typname = 'application_status' AND d.classid = 'pg_type'::regclass) THEN
    DROP TYPE IF EXISTS application_status;
  END IF;
END
$do$;

DO $do$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type t
                 JOIN pg_depend d ON d.refobjid = t.oid
                 WHERE t.typname = 'application_type' AND d.classid = 'pg_type'::regclass) THEN
    DROP TYPE IF EXISTS application_type;
  END IF;
END
$do$;