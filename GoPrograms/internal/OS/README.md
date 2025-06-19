🧠 What is a Semaphore in OS?
A semaphore is a synchronization mechanism used in operating systems to manage concurrent access to shared resources (like files, memory, printers) by multiple processes or threads.

Semaphores help prevent race conditions, deadlocks, and inconsistent states in multithreaded environments.

🏷️ Types of Semaphores
1. Counting Semaphore
Value can range over an unrestricted domain (typically 0 to N).
Used to control access to a resource with multiple instances (like a pool of connections).

2. Binary Semaphore (also called a Mutex)
Value is either 0 or 1.
Works like a simple lock/unlock mechanism.
Used for mutual exclusion (only one process can enter the critical section).

🧮 How Semaphores Work
A semaphore has:
An integer value (initial count)
Two atomic operations:

✅ Real Use Cases
Ensuring mutual exclusion in critical sections
Controlling access to resource pools
Order enforcement (process A must run before B)
Handling producer-consumer problems
Preventing deadlocks (if used carefully)


🧠 Race Condition in OS
A Race Condition in Operating Systems occurs when two or more processes or threads access shared data concurrently, and the final result depends on the timing of their execution.
If the processes are not properly synchronized, it can lead to unexpected behavior or bugs.

🧠 Real-World Analogy
Imagine two people writing on the same whiteboard at the same time — if they don’t coordinate, the result will be messy or incorrect.

🔥 Classic Example: Race Condition in Code
Shared Counter Example (in C-like pseudocode):

int counter = 0;
Thread 1:                Thread 2:
temp = counter;          temp = counter;
temp = temp + 1;         temp = temp + 1;
counter = temp;          counter = temp;

You expect counter to increase by 2.
But due to overlapping execution, it may increase only by 1. This is a race condition.

✅ When Do Race Conditions Happen?
Accessing shared memory without proper locks
I/O operations (e.g., logging, file writing)
Poor use of threads in parallel programs
Interrupts pre-empting system calls

🔐 How to Prevent Race Conditions
1. Mutual Exclusion
Use mutexes, locks, or semaphores to ensure only one thread accesses the shared resource at a time.
2. Atomic Operations
Use atomic instructions (like atomic_add) to modify shared variables safely.
3. Critical Sections
Protect the code that accesses shared data inside a critical section.
4. Thread Synchronization Primitives
Use tools like:
Mutexes
Semaphores
Monitors
Condition variables