package terrain

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestTerrain(t *testing.T) {
	size := 1024

	lookup := make(map[Tile]color.NRGBA)
	lookup[tileWater] = color.NRGBA{0x00, 0x77, 0xbe, 0xff}
	lookup[tileFrozenWater] = color.NRGBA{0xaa, 0xdf, 0xff, 0xff}
	lookup[tileTropicalRainForest] = color.NRGBA{0x00, 0x42, 0x35, 0xff}
	lookup[tileTropicalForest] = color.NRGBA{0x00, 0x75, 0x5e, 0xff}
	lookup[tileGrassland] = color.NRGBA{0x26, 0x8b, 0x07, 0xff}
	lookup[tileDesert] = color.NRGBA{0xc2, 0xb2, 0x80, 0xff}
	lookup[tileTemperateRainForest] = color.NRGBA{0x1a, 0x37, 0x36, 0xff}
	lookup[tileTemperateForest] = color.NRGBA{0x1a, 0x47, 0x46, 0xff}
	lookup[tileTemperateDesert] = color.NRGBA{0xc2, 0xb2, 0x80, 0xff}
	lookup[tileTaiga] = color.NRGBA{0xcd, 0xd6, 0x81, 0xff}
	lookup[tileShrubland] = color.NRGBA{0x26, 0x6b, 0x07, 0xff}
	lookup[tileSnow] = color.NRGBA{0xff, 0xff, 0xff, 0xff}
	lookup[tileTundra] = color.NRGBA{0xb9, 0xbb, 0xdf, 0xff}
	lookup[tileBare] = color.NRGBA{0xa0, 0xa0, 0xa0, 0xff}
	lookup[tileScorched] = color.NRGBA{0x46, 0x40, 0x3f, 0xff}

	g := NewGenerator(0)

	img := image.NewNRGBA(image.Rect(0, 0, size, size))

	for x := 0; x < size/32; x++ {
		for y := 0; y < size/32; y++ {
			chunk := g.NewChunk(int32(x), int32(y))

			for xp := 0; xp < 32; xp++ {
				for yp := 0; yp < 32; yp++ {
					img.Set(x*32+xp, y*32+yp, lookup[chunk.Data[xp][yp]])
				}
			}
		}
	}

	f, _ := os.Create("terrain.png")
	png.Encode(f, img)
	f.Close()
}
