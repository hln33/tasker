export interface Task {
	id: string;
	title: string;
	description: string;
	status: string;
	priority: string;
}

export interface CreateTaskInput {
	title: string;
	description?: string;
	status?: string;
	priority?: string;
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
