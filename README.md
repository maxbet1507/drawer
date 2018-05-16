# drawer
temporary values store for golang

## インストール

```
go get github.com/maxbet1507/drawer
```

## 説明

とりあえず値を一箇所にまとめて放り込んで、型ベースで取り出す箱です。
実直に型毎のスライスで実装するのが面倒な場合に、便利かもしれません。

```go
func Example() {
	d := drawer.New()

	d.Push(123, "hello", true)
	d.Push(fmt.Errorf("world"), false, 456)

	rv1 := []int{}
	d.Pull(&rv1)
	fmt.Println(rv1)

	rv2 := []string{}
	d.Pull(&rv2)
	fmt.Println(rv2)

	rv3 := []error{}
	d.Pull(&rv3)
	fmt.Println(rv3)

	// Output:
	// [123 456]
	// [hello]
	// [world]
}
```
