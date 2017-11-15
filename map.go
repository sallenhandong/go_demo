//Go 字典

package main

import (
	"fmt"
)

func main() {

	dictionary := make(map[string]int)

	dictionary["k1"] = 7
	dictionary["k2"] = 10

	//输出字典
	fmt.Println("map:", dictionary)
	//获取一个键的值
	name := dictionary["k1"]
	fmt.Println("name:", name)

	//内置函数len 返回字典元素的个数
	fmt.Println("len:", len(dictionary))

	//内置函数delete 从字典删除一个键对应的值
	delete(dictionary, "k2")
	fmt.Println("map:", dictionary)

	// 根据键来获取值有一个可选的返回值，这个返回值表示字典中是否
	// 存在该键，如果存在为true，返回对应值，否则为false，返回零值
	// 有的时候需要根据这个返回值来区分返回结果到底是存在的值还是零值
	// 比如字典不存在键x对应的整型值，返回零值就是0，但是恰好字典中有
	// 键y对应的值为0，这个时候需要那个可选返回值来判断是否零值。

	_, ok := dictionary["k2"]
	fmt.Println("ok:", ok)

	//可以使用 ":=" 同时定义和初始化一个字典
	myMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", myMap)

}
