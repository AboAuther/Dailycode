package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

func main() {
	/*var score [10]int
	var i int
	for i=0;i<10;i++{
		score[i]=i+10
	}
	for i=0;i<10;i++{
		fmt.Printf("Element[%d]=%d\n",i,score[i])
	}*/

	//统计字符串中的汉字数量 用unicode.Han函数
	/*
		//s1:="hello沙河小王子"
		//num :=0
		//for _,char :=range s1{
		//	if unicode.Is(unicode.Han,char){
		//		num++;
		//	}
		//}
		//fmt.Printf("汉字数量为：%d",num)*/

	//有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字
	/*	var nums  = []int {5,7,8,8,9,7,9}
		i := 0
		for j := 0; j < len(nums); j ++ {
			i = i ^ nums[j]
		}
		fmt.Println(i)*/

	//for循环打印99乘法表
	/*for i:=1;i<10;i++{
		for j:=i;j<10;j++{
			fmt.Printf("%d * %d =%d\t",i,j,i*j)
		}
		fmt.Printf("\n")
	}*/
	//匿名变量测试
	/*	a := [5]int{1,3,5,7,9}
		sum:=0
		for _,v := range a{
			sum+=v
		}
		fmt.Print(sum)*/

	/*	a := []int{1,3,5,7,9}
		sum:=8
		for i:=0;i<len(a);i++{
			for j:=i;j<len(a);j++{
				if a[i]+a[j]==sum {
					println(i,j)
				}
			}
		}*/

	/*	str1 := "asSASA ddd dsjkdsjs dk"
		fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
		fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))
		str2 := "asSASA ddd dsjkdsjsこん dk"
		fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
		fmt.Printf("The number of characters in string str2 is %d", utf8.RuneCountInString(str2))*/

	/*	//HasPrefix 判断字符串 s 是否以 prefix 开头：
		//HasSuffix 判断字符串 s 是否以 suffix 结尾：
		var str string = "This is an example of a string"
		fmt.Printf("T/F? Does the string \"%s\" have prefix %s?\n ", str, "Th")
		fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))*/

	// 定义一个字符串，统计字符串中单词出现的顺序
	str := "how do you how how do"
	//定义并初始化map
	count := make(map[string]int)
	//切分字符串
	str2 := strings.Split(str, " ")
	for _, v := range str2 {
		count[v]++
	}
	fmt.Println(count)
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	/*	var a = make([]string, 5, 10)
		for i := 0; i < 10; i++ {
			a = append(a, fmt.Sprintf("%v", i))
		}
		fmt.Println(a)   //	[     0 1 2 3 4 5 6 7 8 9]*/

	//创建一个 map 来保存每周 7 天的名字，将它们打印出来并且测试是否存在 Tuesday 和 Hollyday。
	/*week:=make(map[string]string)
	week["Monday"]="星期一"
	week["Tuesday"]="星期二"
	week["Wednesday"]="星期三"
	week["Thursday"]="星期四"
	week["Friday"]="星期五"
	week["Saturday"]="星期六"
	week["Sunday"]="星期日"
	for k,v :=range week{
		fmt.Println(k,v)
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
	var value string
	var isPresent bool
	value,isPresent=week["Tuesday"]
	fmt.Printf("Is the Tuesday is in week?\n %v\t",isPresent)
	fmt.Printf("%v\n",value)
	_,isPresent=week["holiday"]
	fmt.Printf("Is the holiday is in week?\n %v",isPresent)*/

	/*	a := [...]string{"a", "b", "c", "d"}
		for i := range a {
			fmt.Println("Array item", i, "is", a[i])
		}*/

	/*	items2 := make([]map[int]int, 5)
		for _, item := range items2 {
			item = make(map[int]int, 1) // item is only a copy of the slice element.
			item[1] = 2 // This 'item' will be lost on the next iteration.
		}
		fmt.Printf("Version B: Value of items: %v\n", items2)*/

}
