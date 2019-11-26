// xml.go
package main
import (
	"encoding/xml"
	"fmt"
	"strings"
)
var t, token xml.Token//申明Token，像这样2个一起申明，即使其中一个不使用也是可以的！
var err error
func main34343345() {
	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)//创建解析input字符串的解码器
	for t, err = p.Token(); err == nil; t, err = p.Token() {//p.Token()获取解码器的Token，遍历结束后err == EOF而不是nil
		switch token := t.(type) {
		case xml.StartElement://如果是开始标签
			name := token.Name.Local//获取标签的名字
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {//获取标签的属性，这个属性包含名字和值
				attrName := attr.Name.Local
				attrValue := attr.Value//分别取得属性的名字和值
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
				// ...
			}
		case xml.EndElement://如果是结束标签
			fmt.Println("End of token")
		case xml.CharData://如果是文本字符
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
			// ...
		default://如果什么都不是的话，什么都不做
			// ...
		}
	}

	//输出：
	//	Token name: Person
	//	Token name: FirstName
	//	This is the content: Laura
	//	End of token
	//	Token name: LastName
	//	This is the content: Lynn
	//	End of token
	//	End of token
}