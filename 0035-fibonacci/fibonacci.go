package fibonacci

import (
	"errors"
	"math/big"
)

type FibonacciCalculator struct {
	memo map[int]*big.Int
}

func NewFibonacciCalculator() *FibonacciCalculator {
	return &FibonacciCalculator{
		memo: make(map[int]*big.Int),
	}
}

func (fc *FibonacciCalculator) FibonacciMemoized(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("fibonacci is not defined for negative numbers")
	}

	if result, exists := fc.memo[n]; exists {
		return new(big.Int).Set(result), nil
	}

	if n == 0 {
		result := big.NewInt(0)
		fc.memo[n] = new(big.Int).Set(result)
		return result, nil
	}

	if n == 1 {
		result := big.NewInt(1)
		fc.memo[n] = new(big.Int).Set(result)
		return result, nil
	}

	prev1, err1 := fc.FibonacciMemoized(n - 1)
	if err1 != nil {
		return nil, err1
	}

	prev2, err2 := fc.FibonacciMemoized(n - 2)
	if err2 != nil {
		return nil, err2
	}

	result := new(big.Int).Add(prev1, prev2)
	fc.memo[n] = new(big.Int).Set(result)
	return result, nil
}

func (fc *FibonacciCalculator) GetMemoSize() int {
	return len(fc.memo)
}

func (fc *FibonacciCalculator) ClearMemo() {
	fc.memo = make(map[int]*big.Int)
}

func FibonacciIterative(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("fibonacci is not defined for negative numbers")
	}

	if n == 0 {
		return big.NewInt(0), nil
	}

	if n == 1 {
		return big.NewInt(1), nil
	}

	prev := big.NewInt(0)
	current := big.NewInt(1)

	for i := 2; i <= n; i++ {
		temp := new(big.Int).Add(prev, current)
		prev.Set(current)
		current.Set(temp)
	}

	return current, nil
}

func FibonacciRecursive(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("fibonacci is not defined for negative numbers")
	}

	if n == 0 {
		return big.NewInt(0), nil
	}

	if n == 1 {
		return big.NewInt(1), nil
	}

	prev1, err1 := FibonacciRecursive(n - 1)
	if err1 != nil {
		return nil, err1
	}

	prev2, err2 := FibonacciRecursive(n - 2)
	if err2 != nil {
		return nil, err2
	}

	return new(big.Int).Add(prev1, prev2), nil
}

func FibonacciDP(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("fibonacci is not defined for negative numbers")
	}

	if n == 0 {
		return big.NewInt(0), nil
	}

	if n == 1 {
		return big.NewInt(1), nil
	}

	dp := make([]*big.Int, n+1)
	dp[0] = big.NewInt(0)
	dp[1] = big.NewInt(1)

	for i := 2; i <= n; i++ {
		dp[i] = new(big.Int).Add(dp[i-1], dp[i-2])
	}

	return dp[n], nil
}

func FibonacciDPOptimized(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("fibonacci is not defined for negative numbers")
	}

	if n == 0 {
		return big.NewInt(0), nil
	}

	if n == 1 {
		return big.NewInt(1), nil
	}

	prev2 := big.NewInt(0)
	prev1 := big.NewInt(1)

	for i := 2; i <= n; i++ {
		current := new(big.Int).Add(prev1, prev2)
		prev2.Set(prev1)
		prev1.Set(current)
	}

	return prev1, nil
}

type Matrix2x2 struct {
	a, b, c, d *big.Int
}

func NewMatrix2x2(a, b, c, d int64) *Matrix2x2 {
	return &Matrix2x2{
		a: big.NewInt(a),
		b: big.NewInt(b),
		c: big.NewInt(c),
		d: big.NewInt(d),
	}
}

func (m1 *Matrix2x2) Multiply(m2 *Matrix2x2) *Matrix2x2 {
	result := &Matrix2x2{
		a: new(big.Int),
		b: new(big.Int),
		c: new(big.Int),
		d: new(big.Int),
	}

	temp1 := new(big.Int).Mul(m1.a, m2.a)
	temp2 := new(big.Int).Mul(m1.b, m2.c)
	result.a.Add(temp1, temp2)

	temp1.Mul(m1.a, m2.b)
	temp2.Mul(m1.b, m2.d)
	result.b.Add(temp1, temp2)

	temp1.Mul(m1.c, m2.a)
	temp2.Mul(m1.d, m2.c)
	result.c.Add(temp1, temp2)

	temp1.Mul(m1.c, m2.b)
	temp2.Mul(m1.d, m2.d)
	result.d.Add(temp1, temp2)

	return result
}

func (m *Matrix2x2) Power(n int) *Matrix2x2 {
	if n == 0 {
		return NewMatrix2x2(1, 0, 0, 1)
	}

	if n == 1 {
		return &Matrix2x2{
			a: new(big.Int).Set(m.a),
			b: new(big.Int).Set(m.b),
			c: new(big.Int).Set(m.c),
			d: new(big.Int).Set(m.d),
		}
	}

	if n%2 == 0 {
		half := m.Power(n / 2)
		return half.Multiply(half)
	}

	return m.Multiply(m.Power(n - 1))
}

func FibonacciMatrix(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("fibonacci is not defined for negative numbers")
	}

	if n == 0 {
		return big.NewInt(0), nil
	}

	if n == 1 {
		return big.NewInt(1), nil
	}

	fibMatrix := NewMatrix2x2(1, 1, 1, 0)
	result := fibMatrix.Power(n - 1)

	return result.a, nil
}

