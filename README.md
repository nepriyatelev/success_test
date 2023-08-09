<h1>Тестовое задание</h1>

<h2>Как скачать?</h2>

```bash
git clone https://github.com/nepriyatelev/success_test.git
```

<h2>Как собрать?</h2>

```bash
make
```
Или используйте Makefile для более точного управления сборкой.
- make build - сборка проекта
```bash
make build
```
- make clean - удаление собранных файлов
```bash
make clean
```

<h2>Как запустить?</h2>

При запуске приложения, вы можете указать следующие флаги:
- -clients=N: Количество клиентов для тестирования (по умолчанию 1).
- -log=level: Уровень логирования (debug, info, warn, error, fatal, panic). По умолчанию - info.

пример:
```bash
./test_app -clients=5 -log=debug
```
запуск приложения с 5 клиентами и уровнем логирования debug.

```bash
./test_app
```
запуск приложения с 1 клиентом и уровнем логирования info.

<h2> Также можно запустить Linter </h2>

Предварительно скачав и установив его:
```bash
https://golangci-lint.run/usage/install/#local-installation
```
Запустить линтер:
```bash
run golangci-lint
```