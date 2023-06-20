

饿汉模式
```go

type Singleton struct {
}

var eagerSingleton *Singleton

func init() {
    eagerSingleton = &Singleton{}
}

func GetEagerInstance() *Singleton {
    return eagerSingleton
}

```



懒汉模式
```go
var (
    lazySingleton *Singleton
    once          = &sync.Once{}
)

func GetLazyInstance() *Singleton {
    if lazySingleton == nil {
        once.Do(func() {
            lazySingleton = &Singleton{}
        })
    }
    return lazySingleton
}
```

测试文件
```go
func TestGetInstance(t *testing.T) {
	assert.Equal(t, GetEagerInstance(), GetLazyInstance())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetEagerInstance() != GetEagerInstance() {
				b.Errorf("test fail")
			}
		}
	})
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazyInstance() != GetLazyInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
```

```go
//启动命令
go test -benchmem -bench="." -v
```
