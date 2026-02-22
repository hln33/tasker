// Re-export types inferred from Zod schemas
export type {
	CreateTaskInput,
	TaskIdInput,
	UpdateTaskInput,
	CreateBoardInput,
	TaskStatus,
	TaskPriority
} from './schemas';

// Import the types locally so we can use them in interfaces
import type { TaskStatus, TaskPriority } from './schemas';

// Database response types (these include fields like id, created_at, etc.)
export interface Task {
	id: number;
	title: string;
	description: string;
	status: TaskStatus;
	priority: TaskPriority;
	board_id: number;
}

export interface Board {
	id: number;
	name: string;
	description: string;
	color: string;
	created_at: string;
	updated_at: string;
}
