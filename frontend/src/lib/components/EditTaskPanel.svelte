<script lang="ts">
	import { enhance } from '$app/forms';
	import { page } from '$app/state';
	import type { Task } from '$lib/types';

	let {
		open = false,
		task,
		onClose,
		onEditSuccess
	}: {
		open: boolean;
		task: Task | null;
		onClose: () => void;
		onEditSuccess: () => void;
	} = $props();

	function handleClose() {
		onClose();
	}

	function handleBackdropClick(e: MouseEvent) {
		if (e.target === e.currentTarget) {
			handleClose();
		}
	}
</script>

{#if open}
	<div
		onclick={handleBackdropClick}
		onkeydown={(e) => e.key === 'Escape' && handleClose()}
		role="dialog"
		aria-modal="true"
		tabindex="-1"
		class="fixed inset-0 z-50 flex justify-end bg-black/80 p-4"
	>
		<div
			class="w-full max-w-lg bg-white shadow-xl transition-transform duration-300 ease-in-out"
			class:translate-x-full={!open}
			class:translate-x-0={open}
		>
			<div class="border-b border-gray-200 p-6">
				<h2 class="text-2xl font-bold text-gray-900">Edit Task</h2>
				<p class="mt-1 text-sm text-gray-500">Update task details below</p>
			</div>

			<form
				method="POST"
				action="?/updateTask"
				use:enhance={() => {
					return async ({ result, update }) => {
						if (result.type === 'success') {
							onEditSuccess();
							onClose();
							await update();
						}
					};
				}}
				class="space-y-4 p-6"
			>
				<!-- Hidden input for task ID -->
				<input type="hidden" name="taskId" value={task?.id} />

				{#if page.form?.error}
					<div class="rounded-lg border border-red-200 bg-red-50 p-4">
						<p class="text-sm text-red-800">{page.form.error}</p>
					</div>
				{/if}

				<div>
					<label for="edit-title" class="mb-1 block text-sm font-medium text-gray-700">
						Title <span class="text-red-500">*</span>
					</label>
					<input
						id="edit-title"
						name="title"
						type="text"
						value={page.form?.title ?? task?.title ?? ''}
						class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500"
						placeholder="Enter task title"
						required
					/>
				</div>

				<div>
					<label for="edit-description" class="mb-1 block text-sm font-medium text-gray-700">
						Description
					</label>
					<textarea
						id="edit-description"
						name="description"
						rows="3"
						class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500"
						placeholder="Enter task description (optional)"
					>
						{page.form?.description ?? task?.description ?? ''}
					</textarea>
				</div>

				<div>
					<label for="edit-priority" class="mb-1 block text-sm font-medium text-gray-700">
						Priority
					</label>
					<select
						id="edit-priority"
						name="priority"
						class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500"
					>
						<option value="Low" selected={page.form?.priority === 'Low' || task?.priority === 'Low'}
							>Low</option
						>
						<option
							value="Medium"
							selected={page.form?.priority === 'Medium' || task?.priority === 'Medium'}
						>
							Medium
						</option>
						<option
							value="High"
							selected={page.form?.priority === 'High' || task?.priority === 'High'}>High</option
						>
					</select>
				</div>

				<div class="flex gap-3 pt-4">
					<button
						type="submit"
						class="flex-1 cursor-pointer rounded-lg bg-blue-600 px-4 py-2 font-medium text-white hover:bg-blue-700"
					>
						Save Task
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
