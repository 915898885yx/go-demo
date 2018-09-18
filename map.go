package main

import ("fmt")

func main() {
	var countryCapitalMap map[string]string /*创建集合*/
	countryCapitalMap = make(map[string]string)

	// map插入key-value对，各个国家对应的首都
	countryCapitalMap ["France"] = "paris"
	countryCapitalMap ["Italy"] = "罗马"
	countryCapitalMap ["jspan"] = "东京"
	countryCapitalMap ["India"] = "新德里"

	// 使用键值对输出地图值
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}

	delete(countryCapitalMap, "India")
	fmt.Println("删除后")
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}
	// 查看元素在集合中是否存在
	capital, ok := countryCapitalMap ["美国"]
	if (ok) {
		fmt.Println("美国的首都是", capital)
	} else {
		fmt.Println("美国的首都不存在")
	}
}