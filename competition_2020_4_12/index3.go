package main

import (
	"strings"
)

// func main() {
// 	fmt.Println(entityParser("&amp; is an HTML entity but &ambassador; is not."))
// 	fmt.Println(entityParser("and I quote: &quot;...&quot;"))
// 	fmt.Println(entityParser("x &gt; y &amp;&amp; x &lt; y is always false"))
// 	fmt.Println(entityParser("and I quote: &quot;...&quot;"))
// 	fmt.Println(entityParser("leetcode.com&frasl;problemset&frasl;all"))
// }

// 双引号：字符实体为 &quot; ，对应的字符是 " 。
// 单引号：字符实体为 &apos; ，对应的字符是 ' 。
// 与符号：字符实体为 &amp; ，对应对的字符是 & 。
// 大于号：字符实体为 &gt; ，对应的字符是 > 。
// 小于号：字符实体为 &lt; ，对应的字符是 < 。
// 斜线号：字符实体为 &frasl; ，对应的字符是 / 。

func entityParser(text string) string {
	tempText := strings.Split(text, "&quot;")
	text = strings.Join(tempText, "\"")
	tempText = strings.Split(text, "&apos;")
	text = strings.Join(tempText, "'")
	tempText = strings.Split(text, "&gt;")
	text = strings.Join(tempText, ">")
	tempText = strings.Split(text, "&lt;")
	text = strings.Join(tempText, "<")
	tempText = strings.Split(text, "&frasl;")
	text = strings.Join(tempText, "/")
	tempText = strings.Split(text, "&amp;")
	text = strings.Join(tempText, "&")
	return text
}
