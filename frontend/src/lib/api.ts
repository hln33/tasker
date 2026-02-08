import type { CreateTaskInput, Task } from './types';

const API_BASE_URL = 'http://localhost:8080/api';

export async function createTask(taskData: CreateTaskInput): Promise<Task> {
	const res = await fetch(`${API_BASE_URL}/task`, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify(taskData)
	});

	if (!res.ok) {
		const err = await res.json().catch(() => ({ message: 'Unknown error' }));
		throw new Error(err.message || 'Failed to create task');
	}

	return res.json();
}

export async function deleteTask(taskId: string): Promise<void> {
	const res = await fetch(`${API_BASE_URL}/task/${taskId}`, {
		method: 'DELETE',
		headers: { 'Content-Type': 'application/json' }
	});

	if (!res.ok) {
		const err = await res.json().catch(() => ({ message: 'Unknown error' }));
		throw new Error(err.message || 'Failed to delete task');
	}
}

export async function updateTaskStatus(taskId: string, status: string): Promise<Task> {
	const res = await fetch(`${API_BASE_URL}/task/${taskId}`, {
		method: 'PUT',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ status })
	});
	if (!res.ok) {
		const err = await res.json().catch(() => ({ message: 'Unknown error' }));
		throw new Error(err.message || 'Failed to update task');
	}

	return res.json();
}

export async function updateTask(taskId: string, taskData: Partial<CreateTaskInput>): Promise<Task> {
	const res = await fetch(`${API_BASE_URL}/task/${taskId}`, {
		method: 'PUT',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify(taskData)
	});

	if (!res.ok) {
		const err = await res.json().catch(() => ({ message: 'Unknown error' }));
		throw new Error(err.message || 'Failed to update task');
	}

	return res.json();
}
