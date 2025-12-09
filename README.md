# calculator-go

Pure mathematical calculation library for trading position sizing, risk management, and PnL calculations.

## Overview

`calculator-go` is a standalone Go module that provides position sizing, leverage calculation, risk-reward ratio calculations, and PnL computations. It has no strategy logic - just pure mathematical functions that can be used by strategies, CLI formatters, or any other trading system component.

## Features

- **Position Sizing**: Calculate position size based on account balance and risk percentage
- **Leverage Calculation**: Determine optimal leverage based on notional value
- **Risk-Reward Ratios**: Calculate take profit prices based on RR ratios
- **Price Validation**: Validate entry and stop loss price logic
- **PnL Calculations**: Calculate profit/loss in both nominal and percentage terms
- **Distance Calculations**: Calculate percentage distance to target prices

## Installation

```bash
go get github.com/agatticelli/calculator-go
```

## Usage

```go
import (
    "github.com/agatticelli/calculator-go"
    "github.com/agatticelli/trading-go/broker"
)

// Create calculator instance
calc := calculator.New(125) // max leverage 125x

// Calculate position size based on risk
size := calc.CalculateSize(
    1000.0,           // account balance
    2.0,              // risk percent (2%)
    3950.0,           // entry price
    3900.0,           // stop loss
    broker.SideLong,  // position side
)

// Calculate required leverage
leverage := calc.CalculateLeverage(size, 3950.0, 1000.0, 125)

// Calculate take profit based on risk-reward ratio
tp := calc.CalculateRRTakeProfit(3950.0, 3900.0, 2.0, broker.SideLong)

// Calculate expected PnL
nominal, percent := calc.CalculateExpectedPnL(
    broker.SideLong,
    3950.0,  // entry
    4000.0,  // exit
    0.5,     // size
)

// Validate price logic
err := calc.ValidatePriceLogic(broker.SideLong, 3950.0, 4000.0)
```

## API Reference

### Calculator Methods

#### `New(maxLeverage int) *Calculator`
Creates a new calculator instance with specified max leverage.

#### `CalculateSize(balance, riskPercent, entry, stopLoss float64, side broker.Side) float64`
Calculates position size based on risk.
- Formula: `size = (balance * risk%) / |entry - stopLoss|`

#### `CalculateLeverage(size, price, balance float64, maxLeverage int) int`
Calculates required leverage for a position.
- Formula: `leverage = ceil((size * price) / balance)`

#### `CalculateRRTakeProfit(entry, stopLoss, rrRatio float64, side broker.Side) float64`
Calculates take profit price based on risk-reward ratio.
- LONG: `tp = entry + (|entry - sl| * ratio)`
- SHORT: `tp = entry - (|entry - sl| * ratio)`

#### `ValidatePriceLogic(side broker.Side, entry, current float64) error`
Validates entry price logic to prevent market execution.
- LONG: entry must be below current price
- SHORT: entry must be above current price

#### `ValidateStopLoss(side broker.Side, entry, stopLoss float64) error`
Validates stop loss placement.
- LONG: stop loss must be below entry
- SHORT: stop loss must be above entry

#### `CalculatePnLPercent(side broker.Side, entryPrice, markPrice float64) float64`
Calculates PnL percentage.
- LONG: `((mark - entry) / entry) * 100`
- SHORT: `((entry - mark) / entry) * 100`

#### `CalculateDistanceToPrice(side broker.Side, currentPrice, targetPrice float64) float64`
Calculates percentage distance from current to target price.
- LONG: `((target - current) / current) * 100`
- SHORT: `((current - target) / current) * 100`

#### `CalculateExpectedPnL(side broker.Side, entryPrice, exitPrice, size float64) (nominal, percent float64)`
Calculates expected PnL for a closing order in both nominal and percentage values.

#### `ValidateInputs(side broker.Side, entryPrice, stopLoss, riskPercent, accountEquity float64) error`
Validates all calculation inputs.

## Dependencies

- `github.com/agatticelli/trading-go` - Only for `broker.Side` type definition

## Architecture

This module is part of a 5-module trading system:
- **calculator-go** (this module): Pure mathematical calculations
- **trading-go**: Broker abstraction layer
- **strategy-go**: Trading strategy implementations
- **intent-go**: NLP command processing
- **trading-cli**: Command-line interface

## License

MIT
