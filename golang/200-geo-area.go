package main

import (
	"fmt"
	"math"

	"github.com/golang/geo/s2"
)

func main() {
	a := s2.PointFromLatLng(s2.LatLngFromDegrees(35.659104, 139.703742)) // ヒカリエ
	b := s2.PointFromLatLng(s2.LatLngFromDegrees(35.689634, 139.692101)) // 新宿都庁
	c := s2.PointFromLatLng(s2.LatLngFromDegrees(35.701732, 139.580418)) // 吉祥寺駅
	fmt.Println(s2.PointArea(a, b, c))
	fmt.Println(s2.GirardArea(a, b, c))
	fmt.Println(s2.SignedArea(a, b, c))

	// 直線上の三点 = 0
	a = s2.PointFromLatLng(s2.LatLngFromDegrees(0, 0))
	b = s2.PointFromLatLng(s2.LatLngFromDegrees(0, 90))
	c = s2.PointFromLatLng(s2.LatLngFromDegrees(0, 180))
	fmt.Println(s2.PointArea(a, b, c))

	// 南北極と東経90度(地球の1/4) = PI
	a = s2.PointFromLatLng(s2.LatLngFromDegrees(90, 0))
	b = s2.PointFromLatLng(s2.LatLngFromDegrees(-90, 0))
	c = s2.PointFromLatLng(s2.LatLngFromDegrees(0, 90))
	fmt.Println(s2.PointArea(a, b, c))

	// 南北極と東経180度(地球の1/2) = 2PI
	a = s2.PointFromLatLng(s2.LatLngFromDegrees(90, 0))
	b = s2.PointFromLatLng(s2.LatLngFromDegrees(-90, 0))
	c = s2.PointFromLatLng(s2.LatLngFromDegrees(0, 180))
	fmt.Println(s2.PointArea(a, b, c))
	fmt.Println(s2.PointArea(a, b, c) * math.Pow(6378, 2) * 2) // 地球表面積になるはず
}
