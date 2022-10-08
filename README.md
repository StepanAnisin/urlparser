# urlparser #

Пакет, который считает количество слов "Go" в теле ответа на Get запрос по слайсу строк, содержащих Url

Стоит уточнить, что:

"Golang", "golang" учитываться не будут

но ">Go", "<Go ..." и тд будут

# Локальный импорт #
В go.mod в главной директории:

структура директории:

|YourProject
    |--main
    |--urlparser

require github.com/StepanAnisin/urlparser v0.0.0

replace github.com/StepanAnisin/urlparser => ../urlparser
