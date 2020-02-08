package terrain

//Chunk represents a 32x32 chunk of game world
type Chunk struct {
	X        int32        "json:x"
	Y        int32        "json:y"
	Data     [32][32]Tile "json:data"
	Modified bool         "json:modified"
}

//ID returns a unique identifier for the chunk
func (c *Chunk) ID() int64 {
	//Stores c.Y in the 32 most significant bits and c.X in the 32 least significant bits.
	//This can then be used as a unique identifier for each chunk.
	return int64(c.Y)<<32 | int64(c.X)
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
