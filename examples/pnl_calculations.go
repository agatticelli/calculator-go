package main

import (
	"fmt"
	"strings"

	"github.com/agatticelli/calculator-go"
	"github.com/agatticelli/trading-common-types"
)

// This example demonstrates PnL calculations for open and pending positions
func main() {
	fmt.Println("=== PnL Calculations Example ===\n")

	calc := calculator.New(125)

	// Scenario: LONG position
	entryPrice := 45000.0
	positionSize := 0.5
	currentPrice := 46000.0

	fmt.Println("📊 Open LONG Position:")
	fmt.Printf("  Entry Price: $%.2f\n", entryPrice)
	fmt.Printf("  Position Size: %.2f contracts\n", positionSize)
	fmt.Printf("  Current Price: $%.2f\n\n", currentPrice)

	// Calculate current PnL percentage
	pnlPercent := calc.CalculatePnLPercent(types.SideLong, entryPrice, currentPrice)
	fmt.Printf("✅ Current PnL: %.2f%%\n", pnlPercent)

	// Calculate nominal PnL
	nominalPnL := (currentPrice - entryPrice) * positionSize
	fmt.Printf("✅ Nominal PnL: $%.2f\n", nominalPnL)

	// Calculate distance to take profit
	takeProfitPrice := 47000.0
	distanceToTP := calc.CalculateDistanceToPrice(types.SideLong, currentPrice, takeProfitPrice)
	fmt.Printf("✅ Distance to TP ($%.2f): %.2f%%\n", takeProfitPrice, distanceToTP)

	// Calculate distance to stop loss
	stopLossPrice := 44000.0
	distanceToSL := calc.CalculateDistanceToPrice(types.SideLong, currentPrice, stopLossPrice)
	fmt.Printf("✅ Distance to SL ($%.2f): %.2f%%\n", stopLossPrice, distanceToSL)

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Scenario: Pending close order (calculating expected PnL)
	fmt.Println("\n📋 Pending Close Order Analysis:")

	exitPrices := []float64{46000.0, 47000.0, 48000.0}

	fmt.Printf("\nIf you close %.2f contracts at:\n", positionSize)
	for _, exitPrice := range exitPrices {
		nominal, percent := calc.CalculateExpectedPnL(types.SideLong, entryPrice, exitPrice, positionSize)
		fmt.Printf("  $%.2f → PnL: $%.2f (%.2f%%)\n", exitPrice, nominal, percent)
	}

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Scenario: SHORT position
	shortEntry := 45000.0
	shortSize := 0.5
	shortCurrent := 44000.0

	fmt.Println("\n📊 Open SHORT Position:")
	fmt.Printf("  Entry Price: $%.2f\n", shortEntry)
	fmt.Printf("  Position Size: %.2f contracts\n", shortSize)
	fmt.Printf("  Current Price: $%.2f\n\n", shortCurrent)

	shortPnlPercent := calc.CalculatePnLPercent(types.SideShort, shortEntry, shortCurrent)
	fmt.Printf("✅ Current PnL: %.2f%%\n", shortPnlPercent)

	shortNominalPnL := (shortEntry - shortCurrent) * shortSize
	fmt.Printf("✅ Nominal PnL: $%.2f\n", shortNominalPnL)

	shortTP := 43000.0
	shortDistanceToTP := calc.CalculateDistanceToPrice(types.SideShort, shortCurrent, shortTP)
	fmt.Printf("✅ Distance to TP ($%.2f): %.2f%%\n", shortTP, shortDistanceToTP)

	fmt.Println("\n💡 Note: PnL calculations automatically handle LONG vs SHORT logic")
	fmt.Println("   - LONG: Profit when price goes UP")
	fmt.Println("   - SHORT: Profit when price goes DOWN")
}
