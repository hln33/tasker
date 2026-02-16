import { fail } from '@sveltejs/kit';
import { createTask, deleteTask, updateTask } from '$lib/api';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch }) => {
	const res = await fetch('http://localhost:8080/api/task');
	if (!res.ok) throw new Error(`HTTP error! status: ${res.status}`);

	const tasks = await res.json();

	return { tasks };
};

export const actions: Actions = {
	createTask: async ({ request }) => {
		const formData = await request.formData();
		const title = formData.get('title');
		const description = formData.get('description');
		const status = formData.get('status') || 'TODO';
		const priority = formData.get('priority') || 'Medium';

		if (!title || typeof title !== 'string' || !title.trim()) {
			return fail(400, {
				error: 'Title is required',
				title: title?.toString() || '',
				description: description?.toString() || '',
				status: status.toString(),
				priority: priority.toString()
			});
		}

		try {
			const task = await createTask({
				title: title.trim(),
				description: description?.toString() || '',
				status: status.toString() as 'TODO' | 'In Progress' | 'Done',
				priority: priority.toString() as 'Low' | 'Medium' | 'High'
			});

			return { success: true, task };
		} catch (e) {
			return fail(500, {
				error: e instanceof Error ? e.message : 'Failed to create task',
				title: title.toString(),
				description: description?.toString() || '',
				status: status.toString(),
				priority: priority.toString()
			});
		}
	},

	deleteTask: async ({ request }) => {
		const formData = await request.formData();
		const taskId = formData.get('taskId');

		if (!taskId || typeof taskId !== 'string') {
			return fail(400, {
				error: 'Task ID is required'
			});
		}

		// Parse string to number
		const taskIdNum = parseInt(taskId, 10);
		if (isNaN(taskIdNum)) {
			return fail(400, { error: 'Invalid task ID format' });
		}

		try {
			await deleteTask(taskIdNum);
			return { success: true };
		} catch (e) {
			return fail(500, {
				error: e instanceof Error ? e.message : 'Failed to delete task'
			});
		}
	},

	updateTask: async ({ request }) => {
		const formData = await request.formData();
		const taskId = formData.get('taskId');
		const title = formData.get('title');
		const description = formData.get('description');
		const priority = formData.get('priority') || 'Medium';

		if (!taskId || typeof taskId !== 'string') {
			return fail(400, { error: 'Task ID is required' });
		}

		// Parse string to number
		const taskIdNum = parseInt(taskId, 10);
		if (isNaN(taskIdNum)) {
			return fail(400, { error: 'Invalid task ID format' });
		}

		if (!title || typeof title !== 'string' || !title.trim()) {
			return fail(400, {
				error: 'Title is required',
				title: '',
				description: description?.toString() || '',
				priority: priority.toString()
			});
		}

		try {
			const task = await updateTask(taskIdNum, {
				title: title.trim(),
				description: description?.toString() || '',
				priority: priority.toString() as 'Low' | 'Medium' | 'High'
			});

			return { success: true, task };
		} catch (e) {
			return fail(500, {
				error: e instanceof Error ? e.message : 'Failed to update task',
				title: title.toString(),
				description: description?.toString() || '',
				priority: priority.toString()
			});
		}
	}
};
