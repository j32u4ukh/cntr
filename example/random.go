package main

import (
	"fmt"

	"github.com/j32u4ukh/cntr"
	"github.com/spf13/cobra"
)

const (
	defaultRandomUpper  = 50
	defaultRandomTrials = 10_000
)

func registerRandomCommands(root *cobra.Command) {
	var upper, trials int
	cmd := &cobra.Command{
		Use:   "random",
		Short: "整數區間 [0,upper) 隨機抽樣，EvaluateUniformity 評估（direct / buffered 子指令）",
	}
	cmd.PersistentFlags().IntVarP(&upper, "upper", "u", defaultRandomUpper,
		"區間上界（不包含），抽樣整數為 [0,upper)")
	cmd.PersistentFlags().IntVarP(&trials, "trials", "n", defaultRandomTrials,
		"抽樣筆數（EvaluateUniformity 建議 ≥30）")

	cmd.AddCommand(
		// go run . random direct -n 10000 -u 50
		&cobra.Command{
			Use:   "direct",
			Short: "DirectRandom.Next(upper)；預設 upper=50、trials=10000",
			Run: func(*cobra.Command, []string) {
				if !randomArgsOK(upper, trials) {
					return
				}
				fmt.Printf("[0,%d) DirectRandom，%d 筆\n", upper, trials)
				data := sampleDirectRandom(trials, upper)
				printUniformEval(fmt.Sprintf("DirectRandom.Next(%d)", upper), EvaluateUniformity(data))
			},
		},
		// go run . random buffered -n 10000 -u 50
		&cobra.Command{
			Use:   "buffered",
			Short: "BufferedRandom Init(0,upper)+Next()；預設 upper=50、trials=10000",
			Run: func(*cobra.Command, []string) {
				if !randomArgsOK(upper, trials) {
					return
				}
				fmt.Printf("[0,%d) BufferedRandom，%d 筆\n", upper, trials)
				data := sampleBufferedRandom(trials, upper)
				printUniformEval(fmt.Sprintf("BufferedRandom Init(0,%d)+Next()", upper), EvaluateUniformity(data))
			},
		},
	)
	root.AddCommand(cmd)
}

func randomArgsOK(upper, trials int) bool {
	if upper < 1 {
		fmt.Println("錯誤: --upper / -u 須 ≥1（對應區間 [0,upper)）")
		return false
	}
	if trials < 1 {
		fmt.Println("錯誤: --trials / -n 須 ≥1")
		return false
	}
	return true
}

func sampleDirectRandom(n, upperExclusive int) []int {
	d := &cntr.DirectRandom{}
	out := make([]int, n)
	for i := range out {
		out[i] = d.Next(upperExclusive)
	}
	return out
}

func sampleBufferedRandom(n, upperExclusive int) []int {
	buf := cntr.NewBufferedRandom[int](256)
	buf.Init(0, upperExclusive)
	out := make([]int, n)
	for i := range out {
		out[i] = buf.Next()
	}
	return out
}

func printUniformEval(label string, r DistributionResult) {
	if lenSampleTooSmall(r) {
		fmt.Printf("%s 樣本過少（EvaluateUniformity 需至少 30 筆），無法評估\n", label)
		return
	}
	pass := "未通過（χ² > 臨界）"
	if r.IsUniform {
		pass = "通過（χ² ≤ 臨界，α≈0.05）"
	}
	fmt.Printf("%s  χ²=%.4f  臨界值=%.4f  %s  均勻度得分=%.1f\n",
		label, r.ChiSquare, r.CriticalValue, pass, r.Score)
}

func lenSampleTooSmall(r DistributionResult) bool {
	return r.ChiSquare == 0 && r.CriticalValue == 0 && r.Score == 0 && !r.IsUniform
}

// DistributionResult 評估結果結構
type DistributionResult struct {
	ChiSquare     float64 // 卡方統計量，越小越均勻
	CriticalValue float64 // 臨界值 (alpha=0.05)
	IsUniform     bool    // 是否通過均勻分佈檢定（χ² ≤ 臨界）
	Score         float64 // 均勻度得分 (0-100)，越高越好
}

// EvaluateUniformity 傳入 int 列表，評估其是否服從均勻分配
func EvaluateUniformity(data []int) DistributionResult {
	n := len(data)
	if n < 30 {
		return DistributionResult{}
	}

	min, max := data[0], data[0]
	for _, v := range data {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	numBuckets := 10
	if (max - min + 1) < numBuckets {
		numBuckets = max - min + 1
	}
	if numBuckets < 2 {
		numBuckets = 2
	}

	counts := make([]int, numBuckets)
	rangeSize := float64(max - min + 1)
	for _, v := range data {
		idx := int(float64(v-min) / rangeSize * float64(numBuckets))
		if idx >= numBuckets {
			idx = numBuckets - 1
		}
		if idx < 0 {
			idx = 0
		}
		counts[idx]++
	}

	expected := float64(n) / float64(numBuckets)
	var chiSquare float64
	for _, obs := range counts {
		diff := float64(obs) - expected
		chiSquare += (diff * diff) / expected
	}

	criticalTable := map[int]float64{
		1: 3.84, 2: 5.99, 3: 7.81, 4: 9.49, 5: 11.07,
		6: 12.59, 7: 14.07, 8: 15.51, 9: 16.92, 10: 18.31,
	}
	df := numBuckets - 1
	crit, found := criticalTable[df]
	if !found {
		crit = float64(df) * 1.5
	}

	score := 100.0 * (1.0 - (chiSquare / (chiSquare + crit)))

	return DistributionResult{
		ChiSquare:     chiSquare,
		CriticalValue: crit,
		IsUniform:     chiSquare <= crit,
		Score:         score,
	}
}
