# Golang Concurrency
> Learn about golang concurrency

## What is Concurrency?
"Concurrency is about dealing with lots of thing at once" - Rob Pike

![screenshot](https://github.com/phongdk29101999/golang-concurrency/blob/master/screenshots/concurrency.png)

For example, with 3 tasks displaying, playing music, downloading, the system will divide large tasks into smaller tasks. Then sometimes doing one task, sometimes doing another, and switching between tasks so fast makes us feel like the system is handling many things at a time. In fact, the system can only do one thing at a time. And in Go, each task is handled by a Goroutine.

## Goroutine

- Goroutine is a function that is processed concurrently with other goroutines.
- Can be considered as a small thread, can increase or decrease the memory area according to the needs of use.

![screenshot](https://github.com/phongdk29101999/golang-concurrency/blob/master/screenshots/goroutine.png)

## Sections
- [Goroutine](https://github.com/phongdk29101999/golang-concurrency/tree/master/goroutine)
- [Waitgroup](https://github.com/phongdk29101999/golang-concurrency/tree/master/waitgroup)
- [Channel](https://github.com/phongdk29101999/golang-concurrency/tree/master/channel)
