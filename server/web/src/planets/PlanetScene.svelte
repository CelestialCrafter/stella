<script>
	import { onDestroy, onMount } from 'svelte';
	import { initScene } from './scene.js';
	import { selectedPlanet } from '../stores.js';

	let canvas;
	let getSelectedPlanet = () => null;
	let cleanup = () => {};

	onMount(() => {
		[cleanup, getSelectedPlanet] = initScene(canvas);
	});

	onDestroy(cleanup);
</script>

<canvas
	on:mousedown={() => selectedPlanet.set(getSelectedPlanet()?.name || null)}
	bind:this={canvas}
></canvas>

<style>
	canvas {
		grid-row: 1/3;
		width: 100% !important;
		height: 100% !important;
	}
</style>
