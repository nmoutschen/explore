package terrain

//Tile represents the different types of tiles that can be used in game.
//
//They are stored as a 8-bit unsigned integer (uint8), with the four least significant bits used to store the
//biome type and the four most significant bits used for structure information. For example, you can create a tile
//with a desert biome and a road by using a logical OR operation: terrain.BiomeDesert | terrain.StructureRoad.
//
//Tiles cannot have multiple biomes or structures. Therefore, terrain.BiomeFrozenWater | terrain.BiomeGrassland == terrain.BiomeDesert
//
//Tile types are based on the Whittaker classification with some modifications. The following table is used to map
//discrete humidity and height values to biomes. An array of thresholds map continuous values for height and humidity
//from the noise generator into the discrete values contained in the following table.
//
// | Height\Humidity | 6 (wet)      | 5                     | 4                    | 3                | 2               | 1                | 0 (dry)          |
// | 4 (high)        | Frozen Water | Snow                  | Snow                 | Snow             | Tundra          | Bare             | Scorched         |
// | 3               | Frozen Water | Taiga                 | Taiga                | Shrubland        | Shrubland       | Temperate Desert | Temperate Desert |
// | 2               | Water        | Temperate Rain Forest | Temperate Forest     | Temperate Forest | Grassland       | Grassland        | Temperate Desert |
// | 1 (low)         | Water        | Tropical Rain Forest  | Tropical Rain Forest | Tropical Forest  | Tropical Forest | Grassland        | Desert           |
// | 0               | Water        | Water                 | Water                | Water            | Water           | Water            | Water            |
type Tile uint8

//Tile types
const (
	//Biome types
	BiomeWater               Tile = 0x00
	BiomeFrozenWater         Tile = 0x01
	BiomeTropicalRainForest  Tile = 0x02
	BiomeTropicalForest      Tile = 0x03
	BiomeGrassland           Tile = 0x04
	BiomeDesert              Tile = 0x05
	BiomeTemperateRainForest Tile = 0x06
	BiomeTemperateForest     Tile = 0x07
	BiomeTemperateDesert     Tile = 0x08
	BiomeTaiga               Tile = 0x09
	BiomeShrubland           Tile = 0x0A
	BiomeSnow                Tile = 0x0B
	BiomeTundra              Tile = 0x0C
	BiomeBare                Tile = 0x0D
	BiomeScorched            Tile = 0x0E
	//0x0F is reserved for future use

	//Structure types
	StructureEmpty Tile = 0x00
	StructureRoad  Tile = 0x10
	StructureHouse Tile = 0x20
	//0x30-0xF0 are reserved for future use

	tileMaxHeight   int = 4
	tileMaxHumidity int = 6
	biomeLookupSize int = (tileMaxHeight + 1) * (tileMaxHumidity + 1)
)

var tileHeightThresholds [tileMaxHeight]float64 = [tileMaxHeight]float64{
	2. / 6, 3. / 6, 4. / 6, 5. / 6,
}
var tileHumidityThresholds [tileMaxHumidity]float64 = [tileMaxHumidity]float64{
	1. / 7, 2. / 7, 3. / 7, 4. / 7, 5. / 7, 6. / 7,
}

//BiomeLookup is a lookup map for every height and humidity levels
var BiomeLookup [biomeLookupSize]Tile = [biomeLookupSize]Tile{
	//Height = 0
	BiomeWater, BiomeWater, BiomeWater, BiomeWater, BiomeWater, BiomeWater, BiomeWater,
	//Height = 1
	BiomeDesert, BiomeGrassland, BiomeTropicalForest, BiomeTropicalForest, BiomeTropicalRainForest, BiomeTropicalRainForest, BiomeWater,
	//Height = 2
	BiomeTemperateDesert, BiomeGrassland, BiomeGrassland, BiomeTemperateForest, BiomeTemperateForest, BiomeTemperateRainForest, BiomeWater,
	//Height = 3
	BiomeTemperateDesert, BiomeTemperateDesert, BiomeShrubland, BiomeShrubland, BiomeTaiga, BiomeTaiga, BiomeFrozenWater,
	//Height = 4
	BiomeScorched, BiomeBare, BiomeTundra, BiomeSnow, BiomeSnow, BiomeSnow, BiomeFrozenWater,
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

	return BiomeLookup[iHeight*(tileMaxHumidity+1)+iHumidity]
}
