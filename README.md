# leet-code-go 🦫

This repository is a collection of Go solutions for various LeetCode problems, organized by problem number and title. Each problem is implemented in its own package, following a consistent structure. This project helps improve Go programming skills through algorithm challenges.

## Project Structure

The workspace follows a modular approach with each LeetCode problem stored as its own Go package. The project includes both the solutions and unit tests for the problems.

```bash
leet-code-go/
│
├── go.mod                        # Go module definition
├── go.sum                        # Dependency checksums
├── p344_reverse_string/          # Example of a problem package
│   ├── README.go                 # Problem info
│   ├── solution.go               # Solution implementation
│   └── solution_test.go          # Unit tests for the solution
└── ...                           # Other LeetCode problem directories
```

## Getting Started

### Prerequisites

Make sure you have [Go](https://golang.org/doc/install) installed on your machine.

To verify that Go is installed, run:

```bash
go version
```

### Clone the Repository

```bash
git clone https://github.com/ismaelpamplona/leet-code-go.git
cd leet-code-go
```

### Running the Tests

Each problem has accompanying tests to verify the correctness of the solution. You can run tests for individual problems or for the entire workspace.

To run the tests for a specific problem, use:

```bash
go test ./p344_reverse_string -v
```

To run all tests in the workspace:

```bash
go test ./... -v
```

### Adding a New Problem

1. **Add the problem to the workspace**:

   Use the `go mod tidy` command to ensure your dependencies are correctly managed.

2. **Create a new package**:

   Add a new folder for the problem and implement the solution.

   ```bash
   mkdir pXXX_problem_name
   touch pXXX_problem_name/solution.go pXXX_problem_name/solution_test.go
   ```

3. **Implement your solution** in the `solution.go` file and write corresponding tests in `solution_test.go`.

4. **Run the tests** to verify your solution works as expected.

### Example Problem Solution

Here's a brief overview of how the problem solution files are structured:

#### `p344_reverse_string/README.md`

````markdown
# p344_reverse_string

[https://leetcode.com/problems/reverse-string/](https://leetcode.com/problems/reverse-string/)

## Initial provided code

```go
func reverseString(s []byte)  {

}
```

## First approach - Two Pointers

- n = number of chars
- time complexity: O(n)
- space complexity: O(1)
````

#### `p344_reverse_string/solution.go`

```go
package p344_reverse_string

func reverseString(s []byte) {
    left, right := 0, len(s)-1
    for left < right {
        s[left], s[right] = s[right], s[left]
        left++
        right--
    }
}
```

#### `p344_reverse_string/solution_test.go`

```go
package p344_reverse_string

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input1 := []byte{'h', 'e', 'l', 'l', 'o'}
	expected1 := []byte{'o', 'l', 'l', 'e', 'h'}
	reverseString(input1)
	if !reflect.DeepEqual(input1, expected1) {
		t.Errorf("Test case 1 failed. Got %v, expected %v", input1, expected1)
	}
}

func TestCase2(t *testing.T) {
	input2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	expected2 := []byte{'h', 'a', 'n', 'n', 'a', 'H'}
	reverseString(input2)
	if !reflect.DeepEqual(input2, expected2) {
		t.Errorf("Test case 2 failed. Got %v, expected %v", input2, expected2)
	}
}

```

## Contribution Guidelines

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Write your solution and tests.
4. Make sure all tests pass.
5. Submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
