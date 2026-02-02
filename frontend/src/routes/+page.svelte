<script lang="ts">
	import type { PageData } from './$types';
	import type { Task } from '$lib/types';
	import { invalidate } from '$app/navigation';
	import AddTaskModal from '$lib/components/AddTaskModal.svelte';
	import DeleteTaskModal from '$lib/components/DeleteTaskModal.svelte';
	import TaskCard from '$lib/components/TaskCard.svelte';

	let { data }: { data: PageData } = $props();

	let showAddTaskModal = $state(false);
	let showDeleteModal = $state(false);
	let taskToDelete = $state<Task | null>(null);
	let successMessage = $state('');

	$effect(() => {
		console.log('Data updated:', data);
	});

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
			{@const todoTasks = tasks.filter(t => t.status === 'TODO')}
			{@const inProgressTasks = tasks.filter(t => t.status === 'In Progress')}
			{@const doneTasks = tasks.filter(t => t.status === 'Done')}
			<div class="grid grid-cols-1 gap-6 md:grid-cols-3">
				<!-- TODO Column -->
				<div class="rounded-lg bg-gray-50 p-4">
					<h2 class="mb-4 text-lg font-bold text-gray-700">TODO</h2>
					<div class="flex flex-col gap-4">
						{#each todoTasks as task}
							<TaskCard {task} onDelete={openDeleteModal} />
						{/each}
						{#if todoTasks.length === 0}
							<p class="text-sm italic text-gray-500">No tasks to do</p>
						{/if}
					</div>
				</div>

				<!-- In Progress Column -->
				<div class="rounded-lg bg-blue-50 p-4">
					<h2 class="mb-4 text-lg font-bold text-blue-700">In Progress</h2>
					<div class="flex flex-col gap-4">
						{#each inProgressTasks as task}
							<TaskCard {task} onDelete={openDeleteModal} />
						{/each}
						{#if inProgressTasks.length === 0}
							<p class="text-sm italic text-gray-500">No tasks in progress</p>
						{/if}
					</div>
				</div>

				<!-- Done Column -->
				<div class="rounded-lg bg-green-50 p-4">
					<h2 class="mb-4 text-lg font-bold text-green-700">Done</h2>
					<div class="flex flex-col gap-4">
						{#each doneTasks as task}
							<TaskCard {task} onDelete={openDeleteModal} />
						{/each}
						{#if doneTasks.length === 0}
							<p class="text-sm italic text-gray-500">No completed tasks</p>
						{/if}
					</div>
				</div>
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
			onClose={() => {
				closeDeleteModal();
				handleTaskDeleted();
			}}
		/>
	</div>
</div>
