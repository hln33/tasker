<script lang="ts">
	import type { Task } from '$lib/types';
	import { dndzone } from 'svelte-dnd-action';
	import TaskCard from './TaskCard.svelte';
	import Checkmark from '$lib/icons/Checkmark.svelte';
	import Pencil from '$lib/icons/Pencil.svelte';
	import FastForward from '$lib/icons/FastForward.svelte';

	const TODO = 'TODO';
	const IN_PROGRESS = 'In Progress';
	const DONE = 'Done';

	let {
		type,
		tasks,
		emptyMessage,
		onDelete
	}: {
		type: typeof TODO | typeof IN_PROGRESS | typeof DONE;
		tasks: Task[];
		emptyMessage: string;
		onDelete: (task: Task) => void;
	} = $props();

	const colorClasses = {
		[TODO]: {
			bg: 'bg-gray-50',
			text: 'text-gray-700',
			outline: 'rgba(107, 114, 128, 0.5)',
			background: 'rgba(107, 114, 128, 0.1)'
		},
		[IN_PROGRESS]: {
			bg: 'bg-blue-50',
			text: 'text-blue-700',
			outline: 'rgba(59, 130, 246, 0.5)',
			background: 'rgba(59, 130, 246, 0.1)'
		},
		[DONE]: {
			bg: 'bg-green-50',
			text: 'text-green-700',
			outline: 'rgba(34, 197, 94, 0.5)',
			background: 'rgba(34, 197, 94, 0.1)'
		}
	};

	const currentColor = $derived(() => colorClasses[type]);

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

<div class="h-fit min-h-96 {currentColor().bg} rounded-lg p-4 shadow-md">
	<div class="mb-4 flex items-center gap-2">
		<h2 class="text-lg font-bold {currentColor().text}">{type}</h2>
		{#if type === TODO}
			<Pencil class="{currentColor().text} size-5" />
		{:else if type === IN_PROGRESS}
			<FastForward class="{currentColor().text} size-5" />
		{:else if type === DONE}
			<Checkmark class={currentColor().text}/>
		{/if}
	</div>

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
