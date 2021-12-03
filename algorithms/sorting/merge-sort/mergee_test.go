/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/18 2:23 下午
 **/
package merge_sort

import (
	"DataStruct/algorithms/sorting/utils"
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	list := utils.GenerateArray(10)
	Sort(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}
func benchmarkMergeSort(n int, b *testing.B) {
	list := utils.GenerateArray(n)
	for i := 0; i < b.N; i++ {
		Sort(list)
	}
}

func BenchmarkMergeSort100(b *testing.B)    { benchmarkMergeSort(100, b) }
func BenchmarkMergeSort1000(b *testing.B)   { benchmarkMergeSort(1000, b) }
func BenchmarkMergeSort10000(b *testing.B)  { benchmarkMergeSort(10000, b) }
func BenchmarkMergeSort100000(b *testing.B) { benchmarkMergeSort(100000, b) }
