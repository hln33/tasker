import { fail } from '@sveltejs/kit';
import * as api from '$lib/api';
import { createBoardSchema } from '$lib/schemas';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	try {
		const boards = await api.getBoards();
		return { boards };
	} catch (error) {
		console.error('Failed to load boards:', error);
		return { boards: [] };
	}
};

export const actions: Actions = {
	createBoard: async ({ request }) => {
		const formData = await request.formData();
		const rawData = Object.fromEntries(formData);

		// Validate with Zod
		const result = createBoardSchema.safeParse(rawData);
		if (!result.success) {
			const firstError = result.error.issues[0].message;
			return fail(400, {
				error: firstError,
				name: rawData.name?.toString() || '',
				description: rawData.description?.toString() || '',
				color: rawData.color?.toString() || '#3B82F6'
			});
		}

		try {
			const board = await api.createBoard(result.data);
			return { success: true, board };
		} catch (e) {
			return fail(500, {
				error: e instanceof Error ? e.message : 'Failed to create board',
				name: rawData.name?.toString() || '',
				description: rawData.description?.toString() || '',
				color: rawData.color?.toString() || '#3B82F6'
			});
		}
	}
};
