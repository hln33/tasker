<script lang="ts">
	import type { Task } from '$lib/types';
	import { dndzone } from 'svelte-dnd-action';
	import TaskCard from './TaskCard.svelte';

	let {
		title,
		tasks,
		color,
		emptyMessage,
		onDelete
	}: {
		title: string;
		tasks: Task[];
		color: 'gray' | 'blue' | 'green';
		emptyMessage: string;
		onDelete: (task: Task) => void;
	} = $props();

	const colorClasses = {
		gray: {
			bg: 'bg-gray-50',
			text: 'text-gray-700',
			outline: 'rgba(107, 114, 128, 0.5)',
			background: 'rgba(107, 114, 128, 0.1)'
		},
		blue: {
			bg: 'bg-blue-50',
			text: 'text-blue-700',
			outline: 'rgba(59, 130, 246, 0.5)',
			background: 'rgba(59, 130, 246, 0.1)'
		},
		green: {
			bg: 'bg-green-50',
			text: 'text-green-700',
			outline: 'rgba(34, 197, 94, 0.5)',
			background: 'rgba(34, 197, 94, 0.1)'
		}
	};

	const currentColor = $derived(() => colorClasses[color]);

	const handleDndConsider = (e: CustomEvent) => {
		tasks = e.detail.items;

		const info = e.detail.info;
		if (info.trigger === 'DROPPED_INTO_ANOTHER' || info.trigger === 'DROPPED_INTO_ZONE') {
			console.log('Task moved to Done - Status update stub:', {
				taskId: info.id,
				newStatus: 'Done',
				sourceColumn: info.source
			});
		}
	};

	const handleDndFinalize = (e: CustomEvent) => {
		tasks = e.detail.items;

		const info = e.detail.info;
		if (info.trigger === 'DROPPED_INTO_ANOTHER' || info.trigger === 'DROPPED_INTO_ZONE') {
			console.log('Task moved to Done - Status update stub:', {
				taskId: info.id,
				newStatus: 'Done',
				sourceColumn: info.source
			});
		}
	};
</script>

<div class="{currentColor().bg} rounded-lg p-4 shadow-md">
	<h2 class="mb-4 text-lg font-bold {currentColor().text}">{title}</h2>

	<section
		use:dndzone={{
			items: tasks,
			type: 'task',
			flipDurationMs: 200,
			dropTargetStyle: {
				outline: `${currentColor().outline} solid 3px`,
				backgroundColor: currentColor().background
			}
		}}
		onconsider={(e) => handleDndConsider(e)}
		onfinalize={(e) => handleDndFinalize(e)}
		class="flex min-h-52 flex-col gap-4 rounded-lg p-2 transition-all duration-100 ease-in-out"
	>
		{#each tasks as task (task.id)}
			<TaskCard {task} {onDelete} />
		{/each}
	</section>

	{#if tasks.length === 0}
		<p class="text-sm text-gray-500 italic">{emptyMessage}</p>
	{/if}
</div>
