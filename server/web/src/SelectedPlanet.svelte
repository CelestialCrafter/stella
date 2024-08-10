<script>
	export let planet;
</script>

{#if planet != null}
	{@const { features, values } = planet}
	<section class="selected-planet">
		<span>Type: {features.type}</span><br />
		<span>Neutron Star: {features.star_neutron ? 'yes' : 'no'}</span><br />
		<span>Rings: {features.normal_rings ? 'yes' : 'no'}</span><br />

		{#if features.type === 'normal'}
			<span>Size: {values.normal_size}</span><br />
			<span>Color: rgb {values.normal_color.join(', ')}</span><br />
			{#if features.normal_rings}
				<span>Rings: {values.normal_ring_amount}</span><br />
				<span>Ring Size: {values.normal_ring_size}</span><br />
				<span>Ring Color: rgb {values.normal_ring_colors.join(', ')}</span><br />
				<span>Ring Rotation: xyz {values.normal_ring_rotation.join(', ')}</span><br />
			{/if}
		{:else}
			<span>Brightness: {values.star_brightness}</span><br />
			<span>Size: {values.star_size}</span><br />
			{#if features.star_neutron}
				<span>Color: rgb {values.star_neutron_color.join(', ')}</span><br />
			{/if}
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
