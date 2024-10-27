package distributor

import "challenge2016/region"

type Distributor struct {
	Name     string
	Parent   *Distributor
	Includes []region.Region
	Excludes []region.Region
}
