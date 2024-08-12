<script>
	import { onMount } from 'svelte';
	import { initScene, addPlanets } from './scene.js';
	import { selectedPlanet } from './scene.js';

	export let planets;
	export let selected;

	let canvas;
	onMount(() => {
		addPlanets(Object.values(planets).map(planet => planet.hash));
		initScene(canvas);
	});
</script>

<canvas on:mousedown={() => (selected = selectedPlanet()?.name)} bind:this={canvas}></canvas>
