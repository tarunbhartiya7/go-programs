# TCP Server - Reverse String

## Problem Description

Implement a **TCP server** that:

* Accepts messages as strings from a client
* Reverses each message
* Sends the reversed string back to the client

---

## Task

The provided stub function:

* Accepts an array of strings `messages`
* Starts the TCP server
* Waits until the server is ready
* Sends messages one by one
* Prints the responses received from the server

The server configuration includes:

* `address` → TCP server address
* `maxBufferSize` → Buffer size for reading/writing

---

## Example

### Input:

```id="0o3q1b"
messages = ["adi", "glm", "tuv"]
```

### Output:

```id="n7m2lg"
ida
mlg
vut
```

---

## Function Description

Implement the function:

```id="e0c6gn"
TCPServer(ready chan bool)
```

### Parameters:

* `ready` → A channel that indicates when the server is ready to accept connections

---

## Return Value

* None
* All responses are handled via the TCP connection between client and server

---

## Constraints

* Maximum number of messages: **500**
* Each message:

  * Contains only lowercase English letters
  * Maximum length: **1000 characters**

---

## Input Format (Custom Testing)

```id="av0ix4"
n
messages[0]
messages[1]
...
messages[n-1]
```

---

## Sample Case 0

### Input:

```id="4fmyhx"
3
abc
def
ghi
```

### Output:

```id="y3r0pf"
cba
fed
ihg
```

### Explanation:

* Each string is reversed by the server
* The reversed string is returned to the client

---

## Sample Case 1

### Input:

```id="19b7af"
4
eno
owt
eerht
ruof
```

### Output:

```id="4d9k6t"
one
two
three
four
```

### Explanation:

* Each input string is reversed
* The server sends back the corrected (reversed) string

---
