package main
//2 8 50 40   4 25 20
import (
	"fmt"
	"math/rand"
	"time"
)
func drawcard(star6 float64)(int){
	rand.NewSource(time.Now().UnixNano())
	carddraw := rand.Float64()
	if carddraw < star6{
		return 6
	}
	return 3
}
func trial(try int)(int,float64){
	var now int
	now =0
	var star6 float64
	star6 = 0.02
	star6drawn := 0
	for i:=1;i<=108;i++{
		now++
		if now>50 {
			star6 += 0.02
		}
		card := drawcard(star6)
		if card==6{
			star6drawn += 1
			now = 0
			star6 = 0.02
		}
	}
	var xi float64
	if star6drawn != 0{
		var nonXi float64 = 1
		for i:=1;i<=star6drawn;i++{
			nonXi*=0.65
		}
		xi = 1-nonXi;
	}
	if try%10000==0{
		fmt.Print(try)
		fmt.Print(" 6star:")
		fmt.Print(star6drawn)
		fmt.Print(" Xi:")
		fmt.Println(xi)
	}
	return star6drawn,xi
}
func main(){
	star6:=0
	num6:=0
	var xixi float64 = 0
	for i:=1;i<=10000;i++{
		six,xi:=trial(i)
		if six!=0{
			num6 += six
			star6+= 1
			xixi += xi
		}
	}
	xixi /= 100000
	fmt.Println(star6)
	fmt.Println(num6)
	fmt.Println(xixi)
}