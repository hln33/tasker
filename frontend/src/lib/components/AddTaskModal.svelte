<script lang="ts">
	import { createTask } from '$lib/api';
	import type { CreateTaskInput } from '$lib/types';

	let {
		open = false,
		onClose,
		onAddSuccess
	}: { open: boolean; onClose: () => void; onAddSuccess: () => void } = $props();

	let isSubmitting = $state(false);
	let errorMessage = $state('');
	let formData = $state<CreateTaskInput>({
		title: '',
		description: '',
		status: 'TODO',
		priority: 'Medium'
	});

	function handleClose() {
		if (!isSubmitting) {
			onClose();
			resetForm();
		}
	}

	function resetForm() {
		formData = { title: '', description: '', status: 'TODO', priority: 'Medium' };
		errorMessage = '';
	}

	async function handleSubmit(e: Event) {
		e.preventDefault();

		if (!formData.title.trim()) {
			errorMessage = 'Title is required';
			return;
		}

		isSubmitting = true;
		errorMessage = '';

		try {
			await createTask(formData);
			resetForm();
			onAddSuccess();
      onClose();
		} catch (error) {
			errorMessage = error instanceof Error ? error.message : 'Failed to create task';
		} finally {
			isSubmitting = false;
		}
	}
</script>

{#if open}
	<div class="fixed inset-0 z-50 flex items-center justify-center bg-black/80 p-4">
		<div class="w-full max-w-md rounded-lg bg-white shadow-xl">
			<div class="border-b border-gray-200 p-6">
				<h2 class="text-2xl font-bold text-gray-900">Create New Task</h2>
			</div>

			<form onsubmit={handleSubmit} class="space-y-4 p-6">
				{#if errorMessage}
					<div class="rounded-lg border border-red-200 bg-red-50 p-4">
						<p class="text-sm text-red-800">{errorMessage}</p>
					</div>
				{/if}

				<div>
					<label for="title" class="mb-1 block text-sm font-medium text-gray-700">
						Title <span class="text-red-500">*</span>
					</label>
					<input
						id="title"
						type="text"
						bind:value={formData.title}
						disabled={isSubmitting}
						class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500 disabled:bg-gray-100"
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
						bind:value={formData.description}
						disabled={isSubmitting}
						rows="3"
						class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500 disabled:bg-gray-100"
						placeholder="Enter task description (optional)"
					></textarea>
				</div>

				<div class="grid grid-cols-2 gap-4">
					<div>
						<label for="status" class="mb-1 block text-sm font-medium text-gray-700">
							Status
						</label>
						<select
							id="status"
							bind:value={formData.status}
							disabled={isSubmitting}
							class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500 disabled:bg-gray-100"
						>
							<option value="TODO">TODO</option>
							<option value="In Progress">In Progress</option>
							<option value="Done">Done</option>
						</select>
					</div>

					<div>
						<label for="priority" class="mb-1 block text-sm font-medium text-gray-700"
							>Priority</label
						>
						<select
							id="priority"
							bind:value={formData.priority}
							disabled={isSubmitting}
							class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500 disabled:bg-gray-100"
						>
							<option value="Low">Low</option>
							<option value="Medium">Medium</option>
							<option value="High">High</option>
						</select>
					</div>
				</div>

				<div class="flex gap-3 pt-4">
					<button
						type="submit"
						disabled={isSubmitting}
						class="flex-1 cursor-pointer rounded-lg bg-blue-600 px-4 py-2 font-medium text-white hover:bg-blue-700 disabled:bg-gray-400"
					>
						{isSubmitting ? 'Creating...' : 'Create Task'}
					</button>
					<button
						type="button"
						onclick={handleClose}
						disabled={isSubmitting}
						class="cursor-pointer rounded-lg border border-gray-300 px-4 py-2 font-medium text-gray-700 hover:bg-gray-50"
					>
						Cancel
					</button>
				</div>
			</form>
		</div>
	</div>
{/if}
