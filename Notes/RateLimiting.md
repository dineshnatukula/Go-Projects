Rate Limiting:

Designing a Rate Limiting library in Go is a great systems-level project. Here's a step-by-step guide with design patterns, key decisions, and Go-specific implementations for building a scalable and reusable rate limiter.

✅ Key Features:
Support for fixed window, sliding window, token bucket, or leaky bucket
Easy-to-use API (middleware style, or function call)
In-memory and/or distributed support (e.g., Redis)
Configurable:
Rate (requests per second/minute)
Burst capacity
Per client/IP/user/API key
Optional persistence or sync across nodes (for distributed rate limiting)

| Algorithm          | Description                                  | Use Case               |
| ------------------ | -------------------------------------------- | ---------------------- |
| **Fixed Window**   | Count requests in a fixed time window        | Simple, coarse-grained |
| **Sliding Window** | More accurate per-second rate enforcement    | Precise, heavier       |
| **Token Bucket**   | Allows bursts, refills tokens at steady rate | Best for APIs          |
| **Leaky Bucket**   | Enforces smooth outflow                      | Bandwidth throttling   |

Recommended default: Token Bucket (simple, supports bursts, well-suited to Go channels)









| Feature               | **Token Bucket**                                                     | **Leaky Bucket**                                   |
| --------------------- | -------------------------------------------------------------------- | -------------------------------------------------- |
| **Concept**           | Tokens are added to a bucket at a fixed rate. Requests take a token. | Requests enter a queue. Processed at a fixed rate. |
| **Handles Bursts?**   | ✅ Yes — allows short bursts (up to bucket size)                      | ❌ No — processes at a steady rate only          |
| **Request When Full** | ❌ Rejected if no tokens                                              | ✅ Queued, or dropped if queue is full           |
| **Output Rate**       | Varies, allows bursts if tokens are available                        | Constant output rate                               |
| **Analogy**           | Faucet drips tokens into a bucket. Requests take tokens.             | A funnel leaks water at a fixed rate.              |
| **Latency Impact**    | Low — if tokens exist, processed immediately                         | Can introduce delay due to queuing                 |
| **Implementation**    | More dynamic; tracks time-based refill and usage                     | Simpler; fixed rate dequeuing of requests          |
| **Use Case**          | APIs, microservices, user actions (e.g., login attempts)             | Network traffic shaping, pacing packets            |

🧠 Visual Analogy
Token Bucket: You have a bucket that fills with tokens at a steady rate. Each request "pays" 1 token. If enough tokens are saved, you can send a burst of requests.
Leaky Bucket: You pour requests into a bucket. They leak out at a steady rate. Excess requests overflow and are dropped (or queued).

✅ Summary
| Goal                     | Use This Algorithm |
| ------------------------ | ------------------ |
| Allow bursts             | Token Bucket       |
| Enforce steady rate      | Leaky Bucket       |
| Simple, constant pacing  | Leaky Bucket       |
| Flexibility with control | Token Bucket       |

