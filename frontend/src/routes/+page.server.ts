import { fail } from '@sveltejs/kit';
import * as api from '$lib/api';
import { createTaskSchema, taskIdSchema, updateTaskSchema } from '$lib/schemas';
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
		const rawData = Object.fromEntries(formData);

		// Validate with Zod
		const result = createTaskSchema.safeParse(rawData);
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
			return fail(500, {
				error: e instanceof Error ? e.message : 'Failed to create task',
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
			return fail(400, {
				error: result.error.issues[0].message
			});
		}

		try {
			await api.deleteTask(result.data.taskId);
			return { success: true };
		} catch (e) {
			return fail(500, {
				error: e instanceof Error ? e.message : 'Failed to delete task'
			});
		}
	},

	updateTask: async ({ request }) => {
		const formData = await request.formData();
		const rawData = Object.fromEntries(formData);

		// Validate with Zod (using updateTaskSchema which includes taskId)
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
			return fail(500, {
				error: e instanceof Error ? e.message : 'Failed to update task',
				title: rawData.title?.toString() || '',
				description: rawData.description?.toString() || '',
				priority: rawData.priority?.toString() || 'Medium'
			});
		}
	}
};
