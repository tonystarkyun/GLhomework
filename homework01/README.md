# Homework 01 Solutions

Each problem in `homework01.md` has its own directory with:
- `solution.go`: implementation
- `solution_test.go`: table-driven test cases

## Problem Directories
- `single_number`
- `palindrome_number`
- `valid_parentheses`
- `longest_common_prefix`
- `plus_one`
- `remove_duplicates_sorted_array`
- `merge_intervals`
- `two_sum`

## Test Commands
Run all tests:

```bash
cd homework01
go test ./... -v
```

Run a single problem test:

```bash
cd homework01
go test ./single_number -v
go test ./palindrome_number -v
go test ./valid_parentheses -v
go test ./longest_common_prefix -v
go test ./plus_one -v
go test ./remove_duplicates_sorted_array -v
go test ./merge_intervals -v
go test ./two_sum -v
```
