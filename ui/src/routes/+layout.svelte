<script lang="ts">
	import favicon from '$lib/assets/favicon.svg';
	import ConsentBanner from '$lib/components/ConsentBanner.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

{@render children()}

<ConsentBanner />
