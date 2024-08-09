<script>
	import { onMount } from 'svelte';
	import { initScene, addPlanets, selectedPlanet } from './scene.js';

	let canvas;
	onMount(() => initScene(canvas));
	addPlanets(['312a6ec69acf2c905f40c839af63873f7d2d3a3972f578b30d47f123e0b94932', '7d27a746d6abef020f1f33a9b3115399f45d8f145e6e4e1c1e95cab558d69ad0', '84203a979ef5729707d8f2fb2bc3e06fdcc76a7c1c04e4fa54d31977b0077c27'])

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
