<script>
	export let planet;
</script>

{#if planet != null}
	{@const { features, values } = planet}
	<section class="selected-planet">
		<span>Type: {features.type}</span><br />
		<span>Neutron Star: {features.star_neutron ? 'yes' : 'no'}</span><br />

		{#if features.type === 'normal'}
			<span>Size: {values.normal_size}</span><br />
			<span>Color: rgb {values.normal_color.join(', ')}</span><br />
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
