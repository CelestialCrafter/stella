<script>
	export let dialog;

	import { planets, selectedPlanet } from '../../stores';
	import { onMount } from 'svelte';

	let destination;
	onMount(() => {
		dialog.onclose = async () => {
			const planet = $selectedPlanet;

			const rv = dialog.returnValue;
			if (rv !== 'confirm') return;

			// @TODO add notification when deleted
			await fetch(`/api/planet/transfer/${planet}`, {
				method: 'POST',
				body: JSON.stringify({ destination }),
				headers: {
					'Content-Type': 'application/json;charset=utf-8',
					Authorization: `Bearer ${localStorage.getItem('token')}`
				}
			});
			planets.update(prev => {
				const copy = { ...prev };
				delete copy[planet];
				return copy;
			});
		};
	});
</script>

<dialog bind:this={dialog}>
	<h2>Are you sure you want to transfer {$selectedPlanet}?</h2>
	<span>
		WARNING: double check the new owner id! if you mistype it, the planet will be effectively voided
	</span>
	<form method="dialog">
		<input type="text" bind:value={destination} placeholder="New Owner ID..." required />
		<button value="confirm">Transfer</button>
		<button value="cancel" formnovalidate>Cancel</button>
	</form>
</dialog>

<style lang="scss">
	@use '../../colors.scss';

	h2,
	span {
		color: colors.$danger;
	}
</style>
