package main

import (
	"fmt"
	"strings"

	"github.com/agatticelli/calculator-go"
	"github.com/agatticelli/trading-common-types"
)

// This example demonstrates position sizing and leverage calculation
func main() {
	fmt.Println("=== Position Sizing Example ===\n")

	// Create calculator with max leverage of 125x
	calc := calculator.New(125)

	// Trading scenario
	accountBalance := 1000.0
	riskPercent := 2.0 // Risk 2% of account
	entryPrice := 45000.0
	stopLoss := 44500.0

	// Calculate position size for a LONG position
	fmt.Println("📊 Calculating LONG position:")
	fmt.Printf("  Account Balance: $%.2f\n", accountBalance)
	fmt.Printf("  Risk Percentage: %.1f%%\n", riskPercent)
	fmt.Printf("  Entry Price: $%.2f\n", entryPrice)
	fmt.Printf("  Stop Loss: $%.2f\n", stopLoss)
	fmt.Println()

	size := calc.CalculateSize(accountBalance, riskPercent, entryPrice, stopLoss, types.SideLong)
	fmt.Printf("✅ Position Size: %.4f contracts\n", size)

	// Calculate required leverage
	leverage := calc.CalculateLeverage(size, entryPrice, accountBalance, 125)
	fmt.Printf("✅ Required Leverage: %dx\n", leverage)

	// Calculate notional value
	notionalValue := size * entryPrice
	fmt.Printf("✅ Notional Value: $%.2f\n", notionalValue)

	// Calculate risk amount
	riskAmount := accountBalance * (riskPercent / 100)
	fmt.Printf("✅ Risk Amount: $%.2f (%.1f%% of balance)\n", riskAmount, riskPercent)

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Now calculate for a SHORT position
	shortEntry := 45000.0
	shortStopLoss := 45500.0

	fmt.Println("\n📊 Calculating SHORT position:")
	fmt.Printf("  Account Balance: $%.2f\n", accountBalance)
	fmt.Printf("  Risk Percentage: %.1f%%\n", riskPercent)
	fmt.Printf("  Entry Price: $%.2f\n", shortEntry)
	fmt.Printf("  Stop Loss: $%.2f\n", shortStopLoss)
	fmt.Println()

	shortSize := calc.CalculateSize(accountBalance, riskPercent, shortEntry, shortStopLoss, types.SideShort)
	fmt.Printf("✅ Position Size: %.4f contracts\n", shortSize)

	shortLeverage := calc.CalculateLeverage(shortSize, shortEntry, accountBalance, 125)
	fmt.Printf("✅ Required Leverage: %dx\n", shortLeverage)

	shortNotional := shortSize * shortEntry
	fmt.Printf("✅ Notional Value: $%.2f\n", shortNotional)

	fmt.Println("\n💡 Note: The calculator automatically handles LONG vs SHORT logic")
	fmt.Println("   - LONG: Stop Loss must be below Entry")
	fmt.Println("   - SHORT: Stop Loss must be above Entry")
}
