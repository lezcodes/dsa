package factorial

import (
	"errors"
	"math/big"
)

type FactorialCalculator struct {
	memo map[int]*big.Int
}

func NewFactorialCalculator() *FactorialCalculator {
	return &FactorialCalculator{
		memo: make(map[int]*big.Int),
	}
}

func (fc *FactorialCalculator) FactorialMemoized(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("factorial is not defined for negative numbers")
	}

	if result, exists := fc.memo[n]; exists {
		return new(big.Int).Set(result), nil
	}

	if n == 0 || n == 1 {
		result := big.NewInt(1)
		fc.memo[n] = new(big.Int).Set(result)
		return result, nil
	}

	prevResult, err := fc.FactorialMemoized(n - 1)
	if err != nil {
		return nil, err
	}

	result := new(big.Int).Mul(prevResult, big.NewInt(int64(n)))
	fc.memo[n] = new(big.Int).Set(result)
	return result, nil
}

func (fc *FactorialCalculator) GetMemoSize() int {
	return len(fc.memo)
}

func (fc *FactorialCalculator) ClearMemo() {
	fc.memo = make(map[int]*big.Int)
}

func FactorialIterative(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("factorial is not defined for negative numbers")
	}

	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}

	return result, nil
}

func FactorialRecursive(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("factorial is not defined for negative numbers")
	}

	if n == 0 || n == 1 {
		return big.NewInt(1), nil
	}

	prevResult, err := FactorialRecursive(n - 1)
	if err != nil {
		return nil, err
	}

	return new(big.Int).Mul(prevResult, big.NewInt(int64(n))), nil
}

func FactorialDP(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("factorial is not defined for negative numbers")
	}

	if n == 0 || n == 1 {
		return big.NewInt(1), nil
	}

	dp := make([]*big.Int, n+1)
	dp[0] = big.NewInt(1)
	dp[1] = big.NewInt(1)

	for i := 2; i <= n; i++ {
		dp[i] = new(big.Int).Mul(dp[i-1], big.NewInt(int64(i)))
	}

	return dp[n], nil
}

func FactorialDPOptimized(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("factorial is not defined for negative numbers")
	}

	if n == 0 || n == 1 {
		return big.NewInt(1), nil
	}

	prev := big.NewInt(1)
	current := big.NewInt(1)

	for i := 2; i <= n; i++ {
		current.Mul(prev, big.NewInt(int64(i)))
		prev.Set(current)
	}

	return current, nil
}

func FactorialRange(start, end int) ([]*big.Int, error) {
	if start < 0 || end < 0 {
		return nil, errors.New("factorial is not defined for negative numbers")
	}

	if start > end {
		return nil, errors.New("start must be less than or equal to end")
	}

	results := make([]*big.Int, end-start+1)

	for i := start; i <= end; i++ {
		factorial, err := FactorialIterative(i)
		if err != nil {
			return nil, err
		}
		results[i-start] = factorial
	}

	return results, nil
}

func DoubleFactorial(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("double factorial is not defined for negative numbers")
	}

	result := big.NewInt(1)
	for i := n; i > 0; i -= 2 {
		result.Mul(result, big.NewInt(int64(i)))
	}

	return result, nil
}

func Subfactorial(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("subfactorial is not defined for negative numbers")
	}

	if n == 0 {
		return big.NewInt(1), nil
	}

	if n == 1 {
		return big.NewInt(0), nil
	}

	dp := make([]*big.Int, n+1)
	dp[0] = big.NewInt(1)
	dp[1] = big.NewInt(0)

	for i := 2; i <= n; i++ {
		temp := new(big.Int).Add(dp[i-1], dp[i-2])
		dp[i] = new(big.Int).Mul(big.NewInt(int64(i-1)), temp)
	}

	return dp[n], nil
}

