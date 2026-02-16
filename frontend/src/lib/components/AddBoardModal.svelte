<script lang="ts">
	import { enhance } from '$app/forms';
	import { page } from '$app/state';

	let {
		open = false,
		onClose,
		onAddSuccess
	}: {
		open: boolean;
		onClose: () => void;
		onAddSuccess: () => void;
	} = $props();

	function handleClose() {
		onClose();
	}
</script>

{#if open}
	<div class="fixed inset-0 z-50 flex items-center justify-center bg-black/80 p-4">
		<div class="w-full max-w-md rounded-lg bg-white shadow-xl">
			<div class="border-b border-gray-200 p-6">
				<h2 class="text-2xl font-bold text-gray-900">Create New Board</h2>
			</div>

			<form
				method="POST"
				action="?/createBoard"
				use:enhance={() => {
					return async ({ result, update }) => {
						if (result.type === 'success') {
							onAddSuccess();
							onClose();
						}
						await update();
					};
				}}
				class="space-y-4 p-6"
			>
				{#if page.form?.error}
					<div class="rounded-lg border border-red-200 bg-red-50 p-4">
						<p class="text-sm text-red-800">{page.form.error}</p>
					</div>
				{/if}

				<div>
					<label for="name" class="mb-1 block text-sm font-medium text-gray-700">
						Board Name <span class="text-red-500">*</span>
					</label>
					<input
						id="name"
						name="name"
						type="text"
						value={page.form?.name ?? ''}
						class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500"
						placeholder="e.g., Marketing Campaign"
						required
					/>
				</div>

				<div>
					<label for="description" class="mb-1 block text-sm font-medium text-gray-700">
						Description
					</label>
					<textarea
						id="description"
						name="description"
						rows="3"
						class="w-full rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500"
						placeholder="Describe the purpose of this board (optional)"
					>
						{page.form?.description ?? ''}
					</textarea>
				</div>

				<div>
					<label for="color" class="mb-1 block text-sm font-medium text-gray-700">
						Color
					</label>
					<div class="flex items-center gap-3">
						<input
							id="color"
							name="color"
							type="color"
							value={page.form?.color ?? '#3B82F6'}
							class="h-10 w-16 rounded border border-gray-300 cursor-pointer"
						/>
						<input
							type="text"
							value={page.form?.color ?? '#3B82F6'}
							class="w-24 rounded-lg border border-gray-300 px-3 py-2 font-mono text-sm"
							pattern="^#[0-9A-Fa-f]{6}$"
							placeholder="#3B82F6"
							readonly
						/>
						<span class="text-sm text-gray-500">Pick a board color</span>
					</div>
				</div>

				<div class="flex gap-3 pt-4">
					<button
						type="submit"
						class="flex-1 cursor-pointer rounded-lg bg-blue-600 px-4 py-2 font-medium text-white hover:bg-blue-700"
					>
						Create Board
					</button>
					<button
						type="button"
						onclick={handleClose}
						class="cursor-pointer rounded-lg border border-gray-300 px-4 py-2 font-medium text-gray-700 hover:bg-gray-50"
					>
						Cancel
					</button>
				</div>
			</form>
		</div>
	</div>
{/if}
