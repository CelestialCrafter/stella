<script>
	export let dialog;

	import { planets, selectedPlanet } from '../../stores';
	import { onMount } from 'svelte';

	onMount(() => {
		dialog.onclose = async () => {
			const planet = $selectedPlanet;

			const rv = dialog.returnValue;
			if (rv !== 'confirm') return;

			// @TODO add notification when deleted
			await fetch(`/api/planet/${planet}`, {
				method: 'DELETE',
				headers: {
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
	<h2>Are you sure you want to delete {$selectedPlanet}?</h2>
	<form method="dialog">
		<button value="confirm">Delete</button>
		<button value="cancel">Cancel</button>
	</form>
</dialog>

<style lang="scss">
	@use '../../colors.scss';

	h2 {
		color: colors.$danger;
	}
</style>
