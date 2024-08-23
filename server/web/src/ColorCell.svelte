<script>
	// @NOTE expects rgb value
	export let value;
	// @TODO add notification when color copied
	const copy = () => navigator.clipboard.writeText(value.join(' '));
</script>

<button on:click={copy} class="color-cell" style="--value: {value}">.</button>

<style>
	.color-cell {
		border: none;
		display: inline-block;
		vertical-align: middle;
		padding: 0;
		color: transparent;
		width: 3rem;
		background-color: rgb(var(--value));
	}
</style>
