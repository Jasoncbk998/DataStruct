/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/16 9:24 下午
 **/
package main

import (
	"DataStruct/algorithms/sorting/utils"
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	//list := utils.GetArrayOfSize(10)
	list := utils.GenerateArray(10000)
	sort(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}

}

func benchmarkBubbleSort(n int, b *testing.B) {
	list := utils.GenerateArray(n)
	for i := 0; i < b.N; i++ {
		sort(list)
	}
}

func BenchmarkBubbleSort100(b *testing.B)    { benchmarkBubbleSort(500, b) }
func BenchmarkBubbleSort1000(b *testing.B)   { benchmarkBubbleSort(1000, b) }
func BenchmarkBubbleSort10000(b *testing.B)  { benchmarkBubbleSort(10000, b) }
func BenchmarkBubbleSort100000(b *testing.B) { benchmarkBubbleSort(100000, b) }
