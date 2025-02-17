# Silent Map Access in Go

This repository demonstrates a common but subtle issue in Go: accessing a non-existent key in a map does not result in a runtime error. Instead, it silently returns the zero value for the map's value type.

This behavior can be problematic because it can lead to unexpected results and hard-to-debug bugs.  The example code shows how this can occur and a solution to mitigate this issue.

## How to reproduce:

1. Clone this repository.
2. Navigate to the directory in your terminal.
3. Run `go run bug.go`.

You'll notice that the program doesn't panic, but instead prints 0.

## Solution:

The `bugSolution.go` file demonstrates a safer approach using the comma ok idiom.