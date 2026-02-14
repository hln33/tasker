<script lang="ts">
	import { enhance } from '$app/forms';
	import { page } from '$app/state';
	import type { ActionResult } from '@sveltejs/kit';
	import type { Task } from '$lib/types';

	let {
		open = false,
		task,
		onClose,
		onDeleteSuccess
	}: {
		open: boolean;
		task: Task | null;
		onClose: () => void;
		onDeleteSuccess: () => void;
	} = $props();

	function handleClose() {
		onClose();
	}
</script>

{#if open}
	<div class="bg-opacity-50 fixed inset-0 z-50 flex items-center justify-center bg-black p-4">
		<div class="w-full max-w-md rounded-lg bg-white shadow-xl">
			<div class="border-b border-gray-200 p-6">
				<h2 class="text-2xl font-bold text-gray-900">Delete Task</h2>
			</div>

			<form
				method="POST"
				action="?/deleteTask"
				use:enhance={() => {
					return async ({ result, update }) => {
						if (result.type === 'success') {
							onDeleteSuccess();
							onClose();
						}
            await update();
					};
				}}
				class="space-y-4 p-6"
			>
				<input type="hidden" name="taskId" value={task?.id} />

				{#if page.form?.error}
					<div class="rounded-lg border border-red-200 bg-red-50 p-4">
						<p class="text-sm text-red-800">{page.form.error}</p>
					</div>
				{/if}

				<p class="text-gray-700">
					Are you sure you want to delete the task
					<span class="font-semibold text-gray-900">"{task?.title}"</span>? This action cannot be
					undone.
				</p>

				<div class="flex gap-3 pt-4">
					<button
						type="submit"
						class="flex-1 cursor-pointer rounded-lg bg-red-600 px-4 py-2 font-medium text-white hover:bg-red-700"
					>
						Delete
					</button>
					<button
						type="button"
						onclick={handleClose}
						class="cursor-pointer rounded-lg border border-gray-300 px-4 py-2 font-medium text-gray-700 hover:bg-gray-50"
					>
						Cancel
					</button>
				</div>
			</form>
		</div>
	</div>
{/if}
