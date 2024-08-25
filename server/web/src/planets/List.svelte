<script>
	import { planets, selectedPlanet } from '../stores';
	import Actions from './Actions.svelte';

	$: actionsDisabled = !$selectedPlanet;
</script>

<ul style={actionsDisabled ? 'grid-row: 1/-1' : 'grid-row: 1'}>
	<Actions disabled={actionsDisabled} />
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
	@use '../styles/colors.scss';

	ul {
		list-style: none;
		border-radius: 1rem;
		padding: 1rem;
		background-color: colors.$floating;

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
