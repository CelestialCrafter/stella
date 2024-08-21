<script>
	import { selectedPlanet } from '../stores';
	import ColorCell from '../ColorCell.svelte';

	const displayNumber = num => num.toFixed(2);
</script>

{#if $selectedPlanet}
	{@const { features, values } = $selectedPlanet}
	<ul>
		<li><span>Type</span>{features.type}</li>

		{#if features.type === 'normal'}
			<li><span>Rings</span>{features.normal_rings ? 'yes' : 'no'}</li>
			<li><span>Size</span>{displayNumber(values.normal_size)}</li>
			<li><span>Color</span><ColorCell value={values.normal_color} /></li>
			{#if features.normal_rings}
				<li><span>Rings</span>{values.normal_ring_amount}</li>
				<li><span>Ring Size</span>{displayNumber(values.normal_ring_size)}</li>
				{#each [...Array(values.normal_ring_amount).keys()] as i}
					<li><span>Ring Color {i + 1}</span><ColorCell value={values.normal_ring_colors[i]} /></li>
					<li>
						<span>Ring Rotation {i + 1}</span>xyz {values.normal_ring_rotation[i]
							.map(displayNumber)
							.join(', ')}
					</li>
				{/each}
			{/if}
		{:else if features.type === 'star'}
			<li><span>Neutron Star</span>{features.star_neutron ? 'yes' : 'no'}</li>
			<li><span>Brightness</span>{displayNumber(values.star_brightness)}</li>
			<li><span>Size</span>{displayNumber(values.star_size)}</li>
			{#if features.star_neutron}
				<li><span>Color</span><ColorCell value={values.star_neutron_color} /></li>
			{/if}
		{:else if features.type === 'blackhole'}
			<li><span>Size</span>{values.normal_size}</li>
			<li>
				Blackhole Style: {features.blackhole_style === 'vertical' ? 'vertical' : 'horizontal'}
			</li>
			<li><span>Blackhole Ring Color 1</span><ColorCell value={values.blackhole_colors[0]} /></li>
			<li><span>Blackhole Ring Color 2</span><ColorCell value={values.blackhole_colors[1]} /></li>
		{/if}
	</ul>
{/if}

<style lang="scss">
	@use '../colors.scss';

	ul {
		grid-row: 2;
		list-style: none;
		border-radius: 1rem;
		padding: 1rem;
		background-color: colors.$floating;
	}

	li {
		word-break: break-all;
	}

	span {
		font-weight: bold;
		margin-right: 0.5rem;
	}
</style>
