package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/j32u4ukh/cntr"
	"github.com/spf13/cobra"
)

const defaultWeightedSelectTrials = 10_000

func registerWeightedRandomCommands(root *cobra.Command) {
	var trials int
	cmd := &cobra.Command{
		Use:   "weighted-random",
		Short: "Examples for cntr.WeightedRandom (BST over prefix weights).",
	}
	cmd.PersistentFlags().IntVarP(&trials, "trials", "n", defaultWeightedSelectTrials,
		"加權抽樣次數")
	// go run . weighted-random select -n 10000
	cmd.AddCommand(
		&cobra.Command{
			Use:   "select",
			Short: "Init from map，依 --trials / -n 抽樣並對照預期比例",
			Run: func(*cobra.Command, []string) {
				runWeightedRandomSelectExample(trials)
			},
		},
	)
	root.AddCommand(cmd)
}

func runWeightedRandomSelectExample(n int) {
	if n < 1 {
		fmt.Println("錯誤: --trials / -n 須 ≥1")
		return
	}
	w := cntr.NewWeightedSelector()
	m := map[string]int{
		"apple":    35,
		"banana":   35,
		"coconut":  30,
		"orange":   10,
		"skipped":  0,
		"_ignored": 0,
	}
	w.Init(m)

	expectedProb, totalW := marginalProbFromWeights(m)
	fmt.Printf("總權重=%d（與 WeightedRandom.Init 相同：僅加總 w>0）\n", totalW)
	fmt.Println("各鍵預期機率 P(Select=k)（邊際；w≤0 為 0）：")
	for _, k := range sortedKeys(m) {
		fmt.Printf("  %s: %.4f\n", k, expectedProb[k])
	}

	counts := make(map[string]int, len(m))
	for k := range m {
		counts[k] = 0
	}
	for i := 0; i < n; i++ {
		k := w.Select()
		if _, ok := counts[k]; ok {
			counts[k]++
		}
	}

	fmt.Printf("\n抽樣 %d 次，觀測比例 vs 預期比例（鍵字排序）：\n", n)
	maxAbsDiff := 0.0
	for _, k := range sortedKeys(m) {
		obs := float64(counts[k]) / float64(n)
		exp := expectedProb[k]
		diff := obs - exp
		ad := math.Abs(diff)
		if ad > maxAbsDiff {
			maxAbsDiff = ad
		}
		fmt.Printf("  %-10s 次數=%5d  觀測=%.4f  預期=%.4f  Δ=%+.4f\n", k, counts[k], obs, exp, diff)
	}
	// 樣本夠大時，單鍵 |Δ| 多在數個千分點內；僅作直觀參考
	fmt.Printf("\n各鍵 |觀測−預期| 最大值≈%.4f（樣本越大通常越小）\n", maxAbsDiff)
}

// marginalProbFromWeights 依權重字典計算各鍵邊際機率（與 Init 語意：總權重僅含 w>0）。
func marginalProbFromWeights(weights map[string]int) (prob map[string]float64, total int) {
	for _, w := range weights {
		if w > 0 {
			total += w
		}
	}
	prob = make(map[string]float64, len(weights))
	for k, w := range weights {
		if total == 0 || w <= 0 {
			prob[k] = 0
			continue
		}
		prob[k] = float64(w) / float64(total)
	}
	return prob, total
}

func sortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
