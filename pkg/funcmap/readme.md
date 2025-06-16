在 golang 中用函数名字调用某个函数。

在 golang 中，你不能这样做：

```go
func foobar() {
    // bla...bla
}
funcname := "foobar"
funcname()
```

不过可以：

```go
func foobar() {
    // bla...bla
}
funcs := map[string]func(){
    "foobar": foobar,
}
funcs["foobar"]()
```

这里有一个限制：这个 `map` 仅仅可以用原型是 `func()` 的没有输入参数或返回值的函数。

如果想要用这个方法实现调用不同函数原型的函数，需要用到 `interface{}`。

```go
func foo() {
    // bla...bla
}
func bar(a, b int) {
    // bla...bla
}
funcs := map[string]interface{}{
    "foo": foo,
    "bar": bar,
}
```

但是你不能这样调用 `map` 中的函数：

```go
funcs["foo"]()
```

这无法工作！你不能直接调用存储在空接口中的函数。

```go
func Call(m map[string]interface{}, name string, params ... interface{}) (result []reflect.Value, err error) {
    f := reflect.ValueOf(m[name])
    if len(params) != f.Type().NumIn() {
        err = errors.New("The number of params is not adapted.")
        return
    }
    in := make([]reflect.Value, len(params))
    for k, param := range params {
        in[k] = reflect.ValueOf(param)
    }
    result = f.Call(in)
    return
}
Call(funcs, "foo")
Call(funcs, "bar", 1, 2, 3)
```

将函数的值从空接口中反射出来，然后使用 `reflect.Call` 来传递参数并调用它。

- TODO: 带有返回值的函数；
