package terrain

//Chunk represents a 32x32 chunk of game world
type Chunk struct {
	X        int32
	Y        int32
	Data     [32][32]Tile
	Modified bool
}

//NewChunk generates a new chunk using the default generator
//
//See Generator.NewChunk to understand how the chunk generation algorithm works.
func NewChunk(x int32, y int32) *Chunk {
	return DefaultGenerator.NewChunk(x, y)
}

//NewChunkFromData creates a new Chunk based on existing data rather than using the terrain generator
//
//Use this function if you have pre-generated a chunk and recover it from storage, or if you have user modification for that chunk.
func NewChunkFromData(x int32, y int32, data [32][32]Tile, modified bool) *Chunk {
	return &Chunk{
		X:        x,
		Y:        y,
		Data:     data,
		Modified: modified,
	}
}

//ID returns a unique identifier for the chunk
//
//See ChunkPosToID for more information about the transformation applied.
func (c *Chunk) ID() uint64 {
	return ChunkPosToID(c.X, c.Y)
}

//RawData returns the data as an array of bytes
//
//This is used for transmitting data about a specific chunk in a more efficient format. A JSON representation of a 32x32 2-dimensional
//array with each value containing 1 byte would take between 2113 and 4161 bytes, while this can be sent with 1368 bytes in base64 format.
func (c *Chunk) RawData() (val [32 * 32]byte) {
	for xi := 0; xi < 32; xi++ {
		for yi := 0; yi < 32; yi++ {
			val[xi*32+yi] = byte(c.Data[xi][yi])
		}
	}

	return val
}

//ChunkIDToPos map a chunk ID into x, y coordinates
//
//This is the reverse of ChunkPosToID. See ChunkPosToID for more information about the transformation applied.
func ChunkIDToPos(id uint64) (x, y int32) {
	return int32(id & 0xFFFFFFFF), int32(id >> 32)
}

//ChunkPosToID map x, y coordinates to an unique identifier
//
//This transforms the bits from x and y coordinates into a 64 bits unsigned integer. The y position is stored in the 32 most significant
//bits while the x position is stored in the 32 least significant bit.
//
//If any coordinate is negative, this uses the 2-complement representation. For example, (x, y) coordinates (0, -1) map to an identifier
//of 0xFFFF_FFFF_0000_0000.
//
//This is the reverse of ChunkIDToPos
func ChunkPosToID(x, y int32) (id uint64) {
	xu, yu := uint32(x), uint32(y)
	return uint64(yu)<<32 | uint64(xu)
}
