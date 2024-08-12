<script>
	export let apiKey = null;

	const generateApiKey = async () => {
		const response = await fetch('/api/key/new', {
			method: 'PUT',
			headers: {
				Authorization: `Bearer ${window.localStorage.getItem('token')}`
			}
		}).then(data => data.json());
		apiKey = response.token;
	};
</script>

<section>
	<button on:click={generateApiKey}>Generate API Key</button>
	{#if apiKey}
		<button on:click={() => navigator.clipboard.writeText(apiKey)}>Copy</button>
		<pre>{apiKey}</pre>
	{/if}
</section>

<style>
	pre {
		word-wrap: break-word;
		white-space: pre-wrap;
		width: 30%;
	}
</style>
