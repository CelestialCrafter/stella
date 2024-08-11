<script>
	import { onMount } from 'svelte';
	import { selectedPlanet } from './scene.js';
	import * as jose from 'jose';

	import Scene from './Scene.svelte';
	import SelectedPlanet from './SelectedPlanet.svelte';

	$: planets = {};
	$: selected = null;

	const login = () => window.location.assign('/auth/login');
	onMount(async () => {
		if (!window.localStorage.getItem('token')) return login();

		const token = jose.decodeJwt(localStorage.getItem('token'));
		const response = await fetch(`/api/user/${token.id}`);
		if (!response.ok) return login();
		const user = await response.json();
		planets = user.planets.reduce((acc, x) => ({ ...acc, [x.hash]: x }), {});
	});
</script>

<main>
	{#if Object.keys(planets).length < 1}
		<span>
			Loading...<br />
			(or no planets)
		</span>
	{:else}
		<SelectedPlanet planet={selected ? planets[selected] : null} />
		<Scene
			handleCanvasClick={() => (selected = selectedPlanet())}
			planets={Object.values(planets)}
		/>
	{/if}
</main>

<style>
</style>
