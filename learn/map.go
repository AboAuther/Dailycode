package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]string{
		"beer":   "bière",
		"wine":   "vin",
		"water":  "eau",
		"coffee": "café",
		"thea":   "thé",
	}
)

func Season(month int) (strMonth string) {
	switch {
	case month <= 5 && month >= 3:
		strMonth = "Spring"
	case month >= 6 && month <= 8:
		strMonth = "Summer"
	case month >= 9 && month <= 11:
		strMonth = "Autmn"
	default:
		strMonth = "Winter"
	}
	return
}

type Rope string

func main() {
	fmt.Println("Unsorted:")
	for k, v := range barVal {
		fmt.Printf("key:%v,value:%v\t", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("key:%v,value:%v\t", k, barVal[k])
	}
	//构造一个将英文饮料名映射为法语（或者任意你的母语）的集合；
	//先打印所有的饮料，然后打印原名和翻译后的名字。接下来按照英文名排序后再打印出来。
	/*	drinks:=make([]string,len(barVal))
		fmt.Println("饮料有：")
		i:=0
		for v:=range barVal{
			drinks[i]=v
			fmt.Printf(v+"\t")
			i++
		}
		fmt.Println()
		for k,v:=range barVal{
			fmt.Print(k+"\t"+v+"\n")
		}
		fmt.Println()
		sort.Strings(drinks)
		for _,k:= range drinks{
			fmt.Printf(k+"\t"+barVal[k]+"\n")
		}*/

	/*	fmt.Print()
		var s Rope
		s="I love China"
		fmt.Print(s)
		var i1 = 5
		fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)*/

	//重写本节中生成斐波那契数列的程序并返回两个命名返回值（详见第 6.2 节），即数列中的位置和对应的值，例如 5 与 4，89 与 10。
	/*	result := 0
		loc :=0
		for i := 0; i <= 10; i++ {

			result,loc = fibonacci(i)
			fmt.Printf("fibonacci(%d) is: %d\n", loc, result)
		}*/

	//使用递归函数从 10 打印到 1。
	/*	printOneToTen(10)
	 */
	//实现一个输出前 30 个整数的阶乘的程序。
	/*fmt.Println(iterCount(30,big.NewInt(30)))
	 */
	//在 main 函数中写一个用于打印 Hello World 字符串的匿名函数并赋值给变量 fv，然后调用该函数并打印变量 fv 的类型。
	/*fv:=func(s string){fmt.Println(s)}
	fv("hello world")
	fmt.Printf("Type:%T",fv)*/
	//闭包应用
	/*	var g int
		go func(i int ) {//使用go关键字会创建goroutines协程  导致程序并发执行的顺序不同致使得到的结果为0
			s := 0
			for j := 0; j < i; j++ { s += j }
			g = s
		}(1000)
		print(g)*/
	// 不使用递归但使用闭包改写第 6.6 节中的斐波那契数列程序
	/*var array [30]int
	f:=fibonacci()
	for i:=0;i< len(array);i++{
		array[i]=f()
	}
	fmt.Println(array)*/

}

//重写本节中生成斐波那契数列的程序并返回两个命名返回值（详见第 6.2 节），即数列中的位置和对应的值，例如 5 与 4，89 与 10。
/*func fibonacci(s int)( res int,location int){
	location=s
	if s<=1{
		res=1
	}else{
		res1,_ :=fibonacci(s-1)
		res2,_:=fibonacci(s-2)
		res=res1+res2
	}
	return
}*/
//使用递归函数从 10 打印到 1。
/*func printOneToTen(num int)(int){
	if num<1{
		return 0
	}
	fmt.Println(num)
	return printOneToTen(num-1)
}*/
//实现一个输出前 30 个整数的阶乘的程序。
/*func iterCount(num int ,count *big.Int) *big.Int{
	if num<=1 || count.Int64() == 1{
		return big.NewInt(1)
	}else{
		return count.Mul(count,iterCount(num-1,big.NewInt(count.Int64()-1)))
	}
}*/
// 不使用递归但使用闭包改写第 6.6 节中的斐波那契数列程序
/*func fibonacci() func()int{
	back1,back2:=0,1
	return func () int{
		back1,back2=back2,back1+back2
		return back1
	}

}*/
