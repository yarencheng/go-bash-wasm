<script lang="ts">
	import { onMount } from 'svelte';
	import { fly } from 'svelte/transition';

	let visible = $state(false);

	onMount(() => {
		// Check if user has already agreed
		const agreed = localStorage.getItem('consent_agreed');
		if (agreed === 'true') {
			// Already agreed, update consent immediately
			updateConsent();
		} else {
			// Not yet agreed, show banner after a delay
			setTimeout(() => {
				visible = true;
			}, 1000);
		}
	});

	function updateConsent() {
		// Update Google Analytics consent state
		if (typeof gtag === 'function') {
			gtag('consent', 'update', {
				'analytics_storage': 'granted',
				'ad_storage': 'granted',
				'ad_user_data': 'granted',
				'ad_personalization': 'granted'
			});
		}

		// Update Clarity if available
		if (typeof clarity === 'function') {
			clarity('consent');
		}
	}

	function handleAgree() {
		localStorage.setItem('consent_agreed', 'true');
		visible = false;
		updateConsent();
	}

	function handleDeny() {
		visible = false;
		// We do NOT save the deny state to localStorage, so it will show up again on next load
		// as per the requirement: "ask again when user open this page again"
	}
</script>

{#if visible}
	<div
		class="consent-banner"
		transition:fly={{ y: 50, duration: 800, opacity: 0 }}
		role="dialog"
		aria-labelledby="consent-text"
	>
		<div class="glass-container">
			<div class="content">
				<div class="icon-wrapper">
					<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="analytics-icon"><path d="M3 3v18h18"></path><path d="m19 9-5 5-4-4-3 3"></path></svg>
				</div>
				<p id="consent-text">
					We use analytics tools to understand how people use our terminal and to improve the experience. 
					Data is collected anonymously.
				</p>
			</div>
			<div class="actions">
				<button class="btn btn-deny" onclick={handleDeny}>
					Deny
				</button>
				<button class="btn btn-agree" onclick={handleAgree}>
					Agree
				</button>
			</div>
		</div>
	</div>
{/if}

<style>
	.consent-banner {
		position: fixed;
		bottom: 1.5rem;
		left: 0;
		right: 0;
		display: flex;
		justify-content: center;
		padding: 0 1rem;
		z-index: 10000;
		pointer-events: none;
	}

	.glass-container {
		pointer-events: auto;
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 1.25rem;
		width: 100%;
		max-width: 650px;
		padding: 1.25rem 1.5rem;
		background: rgba(13, 17, 23, 0.8);
		backdrop-filter: blur(16px) saturate(180%);
		-webkit-backdrop-filter: blur(16px) saturate(180%);
		border: 1px solid rgba(88, 166, 255, 0.2);
		border-radius: 16px;
		box-shadow: 0 10px 40px rgba(0, 0, 0, 0.5), inset 0 0 0 1px rgba(255, 255, 255, 0.05);
	}

	@media (min-width: 768px) {
		.glass-container {
			flex-direction: row;
			justify-content: space-between;
			padding: 1rem 1.5rem;
		}
	}

	.content {
		display: flex;
		align-items: center;
		gap: 1rem;
	}

	.icon-wrapper {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 36px;
		height: 36px;
		background: rgba(88, 166, 255, 0.1);
		border-radius: 10px;
		color: #58a6ff;
		flex-shrink: 0;
	}

	p {
		margin: 0;
		font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji";
		font-size: 0.9rem;
		line-height: 1.4;
		color: #c9d1d9;
	}

	.actions {
		display: flex;
		gap: 0.75rem;
		width: 100%;
	}

	@media (min-width: 768px) {
		.actions {
			width: auto;
		}
	}

	.btn {
		flex: 1;
		white-space: nowrap;
		padding: 0.6rem 1.25rem;
		font-size: 0.85rem;
		font-weight: 600;
		border-radius: 8px;
		cursor: pointer;
		transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
		border: 1px solid transparent;
		font-family: inherit;
	}

	@media (min-width: 768px) {
		.btn {
			flex: none;
			min-width: 90px;
		}
	}

	.btn-deny {
		background: rgba(48, 54, 61, 0.5);
		color: #8b949e;
		border-color: rgba(240, 246, 2fc, 0.1);
	}

	.btn-deny:hover {
		background: rgba(48, 54, 61, 0.8);
		color: #c9d1d9;
		border-color: rgba(240, 246, 2fc, 0.2);
	}

	.btn-agree {
		background: #238636;
		color: #ffffff;
		border-color: rgba(240, 246, 2fc, 0.1);
	}

	.btn-agree:hover {
		background: #2ea043;
		transform: translateY(-1px);
		box-shadow: 0 4px 12px rgba(35, 134, 54, 0.3);
	}

	.btn:active {
		transform: translateY(0);
	}

	.analytics-icon {
		filter: drop-shadow(0 0 4px rgba(88, 166, 255, 0.4));
	}
</style>
