<script lang="ts">
	import ManageTodo from '$lib/manage-todo.svelte';
	import SearchBar, { type SearchEvent } from '$lib/search-bar.svelte';
	import { useQuery, useMutation, useQueryClient } from '@sveltestack/svelte-query';
	import TodoBar from '$lib/todo-bar.svelte';
	import type { TodoSelectedEvent } from '$lib/todo-bar.svelte';
	import type { Todo, RequestTodoDto } from '$lib/todo-service';
	import { getTodos, createTodo, updateTodo, deleteTodo } from '$lib/todo-service';
	import { Handler } from '$lib/utils';

	let selectedTodo: Todo | undefined = undefined;

	const queryClient = useQueryClient();

	const todoFilter: RequestTodoDto = {};

	const todoQueryResult = useQuery(['get-todos', todoFilter] as const, Handler(getTodos));

	const handleSearchEvent = (e: CustomEvent<SearchEvent>) => {
		todoFilter.Title = e.detail;
		queryClient.invalidateQueries('get-todos');
	};

	const handleTodoSelected = (e: CustomEvent<TodoSelectedEvent>) => {
		if (!$todoQueryResult.data) {
			return;
		}

		selectedTodo = $todoQueryResult.data.find((todo) => todo.Id === e.detail.Id);
	};

	const createTodoMutation = useMutation(createTodo, {
		onSuccess: () => {
			queryClient.invalidateQueries('get-todos');
		}
	});

	const updateTodoMutation = useMutation(updateTodo, {
		onSuccess: () => {
			queryClient.invalidateQueries('get-todos');
		}
	});

	const deleteTodoMutation = useMutation(deleteTodo, {
		onSuccess: () => {
			queryClient.invalidateQueries('get-todos');
		}
	});

	import { onMount } from 'svelte';

	onMount(async () => {});

	const handleCreateButtonClick = async () => {
		selectedTodo = {
			Id: 0,
			Title: '',
			Description: '',
			CompletedAt: undefined
		} as Todo;
	};

	$: {
		console.log('parend', selectedTodo);
	}
</script>

<div class="flex pt-5 justify-center w-full h-full bg-zinc-800 text-zinc-200 font-sans px-5">
	<div class="max-w-4xl w-full h-full flex flex-col">
		<div class="flex gap-2">
			<div class="flex-1">
				<SearchBar on:search={handleSearchEvent} />
			</div>

			<button
				class="h-full px-4 bg-green-500 text-base rounded-[4px]"
				on:click={handleCreateButtonClick}
			>
				Create New</button
			>
		</div>

		<div class="flex mt-8 gap-5 flex-1 min-h-0">
			<div class="flex-[0.5] overflow-y-scroll scroll-bar">
				{#if $todoQueryResult.data}
					<TodoBar todos={$todoQueryResult.data} on:todo-selected={handleTodoSelected}></TodoBar>
				{/if}
			</div>

			<div class="flex-[0.5] py-2">
				<ManageTodo
					todo={selectedTodo}
					on:create-todo={(event) => {
						$createTodoMutation.mutate(event.detail);
					}}
					on:delete-todo={(event) => {
						$deleteTodoMutation.mutate(event.detail.Id);
					}}
					on:update-todo={(event) => {
						$updateTodoMutation.mutate(event.detail);
					}}
				/>
			</div>
		</div>
	</div>
</div>

<style>
</style>
