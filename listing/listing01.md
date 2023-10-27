Какой самый эффективный способ конкатенации строк?

**Ответ:**

Есть 3 способа объединить строки:
1. Конкатенация с использованием `+`:
```go
  foobar := "foo" + "bar"
```
3. Конкатенация с использованием `strings.Builder`:
```go
  var sb strings.Builder
  sb.WriteString("foo")
  sb.WriteString("bar")

  sb.String()
```
3. Конкатенация с использованием `bytes.Buffer`:
```go
  var bb bytes.Buffer
  bb.WriteString("foo")
  bb.WriteString("bar")
  
  bb.String()
```

Ссылаясь на [данную статью](http://test.com/](https://daryl-ng.medium.com/efficiently-concatenate-strings-in-go-a90dfa5161fd)https://daryl-ng.medium.com/efficiently-concatenate-strings-in-go-a90dfa5161fd), где автор произвёл бенчмарк всех 3-х способов, результат оказался следующий:
- Если требуется производительность: `Конкатенация с использованием strings.Builder`
- Если требуется минимум кода и простота ( придётся пожертвовать производительностью ): `Конкатенация с использованием "+"`
