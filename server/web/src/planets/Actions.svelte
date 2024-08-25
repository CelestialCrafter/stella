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
	<button on:click={() => navigate(`/play/${$selectedPlanet}`)} {disabled}> Play </button>
</section>

<style lang="scss">
	@use '../styles/colors.scss';

	section {
		text-align: center;
		button {
			border-color: colors.$primary;
		}
	}
</style>
