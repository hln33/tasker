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
