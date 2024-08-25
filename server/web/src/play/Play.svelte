<script>
	import { writable } from 'svelte/store';
	import { selectedPlanet } from '../stores';
	import { onDestroy, tick } from 'svelte';
	import { initScene, updateScene } from './scene';

	const state = writable({ type: '', game: 0 });
	let canvas;

	const ws = new WebSocket(
		`ws://${location.host}/api/planet/play?planet=${$selectedPlanet?.hash}&token=${localStorage.getItem('token')}`
	);

	// @TODO add notifications for connect and disconnect
	ws.onmessage = e => {
		const data = JSON.parse(e.data);
		state.set(data);
	};
	document.onkeydown = e => ws.send(e.code);

	let cleanup = () => {};
	let scene = null;
	let font = null;

	const unsubscribe = state.subscribe(newState => {
		switch (newState.type) {
			case 'init':
				tick().then(async () => ([cleanup, scene, font] = await initScene(canvas, newState)));
				break;
			case 'state':
				if (scene && font) updateScene(scene, font, newState);
				break;
			default:
		}
	});

	onDestroy(() => {
		unsubscribe();
		cleanup();
	});
</script>

<canvas bind:this={canvas}></canvas>

<style lang="scss">
	@use '../styles/colors.scss';
	@use '../styles/spacing.scss';

	canvas {
		width: 100% !important;
		height: 99% !important;
	}
</style>
