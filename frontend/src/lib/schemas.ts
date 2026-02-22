import { z } from 'zod';

/**
 * Task status enum - matches backend TaskStatus type
 */
export const taskStatusEnum = z.enum(['TODO', 'In Progress', 'Done']);

/**
 * Task priority enum - matches backend TaskPriority type
 */
export const taskPriorityEnum = z.enum(['Low', 'Medium', 'High']);

/**
 * Schema for creating/updating a task
 */
export const createTaskSchema = z.object({
	title: z.string().min(1, 'Title is required').trim(),
	description: z.string().optional().default(''),
	status: taskStatusEnum.default('TODO'),
	priority: taskPriorityEnum.default('Medium'),
	board_id: z.coerce.number().int().positive('Board ID is required')
});

/**
 * Schema for task ID validation (used in delete/update operations)
 */
export const taskIdSchema = z.object({
	taskId: z.coerce.number().int().positive('Invalid task ID format')
});

/**
 * Schema for updating an existing task
 * All fields are optional since we're doing partial updates
 */
export const updateTaskSchema = z.object({
	taskId: z.coerce.number().int().positive('Invalid task ID format'),
	title: z.string().min(1, 'Title is required').trim().optional(),
	description: z.string().optional(),
	priority: taskPriorityEnum.optional()
});

/**
 * Schema for creating/updating a board
 */
export const createBoardSchema = z.object({
	name: z.string().min(1, 'Board name is required').trim(),
	description: z.string().optional().default(''),
	color: z
		.string()
		.regex(/^#[0-9A-F]{6}$/i, 'Color must be a valid hex color (e.g., #3B82F6)')
		.default('#3B82F6')
});

/**
 * Type exports inferred from schemas
 */
export type CreateTaskInput = z.infer<typeof createTaskSchema>;
export type TaskIdInput = z.infer<typeof taskIdSchema>;
export type UpdateTaskInput = z.infer<typeof updateTaskSchema>;
export type CreateBoardInput = z.infer<typeof createBoardSchema>;
export type TaskStatus = z.infer<typeof taskStatusEnum>;
export type TaskPriority = z.infer<typeof taskPriorityEnum>;
