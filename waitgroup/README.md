# Golang WaitGroup
> Wait for a Group of goroutines to finish

- Step 1: create a variable of type sync.WaitGroup: wg
- Step 2: Add the number of goroutines to wait in the newly initialized variable wg.Add(quantity)
- Step 3: at each goroutine call wg.Done() method before return
- Step 4: call wg.Wait() where you need to wait for the goroutine to finish running