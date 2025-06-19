1. Reducing the startup time of the service. (PTL Worker, Identifying the unwanted topics that are consumed by the service).
2. Implemented Idempotency and Reconciliations at transaction levels.
3. Retry asynchronously for n times based on Fib / Exponential time
4. Resolving the memory leak in one of the services
    In go, there is a pprof that gives us from where exactly the memory is happening.
    Here most of the times, it was like global variable declaration happened in one of the service.
5. Graceful shutdown of the service.
6. Detected that the service not restarted as there is an issue in the logger library.
7. Implemented 0 downtime for couple of the services during the onboardings
8. 


Stop consumption of events on reporting services which improves the perf of the service and kafka cluster
PTL-Worker reducing the startup time from 15-20 minutes to very few seconds
SQL to PL/SQL stored procedures / functions
Common Logging for all the services
Flag Migrations from Devops to service
