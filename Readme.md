Домашнее задание №1

Тема: текстовая игра

Основная цель задания - поупражняться в моделировании объектов. Сверх-сверх детальное решение по шагам будет позже.

Когда-то давно, когда компьютеры были большими, интернет - медленным, а графических ускорителей не было вовсе, уже существовали многопользовательские игры. Они имели текстовый интерфейс и назывались MUD, Multi User Dungeon.

Мы пишем простую игру, которая реагирует на команды игрока, возвращая нужный ответ. Тут можно начать смотреть команды-ответы в файле main_test.go

Игровой мир обычно состоит из комнат, где может происходить какое-то действие.
Так же у нас есть игрок.
Как у игрока, так и у команты есть состояние.
Функция `initGame` делает нового игрока и задаёт ему начальное состояние, а так же сбрасывает состояние комнат.
Функция `handleCommand` получает команду и реализует логику на основе этой команды.
В данной версии можно обойтись глобальными переменными для игрока и мира ( команат ).

Список команд:
* осмотреться - выводит текущее окружение, на основе комнаты в которой находится игрок
* идти - перемещает игрока в другую комнату, если это возможно (с этим будет логика)
* надеть - позволяет надеть рюкзак
* взять - кладет предмет в рюкзак
* применить - применяет предмет из рюкзака к чему-то (а у этого чего-то отрабатывает логика) - используется в открытии двери

Список комнат:
* кухня - тут есть предметы и квест
* коридор - тут есть дверь, для которой нужен ключе
* комната - тут есть предметы
* улица - финальная точка игры :)  

Команда в handleCommand парсится как
$команда $параметр1 $параметр2 $параметр3
https://golang.org/pkg/strings/#Split вам в помощь

Детали по ответам везде смотрите в main_test.go. Очень помогает нарисовать схему программы из комнаты и их свойств, что куда. Прямо квадратиками на бумажке :)

В тестах представлены последовательности команд и получаемый ответ.
Задача - пройти все тесты и сделать правильно.
Под правильным понимается универсально, чтобы можно было без проблем что-то добавить или убрать.
Т.е. бесконечный набор захардкоженных if'ов для всего мира не подойдёт.
Конкретные условия могут быть только внутри конкретной комнаты.
Надо думать в сторону объектов, вызова функций, структур, которые описывают состояние комнаты и игрока, функций которые описывают какой-то интерактив в комнате. Не забывайте что вы можете создать мапу из функций. Или можно реализовать триггер (действие, выполняемое при каком-то событии). Или у структуры поле может иметь тип "функция" - тут надо применять анонимные функции - смотрите 2/functions/firstclass.go

Хардкором (набором if-ов без нормального моделирования структур) это задание пишется за 3 часа. Но хардкором нельзя! Нормально вдумчиво - чуть дольше.

Глобальная мапа с полной командой от юзера - это тоже считается за хардкод. Полного текста команды у вас быть не должно - все команды надо разделять на части, а потом на основе этих частей выполнять какие-то дествия.

Тестовых кейсов много. Прочитайте их внимательно, там есть результаты работы всего что вам надо.
Не стесняйтесь задавать вопросы.
Хитрой логики тут нет, алгоритмов тоже. Только вызов методов, сохранение состояния, условия.

В идеале ваша архитектура с комнатами должна без проблем прежить добавление дополнительных комнат. Внутри общих методов не должно быть завязок на конкретную комнату, т.е. `if room == "кухня" {` нельзя. Все данные для построения ответа должны браться из структуры комнаты и задаваться в месте ее инициализации.

Документация по стандартной библиотеке языка: https://golang.org/pkg/ 

Код надо писать в main.go, если требуется - можно создавать дополнительные файлы.
main_test.go править не надо. Средний объем решения при хардкорной реализации - 250-300 строк, при правильной реализации - 350-450 строк. Но не пытайтесь делать космолет - нет задачи сделать идеальную супер-конфгурируемую архитектуру. Прохождения тестов без костылей с хардкодом достаточно. В сторону тоже не уходите - впереди еще много интересных домашек.

Запускать тесты через `go test -v` находясь в папке `game`.
Если вы ходите реализовать все для запуска через консоль - реализуйте считывание команды в функции main - пример можно посмотреть в программе уникализации. Если вы будете реализовывать в нескольких файлах - то при запуске через `go run` надо передать их все, или написать `.` (точка) вместо имени файла - `go run .` 

В этом задании вам понадобятся:
* функции
* структуры
* методы структур
* анонимные функции и поле-функция в структуре
* слайсы и мапы
* применение пакета strings
* условия и циклы, тут хорошо зайдет switch-case в одном месте

Вам НЕ понадобятся:
* интерфейсы
* пустые интерфейсы
* разделение по пакетам (все надо делать в одном пакете)
* регулярные выражения
