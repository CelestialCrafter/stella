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
	@use '../styles/spacing.scss';

	main {
		background-color: colors.$background;
		display: grid;
		padding: spacing.$padding;
		gap: spacing.$padding;
		grid-template-columns: 25rem 3fr;
		grid-template-rows: 2fr;
	}
</style>
