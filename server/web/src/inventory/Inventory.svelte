<script>
	import { onMount } from 'svelte';
	import { initScene } from './scene.js';
	import { planets, selectedPlanet } from '../stores.js';

	let canvas;
	let getSelectedPlanet = () => null;
	onMount(() => (getSelectedPlanet = initScene(canvas)));
</script>

<canvas on:mousedown={() => selectedPlanet.set(getSelectedPlanet()?.name)} bind:this={canvas}
></canvas>
