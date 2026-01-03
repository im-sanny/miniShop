# Go Project Learning Notes

> Personal learning notes while building this project.
> Not a tutorial. Rules here exist because I broke them before.

---

## A. HTTP, Requests & Basics (Things I Touch Every Day)

### Debugging & Output

- Use `fmt.Println` to see output in the terminal for easy debugging/logging.
- Use `fmt.Fprintf(w, ...)` to send output to the user/browser.

### Request Values Are Always Strings

- HTTP is a **text-based protocol**.
- Path params, query params, and form values are **always strings**.

```go
idStr := r.PathValue("itemId")
id, err := strconv.Atoi(idStr)
```

Reason:

- URL values arrive as strings
- Database IDs are integers
- Convert string → int before comparing or querying DB

### Environment Variables

- Environment variables are **always strings**.
- Convert them to other formats (`int`, `bool`, etc.) when required.

---

## B. Middleware & Function Mechanics

### Variadic Parameters

```go
func Use(middlewares ...Middleware)
```

- `...` after the type means the function accepts **any number** of arguments.
- Inside the function, it behaves like `[]Middleware`.

### Slice Expansion (Spreading)

```go
append(mngr.globalMiddlewares, middlewares...)
```

- `...` after a slice variable expands the slice.
- Required because `append` expects multiple elements, not a slice.

---

## C. Handler Flow & Application Logic

### Create User Handler Flow

- Client sends request with correct data
- `CreateUserHandler`:

  - Decodes JSON into a `User` struct
  - Sends the user to `StoreUser`

- `StoreUser`:

  - Adds a user ID
  - Stores user
  - Returns the updated user

- Handler sends the updated user in the response

---

## D. Go Language Rules & Footguns

### Range Loop Pointer Rule (CRITICAL)

🚫 Never return the address of a range variable:

```go
for _, u := range slice {
    return &u
}
```

✅ Always use:

```go
for i := range slice {
    return &slice[i]
}
```

Reason:

- Range variables are reused
- Taking their address causes subtle bugs

---

### Constructors & Naming Conventions

- Functions that create and return a struct should usually be named:

  ```go
  NewXxx
  ```

- If the package already provides context, don’t repeat it in function names.

- Package name gives context, function names should be short.

- Receiver name is for mechanics, not meaning.

---

### Create Function Rule

🚫 Create should **never** return `(nil, nil)`.

- That state is a bug, not a valid state.

---

## E. Database & SQL Notes

### Postgres Driver Error

Error:

```
sql: unknown driver "postgres" (forgotten import?)
```

Fix:

```go
_ "github.com/lib/pq"
```

Reason:

- `pq` is a pure Go Postgres driver
- Registers itself with `database/sql` via side effects
- Used even if not referenced directly

---

### SQL Timestamp

Instead of:

```sql
TIMESTAMP WITH TIMEZONE DEFAULT CURRENT_TIMESTAMP
```

Use:

```sql
TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
```

---

## F. Method Receivers

### Use Pointer Receivers When:

- Method needs to modify the receiver
- Receiver is large and copying is inefficient
- Consistency is needed with other pointer receiver methods

### Use Value Receivers When:

- Method doesn’t modify the receiver
- Receiver is small
- Method should work on a copy

---

## G. Concurrency: Hard Rules First

### HTTP Handler Rules

- Handlers must be stateless
- Never use globals for request data
- `http.ResponseWriter` is **NOT goroutine-safe**
- Never write to `ResponseWriter` inside goroutines

---

### Mutex Rule

🚫 If you need a mutex in an HTTP handler, your design is wrong.

Mutexes are for:

- Caches
- In-memory maps
- Shared resources

🚫 Not for request data

---

## H. WaitGroup — Correct Usage Only

### Correct Reasons to Use WaitGroup

Use it only if:

- You need results from multiple independent operations
- You wait for all of them
- You don’t write responses inside goroutines
- Results are stored in outer variables

---

### When You SHOULD NOT Use WaitGroup

- Only one thing is concurrent
- You don’t need the result immediately
- You’re just trying to be “idiomatic”
- You write HTTP responses in goroutines
- You share global state

---

### WaitGroup Mental Checklist (MEMORIZE)

Before using `WaitGroup`, ask:

- Do I need results from multiple independent operations?
- Are results stored in outer variables?
- Do I call `wg.Wait()` before reading them?
- Do I avoid `:=` inside goroutines?
- Do I write the HTTP response only once?

If all answers are yes → model is correct.

---

## I. Goroutines, Races & Mental Models

### Goroutines & Return Values

- Goroutines don’t return values
- Functions return values
- When using goroutines:

  - Provide a place to store results
  - Read results only after `Wait()`

---

### Data Race Rule (Forever)

A data race exists if ALL are true:

1. Two or more goroutines
2. Access the same memory location
3. At least one access is a write
4. No synchronization (mutex / channel / atomic)

If goroutines touch shared variables → assume a race.

---

### Concurrency Design Hierarchy

🥇 No shared state (pass data, return results)
🥈 Local ownership (one request = one owner)
🥉 Synchronization (last resort)

If you start at 🥉, you already failed.

---

## J. Choosing the Right Tool

### Primitive Responsibilities

- `WaitGroup` → waits for goroutines
- `Mutex` → protects shared memory
- `Channel` → communicates between goroutines

---

### Decision Table

| Situation             | Use          |
| --------------------- | ------------ |
| High QPS handler      | None         |
| Shared mutable state  | ❌ Redesign  |
| Independent DB calls  | errgroup     |
| Simple parallel tasks | WaitGroup    |
| Pipelines / workers   | Channels     |
| Counters              | atomic       |
| Caching               | mutex (rare) |

---

## K. Golden Rules (Read When Stuck)

- Concurrency ≠ performance
- Remove races, don’t lock them
- WaitGroup waits — it doesn’t protect
- Mutex protects — it doesn’t design
