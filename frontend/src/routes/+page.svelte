<script lang="ts">
	import type { PageData } from './$types';
	import type { Task } from '$lib/types';
	import { invalidate } from '$app/navigation';
	import AddTaskModal from '$lib/components/AddTaskModal.svelte';
	import DeleteTaskModal from '$lib/components/DeleteTaskModal.svelte';

	let { data }: { data: PageData } = $props();

	let showAddTaskModal = $state(false);
	let showDeleteModal = $state(false);
	let taskToDelete = $state<Task | null>(null);
	let successMessage = $state('');

	const getPriorityColor = (priority: string) => {
		switch (priority.toLowerCase()) {
			case 'high':
				return 'bg-red-100 text-red-800';
			case 'medium':
				return 'bg-yellow-300 text-yellow-800';
			case 'low':
				return 'bg-green-100 text-green-800';
			default:
				return 'bg-gray-100 text-gray-800';
		}
	};

	const getStatusColor = (status: string) => {
		switch (status.toLowerCase()) {
			case 'in progress':
				return 'bg-blue-100 text-blue-800';
			case 'completed':
				return 'bg-green-100 text-green-800';
			case 'todo':
				return 'bg-gray-100 text-gray-800';
			default:
				return 'bg-gray-100 text-gray-800';
		}
	};

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
        {#each tasks as task }
          <div class="bg-white rounded-lg shadow-lg">
            <div class="p-6 border-b border-gray-200">
              <div class="flex items-start justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-500">{task.id}</span>
                  <h2 class="text-2xl font-bold text-gray-900 mt-1">{task.title}</h2>
                </div>
                <div class="flex items-center gap-3">
                  <span class="px-3 py-1 rounded-full text-sm font-medium {getPriorityColor(task.priority)}">
                    {task.priority}
                  </span>
                  <button
                    onclick={() => openDeleteModal(task)}
                    class="p-2 text-red-600 hover:bg-red-200 rounded-lg transition-colors cursor-pointer"
                    aria-label="Delete task"
                  >
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="hover:stroke-red-700 w-5 h-5">
                      <path stroke-linecap="round" stroke-linejoin="round" d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0" />
                    </svg>
                  </button>
                </div>
              </div>
            </div>
  
            <div class="p-6">
              <h3 class="text-sm font-medium text-gray-500 mb-2">Description</h3>
              <p class="text-gray-700">{task.description}</p>
            </div>
  
            <div class="p-6 rounded-b-lg">
              <div class="flex items-center">
                <span class="text-sm text-gray-500 mr-2">Status:</span>
                <span class="px-3 py-1 rounded-full text-sm font-medium {getStatusColor(task.status)}">
                  {task.status}
                </span>
              </div>
            </div>
          </div>
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
