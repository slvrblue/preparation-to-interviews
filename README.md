# preparation-to-interviews

###### Примеры решенных задачек с собеседований

1. Необходимо написать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:

- `aaaabccddddde` ->`a4bc2d5e`
- `abcd` -> `abcd`
- `a/` -> "" (некорректная строка)
- `aaa10b` -> "" (некорректная строка)
- "" -> "" (некорректная строка)

В случае если была передана некорректная строка функция должна возвращать ошибку.

Реализация: [compress](https://github.com/boshnyakovich/preparation-to-interviews/tree/main/compress)

2. Необходимо написать Go функцию, осуществляющую запуск каждой задачи

- `Task1` -> `err`
- `Task2` -> `nil`
- `Task3` -> `err`

в своей горутине и параллельно осуществлять чтение этих задач:

`process([Task1, Task2, Task3])` -> `[err1, err2]`

Реализация: [tasks](https://github.com/boshnyakovich/preparation-to-interviews/tree/main/tasks)
