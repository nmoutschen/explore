package terrain

import (
	"github.com/ojrac/opensimplex-go"
)

//Generator represents a terrain generator with a specific seed
type Generator struct {
	seed  int64
	noise opensimplex.Noise
}

//genPass represents details for a single map generation pass
type genPass struct {
	Scale   float64
	Weight  float64
	OffsetX float64
	OffsetY float64
}

//heightPasses contain details about the passes for generating the height map
var heightPasses [4]genPass = [4]genPass{
	{64., 1., 0., 0.},
	{16., 1. / 4, 0x0F8000, 0x7FFFFF},
	{4., 1. / 16, 0x19FFFF, -0xAF0000},
	{1., 1. / 16, 0, 0},
}

//humidityPasses contain details about the passes for generating the humidity map
var humidityPasses [4]genPass = [4]genPass{
	{256., 1., 0., 0.},
	{64., 1. / 4, 0x7FFFFF, 0x7FFFFF},
	{8., 1. / 8, 0x19FFFF, 0xC9FFFF},
	{1., 1. / 16, -0x39FFFF, 0xA9FFFF},
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
	var weight float64
	for _, pass := range heightPasses {
		val += (g.noise.Eval2(x/pass.Scale+pass.OffsetX, y/pass.Scale+pass.OffsetY)/2 + 0.5) * pass.Weight
		weight += pass.Weight
	}

	return val / weight
}

func (g *Generator) getHumidity(x, y float64) (val float64) {
	var weight float64
	for _, pass := range humidityPasses {
		val += (g.noise.Eval2(x/pass.Scale+pass.OffsetX, y/pass.Scale+pass.OffsetY)/2 + 0.5) * pass.Weight
		weight += pass.Weight
	}

	return val / weight
}

//getTile returns a tile at the given x, y coordinate
func (g *Generator) getTile(x, y float64) Tile {
	height := g.getHeight(x, y)
	humidity := g.getHumidity(x, y)

	return GetTile(x, y, height, humidity)
}

//NewChunk generates a new chunk
//
//This uses OpenSimplex noise to generate a height and humidity map for all 32*32 tiles in the chunk, which are then looked up in
//threshold arrays to find the biome corresponding to the tile.
func (g *Generator) NewChunk(x int32, y int32) *Chunk {
	var data [32][32]Tile

	for xi := 0; xi < 32; xi++ {
		for yi := 0; yi < 32; yi++ {
			data[xi][yi] = g.getTile(float64(x)*32+float64(xi), float64(y)*32+float64(yi))
		}
	}

	return &Chunk{
		X:    x,
		Y:    y,
		Data: data,
	}
}

//DefaultGenerator is a default generator
//
//This generator is initalized with a seed of 0.0 to ensure consistent results.
var DefaultGenerator *Generator = NewGenerator(0.0)
