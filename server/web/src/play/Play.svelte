<script>
	import { writable } from 'svelte/store';
	import { selectedPlanet } from '../stores';
	import { onDestroy, onMount, tick } from 'svelte';
	import { initScene, updateScene } from './scene';

	const state = writable({ type: '', game: 0 });
	let canvas;

	const newSocket = () =>
		new WebSocket(
			`ws://${location.host}/api/planet/play?planet=${$selectedPlanet?.hash}&token=${localStorage.getItem('token')}`
		);
	let ws = newSocket();

	ws.onmessage = e => {
		const data = JSON.parse(e.data);
		state.set(data);
	};
	ws.onopen = () => alert('connected');
	ws.onclose = () => alert('disconnected');
	document.onkeydown = e => ws.send(e.code);

	let data = null;
	const unsubscribe = state.subscribe(newState => {
		switch (newState.type) {
			case 'init':
				tick().then(async () => (data = await initScene(canvas, newState)));
			case 'state':
				if (data) updateScene(...data, newState);
		}
	});

	onDestroy(unsubscribe);
</script>

<canvas bind:this={canvas}></canvas>
