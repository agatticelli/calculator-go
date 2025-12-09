package main

import (
	"fmt"
	"strings"

	"github.com/agatticelli/calculator-go"
	"github.com/agatticelli/trading-common-types"
)

// This example demonstrates price validation to prevent common trading errors
func main() {
	fmt.Println("=== Price Validation Example ===\n")

	calc := calculator.New(125)

	currentPrice := 45000.0
	fmt.Printf("Current Market Price: $%.2f\n\n", currentPrice)

	// Test 1: LONG limit order validation
	fmt.Println("📊 Test 1: LONG Limit Order Validation")
	fmt.Println("  Rule: Entry must be BELOW current price (to avoid market execution)")
	fmt.Println()

	validLongEntry := 44500.0
	fmt.Printf("  ✓ Valid LONG entry at $%.2f: ", validLongEntry)
	if err := calc.ValidatePriceLogic(types.SideLong, validLongEntry, currentPrice); err != nil {
		fmt.Printf("❌ %v\n", err)
	} else {
		fmt.Println("✅ Valid")
	}

	invalidLongEntry := 45500.0
	fmt.Printf("  ✗ Invalid LONG entry at $%.2f: ", invalidLongEntry)
	if err := calc.ValidatePriceLogic(types.SideLong, invalidLongEntry, currentPrice); err != nil {
		fmt.Printf("❌ %v\n", err)
	} else {
		fmt.Println("✅ Valid")
	}

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Test 2: SHORT limit order validation
	fmt.Println("\n📊 Test 2: SHORT Limit Order Validation")
	fmt.Println("  Rule: Entry must be ABOVE current price (to avoid market execution)")
	fmt.Println()

	validShortEntry := 45500.0
	fmt.Printf("  ✓ Valid SHORT entry at $%.2f: ", validShortEntry)
	if err := calc.ValidatePriceLogic(types.SideShort, validShortEntry, currentPrice); err != nil {
		fmt.Printf("❌ %v\n", err)
	} else {
		fmt.Println("✅ Valid")
	}

	invalidShortEntry := 44500.0
	fmt.Printf("  ✗ Invalid SHORT entry at $%.2f: ", invalidShortEntry)
	if err := calc.ValidatePriceLogic(types.SideShort, invalidShortEntry, currentPrice); err != nil {
		fmt.Printf("❌ %v\n", err)
	} else {
		fmt.Println("✅ Valid")
	}

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Test 3: Stop Loss validation for LONG
	fmt.Println("\n📊 Test 3: Stop Loss Validation (LONG)")
	fmt.Println("  Rule: Stop Loss must be BELOW entry price")
	fmt.Println()

	longEntry := 45000.0
	validSL := 44500.0
	fmt.Printf("  ✓ Valid SL at $%.2f (entry: $%.2f): ", validSL, longEntry)
	if err := calc.ValidateStopLoss(types.SideLong, longEntry, validSL); err != nil {
		fmt.Printf("❌ %v\n", err)
	} else {
		fmt.Println("✅ Valid")
	}

	invalidSL := 45500.0
	fmt.Printf("  ✗ Invalid SL at $%.2f (entry: $%.2f): ", invalidSL, longEntry)
	if err := calc.ValidateStopLoss(types.SideLong, longEntry, invalidSL); err != nil {
		fmt.Printf("❌ %v\n", err)
	} else {
		fmt.Println("✅ Valid")
	}

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Test 4: Stop Loss validation for SHORT
	fmt.Println("\n📊 Test 4: Stop Loss Validation (SHORT)")
	fmt.Println("  Rule: Stop Loss must be ABOVE entry price")
	fmt.Println()

	shortEntry := 45000.0
	validShortSL := 45500.0
	fmt.Printf("  ✓ Valid SL at $%.2f (entry: $%.2f): ", validShortSL, shortEntry)
	if err := calc.ValidateStopLoss(types.SideShort, shortEntry, validShortSL); err != nil {
		fmt.Printf("❌ %v\n", err)
	} else {
		fmt.Println("✅ Valid")
	}

	invalidShortSL := 44500.0
	fmt.Printf("  ✗ Invalid SL at $%.2f (entry: $%.2f): ", invalidShortSL, shortEntry)
	if err := calc.ValidateStopLoss(types.SideShort, shortEntry, invalidShortSL); err != nil {
		fmt.Printf("❌ %v\n", err)
	} else {
		fmt.Println("✅ Valid")
	}

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Test 5: Complete input validation
	fmt.Println("\n📊 Test 5: Complete Input Validation")
	fmt.Println("  Validates: entry > 0, SL > 0, risk% in (0,100], equity > 0, SL placement")
	fmt.Println()

	// Valid inputs
	fmt.Println("  ✓ Valid inputs:")
	err := calc.ValidateInputs(types.SideLong, 45000.0, 44500.0, 2.0, 1000.0)
	if err != nil {
		fmt.Printf("    ❌ %v\n", err)
	} else {
		fmt.Println("    ✅ All validations passed")
	}

	// Invalid risk percentage
	fmt.Println("\n  ✗ Invalid risk percentage (150%):")
	err = calc.ValidateInputs(types.SideLong, 45000.0, 44500.0, 150.0, 1000.0)
	if err != nil {
		fmt.Printf("    ❌ %v\n", err)
	} else {
		fmt.Println("    ✅ All validations passed")
	}

	// Invalid stop loss placement
	fmt.Println("\n  ✗ Invalid stop loss placement (above entry for LONG):")
	err = calc.ValidateInputs(types.SideLong, 45000.0, 45500.0, 2.0, 1000.0)
	if err != nil {
		fmt.Printf("    ❌ %v\n", err)
	} else {
		fmt.Println("    ✅ All validations passed")
	}

	fmt.Println("\n💡 Use these validation functions to prevent trading errors before submitting orders!")
}
