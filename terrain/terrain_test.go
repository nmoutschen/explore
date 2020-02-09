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

	lookup := map[Tile]color.NRGBA{
		BiomeWater:               color.NRGBA{0x00, 0x77, 0xbe, 0xff},
		BiomeFrozenWater:         color.NRGBA{0xaa, 0xdf, 0xff, 0xff},
		BiomeTropicalRainForest:  color.NRGBA{0x00, 0x42, 0x35, 0xff},
		BiomeTropicalForest:      color.NRGBA{0x00, 0x75, 0x5e, 0xff},
		BiomeGrassland:           color.NRGBA{0x26, 0x8b, 0x07, 0xff},
		BiomeDesert:              color.NRGBA{0xc2, 0xb2, 0x80, 0xff},
		BiomeTemperateRainForest: color.NRGBA{0x1a, 0x37, 0x36, 0xff},
		BiomeTemperateForest:     color.NRGBA{0x1a, 0x47, 0x46, 0xff},
		BiomeTemperateDesert:     color.NRGBA{0xc2, 0xb2, 0x80, 0xff},
		BiomeTaiga:               color.NRGBA{0xcd, 0xd6, 0x81, 0xff},
		BiomeShrubland:           color.NRGBA{0x26, 0x6b, 0x07, 0xff},
		BiomeSnow:                color.NRGBA{0xff, 0xff, 0xff, 0xff},
		BiomeTundra:              color.NRGBA{0xb9, 0xbb, 0xdf, 0xff},
		BiomeBare:                color.NRGBA{0xa0, 0xa0, 0xa0, 0xff},
		BiomeScorched:            color.NRGBA{0x46, 0x40, 0x3f, 0xff},
	}

	g := NewGenerator(0)

	img := image.NewNRGBA(image.Rect(0, 0, size, size))

	for x := 0; x < size/32; x++ {
		for y := 0; y < size/32; y++ {
			chunk := g.NewChunk(int32(x), int32(y))

			for xp := 0; xp < 32; xp++ {
				for yp := 0; yp < 32; yp++ {
					col := lookup[chunk.Data[xp][yp]&0x0F]
					if chunk.Data[xp][yp]&StructureTree == StructureTree {
						col.G -= 0x20
					} else if chunk.Data[xp][yp]&StructurePlant == StructurePlant {
						col.G -= 0x10
						col.R += 0x10
					}
					img.Set(x*32+xp, y*32+yp, col)
				}
			}
		}
	}

	f, _ := os.Create("terrain.png")
	png.Encode(f, img)
	f.Close()
}
