import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
	return twMerge(clsx(inputs));
}

export type WithElementRef<T extends keyof HTMLElementTagNameMap = keyof HTMLElementTagNameMap> = T extends "button"
	? HTMLButtonElement
	: T extends "a"
		? HTMLAnchorElement
		: HTMLElement;

export type WithoutChildrenOrChild<T> = Omit<T, "children" | "child">;
