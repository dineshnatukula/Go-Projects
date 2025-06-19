✅ Worker Pool Concept
A job channel to send work to the pool.
A worker goroutine that listens to the job channel.
A result channel to get results back from the workers.
A WaitGroup to wait until all work is done.

📌 Key Concepts Illustrated
Concurrency using goroutines
Synchronization using sync.WaitGroup
Channels for communication
Worker re-use for handling many jobs with fewer goroutines

What happens <-context.Done is used mulitple times?

dinesh.natukula@TM-dineshnatukula MINGW64 ~/Documents/Testing/cmd/testing
$ go run .
Goroutine 1: context done
Goroutine 2: context done
Main: context done

Summary
✅ Yes, you can safely listen to <-ctx.Done() multiple times.
🚫 Don’t write to it.
⚠️ If you block on <-ctx.Done() before the timeout, you'll wait until it fires once. After that, it's instant.