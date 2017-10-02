package main

import "fmt"

type Location struct {
	street  string
	house   string
	apt     uint8
	city    string
	zip     string
	country string
}

type Vertex struct {
	Lat, Long float64
	Loc       Location
}

// var m map[string]Vertex

func main() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
		Location{
			"Elmstr", "123", 8, "MountainView", "12345", "USA",
		},
	}
	m["Yarik's home"] = Vertex{
		3, 5,
		Location{
			"Borstellstr", "11", 12, "Berlin", "12345", "Germany",
		},
	}
	fmt.Println(m)
}
