package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	s := time.Now()
	elapsed := time.Since(s)
	fmt.Printf("String: %s\n", elapsed.String())
	fmt.Printf("Seconds: %f\n", elapsed.Seconds())
	fmt.Printf("Nanoseconds: %d\n", elapsed.Nanoseconds())
	trunc(elapsed)

	d1, _ := time.ParseDuration("1h15m31.918273645s")
	trunc(d1)
	d2, _ := time.ParseDuration("1h45m31.918273645s")
	trunc(d2)
	fmt.Println(d1 < d2)

	//fmt.Println(time.Duration(float64(time.Millisecond) * 0.0000001).String()) // -> 0.1になる。結果intにできずpanic
	fmt.Println(time.Duration(float64(time.Millisecond) * 0.000001).String())
	fmt.Println(time.Duration(float64(time.Millisecond) * 0.001).String())
	fmt.Println(time.Duration(float64(time.Millisecond) * 1.15).String())

	d := int64(5)
	//fmt.Println(time.Duration(time.Millisecond * d).String()) // mismatch type
	fmt.Println(time.Duration(time.Millisecond * time.Duration(d)).String())

	fmt.Println(time.Duration(time.Hour * 24 * 365 * 292).String())
	//fmt.Println(time.Duration(time.Hour * 24 * 365 * 293).String()) // overflow
	e, _ := time.Parse(time.RFC3339, "9999-12-31T23:59:59Z")
	fmt.Println(e.Sub(s).String())
	fmt.Println(int64(time.Since(e)))
	fmt.Println(int64(time.Until(e)))
	fmt.Println(math.MaxInt64)

}

func trunc(elapsed time.Duration) {
	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}
	fmt.Println("===", elapsed.String())
	for _, t := range trunc {
		fmt.Printf("d.Truncate(%6s) = %s\n", t, elapsed.Truncate(t).String())
	}
	for _, t := range trunc {
		fmt.Printf("d.Round(%6s) = %s\n", t, elapsed.Round(t).String())
	}
}
