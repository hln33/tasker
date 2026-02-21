export interface Task {
	id: number;
	title: string;
	description: string;
	status: string;
	priority: string;
	board_id: number;
}

export interface CreateTaskInput {
	title: string;
	description?: string;
	status?: string;
	priority?: string;
	board_id?: number;
}

export interface Board {
	id: number;
	name: string;
	description: string;
	color: string;
	created_at: string;
	updated_at: string;
}

export interface CreateBoardInput {
	name: string;
	description?: string;
	color?: string;
}
