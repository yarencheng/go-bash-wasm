<script lang="ts">
	import { Terminal } from '@xterm/xterm';
	import { FitAddon } from '@xterm/addon-fit';
	import '@xterm/xterm/css/xterm.css';

	let terminalElement: HTMLDivElement;
	let term: Terminal;
	let fitAddon: FitAddon;
	let wasmReady = $state(false);
	let lineBuffer = '';

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

		// 4. Wire Terminal Data to WASM Stdin
		term.onData((data) => {
			if (!wasmReady) return;

			for (let i = 0; i < data.length; i++) {
				const char = data[i];
				const code = data.charCodeAt(i);

				switch (code) {
					case 13: // Enter
						term.write('\r\n');
						if (typeof (window as any).writeStdin === 'function') {
							(window as any).writeStdin(lineBuffer);
						}
						lineBuffer = '';
						break;
					case 127: // Backspace
					case 8:
						if (lineBuffer.length > 0) {
							lineBuffer = lineBuffer.slice(0, -1);
							term.write('\b \b');
						}
						break;
					case 4: // Ctrl+D
						if (typeof (window as any).closeStdin === 'function') {
							(window as any).closeStdin();
						}
						break;
					case 3: // Ctrl+C
						term.write('^C\r\n');
						lineBuffer = '';
						break;
					default:
						if (code >= 32) {
							lineBuffer += char;
							term.write(char);
						}
						break;
				}
			}
		});

		const handleResize = () => fitAddon.fit();
		window.addEventListener('resize', handleResize);

		return () => {
			window.removeEventListener('resize', handleResize);
			term.dispose();
		};
	});
</script>

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
