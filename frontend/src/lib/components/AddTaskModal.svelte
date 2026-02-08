<script lang="ts">
	import { enhance } from '$app/forms';
	import { page } from '$app/state';
	import type { ActionResult } from '@sveltejs/kit';

	let {
		open = false,
		onClose,
		onAddSuccess,
	}: {
		open: boolean;
		onClose: () => void;
		onAddSuccess: () => void;
	} = $props();

	function handleClose() {
		onClose();
	}
</script>

{#if open}
	<div class="fixed inset-0 z-50 flex items-center justify-center bg-black/80 p-4">
		<div class="w-full max-w-md rounded-lg bg-white shadow-xl">
			<div class="border-b border-gray-200 p-6">
				<h2 class="text-2xl font-bold text-gray-900">Create New Task</h2>
			</div>

			<form
				method="POST"
				action="?/createTask"
				use:enhance={() => {
					return async ({ result, update }) => {
            if (result.type === 'success') {
              onAddSuccess();
              onClose();
            }
            await update();
					};
				}}
				class="space-y-4 p-6"
			>
				{#if page.form?.error}
					<div class="rounded-lg border border-red-200 bg-red-50 p-4">
						<p class="text-sm text-red-800">{page.form.error}</p>
					</div>
				{/if}

				<div>
					<label for="title" class="mb-1 block text-sm font-medium text-gray-700">
						Title <span class="text-red-500">*</span>
					</label>
					<input
						id="title"
						name="title"
						type="text"
						value={page.form?.title ?? ''}
						class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500"
						placeholder="Enter task title"
						required
					/>
				</div>

				<div>
					<label for="description" class="mb-1 block text-sm font-medium text-gray-700">
						Description
					</label>
					<textarea
						id="description"
						name="description"
						rows="3"
						class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500"
						placeholder="Enter task description (optional)"
					>
						{page.form?.description ?? ''}
					</textarea>
				</div>

				<div class="grid grid-cols-2 gap-4">
					<div>
						<label for="status" class="mb-1 block text-sm font-medium text-gray-700">
							Status
						</label>
						<select
							id="status"
							name="status"
							class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500"
						>
							<option value="TODO" selected={page.form?.status === 'TODO' || !page.form?.status}>TODO</option>
							<option value="In Progress" selected={page.form?.status === 'In Progress'}
								>In Progress</option
							>
							<option value="Done" selected={page.form?.status === 'Done'}>Done</option>
						</select>
					</div>

					<div>
						<label for="priority" class="mb-1 block text-sm font-medium text-gray-700"
							>Priority</label
						>
						<select
							id="priority"
							name="priority"
							class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500"
						>
							<option value="Low" selected={page.form?.priority === 'Low'}>Low</option>
							<option value="Medium" selected={page.form?.priority === 'Medium' || !page.form?.priority}
								>Medium</option
							>
							<option value="High" selected={page.form?.priority === 'High'}>High</option>
						</select>
					</div>
				</div>

				<div class="flex gap-3 pt-4">
					<button
						type="submit"
						class="flex-1 cursor-pointer rounded-lg bg-blue-600 px-4 py-2 font-medium text-white hover:bg-blue-700"
					>
						Create Task
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
