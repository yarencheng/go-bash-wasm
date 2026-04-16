<script lang="ts">
	import { Terminal } from '@xterm/xterm';
	import { FitAddon } from '@xterm/addon-fit';
	import '@xterm/xterm/css/xterm.css';

	let terminalElement: HTMLDivElement;
	let term: Terminal;
	let fitAddon: FitAddon;

	$effect(() => {
		if (typeof window !== 'undefined' && terminalElement) {
			term = new Terminal({
				cursorBlink: true,
				theme: {
					background: '#0a0a0a',
					foreground: '#f0f0f0',
					cursor: '#f0f0f0',
					selectionBackground: 'rgba(255, 255, 255, 0.3)',
					black: '#000000',
					red: '#ff5555',
					green: '#50fa7b',
					yellow: '#f1fa8c',
					blue: '#bd93f9',
					magenta: '#ff79c6',
					cyan: '#8be9fd',
					white: '#bfbfbf',
					brightBlack: '#4d4d4d',
					brightRed: '#ff6e67',
					brightGreen: '#5af78e',
					brightYellow: '#f4f99d',
					brightBlue: '#caa9fa',
					brightMagenta: '#ff92d0',
					brightCyan: '#9aedfe',
					brightWhite: '#e6e6e6'
				},
				fontFamily: 'Menlo, Monaco, "Courier New", monospace',
				fontSize: 14,
				lineHeight: 1.2
			});

			fitAddon = new FitAddon();
			term.loadAddon(fitAddon);
			term.open(terminalElement);
			fitAddon.fit();

			term.writeln('\x1b[32mWelcome to Go-Bash-WASM Terminal!\x1b[0m');
			term.writeln('Type something to begin...');
			term.write('\r\n$ ');

			term.onData((data) => {
				const code = data.charCodeAt(0);
				if (code === 13) {
					// Carriage return
					term.write('\r\n$ ');
				} else if (code < 32) {
					// Control characters
					return;
				} else {
					term.write(data);
				}
			});

			const handleResize = () => {
				fitAddon.fit();
			};

			window.addEventListener('resize', handleResize);

			return () => {
				window.removeEventListener('resize', handleResize);
				term.dispose();
			};
		}
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
