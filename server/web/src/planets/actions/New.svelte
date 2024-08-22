<script>
	export let dialog;

	import { planets } from '../../stores';
	import { onMount } from 'svelte';

	const features = {
		type: 'normal',
		normal_rings: false,
		star_neutron: false,
		blackhole_style: 'horizontal'
	};

	onMount(() => {
		dialog.onclose = async () => {
			const rv = dialog.returnValue;
			if (rv !== 'confirm') return;

			features.normal_rings = Boolean(features.normal_rings);
			features.star_neutron = Boolean(features.star_neutron);

			// @TODO add notification when deleted
			const newPlanet = await fetch(`/api/planet/new`, {
				method: 'POST',
				body: JSON.stringify(features),
				headers: {
					'Content-Type': 'application/json;charset=utf-8',
					Authorization: `Bearer ${localStorage.getItem('token')}`
				}
			}).then(res => res.json());

			planets.update(prev => {
				const copy = { ...prev };
				copy[newPlanet.hash] = newPlanet;
				return copy;
			});
		};
	});
</script>

<dialog bind:this={dialog}>
	<form method="dialog">
		<label for="type">Type</label>
		<select id="type" bind:value={features.type}>
			<option value="normal" selected>Normal</option>
			<option value="star">Star</option>
			<option value="blackhole">Black Hole</option>
		</select>

		{#if features.type === 'normal'}
			<label for="normal_rings">Rings?</label>
			<input id="normal_rings" type="checkbox" bind:value={features.normal_rings} />
		{:else if features.type === 'star'}
			<label for="star_neutron">Neutron?</label>
			<input id="star_neutron" type="checkbox" bind:value={features.star_neutron} />
		{:else}
			<label for="blackhole_style">Blackhole Style</label>
			<select id="blackhole_style" bind:value={features.blackhole_style}>
				<option value="horizontal" selected>Horizontal</option>
				<option value="vertical">Vertical</option>
			</select>
		{/if}

		<button value="confirm">Create</button>
		<button value="cancel" formnovalidate>Cancel</button>
	</form>
</dialog>
