<script>
	export let selected;
	let output = '';

	const playSocket = (_, planet) => {
		const newSocket = newPlanet => {
			const ws = new WebSocket(
				`ws://${location.host}/api/planet/play?planet=${newPlanet?.hash}&token=${localStorage.getItem('token')}`
			);
			return ws;
		};
		let ws = newSocket(planet);

		ws.onmessage = e => (output += e.data + '\n');
		ws.onopen = () => (output += 'connected\n');
		ws.onclose = () => (output = '');
		document.onkeydown = e => ws.send(e.code);

		return {
			update: planet => {
				ws.close();
				ws = newSocket(planet);
			},
			destroy: () => {}
		};
	};
</script>

<section>
	<code use:playSocket={selected}>{output}</code>
</section>

<style>
	code {
		white-space: pre-line;
	}
</style>
