# Calculator-Go Examples

This directory contains working examples demonstrating how to use calculator-go in real-world scenarios.

## Prerequisites

```bash
# Initialize the examples module (already done)
go mod init calculator-examples

# Get the calculator-go dependency
go get github.com/agatticelli/calculator-go
```

## Running the Examples

### 1. Position Sizing Example
Demonstrates how to calculate position size and required leverage based on account balance and risk percentage.

```bash
go run position_sizing.go
```

**What it shows:**
- Calculate position size for LONG and SHORT positions
- Determine required leverage based on notional value
- Calculate risk amount and notional value
- Automatic handling of LONG vs SHORT logic

### 2. Risk-Reward Ratio Example
Shows how to calculate take profit levels based on different risk-reward ratios.

```bash
go run risk_reward.go
```

**What it shows:**
- Calculate TP prices for common R:R ratios (1:1, 1.5:1, 2:1, 3:1, 5:1)
- Compare risk vs reward amounts
- Handle both LONG and SHORT positions
- Understand how R:R ratios work in practice

### 3. PnL Calculations Example
Demonstrates profit/loss calculations for open positions and pending orders.

```bash
go run pnl_calculations.go
```

**What it shows:**
- Calculate current PnL (percentage and nominal)
- Calculate distance to take profit and stop loss
- Calculate expected PnL for pending close orders
- Handle both LONG and SHORT PnL logic

### 4. Validation Example
Shows how to validate prices to prevent common trading errors.

```bash
go run validation.go
```

**What it shows:**
- Validate LONG limit orders (entry must be below current price)
- Validate SHORT limit orders (entry must be above current price)
- Validate stop loss placement (below entry for LONG, above for SHORT)
- Validate all inputs (prices, risk percentage, account equity)
- Error messages for invalid configurations

## Run All Examples

```bash
# Position sizing
go run position_sizing.go

# Risk-reward ratios
go run risk_reward.go

# PnL calculations
go run pnl_calculations.go

# Price validation
go run validation.go
```

## Understanding the Output

Each example is annotated with:
- 📊 Scenario descriptions
- ✅ Successful calculations
- ❌ Validation errors
- 💡 Pro tips and best practices

## Using in Your Own Code

These examples show complete, runnable code that you can copy and adapt for your own trading system. Key patterns demonstrated:

1. **Create a calculator instance:**
   ```go
   calc := calculator.New(125) // max leverage 125x
   ```

2. **Calculate position size:**
   ```go
   size := calc.CalculateSize(balance, riskPercent, entry, stopLoss, calculator.SideLong)
   ```

3. **Validate before trading:**
   ```go
   if err := calc.ValidateInputs(side, entry, stopLoss, riskPercent, equity); err != nil {
       return err
   }
   ```

## Integration with Trading Systems

calculator-go is designed to be **dependency-free** and can be integrated into any trading system:

- **CLI formatters**: Use for displaying PnL, distances, and metrics
- **Trading strategies**: Use for position sizing and risk management
- **Backtesting**: Use for calculating historical PnL
- **Risk management**: Use for validating trade parameters
- **Portfolio management**: Use for multi-position calculations

## Further Reading

- [Main README](../README.md) - Full API documentation
- [calculator.go](../calculator.go) - Source code with implementation details
- [MIGRATION_STATUS.md](../../trading-cli/MIGRATION_STATUS.md) - Architecture documentation
