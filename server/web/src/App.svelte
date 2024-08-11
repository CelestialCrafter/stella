<script>
	import { onMount } from 'svelte';
	import { selectedPlanet } from './scene.js';
	import * as jose from 'jose';

	import Scene from './Scene.svelte';
	import SelectedPlanet from './SelectedPlanet.svelte';

	$: planets = {};
	$: selected = null;
	onMount(async () => {
		if (!window.localStorage.getItem('token')) return window.location.assign('/auth/login');

		const token = jose.decodeJwt(localStorage.getItem('token'));
		const user = await (await fetch(`/api/user/${token.id}`)).json();
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
