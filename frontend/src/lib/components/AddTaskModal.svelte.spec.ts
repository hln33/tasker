import { page } from 'vitest/browser';
import { describe, expect, it } from 'vitest';
import { render } from 'vitest-browser-svelte';
import AddTaskModal from './AddTaskModal.svelte';

describe('addTaskModal.svelte', () => {
	it('should render modal when open prop is true', async () => {
		render(AddTaskModal, { open: true, onClose: () => {} });

		// Check for the modal heading
		const heading = page.getByRole('heading', { name: 'Create New Task' });
		await expect.element(heading).toBeInTheDocument();

		// Check for the title input field
		const titleInput = page.getByLabelText('Title');
		await expect.element(titleInput).toBeInTheDocument();

		// Check for the Create Task button
		const submitButton = page.getByRole('button', { name: 'Create Task' });
		await expect.element(submitButton).toBeInTheDocument();
	});

	it('should not render modal when open prop is false', async () => {
		render(AddTaskModal, { open: false, onClose: () => {} });

		// Check that the modal heading is not present
		const heading = page.getByRole('heading', { name: 'Create New Task' });
		await expect.element(heading).not.toBeInTheDocument();
	});
});
