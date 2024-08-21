<!-- @NOTE expects rgb value -->

<script>
	export let value;
	const copy = () => navigator.clipboard.writeText(value.join(' '));
</script>

<button on:click={copy} class="color-cell" style="--value: {value}">.</button>

<style>
	.color-cell {
		display: inline-block;
		vertical-align: middle;
		padding: 0;
		color: transparent;
		width: 3rem;
		background-color: rgb(var(--value));
	}
</style>
