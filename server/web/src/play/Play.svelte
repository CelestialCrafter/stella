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

	let data = null;
	let cleanup = () => {};
	const unsubscribe = state.subscribe(newState => {
		switch (newState.type) {
			case 'init':
				tick().then(async () => ([cleanup, data] = await initScene(canvas, newState)));
				break;
			case 'state':
				if (data) updateScene(...data, newState);
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
