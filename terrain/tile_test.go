package terrain

import (
	"testing"
)

func TestGetTile(t *testing.T) {
	testCases := []struct {
		Height   float64
		Humidity float64
		Tile     Tile
	}{
		{0.0, 0.0, tileWater},
		{1.0, 0.0, tileScorched},
		{1.0, 1.0, tileFrozenWater},
		{0.0, 1.0, tileWater},
	}

	for _, tc := range testCases {
		tile := GetTile(tc.Height, tc.Humidity)
		if tile != tc.Tile {
			t.Errorf("GetTile(%.02f, %.02f) = %x; want %x", tc.Height, tc.Humidity, tile, tc.Tile)
		}
	}
}

func BenchmarkGetTileBestCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetTile(0.0, 0.0)
	}
}

func BenchmarkGetTileWorstCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetTile(1.0, 1.0)
	}
}

func BenchmarkGetTileAvgCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetTile(0.5, 0.5)
	}
}
