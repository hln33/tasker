<script lang="ts">
	import type { PageProps } from './$types';
	import AddBoardModal from '$lib/components/AddBoardModal.svelte';
	
	let { data }: PageProps = $props();

	let showAddBoardModal = $state(false);
	let successMessage = $state('');

	function openAddBoardModal() {
		showAddBoardModal = true;
		successMessage = '';
	}

	function closeAddBoardModal() {
		showAddBoardModal = false;
	}

	function handleBoardCreated() {
		successMessage = 'Board created successfully!';
		setTimeout(() => {
			successMessage = '';
		}, 3000);
	}
</script>

<div class="max-w-7xl">
	<div class="mb-8 flex items-center justify-between">
		<h1 class="text-3xl font-bold text-gray-900">Boards</h1>
		<button
			onclick={openAddBoardModal}
			class="cursor-pointer rounded-lg bg-blue-600 px-4 py-2 font-medium text-white transition-colors hover:bg-blue-700"
		>
			Add Board
		</button>
	</div>

	{#if successMessage}
		<div class="mb-4 rounded-lg border border-green-200 bg-green-50 p-4">
			<p class="text-green-800">{successMessage}</p>
		</div>
	{/if}

	{#await data.boards}
		<div class="rounded-lg bg-white p-6 shadow">
			<p class="text-gray-700">Loading boards...</p>
		</div>
	{:then boards}
		{#if boards.length === 0}
			<div class="rounded-lg bg-white p-12 shadow text-center">
				<p class="text-xl text-gray-600 mb-4">No boards yet</p>
				<button
					onclick={openAddBoardModal}
					class="cursor-pointer rounded-lg bg-blue-600 px-6 py-3 font-medium text-white transition-colors hover:bg-blue-700"
				>
					Create Your First Board
				</button>
			</div>
		{:else}
			<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
				{#each boards as board (board.id)}
					<div
						class="rounded-lg bg-white shadow hover:shadow-lg transition-shadow"
						style="border-left: 4px solid {board.color};"
					>
						<div class="p-6">
							<h2 class="text-xl font-bold text-gray-900 mb-2">{board.name}</h2>
							{#if board.description}
								<p class="text-gray-600 mb-4">{board.description}</p>
							{/if}
							<div class="flex items-center gap-2 text-sm text-gray-500">
								<span
									class="inline-block w-4 h-4 rounded"
									style="background-color: {board.color};"
								></span>
								<span>Created {new Date(board.created_at).toLocaleDateString()}</span>
							</div>
						</div>
					</div>
				{/each}
			</div>
		{/if}
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

	<AddBoardModal
		open={showAddBoardModal}
		onClose={closeAddBoardModal}
		onAddSuccess={handleBoardCreated}
	/>
</div>
