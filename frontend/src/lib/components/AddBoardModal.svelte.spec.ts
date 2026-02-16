import { describe, it, expect, beforeEach, vi } from 'vitest';
import { render } from 'vitest-browser-svelte';
import { page } from '@vitest/browser/context';
import AddBoardModal from './AddBoardModal.svelte';

describe('AddBoardModal.svelte', () => {
	beforeEach(() => {
		vi.clearAllMocks();
	});

	describe('Rendering', () => {
		it('should not render modal when open prop is false', async () => {
			render(AddBoardModal, {
				open: false,
				onClose: () => {},
				onAddSuccess: () => {}
			});

			const heading = page.getByRole('heading', { name: 'Create New Board' });
			await expect.element(heading).not.toBeInTheDocument();
		});

		it('should render modal when open prop is true', async () => {
			render(AddBoardModal, {
				open: true,
				onClose: () => {},
				onAddSuccess: () => {}
			});

			const heading = page.getByRole('heading', { name: 'Create New Board' });
			await expect.element(heading).toBeInTheDocument();
		});

		it('should render all form fields', async () => {
			render(AddBoardModal, {
				open: true,
				onClose: () => {},
				onAddSuccess: () => {}
			});

			// Board name input
			const nameInput = page.getByLabelText('Board Name');
			await expect.element(nameInput).toBeInTheDocument();
			await expect.element(nameInput).toHaveAttribute('required');

			// Description textarea
			const descriptionInput = page.getByLabelText('Description');
			await expect.element(descriptionInput).toBeInTheDocument();

			// Color input
			const colorLabel = page.getByText('Color');
			await expect.element(colorLabel).toBeInTheDocument();
		});

		it('should render Create Board and Cancel buttons', async () => {
			render(AddBoardModal, {
				open: true,
				onClose: () => {},
				onAddSuccess: () => {}
			});

			const createButton = page.getByRole('button', { name: 'Create Board' });
			const cancelButton = page.getByRole('button', { name: 'Cancel' });

			await expect.element(createButton).toBeInTheDocument();
			await expect.element(cancelButton).toBeInTheDocument();
		});
	});

	describe('Input Behavior', () => {
		it('should have name input with required attribute', async () => {
			render(AddBoardModal, {
				open: true,
				onClose: () => {},
				onAddSuccess: () => {}
			});

			const nameInput = page.getByLabelText('Board Name');
			await expect.element(nameInput).toHaveAttribute('required');
		});

		it('should have placeholder text for board name', async () => {
			render(AddBoardModal, {
				open: true,
				onClose: () => {},
				onAddSuccess: () => {}
			});

			const nameInput = page.getByPlaceholder('e.g., Marketing Campaign');
			await expect.element(nameInput).toBeInTheDocument();
		});

		it('should have placeholder text for description', async () => {
			render(AddBoardModal, {
				open: true,
				onClose: () => {},
				onAddSuccess: () => {}
			});

			const descriptionInput = page.getByPlaceholder(
				'Describe the purpose of this board (optional)'
			);
			await expect.element(descriptionInput).toBeInTheDocument();
		});
	});

	describe('Button Interactions', () => {
		it('should call onClose when Cancel button is clicked', async () => {
			let closeCalled = false;
			const onCloseSpy = () => {
				closeCalled = true;
			};

			render(AddBoardModal, {
				open: true,
				onClose: onCloseSpy,
				onAddSuccess: () => {}
			});

			const cancelButton = page.getByRole('button', { name: 'Cancel' });
			await cancelButton.click();

			expect(closeCalled).toBe(true);
		});
	});
});
