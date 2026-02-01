-- Drop indexes
DROP INDEX IF EXISTS idx_tasks_status;
DROP INDEX IF EXISTS idx_tasks_priority;

-- Drop tasks table
DROP TABLE IF EXISTS tasks;
