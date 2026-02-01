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
	<div class="bg-opacity-50 fixed inset-0 z-50 flex items-center justify-center bg-black p-4">
		<div class="w-full max-w-md rounded-lg bg-white shadow-xl">
			<div class="border-b border-gray-200 p-6">
				<h2 class="text-2xl font-bold text-gray-900">Delete Task</h2>
			</div>

			<div class="space-y-4 p-6">
				{#if errorMessage}
					<div class="rounded-lg border border-red-200 bg-red-50 p-4">
						<p class="text-sm text-red-800">{errorMessage}</p>
					</div>
				{/if}

				<p class="text-gray-700">
					Are you sure you want to delete the task
					<span class="font-semibold text-gray-900">"{task?.title}"</span>? This action cannot be undone.
				</p>

				<div class="flex gap-3 pt-4">
					<button
						onclick={handleDelete}
						disabled={isDeleting}
						class="flex-1 cursor-pointer rounded-lg bg-red-600 px-4 py-2 font-medium text-white hover:bg-red-700 disabled:bg-gray-400"
						type="button"
					>
						{isDeleting ? 'Deleting...' : 'Delete'}
					</button>
					<button
						onclick={handleClose}
						disabled={isDeleting}
						class="cursor-pointer rounded-lg border border-gray-300 px-4 py-2 font-medium text-gray-700 hover:bg-gray-200"
						type="button"
					>
						Cancel
					</button>
				</div>
			</div>
		</div>
	</div>
{/if}
