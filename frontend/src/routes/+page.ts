import type { PageLoad } from './$types';

type Task = {
	id: string;
	title: string;
	description: string;
	status: string;
	priority: string;
};

export const load: PageLoad = async ({ fetch }) => {
	try {
		const response = await fetch('http://localhost:8080/api/task');
		if (!response.ok) {
			throw new Error(`HTTP error! status: ${response.status}`);
		}
		const task: Task = await response.json();
		return {
			task,
			error: null
		};
	} catch (e) {
		return {
			task: null,
			error: e instanceof Error ? e.message : 'Failed to fetch task'
		};
	}
};