func IsFactorial(num *big.Int) (bool, int) {
	if num.Sign() <= 0 {
		return false, -1
	}

	if num.Cmp(big.NewInt(1)) == 0 {
		return true, 0
	}

	current := big.NewInt(1)
	n := 1

	for current.Cmp(num) < 0 {
		current.Mul(current, big.NewInt(int64(n)))
		if current.Cmp(num) == 0 {
			return true, n
		}
		n++
		if n > 1000 {
			break
		}
	}

	return false, -1
}

func Run() any {
	calc := NewFactorialCalculator()

	result := make(map[string]any)

	testValues := []int{0, 1, 5, 10, 15, 20}

	iterativeResults := make([]map[string]any, len(testValues))
	recursiveResults := make([]map[string]any, len(testValues))
	dpResults := make([]map[string]any, len(testValues))
	memoizedResults := make([]map[string]any, len(testValues))

	for i, n := range testValues {
		iterResult, iterErr := FactorialIterative(n)
		recursiveResults[i] = map[string]any{
			"n":     n,
			"error": iterErr != nil,
		}
		if iterErr == nil {
			iterativeResults[i] = map[string]any{
				"n":      n,
				"result": iterResult.String(),
				"error":  false,
			}
		} else {
			iterativeResults[i] = map[string]any{
				"n":     n,
				"error": true,
			}
		}

		recResult, recErr := FactorialRecursive(n)
		if recErr == nil {
			recursiveResults[i] = map[string]any{
				"n":      n,
				"result": recResult.String(),
				"error":  false,
			}
		} else {
			recursiveResults[i] = map[string]any{
				"n":     n,
				"error": true,
			}
		}

		dpResult, dpErr := FactorialDP(n)
		if dpErr == nil {
			dpResults[i] = map[string]any{
				"n":      n,
				"result": dpResult.String(),
				"error":  false,
			}
		} else {
			dpResults[i] = map[string]any{
				"n":     n,
				"error": true,
			}
		}

		memoResult, memoErr := calc.FactorialMemoized(n)
		if memoErr == nil {
			memoizedResults[i] = map[string]any{
				"n":      n,
				"result": memoResult.String(),
				"error":  false,
			}
		} else {
			memoizedResults[i] = map[string]any{
				"n":     n,
				"error": true,
			}
		}
	}

	result["iterativeResults"] = iterativeResults
	result["recursiveResults"] = recursiveResults
	result["dpResults"] = dpResults
	result["memoizedResults"] = memoizedResults
	result["memoSize"] = calc.GetMemoSize()

	rangeResults, rangeErr := FactorialRange(0, 10)
	result["rangeError"] = rangeErr != nil
	if rangeErr == nil {
		rangeStrings := make([]string, len(rangeResults))
		for i, val := range rangeResults {
			rangeStrings[i] = val.String()
		}
		result["rangeResults"] = rangeStrings
	}

	doubleFactResult, doubleErr := DoubleFactorial(10)
	result["doubleFactorialError"] = doubleErr != nil
	if doubleErr == nil {
		result["doubleFactorial10"] = doubleFactResult.String()
	}

	subFactResult, subErr := Subfactorial(5)
	result["subfactorialError"] = subErr != nil
	if subErr == nil {
		result["subfactorial5"] = subFactResult.String()
	}

	isFactorial120, factorialN := IsFactorial(big.NewInt(120))
	result["is120Factorial"] = isFactorial120
	result["factorial120N"] = factorialN

	isFactorial100, _ := IsFactorial(big.NewInt(100))
	result["is100Factorial"] = isFactorial100

	negativeResult, negativeErr := FactorialIterative(-5)
	result["negativeError"] = negativeErr != nil
	if negativeErr != nil {
		result["negativeErrorMessage"] = negativeErr.Error()
	}
	if negativeResult != nil {
		result["negativeResult"] = negativeResult.String()
	}

	largeResult, largeErr := FactorialIterative(100)
	result["largeError"] = largeErr != nil
	if largeErr == nil {
		result["factorial100Length"] = len(largeResult.String())
		result["factorial100StartsWith"] = largeResult.String()[:20]
	}

	return result
}
