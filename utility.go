package main

import (
	"math"
	"encoding/json"
)

//Returns exp at level MLvl with curve Curve
func MaxLevelExp(Curve int, MLvl int) (res int) {
	res = int (  float64(Curve) * math.Pow((float64(MLvl)-1.0) / 98.0 , 2.5)  )
	return res
}


func u_PPJson(input interface{}, a, b string) string{
        s, _ := json.MarshalIndent(input, a, b)
        return(string(s))
}
