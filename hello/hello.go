//本章节主要内容
//编写测试
//用参数和返回类型声明函数
//if，const，switch
//声明变量和常量

package main
import "fmt"


func main()  {
	//1.直接打印
	fmt.Println("Hello,world")

	//2.赋值打印
	fmt.Printf("%v\n","Hello,world")

	//3.调函数赋值打印
	//fmt.Printf("%v\n",Hello())
}

//常量
const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Halo,"
const frenchHelloPrefix = "Bonjour,"
const spanish = "Spanish"
const french = "French"

//版本1
func Hello(name string) string {
	if name == ""{
		name ="Hello,"
	}
	return englishHelloPrefix + name
}

//版本2:多参数
func Hello2(name string,language string) string {
	//1.初始版本
	//if language == spanish{
	//	return   spanishHelloPrefix + name
	//}
	//if language == french{
	//	return   frenchHelloPrefix + name
	//}
	//if name == ""{
	//	name ="World"
	//}
	//return englishHelloPrefix + name



	//2.使用switch重构版本
	//if name == ""{
	//	name ="World"
	//}
	//prefix := englishHelloPrefix
	//switch language {
	//case french:
	//	prefix=frenchHelloPrefix
	//case spanish:
	//	prefix=spanishHelloPrefix
	//}
	//return prefix + name


	//3.将代码封装到一个函数中
	return greetingPrefix(language) + name
}

//在我们的函数签名中，我们使用了 命名返回值（prefix string）。
//这将在你的函数中创建一个名为 prefix 的变量。
//它将被分配「零」值。这取决于类型，例如 int 是 0，对于字符串它是 ""。
//你只需调用 return 而不是 return prefix 即可返回所设置的值。
//这将显示在 Go Doc 中，所以它使你的代码更加清晰。
func greetingPrefix(language string)(prefix string)  {
	switch language {
	case french:
		prefix =frenchHelloPrefix
	case spanish:
		prefix =spanishHelloPrefix
	default:
		prefix =englishHelloPrefix
	}
	return
}
