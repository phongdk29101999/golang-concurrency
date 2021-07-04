# Golang Deadlock
> is a state that occurs in a multithreaded environment when two or more processes go into a loop waiting for resources forever.

Deadlock occurs in go when data is sent to the channel but no goroutine is received or no data is sent to the channel for the goroutine to receive.