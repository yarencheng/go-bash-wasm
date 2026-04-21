<script lang="ts">
	import { Terminal } from '@xterm/xterm';
	import { FitAddon } from '@xterm/addon-fit';
	import { WebLinksAddon } from '@xterm/addon-web-links';
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
		term.loadAddon(new WebLinksAddon());
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
		let commandBuffer = '';
		term.onData((data) => {
			if (!wasmReady) return;

			// Capture command for Google Analytics
			for (const char of data) {
				if (char === '\r' || char === '\n') {
					const trimmedCommand = commandBuffer.trim();
					if (trimmedCommand && typeof (window as any).gtag === 'function') {
						(window as any).gtag('event', 'terminal_command', {
							command: trimmedCommand
						});
					}
					commandBuffer = '';
				} else if (char === '\x7f') {
					// Handle backspace
					commandBuffer = commandBuffer.slice(0, -1);
				} else if (char.charCodeAt(0) >= 32) {
					// Append printable characters
					commandBuffer += char;
				}
			}

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
		title: 'Bash Simulator Online - Interactive WASM Linux Terminal (Go)',
		description:
			'Run a full-featured Bash shell in your browser. Powered by Go and WebAssembly, our Linux simulation includes coreutils, file system, and interactive command execution.',
		url: 'https://bash.devops-playground.dev/',
		image: 'https://bash.devops-playground.dev/social-preview.png',
		twitterHandle: '@yarencheng',
		siteName: 'Bash Simulator WASM',
		keywords:
			'Bash Online, Linux Simulator, WebAssembly, Golang, Terminal Emulator, Coreutils, DevOps Tools, Web Shell'
	};

	const schemaData = {
		'@context': 'https://schema.org',
		'@type': 'SoftwareApplication',
		name: 'Bash Simulator WASM',
		alternateName: 'Bash Online Terminal',
		operatingSystem: 'Any browser with WebAssembly support',
		applicationCategory: 'DeveloperApplication',
		description:
			'A robust, interactive Bash shell simulator built with Go and compiled to WebAssembly. Features include coreutils parity, filesystem simulation, and interactive command execution.',
		url: seoData.url,
		image: seoData.image,
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

	const breadcrumbData = {
		'@context': 'https://schema.org',
		'@type': 'BreadcrumbList',
		itemListElement: [
			{
				'@type': 'ListItem',
				position: 1,
				name: 'Home',
				item: seoData.url
			}
		]
	};

	const websiteData = {
		'@context': 'https://schema.org',
		'@type': 'WebSite',
		name: 'Bash Simulator WASM',
		url: seoData.url,
		potentialAction: {
			'@type': 'SearchAction',
			target: {
				'@type': 'EntryPoint',
				urlTemplate: `${seoData.url}?q={search_term_string}`
			},
			'query-input': 'required name=search_term_string'
		}
	};

	const faqData = {
		'@context': 'https://schema.org',
		'@type': 'FAQPage',
		mainEntity: [
			{
				'@type': 'Question',
				name: 'What is Bash Simulator Online?',
				acceptedAnswer: {
					'@type': 'Answer',
					text: 'It is a full-featured Bash shell simulator developed in Go and running entirely in the browser using WebAssembly (WASM). It allows you to practice Linux commands without installing anything.'
				}
			},
			{
				'@type': 'Question',
				name: 'How do I use this interactive Bash shell?',
				acceptedAnswer: {
					'@type': 'Answer',
					text: 'Simply open the website and start typing standard Unix commands like ls, cd, mkdir, and echo. The environment simulates a real filesystem persistent within your browser session.'
				}
			},
			{
				'@type': 'Question',
				name: 'Is this a real Linux kernel running in the browser?',
				acceptedAnswer: {
					'@type': 'Answer',
					text: 'No, it is a high-fidelity simulation of the Bash shell and GNU Coreutils running on a virtual filesystem in WebAssembly, optimized for speed and safety.'
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
	{@html `<script type="application/ld+json">${JSON.stringify(breadcrumbData)}</script>`}
	{@html `<script type="application/ld+json">${JSON.stringify(websiteData)}</script>`}
</svelte:head>

<!-- Visually hidden semantic content for SEO and AEO -->
<div class="sr-only">
	<h1>{seoData.title}</h1>
	<p>{seoData.description}</p>
	<section>
		<h2>Interactive Bash Shell Simulator</h2>
		<p>
			This online terminal provides a high-fidelity simulation of the GNU Bash shell. It is
			designed for developers, students, and DevOps engineers who want to practice Linux
			commands or test scripts in a safe, sandboxed environment directly in their browser.
		</p>
	</section>
	<section>
		<h2>Advanced Technology Stack: Go & WebAssembly (WASM)</h2>
		<p>
			Built with the performance of Golang and the portability of WebAssembly, this simulator
			runs at near-native speeds. The entire shell logic and core utilities are compiled to a
			main.wasm file, which is executed by the browser without any backend server requirements.
		</p>
		<ul>
			<li><strong>True Bash Parity:</strong> Implements core Bash features and syntax.</li>
			<li><strong>Virtual Filesystem:</strong> A fully functional in-memory filesystem (MEMFS).</li>
			<li><strong>GNU Coreutils:</strong> Includes common commands like ls, cd, cat, grep, and more.</li>
			<li><strong>Fast & Responsive:</strong> Instant startup and low-latency interaction.</li>
			<li><strong>Secure Sandbox:</strong> Runs entirely client-side; your data never leaves the browser.</li>
		</ul>
	</section>
	<section>
		<h2>Use Cases for Bash Online</h2>
		<ul>
			<li>Learning Unix/Linux command line basics.</li>
			<li>Testing shell scripts without a virtual machine.</li>
			<li>Quickly running coreutils commands on any device.</li>
			<li>DevOps interview preparation and practice.</li>
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
