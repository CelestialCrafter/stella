<script>
	import { navigate } from 'svelte-routing';
	import { planets, selectedPlanet } from '../stores';

	import Delete from './actions/Delete.svelte';
	import Transfer from './actions/Transfer.svelte';
	import New from './actions/New.svelte';

	$: actionsDisabled = !$selectedPlanet;

	let deleteDialog;
	let transferDialog;
	let newDialog;
</script>

<Delete bind:dialog={deleteDialog} />
<Transfer bind:dialog={transferDialog} />
<New bind:dialog={newDialog} />

<ul style={actionsDisabled ? 'grid-row: 1/-1' : 'grid-row: 1'}>
	<div class="actions">
		<button on:click={() => newDialog.showModal()}>New</button>
		<button on:click={() => deleteDialog.showModal()} disabled={actionsDisabled}>Delete</button>
		<button on:click={() => transferDialog.showModal()} disabled={actionsDisabled}>Transfer</button>
		<button on:click={() => navigate(`/play/${$selectedPlanet}`)} disabled={actionsDisabled}>
			Play
		</button>
	</div>
	<hr />
	{#each Object.values($planets) as { hash, features }}
		<li class={$selectedPlanet === hash ? 'active' : ''}>
			<a on:click={() => selectedPlanet.set(hash)} href={`#${hash}`}>
				{features['nickname'] || hash}
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
