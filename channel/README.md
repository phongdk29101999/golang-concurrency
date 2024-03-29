# Golang Channel

In Go language, a channel is a medium through which a goroutine communicates with another goroutine and this communication is lock-free. Or in other words, a channel is a technique which allows to let one goroutine to send data to another goroutine. By default channel is bidirectional, means the goroutines can send or receive data through the same channel as shown in the below image.

![screenshot](https://github.com/phongdk29101999/golang-concurrency/blob/master/screenshots/channel.png)

## Channel Locking

When the goruntime runs to where it needs to get data from the channel, the code execution will be locked until the goroutine sends data to the channel and the main function receives the data from the channel => no need to waitgroup to wait for the goroutine to finish running.

![screenshot](https://github.com/phongdk29101999/golang-concurrency/blob/master/screenshots/channel-locking.png)
