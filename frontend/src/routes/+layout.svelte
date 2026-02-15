<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import AppSidebar from '$lib/components/AppSidebar.svelte';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<Sidebar.Provider>
	<div class="flex min-h-screen">
		<AppSidebar />
		<main class="flex-1 bg-gray-200 p-6">
			{@render children()}
		</main>
	</div>
</Sidebar.Provider>
