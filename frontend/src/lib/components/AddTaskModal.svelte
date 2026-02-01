<script lang="ts">
	import { createTask } from '$lib/api';
	import type { CreateTaskInput } from '$lib/types';

	let { open = false, onClose }: { open: boolean; onClose: () => void } = $props();

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
			onClose();
		} catch (error) {
			errorMessage = error instanceof Error ? error.message : 'Failed to create task';
		} finally {
			isSubmitting = false;
		}
	}
</script>

{#if open}
	<div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
		<div class="bg-white rounded-lg shadow-xl max-w-md w-full">
			<div class="p-6 border-b border-gray-200">
				<h2 class="text-2xl font-bold text-gray-900">Create New Task</h2>
			</div>

			<form onsubmit={handleSubmit} class="p-6 space-y-4">
				{#if errorMessage}
					<div class="bg-red-50 border border-red-200 rounded-lg p-4">
						<p class="text-red-800 text-sm">{errorMessage}</p>
					</div>
				{/if}

				<div>
					<label for="title" class="block text-sm font-medium text-gray-700 mb-1">
						Title <span class="text-red-500">*</span>
					</label>
					<input
						id="title"
						type="text"
						bind:value={formData.title}
						disabled={isSubmitting}
						class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-100"
						placeholder="Enter task title"
						required
					/>
				</div>

				<div>
					<label for="description" class="block text-sm font-medium text-gray-700 mb-1"
						>Description</label
					>
					<textarea
						id="description"
						bind:value={formData.description}
						disabled={isSubmitting}
						rows="3"
						class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-100"
						placeholder="Enter task description (optional)"
					></textarea>
				</div>

				<div class="grid grid-cols-2 gap-4">
					<div>
						<label for="status" class="block text-sm font-medium text-gray-700 mb-1">Status</label>
						<select
							id="status"
							bind:value={formData.status}
							disabled={isSubmitting}
							class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-100"
						>
							<option value="TODO">TODO</option>
							<option value="In Progress">In Progress</option>
							<option value="Completed">Completed</option>
						</select>
					</div>

					<div>
						<label for="priority" class="block text-sm font-medium text-gray-700 mb-1"
							>Priority</label
						>
						<select
							id="priority"
							bind:value={formData.priority}
							disabled={isSubmitting}
							class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-100"
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
						class="flex-1 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 font-medium disabled:bg-gray-400"
					>
						{isSubmitting ? 'Creating...' : 'Create Task'}
					</button>
					<button
						type="button"
						onclick={handleClose}
						disabled={isSubmitting}
						class="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 font-medium"
					>
						Cancel
					</button>
				</div>
			</form>
		</div>
	</div>
{/if}
