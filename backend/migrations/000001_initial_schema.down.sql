-- Rollback: Drop all tables and indexes

-- Drop tasks table (indexes will be dropped automatically)
DROP TABLE IF EXISTS tasks CASCADE;

-- Drop boards table
DROP TABLE IF EXISTS boards CASCADE;
