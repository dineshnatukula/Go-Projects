Go (Golang) encourages simplicity, composition, and clarity over heavy object-oriented patterns. That said, several design patterns still apply well — but often in a Go-idiomatic way.

Here’s a list of the most useful and idiomatic design patterns in Go, grouped by category:

✅ 1. Creational Patterns
Factory Pattern
1. Creates instances without exposing construction logic.
2. Used often for creating interfaces with different implementations.

Singleton Pattern
Ensures only one instance of a type is created.

✅ 2. Structural Patterns
2.1 Adapter Pattern
Wraps one interface to match another.
    🧠 Why it's useful
        1. You can use third-party or legacy code without rewriting it.
        2. Keeps your codebase decoupled and interface-based.
        3. Great for plugging in components with mismatched APIs.

2.2 Decorator Pattern
Adds behavior without modifying the original object.

✅ 3. Behavioral Patterns
3.1 Strategy Pattern
Defines a family of interchangeable behaviors.

3.2 Observer Pattern
One-to-many dependency, often used for event systems.

3.3 Command Pattern
Encapsulates requests as objects.

✅ 4. Go-Specific Idioms & Patterns
4.1 Functional Options
Flexible and clean way to configure structs.

4.2 Interface Segregation + Composition
Go favors small interfaces + composition over inheritance.

🔚 Summary of Best Patterns in Go
Pattern	                Go Usage Example
Factory                 Create interfaces with variants
Singleton               Use sync.Once
Adapter                 Convert one interface to another
Decorator               Add features to interfaces dynamically
Strategy                Plug-in behaviors at runtime
Observer                Event-based notifications
Command                 Encapsulate actions
Functional Options      Go idiom for flexible configuration
Interface Composition   Go’s substitute for inheritance