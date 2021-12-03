/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/6 2:15 下午
 **/
package main

import (
	"fmt"
	"sort"
)

/**
有n件物品和一个最大承重为W的背包,每件物品的重量是w,价值是v
保证总重量不超过w的前提下,将哪些物品放进背包,使总价值最大
每个物品只有1件,所以叫01背包
三种选择方案
	价值主导:优先选择价值最高的物品放进背包
	重量主导:优先选择重量最轻的物品放进背包
	价值密度主导:价值/重量,价值密度越高的优先放进去
*/

type Article struct {
	weight int
	value  int
	//valueDensity float64
}

func main() {
	articles := []Article{
		{35, 10},
		{35, 10},
		{60, 30},
		{40, 35},
		{25, 30},
		{30, 40},
		{50, 50},
		{10, 40},
	}
	sort.Slice(articles, func(i, j int) bool {
		if articles[i].value < articles[j].value {
			return false
		}
		return true
	})
	Select("价值主导", articles)

}

func Select(title string, articles []Article) {

	capacity, weight, value := 150, 0, 0
	selectArticles := make([]Article, len(articles))
	for i := 0; i < len(articles) && weight < capacity; i++ {
		newWeight := weight + articles[i].weight
		if newWeight <= capacity {
			weight = newWeight
			value += articles[i].value
			selectArticles = append(selectArticles, articles[i])
		}
	}
	fmt.Println("title:", title)
	fmt.Println("总价值:", value)
	for i := 0; i < len(selectArticles); i++ {
		if selectArticles[i].value > 0 {
			fmt.Println(selectArticles[i])
		}
	}
	fmt.Println("---------")
}

func compareArticle(title string, articles *[]Article) {
	switch title {
	case "价值主导":
		sort.Slice(articles, func(i, j int) bool {
			if (*articles)[i].value > (*articles)[j].value {
				return true
			}
			return false
		})
	case "重量主导":
		sort.Slice(articles, func(i, j int) bool {
			if (*articles)[i].weight > (*articles)[j].weight {
				return true
			}
			return false
		})
	case "价值密度":
		sort.Slice(articles, func(i, j int) bool {
			if (*articles)[i].getValueDensity() > (*articles)[j].getValueDensity() {
				return true
			}
			return false
		})
	}

}

func (article *Article) getValueDensity() int {
	return article.value * 1.0 / article.weight
}
