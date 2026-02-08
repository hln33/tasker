<script lang="ts">
	import type { PageData } from './$types';
	import type { Task } from '$lib/types';
	import { invalidate } from '$app/navigation';
	import AddTaskModal from '$lib/components/AddTaskModal.svelte';
	import DeleteTaskModal from '$lib/components/DeleteTaskModal.svelte';
	import EditTaskPanel from '$lib/components/EditTaskPanel.svelte';
	import TaskColumn from '$lib/components/TaskColumn.svelte';

	let { data }: { data: PageData } = $props();

	let showAddTaskModal = $state(false);
	let showDeleteModal = $state(false);
	let taskToDelete = $state<Task | null>(null);
	let taskToEdit = $state<Task | null>(null);
	let isEditPanelOpen = $state(false);
	let successMessage = $state('');

	function openAddTaskModal() {
		showAddTaskModal = true;
		successMessage = '';
	}

	function closeAddTaskModal() {
		showAddTaskModal = false;
	}

	async function handleTaskCreated() {
		await invalidate('http://localhost:8080/api/task');
		successMessage = 'Task created successfully!';
		setTimeout(() => {
			successMessage = '';
		}, 3000);
	}

	function openDeleteModal(task: Task) {
		taskToDelete = task;
		showDeleteModal = true;
		successMessage = '';
	}

	function closeDeleteModal() {
		showDeleteModal = false;
		taskToDelete = null;
	}

	async function handleTaskDeleted() {
		await invalidate('http://localhost:8080/api/task');
		successMessage = 'Task deleted successfully!';
		setTimeout(() => {
			successMessage = '';
		}, 3000);
	}

	function openEditPanel(task: Task) {
		taskToEdit = task;
		isEditPanelOpen = true;
	}

	function closeEditPanel() {
		taskToEdit = null;
		isEditPanelOpen = false;
	}

	async function handleTaskEdited() {
		await invalidate('http://localhost:8080/api/task');
		successMessage = 'Task updated successfully!';
		setTimeout(() => {
			successMessage = '';
		}, 3000);
		closeEditPanel();
	}
</script>

<div class="min-h-screen bg-gray-200 px-4 py-12">
	<div class="mx-auto max-w-7xl">
		<div class="mb-8 flex items-center justify-between">
			<h1 class="text-3xl font-bold text-gray-900">Tasker</h1>
			<button
				onclick={openAddTaskModal}
				class="cursor-pointer rounded-lg bg-blue-600 px-4 py-2 font-medium text-white transition-colors hover:bg-blue-700"
			>
				Add Task
			</button>
		</div>

		{#if successMessage}
			<div class="mb-4 rounded-lg border border-green-200 bg-green-50 p-4">
				<p class="text-green-800">{successMessage}</p>
			</div>
		{/if}

		{#await data.tasks}
			<div class="rounded-lg bg-white p-6 shadow">
				<p class="text-gray-700">Loading task...</p>
			</div>
		{:then tasks}
			<div class="grid grid-cols-1 gap-6 md:grid-cols-3">
				<TaskColumn
					type="TODO"
					tasks={tasks.filter((t) => t.status === 'TODO')}
					emptyMessage="No tasks to do"
					onDelete={openDeleteModal}
					onEdit={openEditPanel}
				/>

				<TaskColumn
					type="In Progress"
					tasks={tasks.filter((t) => t.status === 'In Progress')}
					emptyMessage="No tasks in progress"
					onDelete={openDeleteModal}
					onEdit={openEditPanel}
				/>

				<TaskColumn
					type="Done"
					tasks={tasks.filter((t) => t.status === 'Done')}
					emptyMessage="No completed tasks"
					onDelete={openDeleteModal}
					onEdit={openEditPanel}
				/>
			</div>
		{:catch error}
			<div class="rounded-lg bg-white p-6 shadow">
				<div class="rounded-lg border border-red-200 bg-red-50 p-4">
					<p class="font-medium text-red-800">Error</p>
					<p class="mt-1 text-red-600">{error.message}</p>
				</div>
				<p class="mt-4 text-sm text-gray-500">
					Make sure the backend is running on http://localhost:8080
				</p>
			</div>
		{/await}

		<AddTaskModal
			open={showAddTaskModal}
			onClose={() => {
				closeAddTaskModal();
				handleTaskCreated();
			}}
		/>
		<DeleteTaskModal
			open={showDeleteModal}
			task={taskToDelete}
			onClose={closeDeleteModal}
      onDeleteSuccess={handleTaskDeleted}
		/>
		<EditTaskPanel
			open={isEditPanelOpen}
			task={taskToEdit}
			onClose={closeEditPanel}
			onEditSuccess={handleTaskEdited}
		/>
	</div>
</div>

<style>
	/* Only for library-specific classes not covered by Tailwind utilities */
	:global(.shadow-element) {
		background: #e5e7eb !important;
		border: 2px dashed #9ca3af !important;
		opacity: 0.6;
	}
</style>
