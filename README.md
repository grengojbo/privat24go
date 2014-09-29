privat24go [![Coverage Status](https://coveralls.io/repos/grengojbo/privat24go/badge.png)](https://coveralls.io/r/grengojbo/privat24go) [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/grengojbo/privat24go/blob/master/LICENSE) [![Build Status](https://travis-ci.org/grengojbo/privat24go.svg?branch=master)](https://travis-ci.org/grengojbo/privat24go)
====

## Описание

Конвертация выписок Приват24 (Excel) в структуру Go
Структура выписки Privat24 Юр.лицо [Ordering](https://github.com/grengojbo/privat24go/blob/master/ordering.go).
После обработки можно применить [фильтры](https://github.com/grengojbo/privat24go/blob/master/filters.go) по обработке *Назначение платежа*.

## Использование

Параметр **name** путь к файлу выписки *c2bstatements*

Параметр **f** загружать:
    - 0:  все данные
    - 1: только поступления
    - 2: только выплаты

Параметр **cl** очисть поле *Назначение платежа* от мусора

<количество записей>, <массив выписки>, <ошибка> := LoadXlsFile(name string, f int, cl bool)

Запуск тестов
```bash
$ make test
```

## Установка

Для установки, использовать `go get`:

```bash
$ go get -v github.com/kr/godep
$ go get -d github.com/grengojbo/privat24go
```

## Вклад

1. Fork ([https://github.com/grengojbo/privat24go/fork](https://github.com/grengojbo/privat24go/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## Автор

[Oleg Dolya](https://github.com/grengojbo)