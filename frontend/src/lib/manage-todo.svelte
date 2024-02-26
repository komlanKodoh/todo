<script context="module" lang="ts">
	export type CreateTodoEvent = {
		Title: string;
		Description: string;
	};

	export type DeleteTodoEvent = {
		Id: number;
	};

	export type UpdateTodoEvent = {
		Id: number;
		Title: string;
		Description: string;
		Completed?: boolean;
	};
</script>

<script lang="ts">
	import { createEventDispatcher, onMount } from 'svelte';
	import type { Todo } from './todo-service';

	export let todo: Partial<Todo> = {
		Id: 0,
		Title: '',
		CompletedAt: undefined
	};

	const dispatcher = createEventDispatcher<{
		'create-todo': CreateTodoEvent;
		'delete-todo': DeleteTodoEvent;
		'update-todo': UpdateTodoEvent;
	}>();

	function handleCreateTodo(e: Event) {
		e.preventDefault();
		dispatcher('create-todo', {
			Title: todo.Title ?? '',
			Description: todo.Description ?? ''
		});
	}

	function handleDeleteTodo(e: Event) {
		e.preventDefault();

		if (todo.Id === undefined) {
			return;
		}

		dispatcher('delete-todo', {
			Id: todo.Id
		});
	}

	function handleUpdateTodo(e: Event) {
		e.preventDefault();

		if (todo.Id === undefined || todo.Title === undefined || todo.Description === undefined) {
			return;
		}

		dispatcher('update-todo', {
			Id: todo.Id,
			Title: todo.Title ?? '',
			Description: todo.Description ?? '',
		});
	}
</script>

{#if todo.Id !== 0}
	<div class="h-full flex flex-col min-h-0">
		<h1 class="font-semibold">UpdateTodo New Todo</h1>
		<form class="flex flex-col min-h-0 h-full">
			<input
				class="list-none px-2 py-2 my-2 w-full bg-zinc-700 focus:bg-zinc-500 rounded-md outline-none outline-gre"
				type="text"
				name="title"
				bind:value={todo.Title}
			/>
			<textarea
				class="list-none px-2 py-2 my-2 h-full max-h-96 w-full bg-zinc-700 focus:bg-zinc-500 rounded-md outline-none outline-gre"
				id="description"
				name="description"
				bind:value={todo.Description}
			/>


			<div class="flex justify-end gap-2">
				<button class="px-2 py-1 bg-red-500 rounded-[4px]" type="submit" on:click={handleDeleteTodo}
					>Delete</button
				>
				<button
					class="px-2 py-1 bg-green-500 rounded-[4px]"
					type="submit"
					on:click={handleUpdateTodo}>Update</button
				>
			</div>
		</form>
	</div>
{:else}
	<div class="h-full flex flex-col min-h-0">
		<h1 class="font-semibold">Create New Todo</h1>
		<form class="flex flex-col min-h-0 h-full">
			<input
				class="list-none px-2 py-2 my-2 w-full bg-zinc-700 focus:bg-zinc-500 rounded-md outline-none outline-gre"
				type="text"
				name="title"
				bind:value={todo.Title}
			/>
			<textarea
				class="list-none px-2 py-2 my-2 h-full max-h-96  w-full bg-zinc-700 focus:bg-zinc-500 rounded-md outline-none outline-gre"
				id="description"
				name="description"
				bind:value={todo.Description}
			/>

			<div class="flex justify-end gap-2">
				<button class="px-2 py-1 bg-red-500 rounded-[4px]" type="submit">Cancel</button>
				<button
					class="px-2 py-1 bg-green-500 rounded-[4px]"
					type="submit"
					on:click={handleCreateTodo}>Create</button
				>
			</div>
		</form>
	</div>
{/if}
