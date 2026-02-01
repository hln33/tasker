import { page } from 'vitest/browser';
import { describe, expect, it } from 'vitest';
import { render } from 'vitest-browser-svelte';
import TaskCard from './TaskCard.svelte';
import type { Task } from '$lib/types';

describe('TaskCard.svelte', () => {
	const mockTask: Task = {
		id: 'TASK-001',
		title: 'Test Task Title',
		description: 'This is a test task description',
		status: 'TODO',
		priority: 'High'
	};

	const onDelete = () => {};

	it('should render task ID', async () => {
		render(TaskCard, { task: mockTask, onDelete });

		const taskId = page.getByText('TASK-001');
		await expect.element(taskId).toBeInTheDocument();
	});

	it('should render task title', async () => {
		render(TaskCard, { task: mockTask, onDelete });

		const title = page.getByRole('heading', { name: 'Test Task Title', level: 2 });
		await expect.element(title).toBeInTheDocument();
	});

	it('should render task description', async () => {
		render(TaskCard, { task: mockTask, onDelete });

		const descriptionLabel = page.getByRole('heading', { name: 'Description' });
		await expect.element(descriptionLabel).toBeInTheDocument();

		const descriptionText = page.getByText('This is a test task description');
		await expect.element(descriptionText).toBeInTheDocument();
	});

	it('should render priority badge', async () => {
		render(TaskCard, { task: mockTask, onDelete });

		const priorityBadge = page.getByText('High');
		await expect.element(priorityBadge).toBeInTheDocument();
	});

	it('should render status badge', async () => {
		render(TaskCard, { task: mockTask, onDelete });

		const statusLabel = page.getByText('Status:');
		await expect.element(statusLabel).toBeInTheDocument();

		const statusBadge = page.getByText('TODO');
		await expect.element(statusBadge).toBeInTheDocument();
	});

	it('should render delete button', async () => {
		render(TaskCard, { task: mockTask, onDelete });

		const deleteButton = page.getByRole('button', { name: 'Delete task' });
		await expect.element(deleteButton).toBeInTheDocument();
	});

	it('should call onDelete when delete button is clicked', async () => {
		let deletedTask: Task | null = null;
		const onDeleteSpy = (task: Task) => {
			deletedTask = task;
		};

		render(TaskCard, { task: mockTask, onDelete: onDeleteSpy });

		const deleteButton = page.getByRole('button', { name: 'Delete task' });
		await deleteButton.click();

		expect(deletedTask).toEqual(mockTask);
	});

	it('should apply correct priority color classes for High priority', async () => {
		render(TaskCard, { task: mockTask, onDelete });

		const priorityBadge = page.getByText('High');
		await expect.element(priorityBadge).toHaveClass(/bg-red-100/);
		await expect.element(priorityBadge).toHaveClass(/text-red-800/);
	});

	it('should apply correct priority color classes for Medium priority', async () => {
		const mediumPriorityTask: Task = { ...mockTask, priority: 'Medium' };
		render(TaskCard, { task: mediumPriorityTask, onDelete });

		const priorityBadge = page.getByText('Medium');
		await expect.element(priorityBadge).toHaveClass(/bg-yellow-300/);
		await expect.element(priorityBadge).toHaveClass(/text-yellow-800/);
	});

	it('should apply correct priority color classes for Low priority', async () => {
		const lowPriorityTask: Task = { ...mockTask, priority: 'Low' };
		render(TaskCard, { task: lowPriorityTask, onDelete });

		const priorityBadge = page.getByText('Low');
		await expect.element(priorityBadge).toHaveClass(/bg-green-100/);
		await expect.element(priorityBadge).toHaveClass(/text-green-800/);
	});

	it('should apply correct status color classes for TODO status', async () => {
		render(TaskCard, { task: mockTask, onDelete });

		const statusBadge = page.getByText('TODO');
		await expect.element(statusBadge).toHaveClass(/bg-gray-100/);
		await expect.element(statusBadge).toHaveClass(/text-gray-800/);
	});

	it('should apply correct status color classes for In Progress status', async () => {
		const inProgressTask: Task = { ...mockTask, status: 'In Progress' };
		render(TaskCard, { task: inProgressTask, onDelete });

		const statusBadge = page.getByText('In Progress');
		await expect.element(statusBadge).toHaveClass(/bg-blue-100/);
		await expect.element(statusBadge).toHaveClass(/text-blue-800/);
	});

	it('should apply correct status color classes for Completed status', async () => {
		const completedTask: Task = { ...mockTask, status: 'Completed' };
		render(TaskCard, { task: completedTask, onDelete });

		const statusBadge = page.getByText('Completed');
		await expect.element(statusBadge).toHaveClass(/bg-green-100/);
		await expect.element(statusBadge).toHaveClass(/text-green-800/);
	});
});
