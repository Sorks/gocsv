### Example

```csv
id,名字,手机号
1,张三,133xxxx3331
2,李四,133xxxx3332
3,王五,133xxxx3333
```

```go
package main

import (
	"github.com/Sorks/gocsv"
	"log"
)

type User struct {
	Id   int64  `csv:"id"`
	Name string `csv:"名字"`
	// Phone string `csv:"手机号"` // 不需要的字段 直接注释或者不写即可
}

func main() {
	// ... 省略获取文件内容的步骤
	var users []User
	err := gocsv.Unmarshal(bytesContent, &users)
	if err != nil {
		log.Fatalln(err)
	}
}
```