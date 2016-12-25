# go-hash
各种hash函数



- RSHash
- JSHash
- PJWHash
- ELFHash
- BKDRHash
- SDBMHash
- DJBHash
- DEKHash
- APHash
- FNVHash
- MYSQLHash



使用示例：

```go
import "github.com/chen-zyc/go-hash"
import "fmt"

func main() {
	for _, f := range hash.AllFuncs {
		fmt.Println(f("hello world"))
	}
}
```

