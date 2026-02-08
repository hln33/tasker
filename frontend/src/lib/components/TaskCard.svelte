<script lang="ts">
	import type { Task } from '$lib/types';
	import Pencil from '$lib/icons/Pencil.svelte';
	import Trash from '$lib/icons/Trash.svelte';

	let {
		task,
		onDelete,
		onEdit
	}: { task: Task; onDelete: (task: Task) => void; onEdit: (task: Task) => void } = $props();

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

<div
	class="cursor-grab space-y-6 rounded-lg bg-white p-4 shadow-md transition-shadow hover:shadow-sm active:cursor-grabbing"
>
	<div class="border-b border-gray-200 p-0">
		<div class="flex flex-col">
			<div class="flex w-full items-center">
				{#if task.priority !== ''}
					<span
						class="flex justify-start rounded-full px-3 py-1 text-sm font-medium {getPriorityColor(task.priority)}"
					>
						{task.priority}
					</span>
				{/if}

				<div class="ml-auto">
					<button
						onclick={() => onEdit(task)}
						class="cursor-pointer rounded-lg p-2 text-blue-600 transition-colors hover:bg-blue-200"
						aria-label="Edit task"
					>
						<Pencil class="size-5" />
					</button>
					<button
						onclick={() => onDelete(task)}
						class="cursor-pointer rounded-lg p-2 text-red-600 transition-colors hover:bg-red-200"
						aria-label="Delete task"
					>
						<Trash class="size-5" />
					</button>
				</div>
			</div>

			<h2 class="mt-3 text-2xl font-bold text-gray-900">{task.title}</h2>
		</div>
	</div>

	<div>
		<h3 class="mb-2 text-sm font-medium text-gray-500">Description:</h3>
		{#if task.description === ''}
			<p class="text-gray-400 italic">N/A</p>
		{:else}
			<p class="text-gray-700">{task.description}</p>
		{/if}
	</div>

	<div class="flex items-center rounded-b-lg">
		<span class="mr-2 text-sm text-gray-500">Status:</span>
		<span class="rounded-full px-3 py-1 text-sm font-medium {getStatusColor(task.status)}">
			{task.status}
		</span>
	</div>
</div>
