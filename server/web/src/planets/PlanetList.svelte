<script>
	import { planets, selectedPlanet } from '../stores';
	$: actionsDisabled = !$selectedPlanet;
</script>

<ul style={actionsDisabled ? 'grid-row: 1/-1' : 'grid-row: 1'}>
	<div class="actions">
		<button>New</button>
		<button disabled={actionsDisabled}>Delete</button>
		<button disabled={actionsDisabled}>Transfer</button>
		<button disabled={actionsDisabled}>Play</button>
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
