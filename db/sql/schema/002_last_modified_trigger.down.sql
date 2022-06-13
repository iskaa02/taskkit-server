DROP TRIGGER IF EXIST tgr_list_last_modified ON list;
DROP TRIGGER IF EXIST tgr_task_last_modified ON task;

DROP FUNCTION IF EXIST bump_timestamp;
