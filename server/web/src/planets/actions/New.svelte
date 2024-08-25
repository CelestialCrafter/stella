<script>
	export let dialog;

	import { planets, selectedPlanet } from '../../stores';
	import { onMount } from 'svelte';

	const features = {
		type: 'normal',
		normal_rings: false,
		star_neutron: false,
		blackhole_style: 'horizontal',
		nickname: ''
	};

	onMount(() => {
		dialog.onclose = async () => {
			const rv = dialog.returnValue;
			if (rv !== 'confirm') return;

			features.normal_rings = Boolean(features.normal_rings);
			features.star_neutron = Boolean(features.star_neutron);

			// @TODO add notification when created && loading indicator
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

			selectedPlanet.set(newPlanet.hash);
		};
	});
</script>

<dialog bind:this={dialog}>
	<form method="dialog">
		<div>
			<label>
				Nickname
				<input type="text" bind:value={features.nickname} />
			</label>
			<label>
				Type
				<select bind:value={features.type}>
					<option value="normal" selected>Normal</option>
					<option value="star">Star</option>
					<option value="blackhole">Black Hole</option>
				</select>
			</label>

			{#if features.type === 'normal'}
				<label>
					Rings?
					<input type="checkbox" bind:value={features.normal_rings} />
				</label>
			{:else if features.type === 'star'}
				<label>
					Neutron?
					<input type="checkbox" bind:value={features.star_neutron} />
				</label>
			{:else}
				<label>
					Blackhole Style
					<select bind:value={features.blackhole_style}>
						<option value="horizontal" selected>Horizontal</option>
						<option value="vertical">Vertical</option>
					</select>
				</label>
			{/if}
		</div>

		<div class="triggers">
			<button value="confirm">Create</button>
			<button value="cancel" formnovalidate>Cancel</button>
		</div>
	</form>
</dialog>

<style lang="scss">
	@use '../../styles/colors.scss';
	dialog {
		max-width: 50%;
	}

	form {
		gap: 1rem;
		display: flex;
		flex-direction: column;
	}

	.triggers button {
		border-color: colors.$primary;
	}
</style>
