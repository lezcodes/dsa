package factorial

import (
	"math/big"
	"testing"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}
}

func TestFactorialIterative(t *testing.T) {
	tests := []struct {
		n        int
		expected string
		hasError bool
	}{
		{0, "1", false},
		{1, "1", false},
		{2, "2", false},
		{3, "6", false},
		{4, "24", false},
		{5, "120", false},
		{10, "3628800", false},
		{-1, "", true},
		{-5, "", true},
	}

	for _, test := range tests {
		result, err := FactorialIterative(test.n)

		if test.hasError {
			if err == nil {
				t.Errorf("Expected error for n=%d, but got none", test.n)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for n=%d: %v", test.n, err)
			} else if result.String() != test.expected {
				t.Errorf("For n=%d, expected %s, got %s", test.n, test.expected, result.String())
			}
		}
	}
}

func TestFactorialRecursive(t *testing.T) {
	tests := []struct {
		n        int
		expected string
		hasError bool
	}{
		{0, "1", false},
		{1, "1", false},
		{2, "2", false},
		{3, "6", false},
		{4, "24", false},
		{5, "120", false},
		{10, "3628800", false},
		{-1, "", true},
		{-5, "", true},
	}

	for _, test := range tests {
		result, err := FactorialRecursive(test.n)

		if test.hasError {
			if err == nil {
				t.Errorf("Expected error for n=%d, but got none", test.n)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for n=%d: %v", test.n, err)
			} else if result.String() != test.expected {
				t.Errorf("For n=%d, expected %s, got %s", test.n, test.expected, result.String())
			}
		}
	}
}

func TestFactorialDP(t *testing.T) {
	tests := []struct {
		n        int
		expected string
		hasError bool
	}{
		{0, "1", false},
		{1, "1", false},
		{2, "2", false},
		{3, "6", false},
		{4, "24", false},
		{5, "120", false},
		{10, "3628800", false},
		{-1, "", true},
		{-5, "", true},
	}

	for _, test := range tests {
		result, err := FactorialDP(test.n)

		if test.hasError {
			if err == nil {
				t.Errorf("Expected error for n=%d, but got none", test.n)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for n=%d: %v", test.n, err)
			} else if result.String() != test.expected {
				t.Errorf("For n=%d, expected %s, got %s", test.n, test.expected, result.String())
			}
		}
	}
}

func TestFactorialDPOptimized(t *testing.T) {
	tests := []struct {
		n        int
		expected string
		hasError bool
	}{
		{0, "1", false},
		{1, "1", false},
		{2, "2", false},
		{3, "6", false},
		{4, "24", false},
		{5, "120", false},
		{10, "3628800", false},
		{-1, "", true},
		{-5, "", true},
	}

	for _, test := range tests {
		result, err := FactorialDPOptimized(test.n)

		if test.hasError {
			if err == nil {
				t.Errorf("Expected error for n=%d, but got none", test.n)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for n=%d: %v", test.n, err)
			} else if result.String() != test.expected {
				t.Errorf("For n=%d, expected %s, got %s", test.n, test.expected, result.String())
			}
		}
	}
}

func TestFactorialMemoized(t *testing.T) {
	calc := NewFactorialCalculator()

	tests := []struct {
		n        int
		expected string
		hasError bool
	}{
		{0, "1", false},
		{1, "1", false},
		{2, "2", false},
		{3, "6", false},
		{4, "24", false},
		{5, "120", false},
		{10, "3628800", false},
		{-1, "", true},
		{-5, "", true},
	}

	for _, test := range tests {
		result, err := calc.FactorialMemoized(test.n)

		if test.hasError {
			if err == nil {
				t.Errorf("Expected error for n=%d, but got none", test.n)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for n=%d: %v", test.n, err)
			} else if result.String() != test.expected {
				t.Errorf("For n=%d, expected %s, got %s", test.n, test.expected, result.String())
			}
		}
	}

	if calc.GetMemoSize() == 0 {
		t.Error("Expected memoization to store results")
	}

	calc.ClearMemo()
	if calc.GetMemoSize() != 0 {
		t.Error("Expected memo to be cleared")
	}
}

