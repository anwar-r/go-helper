package main

import "time"

func TimeIn() time.Time {
	loc, err := time.LoadLocation("")
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc)
}
