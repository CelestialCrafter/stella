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
	@use '../styles/spacing.scss';

	ul {
		list-style: none;
		border-radius: spacing.$radius;
		padding: spacing.$padding;
		background-color: colors.$floating;
		overflow: scroll;

		li {
			border-radius: calc(spacing.$radius * 2);
			padding: calc(spacing.$padding / 2) spacing.$padding;
		}
	}

	a {
		text-decoration: none;
		display: block;
		text-overflow: ellipsis;
		overflow: hidden;
	}
</style>
