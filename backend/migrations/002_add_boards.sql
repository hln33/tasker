-- Migration: Add boards feature
-- This migration creates the boards table and adds board_id to tasks

-- Create boards table
CREATE TABLE IF NOT EXISTS boards (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    color VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Add board_id column to tasks table (nullable initially)
ALTER TABLE tasks ADD COLUMN IF NOT EXISTS board_id VARCHAR(50) REFERENCES boards(id);

-- Create default board
INSERT INTO boards (id, name, description, color, created_at, updated_at)
VALUES ('BOARD-001', 'Default Board', 'Your default task board', '#3B82F6', NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

-- Migrate existing tasks to default board
UPDATE tasks
SET board_id = 'BOARD-001'
WHERE board_id IS NULL;

-- Make board_id required after migration
ALTER TABLE tasks ALTER COLUMN board_id SET NOT NULL;
