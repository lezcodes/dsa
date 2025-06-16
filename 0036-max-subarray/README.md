# max-subarray

## Description

Implementation of the Maximum Subarray Problem using Kadane's Algorithm (Dynamic Programming approach).

Given an integer array, finds the contiguous subarray with the largest sum and returns both the maximum sum and the subarray itself along with its indices.

The algorithm works by maintaining two variables:

- `currentSum`: maximum sum ending at the current position
- `maxSum`: overall maximum sum found so far

At each position, we decide whether to extend the existing subarray or start a new one based on whether the current sum is positive.

## Complexity

- Time Complexity: O(n) - single pass through the array
- Space Complexity: O(1) for the algorithm itself, O(k) for storing the result subarray where k is the length of the maximum subarray

## Usage

```bash
make run n=0036-max-subarray
```

## Testing

```bash
make test n=0036-max-subarray
```
