export const ERROR_MESSAGES = {
	500: {
		title: 'Oops! We dropped the ball!',
		subtitle:
			"Something went wrong on our end. We've made a note to fix this—thanks for your patience!"
	},
	502: {
		title: 'Our server is taking a quick break!',
		subtitle: "It'll be back in a jiffy. Feel free to grab a snack while you wait!"
	},
	503: {
		title: 'So many people, so many tasks!',
		subtitle: "We're getting a lot of love right now. Hang tight and we'll be with you shortly!"
	},
	504: {
		title: 'This is taking longer than a Monday morning...',
		subtitle: "The request timed out. Give it another go—third time's the charm, right?"
	},
	DEFAULT: {
		title: 'Well, that was unexpected!',
		subtitle: "Something went sideways. Don't worry, it's not you—it's us. We're on it!"
	}
} as const;

export type ErrorStatus = keyof typeof ERROR_MESSAGES | (number & {});
