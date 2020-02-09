package terrain

//Chunk represents a 32x32 chunk of game world
type Chunk struct {
	X        int32
	Y        int32
	Data     [32][32]Tile
	Modified bool
}

//NewChunk generates a new chunk using the default generator
func NewChunk(x int32, y int32) *Chunk {
	return DefaultGenerator.NewChunk(x, y)
}

//NewChunkFromData creates a new Chunk based
func NewChunkFromData(x int32, y int32, data [32][32]Tile, modified bool) *Chunk {
	return &Chunk{
		X:        x,
		Y:        y,
		Data:     data,
		Modified: modified,
	}
}

//ID returns a unique identifier for the chunk
func (c *Chunk) ID() uint64 {
	//Stores c.Y in the 32 most significant bits and c.X in the 32 least significant bits.
	//This can then be used as a unique identifier for each chunk.
	x, y := uint32(c.X), uint32(c.Y)
	return uint64(y)<<32 | uint64(x)
}

//RawData returns the data as an array of bytes
func (c *Chunk) RawData() (val [32 * 32]byte) {
	for xi := 0; xi < 32; xi++ {
		for yi := 0; yi < 32; yi++ {
			val[xi*32+yi] = byte(c.Data[xi][yi])
		}
	}

	return val
}

//ChunkIDToPos map a chunk ID into x, y coordinates
func ChunkIDToPos(id uint64) (x, y int32) {
	return int32(id & 0xFFFFFFFF), int32(id >> 32)
}
