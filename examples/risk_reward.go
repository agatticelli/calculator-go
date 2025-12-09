package main

import (
	"fmt"
	"strings"

	"github.com/agatticelli/calculator-go"
	"github.com/agatticelli/trading-common-types"
)

// This example demonstrates risk-reward ratio calculations
func main() {
	fmt.Println("=== Risk-Reward Ratio Example ===\n")

	calc := calculator.New(125)

	// LONG position scenario
	entryPrice := 45000.0
	stopLoss := 44500.0

	fmt.Println("📈 LONG Position Risk-Reward Analysis:")
	fmt.Printf("  Entry Price: $%.2f\n", entryPrice)
	fmt.Printf("  Stop Loss: $%.2f\n", stopLoss)
	fmt.Printf("  Risk per contract: $%.2f\n\n", entryPrice-stopLoss)

	// Calculate take profit levels for different risk-reward ratios
	ratios := []float64{1.0, 1.5, 2.0, 3.0, 5.0}

	fmt.Println("Take Profit levels for different R:R ratios:")
	for _, ratio := range ratios {
		tp := calc.CalculateRRTakeProfit(entryPrice, stopLoss, ratio, types.SideLong)
		reward := tp - entryPrice
		risk := entryPrice - stopLoss
		fmt.Printf("  %.1fR → TP: $%.2f (Risk: $%.2f, Reward: $%.2f)\n", ratio, tp, risk, reward)
	}

	fmt.Println("\n" + strings.Repeat("=", 50))

	// SHORT position scenario
	shortEntry := 45000.0
	shortStopLoss := 45500.0

	fmt.Println("\n📉 SHORT Position Risk-Reward Analysis:")
	fmt.Printf("  Entry Price: $%.2f\n", shortEntry)
	fmt.Printf("  Stop Loss: $%.2f\n", shortStopLoss)
	fmt.Printf("  Risk per contract: $%.2f\n\n", shortStopLoss-shortEntry)

	fmt.Println("Take Profit levels for different R:R ratios:")
	for _, ratio := range ratios {
		tp := calc.CalculateRRTakeProfit(shortEntry, shortStopLoss, ratio, types.SideShort)
		reward := shortEntry - tp
		risk := shortStopLoss - shortEntry
		fmt.Printf("  %.1fR → TP: $%.2f (Risk: $%.2f, Reward: $%.2f)\n", ratio, tp, risk, reward)
	}

	fmt.Println("\n💡 Pro Tip: A 2:1 risk-reward ratio means you target $2 profit for every $1 at risk")
	fmt.Println("   This is a popular ratio for swing trading strategies")
}