func TestFactorialRange(t *testing.T) {
	results, err := FactorialRange(0, 5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []string{"1", "1", "2", "6", "24", "120"}
	if len(results) != len(expected) {
		t.Errorf("Expected %d results, got %d", len(expected), len(results))
	}

	for i, result := range results {
		if result.String() != expected[i] {
			t.Errorf("For index %d, expected %s, got %s", i, expected[i], result.String())
		}
	}

	_, err = FactorialRange(-1, 5)
	if err == nil {
		t.Error("Expected error for negative start")
	}

	_, err = FactorialRange(5, 3)
	if err == nil {
		t.Error("Expected error for start > end")
	}
}

func TestDoubleFactorial(t *testing.T) {
	tests := []struct {
		n        int
		expected string
		hasError bool
	}{
		{0, "1", false},
		{1, "1", false},
		{2, "2", false},
		{3, "3", false},
		{4, "8", false},
		{5, "15", false},
		{6, "48", false},
		{10, "3840", false},
		{-1, "", true},
	}

	for _, test := range tests {
		result, err := DoubleFactorial(test.n)

		if test.hasError {
			if err == nil {
				t.Errorf("Expected error for n=%d, but got none", test.n)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for n=%d: %v", test.n, err)
			} else if result.String() != test.expected {
				t.Errorf("For n=%d, expected %s, got %s", test.n, test.expected, result.String())
			}
		}
	}
}

func TestSubfactorial(t *testing.T) {
	tests := []struct {
		n        int
		expected string
		hasError bool
	}{
		{0, "1", false},
		{1, "0", false},
		{2, "1", false},
		{3, "2", false},
		{4, "9", false},
		{5, "44", false},
		{-1, "", true},
	}

	for _, test := range tests {
		result, err := Subfactorial(test.n)

		if test.hasError {
			if err == nil {
				t.Errorf("Expected error for n=%d, but got none", test.n)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for n=%d: %v", test.n, err)
			} else if result.String() != test.expected {
				t.Errorf("For n=%d, expected %s, got %s", test.n, test.expected, result.String())
			}
		}
	}
}

func TestIsFactorial(t *testing.T) {
	tests := []struct {
		num      int64
		expected bool
		n        int
	}{
		{1, true, 0},
		{2, true, 2},
		{6, true, 3},
		{24, true, 4},
		{120, true, 5},
		{3628800, true, 10},
		{5, false, -1},
		{7, false, -1},
		{100, false, -1},
		{-1, false, -1},
	}

	for _, test := range tests {
		isFactorial, n := IsFactorial(big.NewInt(test.num))

		if isFactorial != test.expected {
			t.Errorf("For num=%d, expected %t, got %t", test.num, test.expected, isFactorial)
		}

		if test.expected && test.num != 1 && n != test.n {
			t.Errorf("For num=%d, expected n=%d, got n=%d", test.num, test.n, n)
		}
	}
}

func TestFactorialConsistency(t *testing.T) {
	calc := NewFactorialCalculator()

	for n := 0; n <= 20; n++ {
		iterResult, iterErr := FactorialIterative(n)
		recResult, recErr := FactorialRecursive(n)
		dpResult, dpErr := FactorialDP(n)
		memoResult, memoErr := calc.FactorialMemoized(n)

		if iterErr != nil || recErr != nil || dpErr != nil || memoErr != nil {
			t.Errorf("Unexpected error for n=%d", n)
			continue
		}

		if iterResult.Cmp(recResult) != 0 {
			t.Errorf("Iterative and recursive results differ for n=%d", n)
		}

		if iterResult.Cmp(dpResult) != 0 {
			t.Errorf("Iterative and DP results differ for n=%d", n)
		}

		if iterResult.Cmp(memoResult) != 0 {
			t.Errorf("Iterative and memoized results differ for n=%d", n)
		}
	}
}

func TestFactorialLargeNumbers(t *testing.T) {
	largeN := 100

	result, err := FactorialIterative(largeN)
	if err != nil {
		t.Errorf("Unexpected error for large n=%d: %v", largeN, err)
	}

	if result == nil {
		t.Errorf("Expected non-nil result for n=%d", largeN)
	}

	resultStr := result.String()
	if len(resultStr) < 100 {
		t.Errorf("Expected very large number for 100!, got length %d", len(resultStr))
	}
}

func TestFactorialEdgeCases(t *testing.T) {
	calc := NewFactorialCalculator()

	result0, err0 := calc.FactorialMemoized(0)
	if err0 != nil || result0.String() != "1" {
		t.Error("Factorial of 0 should be 1")
	}

	result1, err1 := calc.FactorialMemoized(1)
	if err1 != nil || result1.String() != "1" {
		t.Error("Factorial of 1 should be 1")
	}

	_, errNeg := calc.FactorialMemoized(-1)
	if errNeg == nil {
		t.Error("Expected error for negative input")
	}
}

func TestFactorialMemoizationEfficiency(t *testing.T) {
	calc := NewFactorialCalculator()

	n := 20

	_, err := calc.FactorialMemoized(n)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	memoSize := calc.GetMemoSize()
	if memoSize == 0 {
		t.Error("Expected memoization to store results")
	}

	_, err = calc.FactorialMemoized(n)
	if err != nil {
		t.Errorf("Unexpected error on second call: %v", err)
	}

	if calc.GetMemoSize() != memoSize {
		t.Error("Memo size should not change on repeated calls")
	}
}

func BenchmarkFactorialIterative(b *testing.B) {
	for b.Loop() {
		FactorialIterative(20)
	}
}

func BenchmarkFactorialRecursive(b *testing.B) {
	for b.Loop() {
		FactorialRecursive(20)
	}
}

func BenchmarkFactorialDP(b *testing.B) {
	for b.Loop() {
		FactorialDP(20)
	}
}

func BenchmarkFactorialDPOptimized(b *testing.B) {
	for b.Loop() {
		FactorialDPOptimized(20)
	}
}

func BenchmarkFactorialMemoized(b *testing.B) {
	calc := NewFactorialCalculator()

	for b.Loop() {
		calc.FactorialMemoized(20)
	}
}

func BenchmarkFactorialMemoizedCold(b *testing.B) {
	for b.Loop() {
		calc := NewFactorialCalculator()
		calc.FactorialMemoized(20)
	}
}

func BenchmarkFactorialLarge(b *testing.B) {
	for b.Loop() {
		FactorialIterative(100)
	}
}

func BenchmarkDoubleFactorial(b *testing.B) {
	for b.Loop() {
		DoubleFactorial(20)
	}
}

func BenchmarkSubfactorial(b *testing.B) {
	for b.Loop() {
		Subfactorial(20)
	}
}

func BenchmarkIsFactorial(b *testing.B) {
	num := big.NewInt(3628800)

	for b.Loop() {
		IsFactorial(num)
	}
}

func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		Run()
	}
}
