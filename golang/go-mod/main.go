package main

import (
	"fmt"
	"github.com/golang/geo/s2"
)

func main() {
	cell := s2.CellFromLatLng(s2.LatLng{35.659104, 139.703742})
	fmt.Println(cell.ID())
}
