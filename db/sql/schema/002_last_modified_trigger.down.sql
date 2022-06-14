DROP TRIGGER IF EXISTS tgr_list_last_modified ON list;
DROP TRIGGER IF EXISTS tgr_task_last_modified ON task;

DROP FUNCTION IF EXISTS bump_timestamp;
