package distributor

import "challenge2016/region"

func matches(r1, r2 region.Region) bool {
	return (r1.Country == r2.Country || r1.Country == "") &&
		(r1.State == r2.State || r1.State == "") &&
		(r1.City == r2.City || r1.City == "")
}

func (d *Distributor) HasPermission(region region.Region) bool {
	if isExcluded(region, d.Excludes) {
		return false
	}
	if isIncluded(region, d.Includes) {
		if d.Parent != nil {
			return d.Parent.HasPermission(region)
		}
		return true
	}
	return false
}

func isIncluded(region region.Region, includes []region.Region) bool {
	for _, r := range includes {
		if matches(r, region) {
			return true
		}
	}
	return false
}

func isExcluded(region region.Region, excludes []region.Region) bool {
	for _, r := range excludes {
		if matches(r, region) {
			return true
		}
	}
	return false
}
