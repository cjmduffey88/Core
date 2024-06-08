package Core

import "image"

// Core is the interface that the emulator core will have to implement
type Core interface {
	Step()                         // Step all components by one clock cycle
	Render() *image.Image          // Return the current frame buffer
	ClockSpeed() uint              // Return the clock speed of the core
	Init(pathBIOS, pathROM string) // Initialize the core with a BIOS (if any) and ROM
	ROMState() string              // Return the current state of the ROM
	CPUState() string              // Return the current state of the CPU
	APUState() string              // Return the current state of the APU
	GPUState() string              // Return the current state of the GPU
	InputState() string            // Return the current state of the Gamepad/Keyboard
	Reset()                        // Reset the core
	SaveState(path string)         // Save the current state of the core
	LoadState(path string)         // Load a state into the core
	Peek(address uint) uint        // Peek at a memory address (Read) (core will have to convert to and from uint)
	Poke(address uint, value uint) // Poke a memory address (Write) (core will have to convert to and from uint)
}
