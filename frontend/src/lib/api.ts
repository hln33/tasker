import type { CreateTaskInput, Task } from './types';

const API_BASE_URL = 'http://localhost:8080/api';

export async function createTask(taskData: CreateTaskInput): Promise<Task> {
	const response = await fetch(`${API_BASE_URL}/task`, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify(taskData)
	});

	if (!response.ok) {
		const error = await response.json().catch(() => ({ error: 'Unknown error' }));
		throw new Error(error.message || error.error || 'Failed to create task');
	}

	return response.json();
}
