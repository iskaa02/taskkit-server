CREATE OR REPLACE FUNCTION bump_timestamp()
RETURNS trigger AS $$
DECLARE
    previous TIMESTAMP;
    current TIMESTAMP;
BEGIN
    previous :=OLD.last_modified;
    --
    -- This bumps the current timestamp to 1 msec in the future if the previous
    -- timestamp is equal to the current one (or higher if was bumped already).
    --
    -- If a bunch of requests from the same user on the same resource
    -- arrive in the same millisecond, the unicity constraint can raise
    -- an error (operation is cancelled).
    -- See https://github.com/mozilla-services/cliquet/issues/25
    --
    current := clock_timestamp();
    IF previous IS NOT NULL AND previous >= current THEN
        current := previous + INTERVAL '1 milliseconds';
    END IF;
        NEW.last_modified := current;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER tgr_list_last_modified
BEFORE INSERT OR UPDATE ON list
FOR EACH ROW EXECUTE PROCEDURE bump_timestamp();

CREATE TRIGGER tgr_task_last_modified
BEFORE INSERT OR UPDATE ON task
FOR EACH ROW EXECUTE PROCEDURE bump_timestamp();
