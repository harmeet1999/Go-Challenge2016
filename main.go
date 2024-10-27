package main

import (
	"challenge2016/distributor"
	"challenge2016/region"
	"challenge2016/utils"
	"fmt"
)

func main() {
	_, err := utils.LoadRegions("./cities.csv")
	if err != nil {
		fmt.Println("Error loading regions:", err)
		return
	}

	distributor1 := distributor.Distributor{
		Name:     "DISTRIBUTOR1",
		Includes: []region.Region{{Country: "INDIA"}, {Country: "UNITEDSTATES"}},
		Excludes: []region.Region{{State: "KARNATAKA", Country: "INDIA"}, {City: "CHENNAI", State: "TAMILNADU", Country: "INDIA"}},
	}

	distributor2 := distributor.Distributor{
		Name:     "DISTRIBUTOR2",
		Parent:   &distributor1,
		Includes: []region.Region{{Country: "INDIA"}},
		Excludes: []region.Region{{State: "TAMILNADU", Country: "INDIA"}},
	}

	checkRegion := region.Region{City: "BANGALORE", State: "KARNATAKA", Country: "INDIA"}
	fmt.Printf("DISTRIBUTOR1 has permission in BANGALORE-KARNATAKA-INDIA: %v\n", distributor1.HasPermission(checkRegion))
	fmt.Printf("DISTRIBUTOR2 has permission in BANGALORE-KARNATAKA-INDIA: %v\n", distributor2.HasPermission(checkRegion))
}
