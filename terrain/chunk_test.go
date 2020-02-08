package terrain

import (
	"fmt"
	"testing"
)

func TestChunkID(t *testing.T) {
	c := Chunk{
		X: 1,
		Y: 2,
	}

	if c.ID() != 0x0000000200000001 {
		t.Errorf("c.ID() = %x; want %x", c.ID(), 0x0000000200000001)
	}
}

func TestNewChunk(t *testing.T) {
	var x, y int32 = 0, 0
	chunk := NewChunk(x, y)

	if chunk.X != x {
		t.Errorf("chunk.X = %x; want %x", chunk.X, x)
	}
	if chunk.Y != y {
		t.Errorf("chunk.Y = %x; want %x", chunk.Y, y)
	}

	for xi := 0; xi < 32; xi++ {
		fmt.Println(chunk.Data[xi])
	}
}

func TestNewChunkFromData(t *testing.T) {
	x := int32(1)
	y := int32(2)
	data := [32][32]Tile{[32]Tile{1, 2, 3}}
	modified := true

	c := NewChunkFromData(x, y, data, modified)

	if c.X != x {
		t.Errorf("c.X = %x; want %x", c.X, x)
	}
	if c.Y != y {
		t.Errorf("c.Y = %x; want %x", c.Y, y)
	}
	if c.Data != data {
		t.Errorf("c.Data = %x; want %x", c.Data, data)
	}
	if c.Modified != modified {
		t.Errorf("c.Modified = %t; want %t", c.Modified, modified)
	}
}
