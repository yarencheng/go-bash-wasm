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
			lineHeight: 1.5
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

				// go.run() returns a promise but we do NOT await it here.
				go.run(result.instance);

				// Give the Go runtime a tick to initialise and register writeStdin.
				await new Promise((r) => setTimeout(r, 150));

				wasmReady = true;
				term.writeln('\x1b[32mReady! Type a command and press Enter.\x1b[0m\r\n');
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
		title: 'Bash Simulator WASM - Interactive Unix Shell in Browser',
		description:
			'Experience a full-featured Bash shell simulator powered by Go and WebAssembly. Run commands and explore a Linux-like environment directly in your browser.',
		url: 'https://bash.devops-playground.dev/',
		image: 'https://bash.devops-playground.dev/social-preview.png',
		twitterHandle: '@yarencheng'
	};

	const schemaData = {
		'@context': 'https://schema.org',
		'@type': 'SoftwareApplication',
		name: 'Bash Simulator WASM',
		operatingSystem: 'Web Browser',
		applicationCategory: 'DeveloperApplication',
		description: 'A high-performance Bash shell simulator running in WebAssembly.',
		offers: {
			'@type': 'Offer',
			price: '0',
			priceCurrency: 'USD'
		}
	};
</script>

<svelte:head>
	<title>{seoData.title}</title>
	<meta name="description" content={seoData.description} />
	<link rel="canonical" href={seoData.url} />

	<!-- Open Graph / Facebook -->
	<meta property="og:type" content="website" />
	<meta property="og:url" content={seoData.url} />
	<meta property="og:title" content={seoData.title} />
	<meta property="og:description" content={seoData.description} />
	<meta property="og:image" content={seoData.image} />

	<!-- Twitter -->
	<meta property="twitter:card" content="summary_large_image" />
	<meta property="twitter:url" content={seoData.url} />
	<meta property="twitter:title" content={seoData.title} />
	<meta property="twitter:description" content={seoData.description} />
	<meta property="twitter:image" content={seoData.image} />
	<meta name="twitter:creator" content={seoData.twitterHandle} />

	<!-- Structured Data -->
	{@html `<script type="application/ld+json">${JSON.stringify(schemaData)}</script>`}
</svelte:head>

<div class="terminal-container" bind:this={terminalElement}></div>

<style>
	.terminal-container {
		width: 100vw;
		height: 100vh;
		background-color: #0a0a0a;
		padding: 1rem;
		box-sizing: border-box;
	}

	:global(.xterm) {
		padding: 10px;
	}

	:global(.xterm-viewport) {
		background-color: transparent !important;
	}
</style>
