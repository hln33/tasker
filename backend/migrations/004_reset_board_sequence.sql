-- Migration 004: Reset board sequence to fix duplicate key issue
-- This fixes the sequence so new boards can be created with auto-incrementing IDs

-- Reset the sequence to start from the next available ID (max existing ID + 1)
SELECT setval('boards_id_seq', (SELECT COALESCE(MAX(id), 0) + 1 FROM boards));
