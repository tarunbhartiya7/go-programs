# Bursty Rate Limiter - Multiplication Table Server

## Problem Description

Implement a server that generates a **multiplication table sequence** for a given integer by repeatedly adding the value to the previous result, starting from 0.

### Formula:

For a given integer `toAdd`, the sequence is defined as:

```
(i - 1) * toAdd, where i > 0
```

### Example:

For `toAdd = 4`:

```
0, 4, 8, 12, 16, ...
```

---

## Task

The server should:

* Generate the sequence
* Send results in **batches**
* Apply a **rate limiter**:

  * Each batch must be sent with at least **10 milliseconds delay**

---

## Main Function Inputs

The main function accepts:

* `skipBatches` → Number of batches to skip from the beginning
* `printBatches` → Number of batches to print
* `batchSize` → Number of elements per batch
* `toAdd` → The value used to generate t
