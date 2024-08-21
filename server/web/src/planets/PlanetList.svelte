<script>
	import { navigate } from 'svelte-routing';
	import { planets, selectedPlanet } from '../stores';
	import { onMount } from 'svelte';
	$: actionsDisabled = !$selectedPlanet;

	let deleteDialog;

	onMount(() => {
		deleteDialog.onclose = async () => {
			const planet = $selectedPlanet;

			deleteDialog.showModal();
			const rv = deleteDialog.returnValue;
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

<dialog bind:this={deleteDialog}>
	<h2>Are you sure you want to delete {$selectedPlanet}?</h2>
	<form method="dialog">
		<button value="confirm">Delete</button>
		<button value="cancel">Cancel</button>
	</form>
</dialog>

<ul style={actionsDisabled ? 'grid-row: 1/-1' : 'grid-row: 1'}>
	<div class="actions">
		<button>New</button>
		<button on:click={() => deleteDialog.showModal()} disabled={actionsDisabled}>Delete</button>
		<button disabled={actionsDisabled}>Transfer</button>
		<button on:click={() => navigate(`/play/${$selectedPlanet}`)} disabled={actionsDisabled}>
			Play
		</button>
	</div>
	<hr />
	{#each Object.keys($planets) as hash}
		<li class={$selectedPlanet === hash ? 'active' : ''}>
			<a on:click={() => selectedPlanet.set(hash)} href={`#${hash}`}>
				{hash}
			</a>
		</li>
	{/each}
</ul>

<style lang="scss">
	@use '../colors.scss';

	ul {
		list-style: none;
		border-radius: 1rem;
		padding: 1rem;
		background-color: colors.$floating;

		.actions {
			text-align: center;
		}

		li {
			border-radius: 2rem;
			padding: 0.5rem;
		}
	}

	a {
		text-decoration: none;
		display: block;
		text-overflow: ellipsis;
		overflow: hidden;
	}
</style>
