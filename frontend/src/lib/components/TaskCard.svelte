<script lang="ts">
	import type { Task } from '$lib/types';

	let { task, onDelete }: { task: Task; onDelete: (task: Task) => void } = $props();

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
			case 'done':
				return 'bg-green-100 text-green-800';
			case 'todo':
				return 'bg-gray-100 text-gray-800';
			default:
				return 'bg-gray-100 text-gray-800';
		}
	};
</script>

<div class="cursor-grab rounded-lg bg-white shadow-md *:px-4 transition-shadow hover:shadow-sm active:cursor-grabbing">
	<div class="border-b border-gray-200 py-3">
		<div class="flex items-start justify-between">
			<div>
				<span class="text-sm font-medium text-gray-500">{task.id}</span>
				<h2 class="mt-3 text-2xl font-bold text-gray-900">{task.title}</h2>
			</div>

			<div class="flex items-center gap-3">
				<span class="rounded-full px-3 py-1 text-sm font-medium {getPriorityColor(task.priority)}">
					{task.priority}
				</span>
				<button
					onclick={() => onDelete(task)}
					class="cursor-pointer rounded-lg p-2 text-red-600 transition-colors hover:bg-red-200"
					aria-label="Delete task"
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 24 24"
						stroke-width="1.5"
						stroke="currentColor"
						class="h-5 w-5 hover:stroke-red-700"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0"
						/>
					</svg>
				</button>
			</div>
		</div>
	</div>

	<div class="py-6">
		<h3 class="mb-2 text-sm font-medium text-gray-500">Description:</h3>
		{#if task.description === ''}
			<p class="text-gray-400 italic">N/A</p>
		{:else}
			<p class="text-gray-700">{task.description}</p>
		{/if}
	</div>

	<div class="rounded-b-lg py-6">
		<div class="flex items-center">
			<span class="mr-2 text-sm text-gray-500">Status:</span>
			<span class="rounded-full px-3 py-1 text-sm font-medium {getStatusColor(task.status)}">
				{task.status}
			</span>
		</div>
	</div>
</div>
