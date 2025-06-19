✅ Graceful shutdown patterns are useful for:
HTTP servers (http.Server.Shutdown(ctx))
Goroutine pools / background workers
Message queue consumers
Long-running I/O operations

A graceful shutdown for an HTTP server?
🔍 Features:
🧵 Worker pool with graceful exit on context.Done()
🌐 HTTP /enqueue endpoint to simulate user requests
🚦 Signal handling for graceful shutdown (SIGINT, SIGTERM)
🛑 Timeout-based shutdown fallback

curl http://localhost:8080/enqueue
go run main.go

