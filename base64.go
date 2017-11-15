// Go Base64编码
package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// 这里是我们用来演示编码和解码的字符串
	data := "abc123!?$*&()'-=@~"

	// Go支持标准的和兼容URL的base64编码。
	// 我们这里使用标准的base64编码，这个
	// 函数需要一个`[]byte`参数，所以将这个字符串转换为字节数组
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	// 解码一个base64编码可能返回一个错误，
	// 如果你不知道输入是否是正确的base64编码，你需要检测一些解码错误
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()
	// 使用兼容URL的base64编码和解码
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
