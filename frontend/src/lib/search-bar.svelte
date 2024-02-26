<script context="module" lang="ts">
	export type SearchEvent = string;
</script>

<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	let searchquery = '';

	const dispatch = createEventDispatcher<{
		search: string;
	}>();

	function handleSubmit() {
		dispatch('search', searchquery);
	}
</script>

<form class="w-full block h-10" on:submit={handleSubmit}>
	<input
		class="flex-1 focus:bg-zinc-500 px-2 outline-green-600 w-full h-full bg-zinc-700 border border-zinc-500 rounded-[5px]"
		type="text	"
		bind:value={searchquery}
	/>
</form>
