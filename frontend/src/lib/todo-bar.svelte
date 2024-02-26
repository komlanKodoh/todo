<script context="module" lang="ts">
	export type TodoSelectedEvent = {
		Id: number;
	};
</script>

<script lang="ts">
	import type { Todo } from './todo-service';
	import { createEventDispatcher } from 'svelte';

	export let todos: Todo[];

	const dispatcher = createEventDispatcher<{
		'todo-selected': TodoSelectedEvent;
	}>();

	function handleTodoSelected(todo: Todo) {
		dispatcher('todo-selected', { Id: todo.Id });
	}
</script>

{#if todos.length === 0}
	<p>No todos found</p>
{:else}
	{#each todos as todo (todo.Id)}
		<li class="list-none relative px-4 py-2 my-2 bg-zinc-700 rounded-md">
			<button class="absolute w-full h-full top-0 left-0 bottom-0 right-0" on:click={() => handleTodoSelected(todo)}> </button>
			<span class="block font-semibold">
				{todo.Title}
			</span>
			<div class="separator bg-zinc-500"></div>
			<span class="block text-sm">
				{todo.Description}
			</span>
		</li>
	{/each}
{/if}
