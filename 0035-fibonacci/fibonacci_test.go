package fibonacci

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

func TestFibonacciIterative(t *testing.T) {
	tests := []struct {
		n        int
		expected string
		hasError bool
	}{
		{0, "0", false},
		{1, "1", false},
		{2, "1", false},
		{3, "2", false},
		{4, "3", false},
		{5, "5", false},
		{6, "8", false},
		{7, "13", false},
		{8, "21", false},
		{9, "34", false},
		{10, "55", false},
		{-1, "", true},
		{-5, "", true},
	}

	for _, test := range tests {
		result, err := FibonacciIterative(test.n)

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

func TestFibonacciRecursive(t *testing.T) {
	tests := []struct {
		n        int
		expected string
		hasError bool
	}{
		{0, "0", false},
		{1, "1", false},
		{2, "1", false},
		{3, "2", false},
		{4, "3", false},
		{5, "5", false},
		{6, "8", false},
		{7, "13", false},
		{8, "21", false},
		{9, "34", false},
		{10, "55", false},
		{-1, "", true},
		{-5, "", true},
	}

	for _, test := range tests {
		result, err := FibonacciRecursive(test.n)

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

func TestFibonacciDP(t *testing.T) {
	tests := []struct {
		n        int
		expected string
		hasError bool
	}{
		{0, "0", false},
		{1, "1", false},
		{2, "1", false},
		{3, "2", false},
		{4, "3", false},
		{5, "5", false},
		{6, "8", false},
		{7, "13", false},
		{8, "21", false},
		{9, "34", false},
		{10, "55", false},
		{-1, "", true},
		{-5, "", true},
	}

	for _, test := range tests {
		result, err := FibonacciDP(test.n)

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

func TestFibonacciDPOptimized(t *testing.T) {
	tests := []struct {
		n        int
		expected string
		hasError bool
	}{
		{0, "0", false},
		{1, "1", false},
		{2, "1", false},
		{3, "2", false},
		{4, "3", false},
		{5, "5", false},
		{6, "8", false},
		{7, "13", false},
		{8, "21", false},
		{9, "34", false},
		{10, "55", false},
		{-1, "", true},
		{-5, "", true},
	}

	for _, test := range tests {
		result, err := FibonacciDPOptimized(test.n)

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

func TestFibonacciMemoized(t *testing.T) {
	calc := NewFibonacciCalculator()

	tests := []struct {
		n        int
		expected string
		hasError bool
	}{
		{0, "0", false},
		{1, "1", false},
		{2, "1", false},
		{3, "2", false},
		{4, "3", false},
		{5, "5", false},
		{6, "8", false},
		{7, "13", false},
		{8, "21", false},
		{9, "34", false},
		{10, "55", false},
		{-1, "", true},
		{-5, "", true},
	}

	for _, test := range tests {
		result, err := calc.FibonacciMemoized(test.n)

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

func TestFibonacciMatrix(t *testing.T) {
	tests := []struct {
		n        int
		expected string
		hasError bool
	}{
		{0, "0", false},
		{1, "1", false},
		{2, "1", false},
		{3, "2", false},
		{4, "3", false},
		{5, "5", false},
		{6, "8", false},
		{7, "13", false},
		{8, "21", false},
		{9, "34", false},
		{10, "55", false},
		{-1, "", true},
		{-5, "", true},
	}

	for _, test := range tests {
		result, err := FibonacciMatrix(test.n)

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

func TestMatrix2x2Operations(t *testing.T) {
	m1 := NewMatrix2x2(1, 1, 1, 0)
	m2 := NewMatrix2x2(1, 1, 1, 0)

	result := m1.Multiply(m2)

	if result.a.String() != "2" || result.b.String() != "1" ||
		result.c.String() != "1" || result.d.String() != "1" {
		t.Errorf("Matrix multiplication failed: got [%s %s; %s %s]",
			result.a.String(), result.b.String(), result.c.String(), result.d.String())
	}

	identity := NewMatrix2x2(1, 0, 0, 1)
	power0 := m1.Power(0)

	if power0.a.Cmp(identity.a) != 0 || power0.b.Cmp(identity.b) != 0 ||
		power0.c.Cmp(identity.c) != 0 || power0.d.Cmp(identity.d) != 0 {
		t.Error("Matrix power 0 should return identity matrix")
	}

	power1 := m1.Power(1)
	if power1.a.Cmp(m1.a) != 0 || power1.b.Cmp(m1.b) != 0 ||
		power1.c.Cmp(m1.c) != 0 || power1.d.Cmp(m1.d) != 0 {
		t.Error("Matrix power 1 should return original matrix")
	}
}

func TestFibonacciRange(t *testing.T) {
	results, err := FibonacciRange(0, 10)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []string{"0", "1", "1", "2", "3", "5", "8", "13", "21", "34", "55"}
	if len(results) != len(expected) {
		t.Errorf("Expected %d results, got %d", len(expected), len(results))
	}

	for i, result := range results {
		if result.String() != expected[i] {
			t.Errorf("For index %d, expected %s, got %s", i, expected[i], result.String())
		}
	}

	_, err = FibonacciRange(-1, 5)
	if err == nil {
		t.Error("Expected error for negative start")
	}

	_, err = FibonacciRange(5, 3)
	if err == nil {
		t.Error("Expected error for start > end")
	}
}

func TestLucasNumbers(t *testing.T) {
	tests := []struct {
		n        int
		expected string
		hasError bool
	}{
		{0, "2", false},
		{1, "1", false},
		{2, "3", false},
		{3, "4", false},
		{4, "7", false},
		{5, "11", false},
		{6, "18", false},
		{10, "123", false},
		{-1, "", true},
	}

	for _, test := range tests {
		result, err := LucasNumbers(test.n)

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

func TestTribonacciNumbers(t *testing.T) {
	tests := []struct {
		n        int
		expected string
		hasError bool
	}{
		{0, "0", false},
		{1, "1", false},
		{2, "1", false},
		{3, "2", false},
		{4, "4", false},
		{5, "7", false},
		{6, "13", false},
		{7, "24", false},
		{8, "44", false},
		{9, "81", false},
		{10, "149", false},
		{-1, "", true},
	}

	for _, test := range tests {
		result, err := TribonacciNumbers(test.n)

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

func TestIsFibonacci(t *testing.T) {
	tests := []struct {
		num      int64
		expected bool
		n        int
	}{
		{0, true, 0},
		{1, true, 1},
		{2, true, 3},
		{3, true, 4},
		{5, true, 5},
		{8, true, 6},
		{13, true, 7},
		{21, true, 8},
		{34, true, 9},
		{55, true, 10},
		{4, false, -1},
		{6, false, -1},
		{7, false, -1},
		{100, false, -1},
		{-1, false, -1},
	}

	for _, test := range tests {
		isFibonacci, n := IsFibonacci(big.NewInt(test.num))

		if isFibonacci != test.expected {
			t.Errorf("For num=%d, expected %t, got %t", test.num, test.expected, isFibonacci)
		}

		if test.expected && test.num != 1 && n != test.n {
			t.Errorf("For num=%d, expected n=%d, got n=%d", test.num, test.n, n)
		}
	}
}

func TestGoldenRatio(t *testing.T) {
	ratio := GoldenRatio()
	if ratio == nil {
		t.Error("Expected non-nil golden ratio")
	}

	ratioStr := ratio.String()
	if len(ratioStr) < 10 {
		t.Errorf("Expected detailed golden ratio, got %s", ratioStr)
	}

	if !contains(ratioStr, "1.618") {
		t.Errorf("Expected golden ratio to start with 1.618, got %s", ratioStr)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}

func TestFibonacciConsistency(t *testing.T) {
	calc := NewFibonacciCalculator()

	for n := 0; n <= 30; n++ {
		iterResult, iterErr := FibonacciIterative(n)
		dpResult, dpErr := FibonacciDP(n)
		memoResult, memoErr := calc.FibonacciMemoized(n)
		matrixResult, matrixErr := FibonacciMatrix(n)

		if iterErr != nil || dpErr != nil || memoErr != nil || matrixErr != nil {
			t.Errorf("Unexpected error for n=%d", n)
			continue
		}

		if iterResult.Cmp(dpResult) != 0 {
			t.Errorf("Iterative and DP results differ for n=%d", n)
		}

		if iterResult.Cmp(memoResult) != 0 {
			t.Errorf("Iterative and memoized results differ for n=%d", n)
		}

		if iterResult.Cmp(matrixResult) != 0 {
			t.Errorf("Iterative and matrix results differ for n=%d", n)
		}
	}
}

func TestFibonacciLargeNumbers(t *testing.T) {
	largeN := 100

	result, err := FibonacciIterative(largeN)
	if err != nil {
		t.Errorf("Unexpected error for large n=%d: %v", largeN, err)
	}

	if result == nil {
		t.Errorf("Expected non-nil result for n=%d", largeN)
	}

	resultStr := result.String()
	if len(resultStr) < 20 {
		t.Errorf("Expected very large number for Fib(100), got length %d", len(resultStr))
	}
}

func TestFibonacciEdgeCases(t *testing.T) {
	calc := NewFibonacciCalculator()

	result0, err0 := calc.FibonacciMemoized(0)
	if err0 != nil || result0.String() != "0" {
		t.Error("Fibonacci of 0 should be 0")
	}

	result1, err1 := calc.FibonacciMemoized(1)
	if err1 != nil || result1.String() != "1" {
		t.Error("Fibonacci of 1 should be 1")
	}

	_, errNeg := calc.FibonacciMemoized(-1)
	if errNeg == nil {
		t.Error("Expected error for negative input")
	}
}

func TestFibonacciMemoizationEfficiency(t *testing.T) {
	calc := NewFibonacciCalculator()

	n := 30

	_, err := calc.FibonacciMemoized(n)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	memoSize := calc.GetMemoSize()
	if memoSize == 0 {
		t.Error("Expected memoization to store results")
	}

	_, err = calc.FibonacciMemoized(n)
	if err != nil {
		t.Errorf("Unexpected error on second call: %v", err)
	}

	if calc.GetMemoSize() != memoSize {
		t.Error("Memo size should not change on repeated calls")
	}
}

func TestFibonacciMatrixLarge(t *testing.T) {
	n := 50

	iterResult, iterErr := FibonacciIterative(n)
	matrixResult, matrixErr := FibonacciMatrix(n)

	if iterErr != nil || matrixErr != nil {
		t.Errorf("Unexpected error for n=%d", n)
		return
	}

	if iterResult.Cmp(matrixResult) != 0 {
		t.Errorf("Matrix and iterative results differ for large n=%d", n)
	}
}

func BenchmarkFibonacciIterative(b *testing.B) {
	for b.Loop() {
		FibonacciIterative(30)
	}
}

func BenchmarkFibonacciRecursive(b *testing.B) {
	for b.Loop() {
		FibonacciRecursive(20)
	}
}

func BenchmarkFibonacciDP(b *testing.B) {
	for b.Loop() {
		FibonacciDP(30)
	}
}

func BenchmarkFibonacciDPOptimized(b *testing.B) {
	for b.Loop() {
		FibonacciDPOptimized(30)
	}
}

func BenchmarkFibonacciMemoized(b *testing.B) {
	calc := NewFibonacciCalculator()

	for b.Loop() {
		calc.FibonacciMemoized(30)
	}
}

func BenchmarkFibonacciMemoizedCold(b *testing.B) {
	for b.Loop() {
		calc := NewFibonacciCalculator()
		calc.FibonacciMemoized(30)
	}
}

func BenchmarkFibonacciMatrix(b *testing.B) {
	for b.Loop() {
		FibonacciMatrix(30)
	}
}

func BenchmarkFibonacciLarge(b *testing.B) {
	for b.Loop() {
		FibonacciIterative(100)
	}
}

func BenchmarkLucasNumbers(b *testing.B) {
	for b.Loop() {
		LucasNumbers(30)
	}
}

func BenchmarkTribonacciNumbers(b *testing.B) {
	for b.Loop() {
		TribonacciNumbers(30)
	}
}

func BenchmarkIsFibonacci(b *testing.B) {
	num := big.NewInt(55)

	for b.Loop() {
		IsFibonacci(num)
	}
}

func BenchmarkGoldenRatio(b *testing.B) {
	for b.Loop() {
		GoldenRatio()
	}
}

func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		Run()
	}
}
