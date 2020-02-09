package terrain

import (
	"fmt"
	"testing"
)

func BenchmarkNewChunk(b *testing.B) {
	var y int32

	for x := int32(0); x < int32(b.N); x++ {
		NewChunk(x, y)
	}
}

func ExampleNewChunk() {
	var x, y int32

	x = 0x00000010
	y = 0x0000FF8E

	chunk := NewChunk(x, y)

	fmt.Printf("Pos: (%x, %x)\n", chunk.X, chunk.Y)
	for xi := 0; xi < 32; xi++ {
		fmt.Println(chunk.Data[xi])
	}
	//Output:
	//Pos: (10, ff8e)
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 3 3 3 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 3 3 4 3]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 3 4 3]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 3 3 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 4 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 3 3 3 3 3 4 4 4 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 3 3 4 4 4 4 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 3 3 4 4 4 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 4 4 4 4 4 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 4 4 4 4 4 4 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 4 4 4 4 4 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 4 4 4 4 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 4 4 4 4 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 4 4 4 4 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 4 4 4 4 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 4 4 4 4 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 4 4 4 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 4 4 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 4 4 3 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 4 4 4 4]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 3 4 3 3]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 4 3 3 3]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 4 3 3]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3]
	//[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3]
}

func TestNewChunk(t *testing.T) {
	testCases := []struct {
		X, Y int32
		ID   uint64
	}{
		{0x00000000, 0x00000000, 0x0000_0000_0000_0000},
		//+1
		{0x00000001, 0x00000000, 0x0000_0000_0000_0001},
		{0x00000000, 0x00000001, 0x0000_0001_0000_0000},
		{0x00000001, 0x00000001, 0x0000_0001_0000_0001},
		//-1
		{-0x00000001, 0x00000000, 0x0000_0000_FFFF_FFFF},
		{0x00000000, -0x00000001, 0xFFFF_FFFF_0000_0000},
		{-0x00000001, -0x00000001, 0xFFFF_FFFF_FFFF_FFFF},
		//Max int32
		{0x7FFFFFFF, 0x00000000, 0x0000_0000_7FFF_FFFF},
		{0x00000000, 0x7FFFFFFF, 0x7FFF_FFFF_0000_0000},
		{0x7FFFFFFF, 0x7FFFFFFF, 0x7FFF_FFFF_7FFF_FFFF},
		//Min int32
		{-0x80000000, 0x00000000, 0x0000_0000_8000_0000},
		{0x00000000, -0x80000000, 0x8000_0000_0000_0000},
		{-0x80000000, -0x80000000, 0x8000_0000_8000_0000},
	}

	for i, tc := range testCases {
		chunk := NewChunk(tc.X, tc.Y)

		if chunk.X != tc.X {
			t.Errorf("chunk.X = %08x; want %08x for test case %d", chunk.X, tc.X, i)
		}
		if chunk.Y != tc.Y {
			t.Errorf("chunk.Y = %08x; want %08x for test case %d", chunk.Y, tc.Y, i)
		}
		if chunk.ID() != tc.ID {
			t.Errorf("chunk.ID() = %016x; want %016x for test case %d", chunk.ID(), tc.ID, i)
		}
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

func TestChunkID(t *testing.T) {
	c := Chunk{
		X: 1,
		Y: 2,
	}

	if c.ID() != 0x0000000200000001 {
		t.Errorf("c.ID() = %x; want %x", c.ID(), 0x0000000200000001)
	}
}

func TestChunkRawData(t *testing.T) {
	c := NewChunk(0, 0)

	raw := c.RawData()

	for xi := 0; xi < 32; xi++ {
		for yi := 0; yi < 32; yi++ {
			if byte(c.Data[xi][yi]) != raw[xi*32+yi] {
				t.Errorf("raw[%d] = %x; want %x", xi*32+yi, raw[xi*32+yi], c.Data[xi][yi])
			}
		}
	}
}

func TestChunkIDToPos(t *testing.T) {
	testCases := []struct {
		X, Y int32
		ID   uint64
	}{
		{0x00000000, 0x00000000, 0x0000_0000_0000_0000},
		//+1
		{0x00000001, 0x00000000, 0x0000_0000_0000_0001},
		{0x00000000, 0x00000001, 0x0000_0001_0000_0000},
		{0x00000001, 0x00000001, 0x0000_0001_0000_0001},
		//-1
		{-0x00000001, 0x00000000, 0x0000_0000_FFFF_FFFF},
		{0x00000000, -0x00000001, 0xFFFF_FFFF_0000_0000},
		{-0x00000001, -0x00000001, 0xFFFF_FFFF_FFFF_FFFF},
		//Max int32
		{0x7FFFFFFF, 0x00000000, 0x0000_0000_7FFF_FFFF},
		{0x00000000, 0x7FFFFFFF, 0x7FFF_FFFF_0000_0000},
		{0x7FFFFFFF, 0x7FFFFFFF, 0x7FFF_FFFF_7FFF_FFFF},
		//Min int32
		{-0x80000000, 0x00000000, 0x0000_0000_8000_0000},
		{0x00000000, -0x80000000, 0x8000_0000_0000_0000},
		{-0x80000000, -0x80000000, 0x8000_0000_8000_0000},
	}

	for i, tc := range testCases {
		x, y := ChunkIDToPos(tc.ID)
		if tc.X != x || tc.Y != y {
			t.Errorf("ChunkIDToPos(%016x) = (%08x, %08x); want (%08x, %08x) for test case %d", tc.ID, x, y, tc.X, tc.Y, i)
		}
	}
}
