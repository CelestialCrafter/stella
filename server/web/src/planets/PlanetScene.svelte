<script>
	import { onDestroy, onMount } from 'svelte';
	import { initScene } from './scene.js';
	import { planets, selectedPlanet } from '../stores.js';

	let canvas;
	let getSelectedPlanet = () => null;
	let cleanup = () => {};

	onMount(() => {
		const fns = initScene(canvas);
		cleanup = fns[0];
		getSelectedPlanet = fns[1];
	});

	onDestroy(cleanup);
</script>

<canvas
	on:mousedown={() => selectedPlanet.set($planets[getSelectedPlanet()?.name] || null)}
	bind:this={canvas}
></canvas>

<style>
	canvas {
		grid-row: 1/3;
		width: 100% !important;
		height: 100% !important;
	}
</style>
