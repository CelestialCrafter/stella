<script>
	import { onMount } from 'svelte';
	import { selectedPlanet } from './scene.js';

	import Scene from './Scene.svelte';
	import SelectedPlanet from './SelectedPlanet.svelte';

	$: planets = {};
	$: selected = null;
	onMount(async () => {
		const userId = 'google-100735534519069903161';
		const planetsArray = await (await fetch(`/api/planets/${userId}`)).json();
		planets = planetsArray.reduce((acc, x) => ({ ...acc, [x.hash]: x }), {});
	});
</script>

<main>
	{#if Object.keys(planets).length < 1}
		<span>
			Loading...<br />
			(or no planets)
		</span>
	{:else}
		<SelectedPlanet planet={selected ? planets[selected] : null} />
		<Scene
			handleCanvasClick={() => (selected = selectedPlanet())}
			planets={Object.values(planets)}
		/>
	{/if}
</main>

<style>
</style>
