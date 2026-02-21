<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import AppSidebar from '$lib/components/AppSidebar.svelte';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<Sidebar.Provider>
	<AppSidebar />
	<main class="min-h-screen w-full bg-gray-200 p-6">
		{@render children()}
	</main>
</Sidebar.Provider>
