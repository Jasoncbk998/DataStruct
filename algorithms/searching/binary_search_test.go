/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/19 11:59 上午
 **/
package searching

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	sorted := make([]int, 100)
	for i := 0; i < 100; i++ {
		sorted[i] = 2 * i
	}

	for i := 0; i < 100; i++ {
		index := Search(sorted, 2*i)
		if index != i {
			fmt.Println(index)
			t.Error()
		}
	}
	if Search(sorted, 3) != -1 {
		t.Error()
	}

}
