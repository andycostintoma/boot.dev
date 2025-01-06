import time

# Recursive Fibonacci
def fib(n):
    if n == 0:
        return 0
    if n == 1:
        return 1
    return fib(n - 1) + fib(n - 2)

# Iterative Fibonacci (polynomial time)
def fib_polynomial(n):
    if n <= 1:
        return n
    current = 0
    parent = 1
    grandparent = 0
    for _ in range(0, n - 1):
        current = parent + grandparent
        grandparent = parent
        parent = current
    return current

n = 30

# Time for recursive function
start_time_recursive = time.time()
result_recursive = fib(n)
end_time_recursive = time.time()

# Time for iterative function
start_time_iterative = time.time()
result_iterative = fib_polynomial(n)
end_time_iterative = time.time()

# Display results
print(f"Recursive fib({n}): {result_recursive}, Time: {end_time_recursive - start_time_recursive:.6f} seconds")
print(f"Iterative fib_polynomial({n}): {result_iterative}, Time: {end_time_iterative - start_time_iterative:.6f} seconds")
