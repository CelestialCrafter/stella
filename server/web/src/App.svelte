<script>
	import { onDestroy, onMount, tick } from 'svelte';
	import * as jose from 'jose';

	import ApiKey from './ApiKey.svelte';
	import Play from './play/Play.svelte';
	import Inventory from './planets/Inventory.svelte';

	import { planets, selectedPlanet } from './stores';
	import { navigate, Route, Router } from 'svelte-routing';

	const unsubscribe = selectedPlanet.subscribe(hash => {
		if (!hash) return (window.location.hash = '');
		window.location.hash = `#${hash}`;
	});

	const hashChange = () => {
		const hash = window.location.hash.split('#')[1];
		if ($selectedPlanet == hash) return;
	};

	const login = () => window.location.assign('/auth/login');
	onMount(() =>
		(async () => {
			if (!window.localStorage.getItem('token')) return login();
			const token = jose.decodeJwt(window.localStorage.getItem('token'));
			if (token.exp < Date.now() / 1000) return login();
			const response = await fetch(`/api/user/${token.id}`);
			if (!response.ok) return login();

			const user = await response.json();
			planets.set(user.planets.reduce((acc, x) => ({ ...acc, [x.hash]: x }), {}));

			const hash = window.location.hash;
			if (Object.keys($planets).includes(hash)) selectedPlanet.set(hash);
		})()
	);

	window.onhashchange = hashChange;
	onDestroy(() => {
		unsubscribe();
		window.removeEventListener('hashchange', hashChange);
	});
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
		<Route path="/">
			{tick().then(() => navigate('/app/inventory'))}
		</Route>
		<Route path="/inventory" component={Inventory} />
		<Route path="/play/:id" component={Play} let:params />
	{/if}
</Router>
