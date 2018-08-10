package main

import (
	"fmt"

	"github.com/golang/geo/s2"
)

func main() {
	a := s2.PointFromLatLng(s2.LatLng{35.659104, 139.703742}) // ヒカリエ
	b := s2.PointFromLatLng(s2.LatLng{35.689634, 139.692101}) // 新宿都庁
	c := s2.PointFromLatLng(s2.LatLng{35.701732, 139.580418}) // 吉祥寺駅
	fmt.Println(s2.PointArea(a, b, c))
	fmt.Println(s2.GirardArea(a, b, c))
	fmt.Println(s2.SignedArea(a, b, c))
}
