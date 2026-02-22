<script lang="ts">
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import { page } from '$app/state';
	import AllBoards from '$lib/icons/AllBoards.svelte';
	import AllTasks from '$lib/icons/AllTasks.svelte';

	const currentPath = $derived(page.url.pathname);

	const navItems = [
		{ title: 'All Tasks', url: '/', icon: AllTasks },
		{ title: 'Boards', url: '/boards', icon: AllBoards }
		// { title: 'Calendar', url: '/calendar', icon: Calendar },
		// { title: 'Projects', url: '/projects', icon: Folder },
		// { title: 'Settings', url: '/settings', icon: Settings },
	];
</script>

<Sidebar.Root side="left" variant="sidebar" collapsible="none">
	<Sidebar.Content>
		<Sidebar.Group>
			<Sidebar.GroupLabel>Navigation</Sidebar.GroupLabel>
			<Sidebar.GroupContent>
				<Sidebar.Menu>
					{#each navItems as item (item.title)}
						{@const Icon = item.icon}
						<Sidebar.MenuItem>
							<Sidebar.MenuButton isActive={currentPath === item.url} tooltipContent={item.title}>
								{#snippet child({ props }: { props: Record<string, unknown> })}
									<a href={item.url} {...props}>
										<Icon class="size-4" />
										<span>{item.title}</span>
									</a>
								{/snippet}
							</Sidebar.MenuButton>
						</Sidebar.MenuItem>
					{/each}
				</Sidebar.Menu>
			</Sidebar.GroupContent>
		</Sidebar.Group>
	</Sidebar.Content>
	<Sidebar.Rail />
</Sidebar.Root>
