<script>
	import { onMount } from 'svelte';
	import * as jose from 'jose';

	import SelectedPlanet from './SelectedPlanet.svelte';
	import ApiKey from './ApiKey.svelte';
	import Scene from './Scene.svelte';

	$: planets = {};
	$: selected = null;

	const login = () => window.location.assign('/auth/login');
	onMount(async () => {
		if (!window.localStorage.getItem('token')) return login();

		const token = jose.decodeJwt(window.localStorage.getItem('token'));
		const response = await fetch(`/api/user/${token.id}`);
		if (!response.ok) return login();
		const user = await response.json();
		planets = user.planets.reduce((acc, x) => ({ ...acc, [x.hash]: x }), {});
	});
</script>

<main>
	{#if Object.keys(planets).length < 1}
		<p>
			Loading...<br />
			(or no planets)
		</p>
		<br />
		<ApiKey />
	{:else}
		<SelectedPlanet planet={selected ? planets[selected] : null} />
		<Scene bind:selected planets={Object.values(planets)} />
	{/if}
</main>
