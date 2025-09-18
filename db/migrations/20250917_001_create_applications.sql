-- +goose Up

-- +goose StatementBegin
DO $do$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'application_type') THEN
    CREATE TYPE application_type AS ENUM ('passport', 'certificate');
  END IF;

  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'application_status') THEN
    CREATE TYPE application_status AS ENUM ('draft', 'submitted', 'approved', 'rejected');
  END IF;
END
$do$;
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS applications (
  id            BIGSERIAL PRIMARY KEY,
  citizen_name  TEXT NOT NULL,
  document_type application_type NOT NULL,
  data          JSONB NOT NULL DEFAULT '{}'::jsonb,
  status        application_status NOT NULL DEFAULT 'draft',
  created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_applications_status        ON applications (status);
CREATE INDEX IF NOT EXISTS idx_applications_document_type ON applications (document_type);
CREATE INDEX IF NOT EXISTS idx_applications_created_at    ON applications (created_at);

-- +goose StatementBegin
DO $do$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_proc WHERE proname = 'set_updated_at') THEN
    CREATE OR REPLACE FUNCTION set_updated_at()
    RETURNS TRIGGER AS $f$
    BEGIN
      NEW.updated_at = NOW();
      RETURN NEW;
    END;
    $f$ LANGUAGE plpgsql;
  END IF;
END
$do$;
-- +goose StatementEnd

DROP TRIGGER IF EXISTS trg_applications_updated_at ON applications;
CREATE TRIGGER trg_applications_updated_at
BEFORE UPDATE ON applications
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- +goose Down

DROP TRIGGER IF EXISTS trg_applications_updated_at ON applications;
DROP FUNCTION IF EXISTS set_updated_at();
DROP TABLE IF EXISTS applications;

-- +goose StatementBegin
DO $do$
BEGIN
  IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'application_status') THEN
    DROP TYPE application_status;
  END IF;
  IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'application_type') THEN
    DROP TYPE application_type;
  END IF;
END
$do$;
-- +goose StatementEnd
