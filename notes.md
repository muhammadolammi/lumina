
Learnings

In the worker when we comment out the safety update and just update the count directly (stats.logsProcessed++), we introduce a data race condition
We can capture this by running "go run -race ."


Why is time.After inside a loop potentially dangerous for memory?

Hint: Every time the loop repeats, a new timer is created. If the loop is fast, you'll leak memory with thousands of timers. (We’ll solve this properly when we get to Context).