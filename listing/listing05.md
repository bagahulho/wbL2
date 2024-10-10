Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Будет выведено: error

Это произойдет потому, что у интерфейса error будет определена ссылка на тип для интерфейса, 
так как функция test возвращает указатель на структуру CustomError

```