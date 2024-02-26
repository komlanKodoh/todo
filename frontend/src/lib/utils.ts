import type { QueryFunctionContext } from '@sveltestack/svelte-query';
import type { Task } from 'vitest';

export class Url {
	public static api(url: string): string {
		return `${import.meta.env.VITE_BACKEND_URL}${url}`;
	}
}

export function Handler<TArg extends any[], TReturn>(
	handler: (...args: TArg) => Promise<TReturn>
): (ctx: QueryFunctionContext<TArg, TReturn>) => Promise<TReturn> {
	return async (context) => {
		return handler(...(context.queryKey as unknown as TArg));
	};
}
