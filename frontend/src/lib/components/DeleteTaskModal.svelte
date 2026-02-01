<script lang="ts">
	import { deleteTask } from '$lib/api';
	import type { Task } from '$lib/types';

	let {
		open = false,
		task,
		onClose
	}: {
		open: boolean;
		task: Task | null;
		onClose: () => void;
	} = $props();

	let isDeleting = $state(false);
	let errorMessage = $state('');

	async function handleDelete() {
		if (!task || isDeleting) return;

		isDeleting = true;
		errorMessage = '';

		try {
			await deleteTask(task.id);
			onClose();
		} catch (error) {
			errorMessage = error instanceof Error ? error.message : 'Failed to delete task';
		} finally {
			isDeleting = false;
		}
	}

	function handleClose() {
		if (!isDeleting) {
			onClose();
			errorMessage = '';
		}
	}
</script>

{#if open}
	<div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
		<div class="bg-white rounded-lg shadow-xl max-w-md w-full">
			<div class="p-6 border-b border-gray-200">
				<h2 class="text-2xl font-bold text-gray-900">Delete Task</h2>
			</div>

			<div class="p-6 space-y-4">
				{#if errorMessage}
					<div class="bg-red-50 border border-red-200 rounded-lg p-4">
						<p class="text-red-800 text-sm">{errorMessage}</p>
					</div>
				{/if}

				<p class="text-gray-700">
					Are you sure you want to delete the task
					<span class="font-semibold text-gray-900">"{task?.title}"</span
					>? This action cannot be undone.
				</p>

				<div class="flex gap-3 pt-4">
					<button
						onclick={handleDelete}
						disabled={isDeleting}
						class="flex-1 px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 font-medium disabled:bg-gray-400 cursor-pointer"
						type="button"
					>
						{isDeleting ? 'Deleting...' : 'Delete'}
					</button>
					<button
						onclick={handleClose}
						disabled={isDeleting}
						class="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-200 font-medium cursor-pointer"
						type="button"
					>
						Cancel
					</button>
				</div>
			</div>
		</div>
	</div>
{/if}
