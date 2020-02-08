package terrain

import (
	"github.com/ojrac/opensimplex-go"
)

//Generator represents a terrain generator
type Generator struct {
	seed  int64
	noise opensimplex.Noise
}

//NewGenerator creates a new generator with the given seed
func NewGenerator(seed int64) *Generator {
	g := &Generator{
		seed:  seed,
		noise: opensimplex.New(seed),
	}

	return g
}

//getHeight returns the height at a specific position
func (g *Generator) getHeight(x, y float64) (val float64) {
	val = g.noise.Eval2(x, y)

	//Normalize to [0;1]
	val = val/2 + 0.5
	return val
}

//getTile returns a tile at the given x, y coordinate
func (g *Generator) getTile(x, y float64) Tile {
	height := g.getHeight(x, y)
	//TODO make the humidity dynamic
	return GetTile(height, 0.0)
}

//NewChunk returns a new chunk using the provided generator
func (g *Generator) NewChunk(x int32, y int32) *Chunk {
	var data [32][32]Tile
	for xi := 0; xi < 32; xi++ {
		for yi := 0; yi < 32; yi++ {
			data[xi][yi] = g.getTile(float64(x)+float64(xi)/32, float64(y)+float64(yi)/32)
		}
	}

	return &Chunk{
		X:    x,
		Y:    y,
		Data: data,
	}
}

//DefaultGenerator is a default generator
var DefaultGenerator *Generator = NewGenerator(0.0)
