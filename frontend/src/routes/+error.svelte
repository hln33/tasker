<script lang="ts">
	import { page } from '$app/state';
	import { ERROR_MESSAGES } from '$lib/constants/errorMessages';

	const errorMessage = ERROR_MESSAGES[page.status as keyof typeof ERROR_MESSAGES] || ERROR_MESSAGES.DEFAULT;
</script>

<div class="flex min-h-screen items-center justify-center bg-gray-200 px-4">
	<div class="max-w-md rounded-lg bg-white p-8 text-center shadow-lg">
		<h1 class="mb-2 text-2xl font-bold text-gray-900">{errorMessage.title}</h1>
		<p class="mb-6 text-gray-600">{errorMessage.subtitle}</p>
		<div class="mb-6 rounded-lg border border-red-200 bg-red-50 p-4">
			<p class="text-sm font-medium text-red-800">Error {page.status}</p>
			<p class="mt-1 text-sm text-red-600">{page.error?.message}</p>
		</div>
		<button
			onclick={() => window.location.reload()}
			class="cursor-pointer rounded-lg bg-blue-600 px-6 py-2 font-medium text-white transition-colors hover:bg-blue-700"
		>
			Try Again
		</button>
	</div>
</div>
