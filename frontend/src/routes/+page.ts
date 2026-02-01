import type { PageLoad } from './$types';
import type { Task } from '$lib/types';

export const load: PageLoad = async ({ fetch }) => {
	const response = await fetch('http://localhost:8080/api/task');
	if (!response.ok) {
		throw new Error(`HTTP error! status: ${response.status}`);
	}

	const tasks: Promise<Task[]> = response.json();
	return {
		tasks,
		error: null
	};
};
