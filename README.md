# urlparser #

Пакет, который считает количество слов "Go" в теле ответа на Get запрос по слайсу строк, содержащих Url

Стоит уточнить, что будут учитываться точные совпадения. Например:

"Golang", "golang" учитываться не будут

но "Go", ">Go", "<Go ..." и тд будут

# Локальный импорт #

структура директории:
```
|YourProject
    |--main
    |--urlparser
```

В go.mod в директории YourProject/main:

Заменить путь к пакету на локальный

```
require github.com/StepanAnisin/urlparser v0.0.0

replace github.com/StepanAnisin/urlparser => ../urlparser
```

# Импорт через github #

В модуле, в котором используется пакет добавить:
```
import (
	"github.com/StepanAnisin/urlparser"
)
```

Далее:

```
go get -u github.com/StepanAnisin/urlparser
```
или

```
go mod tidy
```
# Вызов функции, принимающей слайс, содержащий Url (в формате string) #

Пример:
```
urls := [...]string{"https://golang.org", "https://golang.org"}
slice := urls[:]
result := urlparser.GetGoWordCountByUrls(slice)
```
