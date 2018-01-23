package tools

//标准包里的Join是实在不给力，效果太烂，只好自己写一个了
func StringJoin(s []string, split string) (result string){
	result = ""
	for i,_ := range s {
		if s[i] != "" {
			result += s[i]+split
		}
	}
	result = string([]rune(result)[:len(result)-1])
	return
}