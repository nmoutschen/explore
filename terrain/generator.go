package terrain

import (
	"github.com/ojrac/opensimplex-go"
)

//Generator represents a terrain generator
type Generator struct {
	seed  int64
	noise opensimplex.Noise
}

const (
	heightPass1Scale   float64 = 64
	heightPass2Scale   float64 = 16
	heightPass3Scale   float64 = 4
	heightPass4Scale   float64 = 1
	heightPass1OffsetX float64 = 0
	heightPass1OffsetY float64 = 0
	heightPass2OffsetX float64 = 0x0F8000
	heightPass2OffsetY float64 = 0x7FFFFF
	heightPass3OffsetX float64 = 0x19FFFF
	heightPass3OffsetY float64 = -0xAF0000
	heightPass4OffsetX float64 = 0
	heightPass4OffsetY float64 = 0
	heightPass1Weight  float64 = 1.
	heightPass2Weight  float64 = 1. / 4
	heightPass3Weight  float64 = 1. / 16
	heightPass4Weight  float64 = 1. / 16

	humidityPass1Scale   float64 = 256
	humidityPass2Scale   float64 = 64
	humidityPass3Scale   float64 = 4
	humidityPass1OffsetX float64 = 0
	humidityPass1OffsetY float64 = 0
	humidityPass2OffsetX float64 = 0x7FFFFFFF
	humidityPass2OffsetY float64 = 0x7FFFFFFF
	humidityPass3OffsetX float64 = 0x19FFFFFF
	humidityPass3OffsetY float64 = 0xC9FFFFFF
	humidityPass1Weight  float64 = 1.
	humidityPass2Weight  float64 = 1. / 2
	humidityPass3Weight  float64 = 1. / 4
)

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
	val1 := (g.noise.Eval2(x/heightPass1Scale+heightPass1OffsetX, y/heightPass1Scale+heightPass1OffsetY)/2 + 0.5) * heightPass1Weight
	val2 := (g.noise.Eval2(x/heightPass2Scale+heightPass2OffsetX, y/heightPass2Scale+heightPass2OffsetY)/2 + 0.5) * heightPass2Weight
	val3 := (g.noise.Eval2(x/heightPass3Scale+heightPass3OffsetX, y/heightPass3Scale+heightPass3OffsetY)/2 + 0.5) * heightPass3Weight
	val4 := (g.noise.Eval2(x/heightPass4Scale+heightPass4OffsetX, y/heightPass4Scale+heightPass4OffsetX)/2 + 0.5) * heightPass4Weight

	val = (val1 + val2 + val3 + val4) / (heightPass1Weight + heightPass2Weight + heightPass3Weight + heightPass4Weight)
	return val
}

func (g *Generator) getHumidity(x, y float64) (val float64) {
	val1 := (g.noise.Eval2(x/humidityPass1Scale+humidityPass1OffsetX, y/humidityPass1Scale+humidityPass1OffsetY)/2 + 0.5) * humidityPass1Weight
	val2 := (g.noise.Eval2(x/humidityPass2Scale+humidityPass2OffsetX, y/humidityPass2Scale+humidityPass2OffsetY)/2 + 0.5) * humidityPass2Weight
	val3 := (g.noise.Eval2(x/humidityPass3Scale+humidityPass3OffsetX, y/humidityPass3Scale+humidityPass3OffsetY)/2 + 0.5) * humidityPass3Weight

	val = (val1 + val2 + val3) / (humidityPass1Weight + humidityPass2Weight + humidityPass3Weight)
	return val
}

//getTile returns a tile at the given x, y coordinate
func (g *Generator) getTile(x, y float64) Tile {
	height := g.getHeight(x, y)
	humidity := g.getHumidity(x, y)

	return GetTile(height, humidity)
}

//NewChunk returns a new chunk using the provided generator
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
var DefaultGenerator *Generator = NewGenerator(0.0)
