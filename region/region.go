package region

type Region struct {
	Country string
	State   string
	City    string
}

func Matches(r1, r2 Region) bool {
	return (r1.Country == r2.Country || r2.Country == "") &&
		(r1.State == r2.State || r2.State == "") &&
		(r1.City == r2.City || r2.City == "")
}
