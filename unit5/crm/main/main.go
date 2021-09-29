package main //1.package进行包的声明，建议：包的声明这个包和所在的文件夹同名

//4.如果有多个包，建议一次性导入
import (
	"fmt"
	"unit5/crm/dbutils"
)

//2.main包是程序的入口包，一般main函数放在这个包下面

//3.包名是从$GOPATH/src/后开始计算的，使用/进行路径分隔

func main() {
	fmt.Println("你好，这是main函数的执行")
	dbutils.GetConn() //4.在函数调用的时候前面要定位到所在的包
}
