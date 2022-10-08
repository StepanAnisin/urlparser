# urlparser #

Пакет, который считает количество слов "Go" в теле ответа на Get запрос по слайсу строк, содержащих Url

Стоит уточнить, что:

"Golang", "golang" учитываться не будут

но ">Go", "<Go ..." и тд будут

# Локальный импорт #
В go.mod в главной директории:

структура директории:
```
|YourProject
    |--main
    |--urlparser
```

require github.com/StepanAnisin/urlparser v0.0.1

replace github.com/StepanAnisin/urlparser => ../urlparser

# Импорт через github #

```
go get -u github.com/StepanAnisin/urlparser
```

В модуле, в котором используется пакет
```
import (
	"github.com/StepanAnisin/urlparser"
)
```
