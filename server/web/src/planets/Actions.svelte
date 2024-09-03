<script>
	export let disabled;

	import { navigate } from 'svelte-routing';

	import Delete from './actions/Delete.svelte';
	import Transfer from './actions/Transfer.svelte';
	import New from './actions/New.svelte';
	import { selectedPlanet } from '../stores';

	let deleteDialog;
	let transferDialog;
	let newDialog;
</script>

<Delete bind:dialog={deleteDialog} />
<Transfer bind:dialog={transferDialog} />
<New bind:dialog={newDialog} />

<section>
	<button on:click={() => newDialog.showModal()}>New</button>
	<button on:click={() => deleteDialog.showModal()} {disabled}>Delete</button>
	<button on:click={() => transferDialog.showModal()} {disabled}>Transfer</button>
</section>

<style lang="scss">
	@use '../styles/colors.scss';
	@use '../styles/spacing.scss';

	section {
		text-align: center;
		gap: calc(spacing.$padding / 2);
		display: grid;
		justify-content: center;
		grid-template-columns: repeat(auto-fill, 40%);

		button {
			border-color: colors.$primary;
		}
	}
</style>
