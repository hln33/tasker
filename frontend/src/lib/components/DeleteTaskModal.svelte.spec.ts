import { page } from 'vitest/browser';
import { describe, expect, it } from 'vitest';
import { render } from 'vitest-browser-svelte';
import DeleteTaskModal from './DeleteTaskModal.svelte';
import type { Task } from '$lib/types';

describe('DeleteTaskModal.svelte', () => {
	const mockTask: Task = {
		id: 'TASK-001',
		title: 'Test Task to Delete',
		description: 'This is a test task',
		status: 'TODO',
		priority: 'Medium'
	};

	it('should render modal when open prop is true', async () => {
		render(DeleteTaskModal, { open: true, task: mockTask, onClose: () => {} });

		// Check for the modal heading
		const heading = page.getByRole('heading', { name: 'Delete Task' });
		await expect.element(heading).toBeInTheDocument();

		// Check for the task title in the confirmation message
		const taskTitle = page.getByText(/"Test Task to Delete"/);
		await expect.element(taskTitle).toBeInTheDocument();

		// Check for the Delete button
		const deleteButton = page.getByRole('button', { name: 'Delete' });
		await expect.element(deleteButton).toBeInTheDocument();

		// Check for the Cancel button
		const cancelButton = page.getByRole('button', { name: 'Cancel' });
		await expect.element(cancelButton).toBeInTheDocument();
	});

	it('should not render modal when open prop is false', async () => {
		render(DeleteTaskModal, { open: false, task: mockTask, onClose: () => {} });

		// Check that the modal heading is not present
		const heading = page.getByRole('heading', { name: 'Delete Task' });
		await expect.element(heading).not.toBeInTheDocument();
	});

	it('should display task title correctly', async () => {
		render(DeleteTaskModal, { open: true, task: mockTask, onClose: () => {} });

		// Check for the task title in the confirmation message
		const taskTitle = page.getByText('Test Task to Delete', { exact: false });
		await expect.element(taskTitle).toBeInTheDocument();
	});

	it('should show confirmation message', async () => {
		render(DeleteTaskModal, { open: true, task: mockTask, onClose: () => {} });

		// Check for the confirmation message text
		const confirmationMessage = page.getByText(/Are you sure you want to delete the task/);
		await expect.element(confirmationMessage).toBeInTheDocument();

		// Check for the "cannot be undone" warning
		const warningText = page.getByText(/This action cannot be undone/);
		await expect.element(warningText).toBeInTheDocument();
	});

	it('should handle null task gracefully', async () => {
		render(DeleteTaskModal, { open: true, task: null, onClose: () => {} });

		// Modal should still render even with null task
		const heading = page.getByRole('heading', { name: 'Delete Task' });
		await expect.element(heading).toBeInTheDocument();
	});
});
