package terrain

//Tile represents the different types of tiles that can be used in game
//
//Tile types are based on the Whittaker classification with some modifications.
//
// | Height\Moisture | 6 (wet)      | 5                     | 4                    | 3                | 2               | 1                | 0 (dry)          |
// | 4 (high)        | Frozen Water | Snow                  | Snow                 | Snow             | Tundra          | Bare             | Scorched         |
// | 3               | Frozen Water | Taiga                 | Taiga                | Shrubland        | Shrubland       | Temperate Desert | Temperate Desert |
// | 2               | Water        | Temperate Rain Forest | Temperate Forest     | Temperate Forest | Grassland       | Grassland        | Temperate Desert |
// | 1 (low)         | Water        | Tropical Rain Forest  | Tropical Rain Forest | Tropical Forest  | Tropical Forest | Grassland        | Desert           |
// | 0               | Water        | Water                 | Water                | Water            | Water           | Water            | Water            |
type Tile uint8

//Tile types
const (
	tileWater               Tile = iota
	tileFrozenWater         Tile = iota
	tileTropicalRainForest  Tile = iota
	tileTropicalForest      Tile = iota
	tileGrassland           Tile = iota
	tileDesert              Tile = iota
	tileTemperateRainForest Tile = iota
	tileTemperateForest     Tile = iota
	tileTemperateDesert     Tile = iota
	tileTaiga               Tile = iota
	tileShrubland           Tile = iota
	tileSnow                Tile = iota
	tileTundra              Tile = iota
	tileBare                Tile = iota
	tileScorched            Tile = iota

	tileMaxHeight   int = 4
	tileMaxHumidity int = 6
	tileLookupSize  int = (tileMaxHeight + 1) * (tileMaxHumidity + 1)
)

var tileHeightThresholds [tileMaxHeight]float64 = [tileMaxHeight]float64{1. / 5, 2. / 5, 3. / 5, 4. / 5}
var tileHumidityThresholds [tileMaxHumidity]float64 = [tileMaxHumidity]float64{1. / 7, 2. / 7, 3. / 7, 4. / 7, 5. / 7, 6. / 7}

//TileLookup is a lookup map for every height and humidity levels
var TileLookup [tileLookupSize]Tile = [tileLookupSize]Tile{
	//Height = 0
	tileWater, tileWater, tileWater, tileWater, tileWater, tileWater, tileWater,
	//Height = 1
	tileDesert, tileGrassland, tileTropicalForest, tileTropicalForest, tileTropicalRainForest, tileTropicalRainForest, tileWater,
	//Height = 2
	tileTemperateDesert, tileGrassland, tileGrassland, tileTemperateForest, tileTemperateForest, tileTemperateRainForest, tileWater,
	//Height = 3
	tileTemperateDesert, tileTemperateDesert, tileShrubland, tileShrubland, tileTaiga, tileTaiga, tileFrozenWater,
	//Height = 4
	tileScorched, tileBare, tileTundra, tileSnow, tileSnow, tileSnow, tileFrozenWater,
}

//GetTile returns a tile for the given height and humidity
//
//All parameters are bounded between 0 and 1. Values below zero are treated like zeroes and values above one are treated as ones.
func GetTile(height, humidity float64) Tile {
	//Assign the maximum values at first
	iHeight := tileMaxHeight
	iHumidity := tileMaxHumidity

	//Find the height
	for h, val := range tileHeightThresholds {
		if height < val {
			iHeight = h
			break
		}
	}

	//Find the humidity level
	for h, val := range tileHumidityThresholds {
		if humidity < val {
			iHumidity = h
			break
		}
	}

	return TileLookup[iHeight*(tileMaxHumidity+1)+iHumidity]
}
