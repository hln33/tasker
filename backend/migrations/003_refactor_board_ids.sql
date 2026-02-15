-- Migration 003: Refactor board IDs from manual string IDs to auto-generated integer IDs
-- This migration drops and recreates the boards and tasks tables to use SERIAL for board IDs

-- Drop and recreate boards table with SERIAL
DROP TABLE tasks CASCADE;
DROP TABLE boards CASCADE;

-- Recreate boards with SERIAL
CREATE TABLE boards (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    color VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Recreate tasks with INTEGER board_id
CREATE TABLE tasks (
    id VARCHAR(50) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'TODO',
    priority VARCHAR(50) NOT NULL DEFAULT 'Medium',
    board_id INTEGER NOT NULL DEFAULT 1 REFERENCES boards(id) ON DELETE RESTRICT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create default board
INSERT INTO boards (id, name, description, color, created_at, updated_at)
VALUES (1, 'Default Board', 'Your default task board', '#3B82F6', NOW(), NOW());
