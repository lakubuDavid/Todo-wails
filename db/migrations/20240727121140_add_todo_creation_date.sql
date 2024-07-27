-- migrate:up
ALTER TABLE todo
ADD creationDate DATETIME;

-- migrate:down
ALTER TABLE todo
DROP COLUMN CreationDate
