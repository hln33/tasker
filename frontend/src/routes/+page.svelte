<script lang="ts">
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();

	const getPriorityColor = (priority: string) => {
		switch (priority.toLowerCase()) {
			case 'high':
				return 'bg-red-100 text-red-800';
			case 'medium':
				return 'bg-yellow-100 text-yellow-800';
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
</script>

<div class="min-h-screen bg-gray-200 py-12 px-4">
	<div class="max-w-2xl mx-auto">
		<h1 class="text-3xl font-bold text-gray-900 mb-8">Tasker</h1>

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
                <span class="px-3 py-1 rounded-full text-sm font-medium {getPriorityColor(task.priority)}">
                  {task.priority}
                </span>
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
	</div>
</div>
