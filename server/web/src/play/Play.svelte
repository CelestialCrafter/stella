<script>
	import { writable } from 'svelte/store';
	import { selectedPlanet } from './stores';
	import { onMount } from 'svelte';
	import { initScene } from './scene';

	const state = writable({ game: null });
	let canvas;

	const newSocket = () => {
		const ws = new WebSocket(
			`ws://${location.host}/api/planet/play?planet=${$selectedPlanet?.hash}&token=${localStorage.getItem('token')}`
		);
		return ws;
	};
	let ws = newSocket();

	ws.onmessage = e => {
		const data = JSON.parse(e.data);
		delete data.type;
		state.update(prev => ({ ...prev, ...data }));
	};
	ws.onopen = () => alert('connected');
	ws.onclose = () => alert('disconnected');
	document.onkeydown = e => ws.send(e.code);

	onMount(() => initScene(canvas));
</script>

<canvas bind:this={canvas}>{JSON.stringify($state)}</canvas>
