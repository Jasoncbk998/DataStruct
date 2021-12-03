/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/10/28 8:43 下午
 **/
package main

func rotate(matrix [][]int) {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}
	for i, row := range matrix {
		for j, v := range row {
			tmp[j][n-1-i] = v
		}
	}
	copy(matrix, tmp) // 拷贝 tmp 矩阵每行的引用
}
