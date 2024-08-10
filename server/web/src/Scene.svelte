<script>
	import { onMount } from 'svelte';
	import { initScene, addPlanets } from './scene.js';
	export let handleCanvasClick;
	export let planets;

	let canvas;
	onMount(() => {
		addPlanets(Object.values(planets).map(planet => planet.hash));
		initScene(canvas);
	});
</script>

<canvas on:mousedown={handleCanvasClick} bind:this={canvas}></canvas>
