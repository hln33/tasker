import { fail } from '@sveltejs/kit';
import * as api from '$lib/api';
import { createTaskSchema, taskIdSchema, updateTaskSchema } from '$lib/schemas';
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
		const rawData = Object.fromEntries(formData);

		// Add board_id from URL params
		const dataWithBoardId = { ...rawData, board_id: boardId };

		// Validate with Zod
		const result = createTaskSchema.safeParse(dataWithBoardId);

		if (!result.success) {
			const firstError = result.error.issues[0].message;
			return fail(400, {
				error: firstError,
				title: rawData.title?.toString() || '',
				description: rawData.description?.toString() || '',
				status: rawData.status?.toString() || 'TODO',
				priority: rawData.priority?.toString() || 'Medium'
			});
		}

		try {
			const task = await api.createTask(result.data);
			return { success: true, task };
		} catch (e) {
			const errorMessage = e instanceof Error ? e.message : 'Failed to create task';
			return fail(500, {
				error: errorMessage,
				title: rawData.title?.toString() || '',
				description: rawData.description?.toString() || '',
				status: rawData.status?.toString() || 'TODO',
				priority: rawData.priority?.toString() || 'Medium'
			});
		}
	},

	deleteTask: async ({ request }) => {
		const formData = await request.formData();
		const rawData = Object.fromEntries(formData);

		// Validate with Zod
		const result = taskIdSchema.safeParse(rawData);
		if (!result.success) {
			return fail(400, { error: result.error.issues[0].message });
		}

		try {
			await api.deleteTask(result.data.taskId);
			return { success: true };
		} catch (e) {
			const errorMessage = e instanceof Error ? e.message : 'Failed to delete task';
			return fail(500, { error: errorMessage });
		}
	},

	updateTask: async ({ request }) => {
		const formData = await request.formData();
		const rawData = Object.fromEntries(formData);

		// Validate with Zod
		const result = updateTaskSchema.safeParse(rawData);
		if (!result.success) {
			const firstError = result.error.issues[0].message;
			return fail(400, {
				error: firstError,
				title: rawData.title?.toString() || '',
				description: rawData.description?.toString() || '',
				priority: rawData.priority?.toString() || 'Medium'
			});
		}

		const { taskId, ...taskData } = result.data;

		try {
			const task = await api.updateTask(taskId, taskData);
			return { success: true, task };
		} catch (e) {
			const errorMessage = e instanceof Error ? e.message : 'Failed to update task';
			return fail(500, {
				error: errorMessage,
				title: rawData.title?.toString() || '',
				description: rawData.description?.toString() || '',
				priority: rawData.priority?.toString() || 'Medium'
			});
		}
	}
};
