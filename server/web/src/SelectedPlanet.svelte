<script>
	import { selectedPlanet } from './stores';
</script>

{#if $selectedPlanet != null}
	{@const { features, values } = $selectedPlanet}
	<section class="selected-planet">
		<span>Type: {features.type}</span><br />

		{#if features.type === 'normal'}
			<span>Rings: {features.normal_rings ? 'yes' : 'no'}</span><br />
			<span>Size: {values.normal_size}</span><br />
			<span>Color: rgb {values.normal_color.join(', ')}</span><br />
			{#if features.normal_rings}
				<span>Rings: {values.normal_ring_amount}</span><br />
				<span>Ring Size: {values.normal_ring_size}</span><br />
				<span>Ring Color: rgb {values.normal_ring_colors.join(', ')}</span><br />
				<span>Ring Rotation: xyz {values.normal_ring_rotation.join(', ')}</span><br />
			{/if}
		{:else if features.type === 'star'}
			<span>Neutron Star: {features.star_neutron ? 'yes' : 'no'}</span><br />
			<span>Brightness: {values.star_brightness}</span><br />
			<span>Size: {values.star_size}</span><br />
			{#if features.star_neutron}
				<span>Color: rgb {values.star_neutron_color.join(', ')}</span><br />
			{/if}
		{:else if features.type === 'blackhole'}
			<span>Size: {values.normal_size}</span><br />
			<span>
				Blackhole Style: {features.blackhole_style === 'vertical' ? 'vertical' : 'horizontal'}
			</span><br />
			<span>Blackhole Ring Color 1: rgb {values.blackhole_colors[0].join(', ')}</span><br />
			<span>Blackhole Ring Color 2: rgb {values.blackhole_colors[1].join(', ')}</span><br />
		{/if}
	</section>
{/if}

<style>
	.selected-planet {
		width: 40%;
		position: fixed;
		background-color: white;
		z-index: 10;
		padding: 1rem;
	}
</style>
