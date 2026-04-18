<script lang="ts">
	import { Terminal } from '@xterm/xterm';
	import { FitAddon } from '@xterm/addon-fit';
	import '@xterm/xterm/css/xterm.css';

	let terminalElement: HTMLDivElement;
	let term: Terminal;
	let fitAddon: FitAddon;
	let wasmReady = $state(false);

	$effect(() => {
		if (typeof window === 'undefined' || !terminalElement) return;

		// 1. Initialize Terminal with a premium theme
		term = new Terminal({
			cursorBlink: true,
			convertEol: true,
			scrollback: 5000,
			theme: {
				background: '#0d1117',
				foreground: '#c9d1d9',
				cursor: '#58a6ff',
				cursorAccent: '#0d1117',
				selectionBackground: 'rgba(88,166,255,0.25)',
				black: '#484f58',
				red: '#ff7b72',
				green: '#3fb950',
				yellow: '#d29922',
				blue: '#58a6ff',
				magenta: '#bc8cff',
				cyan: '#39c5cf',
				white: '#b1bac4',
				brightBlack: '#6e7681',
				brightRed: '#ffa198',
				brightGreen: '#56d364',
				brightYellow: '#e3b341',
				brightBlue: '#79c0ff',
				brightMagenta: '#d2a8ff',
				brightCyan: '#56d4dd',
				brightWhite: '#f0f6fc'
			},
			fontFamily: '"Fira Code", "Cascadia Code", Menlo, Monaco, "Courier New", monospace',
			fontSize: 14,
			lineHeight: 1.1
		});

		fitAddon = new FitAddon();
		term.loadAddon(fitAddon);
		term.open(terminalElement);
		fitAddon.fit();

		term.writeln('\x1b[33mLoading main.wasm...\x1b[0m');

		// 2. Intercept Go stdout/stderr BEFORE go.run()
		const decoder = new TextDecoder();
		const fs = (globalThis as any).fs;
		if (fs) {
			const origWriteSync = fs.writeSync.bind(fs);
			fs.writeSync = function (fd: number, buf: Uint8Array) {
				if (fd === 1 || fd === 2) {
					term.write(decoder.decode(buf));
					return buf.length;
				}
				return origWriteSync(fd, buf);
			};

			const origWrite = fs.write.bind(fs);
			fs.write = function (
				fd: number,
				buf: Uint8Array,
				offset: number,
				length: number,
				position: number,
				callback: (err: Error | null, n: number) => void
			) {
				if (fd === 1 || fd === 2) {
					const n = fs.writeSync(fd, buf.subarray(offset, offset + length));
					callback(null, n);
					return;
				}
				return origWrite(fd, buf, offset, length, position, callback);
			};
		}

		// 3. Load & run WASM
		const initWasm = async () => {
			try {
				const Go = (globalThis as any).Go;
				if (!Go) {
					throw new Error('Go constructor not found. Is wasm_exec.js loaded?');
				}
				const go = new Go();
				const result = await WebAssembly.instantiateStreaming(fetch('/main.wasm'), go.importObject);

				term.writeln('\x1b[32mReady! Type a command and press Enter.\x1b[0m');

				// go.run() returns a promise but we do NOT await it here.
				go.run(result.instance);

				// Give the Go runtime a tick to initialise and register writeStdin.
				await new Promise((r) => setTimeout(r, 150));

				wasmReady = true;
				term.focus();
			} catch (err) {
				term.writeln(`\x1b[31mFailed to load WASM: ${err}\x1b[0m`);
				console.error(err);
			}
		};

		initWasm();

		// 4. Wire Terminal Data to WASM Stdin (Raw Mode)
		term.onData((data) => {
			if (!wasmReady) return;
			if (typeof (window as any).writeStdin === 'function') {
				(window as any).writeStdin(data);
			}
		});

		const handleResize = () => fitAddon.fit();
		window.addEventListener('resize', handleResize);

		return () => {
			window.removeEventListener('resize', handleResize);
			term.dispose();
		};
	});

	const seoData = {
		title: 'Bash Simulator WASM (Golang) - Interactive Unix Shell',
		description:
			'A high-performance Bash shell simulator powered by Go and WebAssembly. Run interactive Linux commands directly in your browser with our WASM-based terminal.',
		url: 'https://bash.devops-playground.dev/',
		image: 'https://bash.devops-playground.dev/social-preview.png',
		twitterHandle: '@yarencheng',
		siteName: 'Bash Simulator WASM',
		keywords: 'Bash, WASM, Golang, WebAssembly, Terminal, Shell Simulator, Linux Online, DevOps'
	};

	const schemaData = {
		'@context': 'https://schema.org',
		'@type': 'SoftwareApplication',
		name: 'Bash Simulator WASM',
		alternateName: 'Bash WASM(Golang)',
		operatingSystem: 'Any browser with WebAssembly support',
		applicationCategory: 'DeveloperApplication',
		description:
			'A robust, interactive Bash shell simulator built with Go and compiled to WebAssembly. Features include coreutils parity, filesystem simulation, and interactive command execution.',
		offers: {
			'@type': 'Offer',
			price: '0',
			priceCurrency: 'USD'
		},
		author: {
			'@type': 'Person',
			name: 'aren',
			url: 'https://github.com/yarencheng'
		},
		softwareVersion: '1.0.0',
		keywords: seoData.keywords
	};

	const faqData = {
		'@context': 'https://schema.org',
		'@type': 'FAQPage',
		mainEntity: [
			{
				'@type': 'Question',
				name: 'What is Bash Simulator WASM?',
				acceptedAnswer: {
					'@type': 'Answer',
					text: 'It is a full-featured Bash shell simulator developed in Go and running entirely in the browser using WebAssembly (WASM).'
				}
			},
			{
				'@type': 'Question',
				name: 'How do I use this Bash shell?',
				acceptedAnswer: {
					'@type': 'Answer',
					text: 'Simply open the website and start typing standard Unix commands like ls, cd, mkdir, and echo. The environment is persistent within your browser session.'
				}
			},
			{
				'@type': 'Question',
				name: 'Is this a real Linux kernel?',
				acceptedAnswer: {
					'@type': 'Answer',
					text: 'No, it is a high-fidelity simulation of the Bash shell and Coreutils running on a virtual filesystem in WebAssembly, not a full Linux kernel.'
				}
			}
		]
	};
