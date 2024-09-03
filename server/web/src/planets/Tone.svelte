<script>
	export let hash;
	
	let audio;
	let volume;
	let paused = true;
</script>

<audio
	src={`/sounds/${hash}.wav`}
	bind:volume
	bind:this={audio}
	bind:paused
	loop
></audio>

<div>
	<button on:click={() => paused ? audio.play() : audio.pause()}>
		{paused ? 'Play' : 'Pause'}
	</button>
	<input type="range" min="0" max="1" bind:value={volume} step="0.01">
</div>

<style>
	input {
		width: 20%;
	}
</style>
