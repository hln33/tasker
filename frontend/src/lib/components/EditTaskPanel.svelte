<script lang="ts">
	import { updateTask } from '$lib/api';
	import type { CreateTaskInput, Task } from '$lib/types';

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

	let isSubmitting = $state(false);
	let errorMessage = $state('');
	let formData = $state<Partial<CreateTaskInput>>({
		title: '',
		description: '',
		priority: 'Medium'
	});

	// Update form data when task changes
	$effect(() => {
		if (task) {
			formData = {
				title: task.title,
				description: task.description || '',
				priority: task.priority
			};
		}
	});

	function handleClose() {
		if (!isSubmitting) {
			onClose();
			resetForm();
		}
	}

	function resetForm() {
		formData = { title: '', description: '', priority: 'Medium' };
		errorMessage = '';
	}

	async function handleSubmit(e: Event) {
		e.preventDefault();

		if (!formData.title?.trim()) {
			errorMessage = 'Title is required';
			return;
		}
		if (!task) {
			errorMessage = 'No task to edit';
			return;
		}

		isSubmitting = true;
		errorMessage = '';

		try {
			await updateTask(task.id, formData);
			resetForm();
			onEditSuccess();
		} catch (error) {
			errorMessage = error instanceof Error ? error.message : 'Failed to update task';
		} finally {
			isSubmitting = false;
		}
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
		class="bg-opacity-50 fixed inset-0 z-50 flex justify-end bg-black p-4"
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

			<form onsubmit={handleSubmit} class="space-y-4 p-6">
				{#if errorMessage}
					<div class="rounded-lg border border-red-200 bg-red-50 p-4">
						<p class="text-sm text-red-800">{errorMessage}</p>
					</div>
				{/if}

				<div>
					<label for="edit-title" class="mb-1 block text-sm font-medium text-gray-700">
						Title <span class="text-red-500">*</span>
					</label>
					<input
						id="edit-title"
						type="text"
						bind:value={formData.title}
						disabled={isSubmitting}
						class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500 disabled:bg-gray-100"
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
						bind:value={formData.description}
						disabled={isSubmitting}
						rows="3"
						class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500 disabled:bg-gray-100"
						placeholder="Enter task description (optional)"
					></textarea>
				</div>

				<div>
					<label for="edit-priority" class="mb-1 block text-sm font-medium text-gray-700">
						Priority
					</label>
					<select
						id="edit-priority"
						bind:value={formData.priority}
						disabled={isSubmitting}
						class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500 disabled:bg-gray-100"
					>
						<option value="Low">Low</option>
						<option value="Medium">Medium</option>
						<option value="High">High</option>
					</select>
				</div>

				<div class="flex gap-3 pt-4">
					<button
						type="submit"
						disabled={isSubmitting}
						class="flex-1 cursor-pointer rounded-lg bg-blue-600 px-4 py-2 font-medium text-white hover:bg-blue-700 disabled:bg-gray-400"
					>
						{isSubmitting ? 'Saving...' : 'Save Task'}
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
