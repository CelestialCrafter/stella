<script>
	import { onDestroy, onMount } from 'svelte';
	import { initScene } from './scene.js';
	import { orbit, selectedPlanet } from '../stores.js';

	let canvas;
	let getSelectedPlanet = () => null;
	let cleanup = () => {};

	onMount(() => {
		[cleanup, getSelectedPlanet] = initScene(canvas);
	});

	onDestroy(cleanup);

	let orbitSpeed = 1;
	let orbitDistance = 30;

	$: orbit.set([orbitDistance, orbitSpeed]);
</script>

<div class="orbit">
	<label>
		Distance
		<input type="range" min="30" max="100" step="0.2" bind:value={orbitDistance} />
	</label>

	<label>
		Speed
		<input type="range" min="0" max="10" step="0.2" bind:value={orbitSpeed} />
	</label>
</div>

<canvas
	on:mousedown={() => selectedPlanet.set(getSelectedPlanet()?.name || null)}
	bind:this={canvas}
></canvas>

<style lang="scss">
	@use '../styles/colors.scss';
	@use '../styles/spacing.scss';

	.orbit {
		position: absolute;
		z-index: 10;
		right: spacing.$padding;
		background-color: colors.$floating;
		padding: spacing.$padding;
		border-radius: spacing.$radius;
	}

	canvas {
		grid-row: 1/3;
		width: 100% !important;
		height: 100% !important;
		border-radius: spacing.$radius;
	}
</style>
