<script>
	import PlanetInfo from './PlanetInfo.svelte';
	import PlanetScene from './PlanetScene.svelte';
	import PlanetList from './PlanetList.svelte';
</script>

<main>
	<PlanetList />
	<PlanetScene />
	<PlanetInfo />
</main>

<style lang="scss">
	@use '../colors.scss';
	main {
		background-color: colors.$background;
		display: grid;
		padding: 1rem;
		gap: 1rem;
		grid-template-columns: 25rem 3fr;
		grid-template-rows: 2fr 1fr;
	}
</style>
