CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE todo (
  id INTEGER PRIMARY KEY,
  content text NOT NULL,
  done BOOLEAN NOT NULL DEFAULT FALSE
, creationDate DATETIME);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20240727074719'),
  ('20240727121140');
