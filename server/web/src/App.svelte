<script>
	import { onMount } from 'svelte';
	import { initScene, addPlanets, selectedPlanet } from './scene.js';

	let canvas;
	onMount(async () => {
		const userId = 'google-100735534519069903161';
		const planets = await (await fetch(`/api/planets/${userId}`)).json();

		addPlanets(planets.map(planet => planet.hash));
		initScene(canvas);
	});

	const handleCanvasClick = () => {
		const selected = selectedPlanet();
		selected != null && console.log(selected.name);
	};
</script>

<main>
	<canvas on:click={handleCanvasClick} bind:this={canvas}></canvas>
</main>

<style>
</style>
