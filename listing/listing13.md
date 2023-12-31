Что выведет данная программа и почему?

```go
  func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
  }
  func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
  }
```

**Ответ:**

`[100 2 3 4 5]`

При вызове функции `someAction`, она изменяет элемент с индексом `0` в слайсе `a` на значение `100`. 
Это происходит, потому что слайсы в **Go** передаются по ссылке, и функция `someAction` имеет доступ к тем же данным, на которые указывает слайс `a`.

Однако, при добавлении элемента в слайс `v` с помощью `append`, создается новый слайс, который больше по размеру и не связан с исходным слайсом `a`. 
Поэтому изменения, произведенные внутри функции `someAction`, касаются только исходного слайса `a` в том случае, когда изменяется существующий элемент.

