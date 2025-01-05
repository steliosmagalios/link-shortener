CREATE TABLE IF NOT EXISTS links (
  id BIGSERIAL PRIMARY KEY,
  location text NOT NULL,
  slug text NOT NULL UNIQUE
);

CREATE INDEX IF NOT EXISTS slug_index ON links (slug);
