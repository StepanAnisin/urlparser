# urlparser #

Пакет, который считает количество слов "Go" в теле ответа на Get запрос по слайсу строк, содержащих Url

Стоит уточнить, что:

"Golang", "golang" учитываться не будут

но ">Go", "<Go ..." и тд будут

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
require github.com/StepanAnisin/urlparser v0.0.1

replace github.com/StepanAnisin/urlparser => ../urlparser
```

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
# Вызов функции, принимающей слайс, содержащий Url (в формате string) #

Пример:
```
urls := [...]string{"https://golang.org", "https://golang.org"}
slice := urls[:]
GetGoWordCountByUrls(slice)
```
