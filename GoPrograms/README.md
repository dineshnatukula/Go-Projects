<H1>How Does Context Work?</H1>
<b>Creating a Context:</b>
You typically create a context using context.Background() or context.TODO(), both of which represent a root context.
<b>Canceling Context:</b>
You can derive a new context with a cancellation mechanism from the root context or from another context using context.WithCancel(). You cancel the context when the task is done or should be aborted.
<b>Timeout or Deadline:</b>
You can also create a context with a timeout or deadline using context.WithTimeout() or context.WithDeadline().
Passing Context:</b>
Context is usually passed as the first argument in functions that are part of an operation that might need cancellation, timeout, or shared data.

<H1>Common Functions in the context Package:</H1>
<b>context.Background():</b> Returns an empty context. Typically used at the top level to create a root context.
<b>context.TODO():</b> Used as a placeholder for contexts when you're unsure whether to use Background() or another type of context. It should be replaced when the code is finished.
<b>context.WithCancel():</b> Returns a new context and a cancel function. The context can be canceled manually by calling the cancel function.
<b>context.WithTimeout():</b> Creates a context that will automatically cancel after a specified timeout.
<b>context.WithDeadline():</b> Similar to WithTimeout() but allows you to specify an exact point in time when the context should be canceled.

<H1>Use Cases for Context in Go:</H1>
<b>Timeout Management in Network Requests:</b>
If you're making HTTP requests or connecting to external services, you can use context with a timeout to ensure the request doesn't hang indefinitely.
<b>Managing Long-Running Operations:</b>
In long-running functions or background tasks, you can use context to allow cancellation if the work is no longer needed or the user cancels the operation.
<b>Distributed Systems:</b>
In systems where multiple services interact, context is used to pass information (like trace IDs) and cancel signals between services.

<H1>About Context</H1>
<b>Cancellation & Timeouts:</b> Context helps manage cancellation and timeouts in concurrent programs.
<b>Graceful Shutdown:</b> It allows you to gracefully shut down tasks and clean up resources.
<b>Request Scoped Values:</b> You can carry data across function calls and goroutines, such as authentication tokens or request IDs.
Credits: Chatgpt

<H1>Interface in go:</H1>
✅ Interfaces in Go — Explained Simply
In Go, an interface is a type that specifies a set of method signatures. When a type provides definitions for all the methods declared in an interface, it implicitly implements that interface — no explicit declaration needed.
<H1>🔧 Why Use Interfaces?</H1>
Interfaces allow you to write flexible and decoupled code. You can write functions that accept interface types and work with any type that implements the required methods.