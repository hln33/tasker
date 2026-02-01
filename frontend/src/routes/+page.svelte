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

<div class="min-h-screen bg-gray-200 py-12 px-4">
	<div class="max-w-2xl mx-auto">
		<div class="flex justify-between items-center mb-8">
			<h1 class="text-3xl font-bold text-gray-900">Tasker</h1>
			<button
				onclick={openAddTaskModal}
				class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors font-medium cursor-pointer"
			>
				Add Task
			</button>
		</div>

		{#if successMessage}
			<div class="mb-4 bg-green-50 border border-green-200 rounded-lg p-4">
				<p class="text-green-800">{successMessage}</p>
			</div>
		{/if}

    {#await data.tasks}
      <div class="bg-white rounded-lg shadow p-6">
        <p class="text-gray-700">Loading task...</p>
      </div>
    {:then tasks}
      <div class="flex flex-col gap-6">
        {#each tasks as task}
          <TaskCard task={task} onDelete={openDeleteModal} />
        {/each}
      </div>
    {:catch error}
      <div class="bg-white rounded-lg shadow p-6">
        <div class="bg-red-50 border border-red-200 rounded-lg p-4">
          <p class="text-red-800 font-medium">Error</p>
          <p class="text-red-600 mt-1">{error.message}</p>
        </div>
        <p class="text-gray-500 text-sm mt-4">
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
