# Modulo Fibonacci Sequence Server

## Problem Description

A **Fibonacci with modulo sequence** is defined for a given modulo `mod` as:

* `fib[x] = (fib[x - 1] + fib[x - 2]) % mod`, for `x > 2`
* `fib[1] = 1`, `fib[2] = 2`

The sequence starts as:

```
[1, 2, 3, 5, 8, 13, 21, ...]
```

---

## Task

Implement a function that runs a **server** to generate a Fibonacci sequence.

### Requirements:

* The server:

  * Accepts **boolean requests**
  * Returns the **next Fibonacci number**
* Must include a **rate limiter**:

  * Each response must be sent **at least 10 milliseconds after the previous one**

---

## Function Signature

```
ModuloFibonacciSequence(requestChan chan bool, resultChan chan int)
```

### Parameters:

* `requestChan`: Channel of boolean values (used to trigger requests)
* `resultChan`: Channel through which Fibonacci results are returned

---

## Main Function Behavior

The main function takes two inputs:

* `skip`: Number of Fibonacci elements to skip from the beginning
* `total`: Number of Fibonacci elements to
