package defaultregion

import (
	"fmt"
	"strconv"
	"utiloci/internal/data"
)

// ask for default region
func Resq() {

	var regionSeleted string

	for i, regions := range data.RegionsDefault {
		fmt.Println(i, ")", regions)
	}

	//Ask for select region
	fmt.Print("Select Default Region (39):")
	fmt.Scanln(&regionSeleted)

	//conver to init
	numberRegions, err := strconv.Atoi(regionSeleted)
	if err != nil {
		numberRegions = 34
	}

	//return regions
	data.DefaultRegion = data.RegionsDefault[numberRegions]

}
