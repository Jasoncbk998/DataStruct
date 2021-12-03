/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/19 12:46 下午
 **/
package 两数之和

import (
	"DataStruct/algorithms/sorting/utils"
	"fmt"
	"math/rand"
	"testing"
)

func TestSum2Number(t *testing.T) {
	array := utils.GenerateArray(100)
	for i := 0; i < len(array); i++ {
		intn := rand.Intn(len(array))
		number := Sum2Number(array, intn)
		fmt.Println(number)
		t.Error()
	}

}
