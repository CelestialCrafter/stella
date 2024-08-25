<script>
	export let dialog;

	import { planets, selectedPlanet } from '../../stores';
	import { onMount } from 'svelte';

	onMount(() => {
		dialog.onclose = async () => {
			const planet = $selectedPlanet;

			const rv = dialog.returnValue;
			if (rv !== 'confirm') return;

			// @TODO add notification when deleted && loading indicator
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

			let planetKeys = Object.keys($planets);
			selectedPlanet.set(planetKeys[planetKeys.length - 1]);
		};
	});
</script>

<dialog bind:this={dialog}>
	<h2>Are you sure you want to delete {$selectedPlanet}?</h2>
	<form method="dialog" class="triggers">
		<button value="confirm">Delete</button>
		<button value="cancel">Cancel</button>
	</form>
</dialog>

<style lang="scss">
	@use '../../styles/colors.scss';
	@use '../../styles/spacing.scss';

	dialog {
		max-width: 50%;
		&[open] {
			gap: spacing.$padding;
			display: flex;
			flex-direction: column;
		}
	}

	.triggers button {
		border-color: colors.$primary;
	}

	h2 {
		overflow: hidden;
		text-overflow: ellipsis;
		color: colors.$danger;
	}
</style>
