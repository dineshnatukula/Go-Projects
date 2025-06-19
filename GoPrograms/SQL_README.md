explain select distinct trim(sportname) as sportname from config.bbmarketsmapping


select sportname as sportname from config.bbmarketsmapping
Option B: Clean data at input to ensure sportname is already trimmed.
ALTER TABLE bbmarketsmapping ADD COLUMN trimmed_sportname TEXT GENERATED ALWAYS AS (TRIM(sportname)) STORED;
CREATE INDEX ON bbmarketsmapping(trimmed_sportname);

ANALYZE config.bbmarketsmapping