func FibonacciRange(start, end int) ([]*big.Int, error) {
	if start < 0 || end < 0 {
		return nil, errors.New("fibonacci is not defined for negative numbers")
	}

	if start > end {
		return nil, errors.New("start must be less than or equal to end")
	}

	results := make([]*big.Int, end-start+1)

	for i := start; i <= end; i++ {
		fibonacci, err := FibonacciIterative(i)
		if err != nil {
			return nil, err
		}
		results[i-start] = fibonacci
	}

	return results, nil
}

func LucasNumbers(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("lucas numbers are not defined for negative numbers")
	}

	if n == 0 {
		return big.NewInt(2), nil
	}

	if n == 1 {
		return big.NewInt(1), nil
	}

	prev2 := big.NewInt(2)
	prev1 := big.NewInt(1)

	for i := 2; i <= n; i++ {
		current := new(big.Int).Add(prev1, prev2)
		prev2.Set(prev1)
		prev1.Set(current)
	}

	return prev1, nil
}

func TribonacciNumbers(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("tribonacci numbers are not defined for negative numbers")
	}

	if n == 0 {
		return big.NewInt(0), nil
	}

	if n == 1 || n == 2 {
		return big.NewInt(1), nil
	}

	prev3 := big.NewInt(0)
	prev2 := big.NewInt(1)
	prev1 := big.NewInt(1)

	for i := 3; i <= n; i++ {
		current := new(big.Int).Add(prev1, prev2)
		current.Add(current, prev3)
		prev3.Set(prev2)
		prev2.Set(prev1)
		prev1.Set(current)
	}

	return prev1, nil
}

func IsFibonacci(num *big.Int) (bool, int) {
	if num.Sign() < 0 {
		return false, -1
	}

	if num.Cmp(big.NewInt(0)) == 0 {
		return true, 0
	}

	if num.Cmp(big.NewInt(1)) == 0 {
		return true, 1
	}

	prev := big.NewInt(0)
	current := big.NewInt(1)
	n := 1

	for current.Cmp(num) < 0 {
		temp := new(big.Int).Add(prev, current)
		prev.Set(current)
		current.Set(temp)
		n++
		if n > 1000 {
			break
		}
	}

	return current.Cmp(num) == 0, n
}

func GoldenRatio() *big.Float {
	sqrt5 := new(big.Float).SetPrec(256)
	sqrt5.SetFloat64(2.2360679774997896964091736687313)

	one := big.NewFloat(1.0)
	two := big.NewFloat(2.0)

	result := new(big.Float).Add(one, sqrt5)
	result.Quo(result, two)

	return result
}

func Run() any {
	calc := NewFibonacciCalculator()

	result := make(map[string]any)

	testValues := []int{0, 1, 5, 10, 15, 20, 30}

	iterativeResults := make([]map[string]any, len(testValues))
	dpResults := make([]map[string]any, len(testValues))
	memoizedResults := make([]map[string]any, len(testValues))
	matrixResults := make([]map[string]any, len(testValues))

	for i, n := range testValues {
		iterResult, iterErr := FibonacciIterative(n)
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

		dpResult, dpErr := FibonacciDP(n)
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

		memoResult, memoErr := calc.FibonacciMemoized(n)
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

		matrixResult, matrixErr := FibonacciMatrix(n)
		if matrixErr == nil {
			matrixResults[i] = map[string]any{
				"n":      n,
				"result": matrixResult.String(),
				"error":  false,
			}
		} else {
			matrixResults[i] = map[string]any{
				"n":     n,
				"error": true,
			}
		}
	}

	result["iterativeResults"] = iterativeResults
	result["dpResults"] = dpResults
	result["memoizedResults"] = memoizedResults
	result["matrixResults"] = matrixResults
	result["memoSize"] = calc.GetMemoSize()

	rangeResults, rangeErr := FibonacciRange(0, 15)
	result["rangeError"] = rangeErr != nil
	if rangeErr == nil {
		rangeStrings := make([]string, len(rangeResults))
		for i, val := range rangeResults {
			rangeStrings[i] = val.String()
		}
		result["rangeResults"] = rangeStrings
	}

	lucasResult, lucasErr := LucasNumbers(10)
	result["lucasError"] = lucasErr != nil
	if lucasErr == nil {
		result["lucas10"] = lucasResult.String()
	}

	tribonacciResult, tribonacciErr := TribonacciNumbers(10)
	result["tribonacciError"] = tribonacciErr != nil
	if tribonacciErr == nil {
		result["tribonacci10"] = tribonacciResult.String()
	}

	isFib55, fibN := IsFibonacci(big.NewInt(55))
	result["is55Fibonacci"] = isFib55
	result["fibonacci55N"] = fibN

	isFib56, _ := IsFibonacci(big.NewInt(56))
	result["is56Fibonacci"] = isFib56

	negativeResult, negativeErr := FibonacciIterative(-5)
	result["negativeError"] = negativeErr != nil
	if negativeErr != nil {
		result["negativeErrorMessage"] = negativeErr.Error()
	}
	if negativeResult != nil {
		result["negativeResult"] = negativeResult.String()
	}

	largeResult, largeErr := FibonacciIterative(100)
	result["largeError"] = largeErr != nil
	if largeErr == nil {
		result["fibonacci100Length"] = len(largeResult.String())
		result["fibonacci100"] = largeResult.String()
	}

	goldenRatio := GoldenRatio()
	result["goldenRatio"] = goldenRatio.String()

	return result
}
