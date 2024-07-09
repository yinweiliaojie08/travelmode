// 最好是小写字母，并且最好不要带下划线，见名知义
package travelmode

import (
	"fmt"
)

func Priceaircraft(cost interface{}) {
	fmt.Printf("飞机飞的很快，价格有点高，本次花费%d \n", cost)
}

func Pricecycling(cost interface{}) {
	fmt.Printf("骑行有点慢，路上买了点水就行，本次花费%d \n", cost)
}

func Pricetrain(cost interface{}) {
	fmt.Printf("火车性价比很高，没花多少钱，本次花费%d \n", cost)
}
