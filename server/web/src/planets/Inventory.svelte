<script>
	import Info from './Info.svelte';
	import Scene from './Scene.svelte';
	import List from './List.svelte';
</script>

<main>
	<List />
	<Scene />
	<Info />
</main>

<style lang="scss">
	@use '../styles/colors.scss';
	main {
		background-color: colors.$background;
		display: grid;
		padding: 1rem;
		gap: 1rem;
		grid-template-columns: 25rem 3fr;
		grid-template-rows: 2fr 1fr;
	}
</style>
