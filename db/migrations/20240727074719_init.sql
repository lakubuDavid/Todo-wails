-- migrate:up

CREATE TABLE todo (
  id INTEGER PRIMARY KEY,
  content text NOT NULL,
  done BOOLEAN NOT NULL DEFAULT FALSE
);

-- migrate:down

DROP TABLE todo
