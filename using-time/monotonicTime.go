// Time calculations for Applications

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//  method 1

	now := time.Now()
	fmt.Printf("Current time is : %v\n", now.Format("15:04:05"))

	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(10)
	randTime := now.Add(time.Duration(randomInt) * time.Second)
	fmt.Printf("time after %v seconds is : %v\n", randomInt, randTime.Format("15:04:05"))

	time.Sleep(time.Duration(randomInt) * time.Second)

	elapsed := time.Until(randTime)
	fmt.Printf("Time remains %v\n", elapsed)

	total := time.Since(now)
	fmt.Printf("Total elapsed time is :  %v\n", total)

}
