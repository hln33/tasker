import { fail } from '@sveltejs/kit';
import * as api from '$lib/api';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	const boardId = parseInt(params.id, 10);
	if (isNaN(boardId)) {
		throw new Error('Invalid board ID');
	}

	try {
		const [board, tasks] = await Promise.all([api.getBoard(boardId), api.getBoardTasks(boardId)]);

		return { board, tasks };
	} catch (error) {
		console.error('Failed to load board:', error);
		throw error;
	}
};

export const actions: Actions = {
	createTask: async ({ request, params }) => {
		const boardId = parseInt(params.id, 10);
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
			const task = await api.createTask({
				title: title.trim(),
				description: description?.toString() || '',
				status: status.toString() as 'TODO' | 'In Progress' | 'Done',
				priority: priority.toString() as 'Low' | 'Medium' | 'High',
				board_id: boardId
			});

			return { success: true, task };
		} catch (e) {
			const errorMessage = e instanceof Error ? e.message : 'Failed to create task';
			return fail(500, {
				error: errorMessage,
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
			return fail(400, { error: 'Task ID is required' });
		}

		const taskIdNum = parseInt(taskId, 10);
		if (isNaN(taskIdNum)) {
			return fail(400, { error: 'Invalid task ID format' });
		}

		try {
			await api.deleteTask(taskIdNum);
			return { success: true };
		} catch (e) {
			const errorMessage = e instanceof Error ? e.message : 'Failed to delete task';
			return fail(500, { error: errorMessage });
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
			const task = await api.updateTask(taskIdNum, {
				title: title.trim(),
				description: description?.toString() || '',
				priority: priority.toString() as 'Low' | 'Medium' | 'High'
			});

			return { success: true, task };
		} catch (e) {
			const errorMessage = e instanceof Error ? e.message : 'Failed to update task';
			return fail(500, {
				error: errorMessage,
				title: title.toString(),
				description: description?.toString() || '',
				priority: priority.toString()
			});
		}
	}
};
