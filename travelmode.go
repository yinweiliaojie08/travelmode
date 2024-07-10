// 最好是小写字母，并且最好不要带下划线，见名知义
package travelmode

import (
	"fmt"
)

func Aircraft(place string, cost interface{}) {
	fmt.Printf("已到达目的：%s,飞机飞的很快，价格有点高，本次花费%d元。\n", place, cost)
}

func Cycling(place string, cost interface{}) {
	fmt.Printf("已到达目的：%s,骑行有点慢，路上买了点水就行，本次花费%d元。\n", place, cost)
}

func Train(place string, cost interface{}) {
	fmt.Printf("已到达目的：%s,火车性价比很高，没花多少钱，本次花费%d元。\n", place, cost)
}
