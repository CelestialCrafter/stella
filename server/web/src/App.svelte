<script>
	import { onMount } from 'svelte';
	import { initScene, addPlanets } from './scene.js';

	let canvas;
	onMount(() => initScene(canvas));
	addPlanets(['/models/b0712d734d150f60773259bc0bd9321fe884bb6302c608ea350e99fc4b3aaa43.glb'])
</script>

<main>
	<canvas bind:this={canvas}></canvas> 
</main>

<style>
</style>
