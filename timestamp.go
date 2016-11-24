package main

import (
    "time"
    "math/rand"
)

func MakeTimeStamp() time.Time {
    base := time.Now().Unix() - 10000
    return time.Unix( int64(rand.Intn(100000)) + base, 0 )
//    return time.Unix( int64(rand.Intn(100000)) + base, 0).Format(time.RFC3339)
}