</script>

<svelte:head>
	<title>{seoData.title}</title>
	<meta name="description" content={seoData.description} />
	<meta name="keywords" content={seoData.keywords} />
	<link rel="canonical" href={seoData.url} />

	<!-- Open Graph / Facebook -->
	<meta property="og:type" content="website" />
	<meta property="og:url" content={seoData.url} />
	<meta property="og:title" content={seoData.title} />
	<meta property="og:description" content={seoData.description} />
	<meta property="og:image" content={seoData.image} />
	<meta property="og:site_name" content={seoData.siteName} />
	<meta property="og:locale" content="en_US" />

	<!-- Twitter -->
	<meta property="twitter:card" content="summary_large_image" />
	<meta property="twitter:url" content={seoData.url} />
	<meta property="twitter:title" content={seoData.title} />
	<meta property="twitter:description" content={seoData.description} />
	<meta property="twitter:image" content={seoData.image} />
	<meta name="twitter:creator" content={seoData.twitterHandle} />
	<meta name="twitter:site" content={seoData.twitterHandle} />

	<!-- Structured Data -->
	{@html `<script type="application/ld+json">${JSON.stringify(schemaData)}</script>`}
	{@html `<script type="application/ld+json">${JSON.stringify(faqData)}</script>`}
</svelte:head>

<!-- Visually hidden semantic content for SEO and AEO -->
<div class="sr-only">
	<h1>{seoData.title}</h1>
	<p>{seoData.description}</p>
	<section>
		<h2>Key Features of Bash WASM(Golang)</h2>
		<ul>
			<li>Full interactive Bash shell simulation</li>
			<li>Powered by Go and WebAssembly</li>
			<li>Virtual filesystem within the browser</li>
			<li>GNU Coreutils parity (ls, cd, cp, mv, etc.)</li>
			<li>Fast and lightweight execution</li>
		</ul>
	</section>
	<section>
		<h2>Frequently Asked Questions</h2>
		{#each faqData.mainEntity as faq}
			<article>
				<h3>{faq.name}</h3>
				<p>{faq.acceptedAnswer.text}</p>
			</article>
		{/each}
	</section>
</div>

<div class="terminal-container" bind:this={terminalElement}></div>

<style>
	.terminal-container {
		width: 100vw;
		height: 100vh;
		background-color: #0a0a0a;
		padding: 1rem;
		box-sizing: border-box;
	}

	.sr-only {
		position: absolute;
		width: 1px;
		height: 1px;
		padding: 0;
		margin: -1px;
		overflow: hidden;
		clip: rect(0, 0, 0, 0);
		white-space: nowrap;
		border-width: 0;
	}

	:global(.xterm) {
		padding: 10px;
	}

	:global(.xterm-viewport) {
		background-color: transparent !important;
	}
</style>
