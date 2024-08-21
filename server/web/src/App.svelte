<script>
	import { onDestroy, onMount } from 'svelte';
	import * as jose from 'jose';

	import ApiKey from './ApiKey.svelte';
	import Play from './play/Play.svelte';
	import Inventory from './planets/Inventory.svelte';

	import { planets, selectedPlanet } from './stores';
	import { Route, Router } from 'svelte-routing';

	const hashChange = () => {
		const hash = window.location.hash.split('#')[1];
		if (Object.keys($planets).includes(hash)) selectedPlanet.set($planets[hash]);
	};

	const login = () => window.location.assign('/auth/login');
	onMount(() =>
		(async () => {
			if (!window.localStorage.getItem('token')) return login();
			const token = jose.decodeJwt(window.localStorage.getItem('token'));
			const response = await fetch(`/api/user/${token.id}`);
			if (!response.ok) return login();
			const user = await response.json();
			planets.set(user.planets.reduce((acc, x) => ({ ...acc, [x.hash]: x }), {}));
			hashChange();
		})()
	);

	window.onhashchange = hashChange;
	onDestroy(() => window.removeEventListener('hashchange', hashChange));
</script>

<Router basepath="/app">
	{#if Object.keys(planets).length < 1}
		<p>
			Loading...<br />
			(or no planets)
		</p>
		<br />
		<ApiKey />
	{:else}
		<Route path="/inventory" component={Inventory} />
		<Route path="/play" component={Play} />
	{/if}
</Router>
