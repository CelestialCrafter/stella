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
</script>

<div class="orbit">
	<label>
		Distance
		<input
			type="range"
			value="30"
			min="15"
			max="100"
			step="0.2"
			on:change={e => orbit.update(prev => [Number(e.target.value), prev[1]])}
		/>
	</label>

	<label>
		Speed
		<input
			type="range"
			value="1"
			min="0.2"
			max="10"
			step="0.2"
			on:change={e => orbit.update(prev => [prev[0], Number(e.target.value)])}
		/>
	</label>
</div>

<canvas
	on:mousedown={() => selectedPlanet.set(getSelectedPlanet()?.name || null)}
	bind:this={canvas}
></canvas>

<style lang="scss">
	@use '../colors.scss';
	.orbit {
		position: absolute;
		z-index: 10;
		right: 1rem;
		background-color: colors.$floating;
		padding: 1rem;
		border-radius: 1rem;
	}

	canvas {
		grid-row: 1/3;
		width: 100% !important;
		height: 100% !important;
		border-radius: 1rem;
	}
</style>
