/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/17 7:05 下午
 **/
package main

import (
	"DataStruct/algorithms/sorting/utils"
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	array := utils.GenerateArray(10)
	sort(array)
	for i := 0; i < len(array)-2; i++ {
		if array[i] > array[i+1] {
			fmt.Println(array)
			t.Error()
		}
	}
}

func benchmarkInsertionSort(n int, b *testing.B) {
	array := utils.GenerateArray(n)
	for i := 0; i < b.N; i++ {
		sort(array)
	}
}

func BenchmarkInsertionSort100(b *testing.B)    { benchmarkInsertionSort(100, b) }
func BenchmarkInsertionSort1000(b *testing.B)   { benchmarkInsertionSort(1000, b) }
func BenchmarkInsertionSort10000(b *testing.B)  { benchmarkInsertionSort(10000, b) }
func BenchmarkInsertionSort100000(b *testing.B) { benchmarkInsertionSort(100000, b) }
