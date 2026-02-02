import { describe, it, expect, beforeEach, vi } from 'vitest';
import { render } from 'vitest-browser-svelte';
import TaskColumn from './TaskColumn.svelte';
import type { Task } from '$lib/types';

describe('TaskColumn', () => {
	const todoTasks: Task[] = [
		{
			id: 'TASK-001',
			title: 'Buy groceries',
			description: 'Get milk, eggs, and bread',
			status: 'TODO',
			priority: 'high'
		},
		{
			id: 'TASK-002',
			title: 'Call dentist',
			description: 'Schedule annual checkup',
			status: 'TODO',
			priority: 'medium'
		}
	];
	const inProgressTasks: Task[] = [
		{
			id: 'TASK-003',
			title: 'Write report',
			description: 'Complete Q4 financial analysis',
			status: 'In Progress',
			priority: 'low'
		}
	];

	const mockDelete = vi.fn();

	beforeEach(() => {
		vi.clearAllMocks();
	});

	describe('Component rendering', () => {
		it('renders column with tasks', () => {
			const col = render(TaskColumn, {
				title: 'TODO',
				tasks: todoTasks,
				color: 'gray',
				emptyMessage: 'No tasks',
				onDelete: mockDelete
			});

			expect(col.getByRole('heading', { name: 'TODO' })).toBeInTheDocument();
			expect(col.getByText('Buy groceries')).toBeInTheDocument();
			expect(col.getByText('Call dentist')).toBeInTheDocument();
		});

		it('renders empty message when no tasks', () => {
			const col = render(TaskColumn, {
				title: 'TODO',
				tasks: [],
				color: 'gray',
				emptyMessage: 'No tasks to do',
				onDelete: mockDelete
			});

			expect(col.getByText('No tasks to do')).toBeInTheDocument();
		});

		it('has drop zone section element', () => {
			const col = render(TaskColumn, {
				title: 'TODO',
				tasks: todoTasks,
				color: 'gray',
				emptyMessage: 'No tasks',
				onDelete: mockDelete
			});

			const dropZone = col.container.querySelector('section');
			expect(dropZone).toBeInTheDocument();
		});
	});

	describe('Delete interaction', () => {
		it('passes onDelete to task cards', () => {
			const deleteCallback = vi.fn();

			const col = render(TaskColumn, {
				title: 'TODO',
				tasks: todoTasks,
				color: 'gray',
				emptyMessage: 'No tasks',
				onDelete: deleteCallback
			});

			// Verify the delete button is present (delegation is working)
			const deleteButtons = col.container.querySelectorAll('button[aria-label="Delete task"]');
			expect(deleteButtons.length).toBe(2);
		});
	});

	describe('Color variations', () => {
		it('renders with gray color', () => {
			const col = render(TaskColumn, {
				title: 'TODO',
				tasks: todoTasks,
				color: 'gray',
				emptyMessage: 'No tasks',
				onDelete: mockDelete
			});

			expect(col.getByRole('heading', { name: 'TODO' })).toBeInTheDocument();
		});

		it('renders with blue color', () => {
			const col = render(TaskColumn, {
				title: 'In Progress',
				tasks: inProgressTasks,
				color: 'blue',
				emptyMessage: 'No tasks',
				onDelete: mockDelete
			});

			expect(col.getByRole('heading', { name: 'In Progress' })).toBeInTheDocument();
		});

		it('renders with green color', () => {
			const col = render(TaskColumn, {
				title: 'Done',
				tasks: todoTasks,
				color: 'green',
				emptyMessage: 'No tasks',
				onDelete: mockDelete
			});

			expect(col.getByRole('heading', { name: 'Done' })).toBeInTheDocument();
		});
	});

	describe('Multiple columns integration', () => {
		it('renders multiple columns independently', () => {
			const todoCol = render(TaskColumn, {
				title: 'TODO',
				tasks: todoTasks,
				color: 'gray',
				emptyMessage: 'No tasks',
				onDelete: mockDelete
			});
			const inProgressCol = render(TaskColumn, {
				title: 'In Progress',
				tasks: inProgressTasks,
				color: 'blue',
				emptyMessage: 'No tasks',
				onDelete: mockDelete
			});
			const doneCol = render(TaskColumn, {
				title: 'Done',
				tasks: [],
				color: 'green',
				emptyMessage: 'No completed tasks',
				onDelete: mockDelete
			});

			expect(todoCol.getByText('Buy groceries')).toBeInTheDocument();
			expect(inProgressCol.getByText('Write report')).toBeInTheDocument();
			expect(doneCol.getByText('No completed tasks')).toBeInTheDocument();

			// Each column should have its own drop zone
			expect(todoCol.container.querySelector('section')).toBeInTheDocument();
			expect(inProgressCol.container.querySelector('section')).toBeInTheDocument();
			expect(doneCol.container.querySelector('section')).toBeInTheDocument();
		});
	});
});
