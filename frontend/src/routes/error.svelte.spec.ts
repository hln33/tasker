import { page } from 'vitest/browser';
import { describe, expect, it, vi } from 'vitest';
import { render } from 'vitest-browser-svelte';
import ErrorPage from './+error.svelte';
import { ERROR_MESSAGES } from '$lib/constants/errorMessages';

// Mock the $app/state module
vi.mock('$app/state', () => {
	let mockStatus = 500;
	let mockError: { message: string } | null = { message: 'Internal Server Error' };

	return {
		page: {
			get status() {
				return mockStatus;
			},
			set status(value: number) {
				mockStatus = value;
			},
			get error() {
				return mockError;
			},
			set error(value: { message: string } | null) {
				mockError = value;
			}
		},
		setPage: (status: number, error: { message: string } | null) => {
			mockStatus = status;
			mockError = error;
		}
	};
});

async function setupMockPage(status: number, message: string) {
	const { setPage } = (await import('$app/state')) as unknown as {
		setPage: (status: number, error: { message: string }) => void;
	};
	setPage(status, { message });
}

describe('/+error.svelte', () => {
	it('should render 500 error with message', async () => {
		await setupMockPage(500, 'Internal Server Error');
		render(ErrorPage);

		const heading = page.getByRole('heading', {
			name: ERROR_MESSAGES[500].title
		});
		await expect.element(heading).toBeInTheDocument();

		const errorMessage = page.getByText('Error 500');
		await expect.element(errorMessage).toBeInTheDocument();

		const errorDetail = page.getByText('Internal Server Error');
		await expect.element(errorDetail).toBeInTheDocument();
	});

	it('should render 502 error with message', async () => {
		await setupMockPage(502, 'Bad Gateway');
		render(ErrorPage);

		const heading = page.getByRole('heading', {
			name: ERROR_MESSAGES[502].title
		});
		await expect.element(heading).toBeInTheDocument();

		const errorMessage = page.getByText('Error 502');
		await expect.element(errorMessage).toBeInTheDocument();
	});

	it('should render 503 error with message', async () => {
		await setupMockPage(503, 'Service Unavailable');
		render(ErrorPage);

		const heading = page.getByRole('heading', {
			name: ERROR_MESSAGES[503].title
		});
		await expect.element(heading).toBeInTheDocument();

		const errorMessage = page.getByText('Error 503');
		await expect.element(errorMessage).toBeInTheDocument();
	});

	it('should render 504 error with message', async () => {
		await setupMockPage(504, 'Gateway Timeout');
		render(ErrorPage);

		const heading = page.getByRole('heading', {
			name: ERROR_MESSAGES[504].title
		});
		await expect.element(heading).toBeInTheDocument();

		const errorMessage = page.getByText('Error 504');
		await expect.element(errorMessage).toBeInTheDocument();
	});

	it('should render default message for unknown status code', async () => {
		await setupMockPage(418, "I'm a teapot");
		render(ErrorPage);

		const heading = page.getByRole('heading', {
			name: ERROR_MESSAGES.DEFAULT.title
		});
		await expect.element(heading).toBeInTheDocument();

		const errorMessage = page.getByText("I'm a teapot");
		await expect.element(errorMessage).toBeInTheDocument();
	});

	it('should render Try Again button', async () => {
		await setupMockPage(500, 'Internal Server Error');
		render(ErrorPage);

		const button = page.getByRole('button', { name: 'Try Again' });
		await expect.element(button).toBeInTheDocument();
	});
});
