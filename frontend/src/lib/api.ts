import type { CreateTaskInput, Task } from './types';

const API_BASE_URL = 'http://localhost:8080/api';

export async function createTask(taskData: CreateTaskInput): Promise<Task> {
	const res = await fetch(`${API_BASE_URL}/task`, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify(taskData)
	});

	if (!res.ok) {
		const error = await res.json().catch(() => ({ error: 'Unknown error' }));
		throw new Error(error.message || error.error || 'Failed to create task');
	}

	return res.json();
}

export async function deleteTask(taskId: string): Promise<void> {
	const res = await fetch(`${API_BASE_URL}/task/${taskId}`, {
		method: 'DELETE',
		headers: { 'Content-Type': 'application/json' }
	});

	if (!res.ok) {
		const error = await res.json().catch(() => ({ error: 'Unknown error' }));
		throw new Error(error.message || error.error || 'Failed to delete task');
	}
}
