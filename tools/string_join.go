package tools

import "strings"

//import "strings"

//标准包里的Join是实在不给力，效果太烂，只好自己写一个了
func StringJoin(s []string, split string) (result string){
	//result = ""
	for _,i := range s {
		if i != "" {
			result += i+split
		}
	}
	strings.Trim(result,",")
	//result = string([]byte(result)[:len(result)-1])
	return
}