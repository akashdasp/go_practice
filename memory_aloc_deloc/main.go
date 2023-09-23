package main

import (
	"fmt"
	"unsafe"
)

type MemoryBlock struct {
	Size uintptr
	Next *MemoryBlock
}

type MemoryAllocator struct {
	Head *MemoryBlock
}

func NewMemoryAllocator() *MemoryAllocator {
	return &MemoryAllocator{}
}

func (ma *MemoryAllocator) Allocate(size uintptr) unsafe.Pointer {
	if size == 0 {
		return nil
	}

	alignedSize := (size + unsafe.Sizeof(MemoryBlock{})) &^ (unsafe.Sizeof(uintptr(0)) - 1)

	newBlock := &MemoryBlock{
		Size: alignedSize,
		Next: ma.Head,
	}

	blockPtr := unsafe.Pointer(newBlock)

	ma.Head = newBlock

	// Return a pointer to the allocated memory (skip block header)
	return unsafe.Pointer(uintptr(blockPtr) + unsafe.Sizeof(MemoryBlock{}))
}

func (ma *MemoryAllocator) Deallocate(ptr unsafe.Pointer) {
	if ptr == nil {
		return
	}

	// Move the pointer back to the block header
	blockPtr := unsafe.Pointer(uintptr(ptr) - unsafe.Sizeof(MemoryBlock{}))

	// Remove the block from the list
	if ma.Head == (*MemoryBlock)(blockPtr) {
		ma.Head = ma.Head.Next
		return
	}

	current := ma.Head
	for current.Next != nil {
		if current.Next == (*MemoryBlock)(blockPtr) {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}

	// The block was not found in the list; this is an error.
	panic("Deallocating unknown block.")
}

func main() {
	allocator := NewMemoryAllocator()

	// Allocate memory
	ptr1 := allocator.Allocate(32)
	ptr2 := allocator.Allocate(64)

	// Deallocate memory
	allocator.Deallocate(ptr1)
	allocator.Deallocate(ptr2)

	// Attempting to deallocate an unknown block will panic
	// allocator.Deallocate(ptr1) // This will panic

	// Allocate more memory
	ptr3 := allocator.Allocate(128)

	// Clean up
	allocator.Deallocate(ptr3)

	fmt.Println("Memory allocation and deallocation completed.")
}
