<script>
	export let selected;
	let output = '';

	const playSocket = (_, planet) => {
		const newSocket = newPlanet =>
			new WebSocket(`ws://${location.host}/api/planet/play?id=${newPlanet?.hash}`);
		let ws = newSocket(planet);

		ws.onmessage = e => (output += e.data + '\n');
		ws.onopen = () => (output += 'connected\n');
		ws.onclose = () => (output = '');

		setInterval(() => ws.send('bababab'), 1000);

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
