import { fail } from '@sveltejs/kit';
import * as api from '$lib/api';
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
		const name = formData.get('name');
		const description = formData.get('description') || '';
		const color = formData.get('color') || '#3B82F6';

		if (!name || typeof name !== 'string' || !name.trim()) {
			return fail(400, {
				error: 'Board name is required',
				name: name?.toString() || '',
				description: description.toString(),
				color: color.toString()
			});
		}

		try {
			const board = await api.createBoard({
				name: name.trim(),
				description: description.toString(),
				color: color.toString()
			});

			return { success: true, board };
		} catch (e) {
			return fail(500, {
				error: e instanceof Error ? e.message : 'Failed to create board',
				name: name.toString(),
				description: description.toString(),
				color: color.toString()
			});
		}
	}
};
