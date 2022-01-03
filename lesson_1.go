package main

//import "fmt"
import (
	"fmt"
	"os"   // 用來接收 terminal 的參數
	"path" // import 多個套件的話，中間不是隔行就是分號
)

func main() {
	fmt.Println("Hello, World16!")

	// declaration 宣告變數
	var number int // 初始為 0
	fmt.Println("plus one is", number+1)

	var dir, file string                           // 多重宣告
	dir, file = path.Split("website/css/main.css") // path: 方便用來處理 url string 路徑的套件
	// _, file = path.Split("website/css/main.css") // 可以用底線拋棄部分回傳值，底線不用宣告
	fmt.Println("dir: ", dir)   // dir:  website/css/
	fmt.Println("file: ", file) // file:  main.css
	//------------------------------------------------------------------------------

	// := 短宣告
	// 習慣上若"知道初始值"就用短宣告(後面不需要 type)；不知道的話用上面那種一般宣告
	width, color := 50, "red"
	fmt.Println("the width of table is", width, " and the color is ", color)

	//-----------------------------------------------------------------------------

	// type 轉換
	// 可用 type(變數) 的方式轉換型別，是破壞性轉換
	speed, force := 100, 2.5
	speed = speed * int(force)      // 若運算型別不一致會 error
	fmt.Println("speed is ", speed) // speed is  200

	//------------------------------------------------------------------------------

	// 需要封裝成 .exe 之後直接使用並給予三個參數
	// 接收 terminal 的參數。os 套件。呼叫 os.Arg[] 來使用參數
	fmt.Printf("%#v\n", os.Args)             // []string{"C:\\Users\\user\\AppData\\Local\\Temp\\go-build204583959\\b001\\exe\\lesson_1.exe", "how", "are", "you"}
	fmt.Println(os.Args)                     // [C:\Users\user\AppData\Local\Temp\go-build204583959\b001\exe\lesson_1.exe how are you]
	fmt.Println("Path", os.Args[0])          // Path C:\Users\user\AppData\Local\Temp\go-build204583959\b001\exe\lesson_1.exe
	fmt.Println("1st Argument:", os.Args[1]) // 1st Argument: how
	fmt.Println("2nd Argument:", os.Args[2]) // 2nd Argument: are
	fmt.Println("3rd Argument:", os.Args[3]) // 3rd Argument: you

	//-----------------------------------------------------------------------------
	// 封裝程式: go build 檔名.go

}
