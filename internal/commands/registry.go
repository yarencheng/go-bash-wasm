package commands

import (
	"fmt"
	"sync"
)

// Registry maintains a mapping of command names to their handlers.
type Registry struct {
	mu       sync.RWMutex
	commands map[string]Command
	disabled map[string]bool
}

func New() *Registry {
	return &Registry{
		commands: make(map[string]Command),
		disabled: make(map[string]bool),
	}
}

// Register adds a command to the registry.
func (r *Registry) Register(cmd Command) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	name := cmd.Name()
	if _, exists := r.commands[name]; exists {
		return fmt.Errorf("command already registered: %s", name)
	}

	r.commands[name] = cmd
	return nil
}

// Get retrieves a command from the registry by name.
func (r *Registry) Get(name string) (Command, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	cmd, exists := r.commands[name]
	if !exists {
		return nil, false
	}
	isDisabled := r.disabled[name]
	return cmd, !isDisabled
}

// List returns all registered command names.
func (r *Registry) List() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	names := make([]string, 0, len(r.commands))
	for name := range r.commands {
		names = append(names, name)
	}
	return names
}

// Enable enables a command.
func (r *Registry) Enable(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.commands[name]; !exists {
		return fmt.Errorf("command not found: %s", name)
	}
	r.disabled[name] = false
	return nil
}

// Disable disables a command.
func (r *Registry) Disable(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.commands[name]; !exists {
		return fmt.Errorf("command not found: %s", name)
	}
	r.disabled[name] = true
	return nil
}

// IsEnabled checks if a command is enabled.
func (r *Registry) IsEnabled(name string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if _, exists := r.commands[name]; !exists {
		return false
	}
	return !r.disabled[name]
}
