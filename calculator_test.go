package calculator

import (
	"math"
	"testing"

	"github.com/agatticelli/trading-common-types"
)

func TestCalculateSize(t *testing.T) {
	calc := New(125)

	tests := []struct {
		name       string
		balance    float64
		riskPct    float64
		entry      float64
		stopLoss   float64
		side       types.Side
		wantSize   float64
	}{
		{
			name:     "LONG position 2% risk",
			balance:  1000.0,
			riskPct:  2.0,
			entry:    45000.0,
			stopLoss: 44500.0,
			side:     types.SideLong,
			wantSize: 0.04, // (1000 * 0.02) / (45000 - 44500) = 20 / 500 = 0.04
		},
		{
			name:     "SHORT position 2% risk",
			balance:  1000.0,
			riskPct:  2.0,
			entry:    45000.0,
			stopLoss: 45500.0,
			side:     types.SideShort,
			wantSize: 0.04, // (1000 * 0.02) / (45500 - 45000) = 20 / 500 = 0.04
		},
		{
			name:     "LONG position 1% risk tight SL",
			balance:  5000.0,
			riskPct:  1.0,
			entry:    3000.0,
			stopLoss: 2950.0,
			side:     types.SideLong,
			wantSize: 1.0, // (5000 * 0.01) / (3000 - 2950) = 50 / 50 = 1.0
		},
		{
			name:     "LONG position 3% risk",
			balance:  10000.0,
			riskPct:  3.0,
			entry:    50000.0,
			stopLoss: 49000.0,
			side:     types.SideLong,
			wantSize: 0.3, // (10000 * 0.03) / (50000 - 49000) = 300 / 1000 = 0.3
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSize := calc.CalculateSize(tt.balance, tt.riskPct, tt.entry, tt.stopLoss, tt.side)
			if math.Abs(gotSize-tt.wantSize) > 0.0001 {
				t.Errorf("CalculateSize() = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}

func TestCalculateLeverage(t *testing.T) {
	calc := New(125)

	tests := []struct {
		name        string
		size        float64
		price       float64
		balance     float64
		maxLeverage int
		want        int
	}{
		{
			name:        "Low leverage position",
			size:        0.1,
			price:       45000.0,
			balance:     10000.0,
			maxLeverage: 125,
			want:        1, // (0.1 * 45000) / 10000 = 0.45, rounds to 1
		},
		{
			name:        "Medium leverage position",
			size:        1.0,
			price:       45000.0,
			balance:     10000.0,
			maxLeverage: 125,
			want:        5, // (1.0 * 45000) / 10000 = 4.5, rounds to 5
		},
		{
			name:        "High leverage position",
			size:        2.0,
			price:       45000.0,
			balance:     1000.0,
			maxLeverage: 125,
			want:        90, // (2.0 * 45000) / 1000 = 90
		},
		{
			name:        "Exceeds max leverage",
			size:        10.0,
			price:       45000.0,
			balance:     1000.0,
			maxLeverage: 125,
			want:        125, // Would be 450, capped at 125
		},
		{
			name:        "Zero leverage case",
			size:        0.01,
			price:       1000.0,
			balance:     100000.0,
			maxLeverage: 125,
			want:        1, // Below 1, rounds to 1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.CalculateLeverage(tt.size, tt.price, tt.balance, tt.maxLeverage)
			if got != tt.want {
				t.Errorf("CalculateLeverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRRTakeProfit(t *testing.T) {
	calc := New(125)

	tests := []struct {
		name     string
		entry    float64
		stopLoss float64
		rrRatio  float64
		side     types.Side
		want     float64
	}{
		{
			name:     "LONG 2:1 RR",
			entry:    45000.0,
			stopLoss: 44500.0,
			rrRatio:  2.0,
			side:     types.SideLong,
			want:     46000.0, // 45000 + (500 * 2) = 46000
		},
		{
			name:     "SHORT 2:1 RR",
			entry:    45000.0,
			stopLoss: 45500.0,
			rrRatio:  2.0,
			side:     types.SideShort,
			want:     44000.0, // 45000 - (500 * 2) = 44000
		},
		{
			name:     "LONG 3:1 RR",
			entry:    3000.0,
			stopLoss: 2950.0,
			rrRatio:  3.0,
			side:     types.SideLong,
			want:     3150.0, // 3000 + (50 * 3) = 3150
		},
		{
			name:     "LONG 1.5:1 RR",
			entry:    50000.0,
			stopLoss: 49000.0,
			rrRatio:  1.5,
			side:     types.SideLong,
			want:     51500.0, // 50000 + (1000 * 1.5) = 51500
		},
		{
			name:     "SHORT 1:1 RR",
			entry:    60000.0,
			stopLoss: 61000.0,
			rrRatio:  1.0,
			side:     types.SideShort,
			want:     59000.0, // 60000 - (1000 * 1) = 59000
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.CalculateRRTakeProfit(tt.entry, tt.stopLoss, tt.rrRatio, tt.side)
			if math.Abs(got-tt.want) > 0.01 {
				t.Errorf("CalculateRRTakeProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidatePriceLogic(t *testing.T) {
	calc := New(125)

	tests := []struct {
		name      string
		side      types.Side
		entry     float64
		current   float64
		wantError bool
	}{
		{
			name:      "LONG valid - entry below current",
			side:      types.SideLong,
			entry:     44000.0,
			current:   45000.0,
			wantError: false,
		},
		{
			name:      "LONG invalid - entry above current",
			side:      types.SideLong,
			entry:     46000.0,
			current:   45000.0,
			wantError: true,
		},
		{
			name:      "SHORT valid - entry above current",
			side:      types.SideShort,
			entry:     46000.0,
			current:   45000.0,
			wantError: false,
		},
		{
			name:      "SHORT invalid - entry below current",
			side:      types.SideShort,
			entry:     44000.0,
			current:   45000.0,
			wantError: true,
		},
		{
			name:      "LONG invalid - entry equals current",
			side:      types.SideLong,
			entry:     45000.0,
			current:   45000.0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := calc.ValidatePriceLogic(tt.side, tt.entry, tt.current)
			if (err != nil) != tt.wantError {
				t.Errorf("ValidatePriceLogic() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestValidateStopLoss(t *testing.T) {
	calc := New(125)

	tests := []struct {
		name      string
		side      types.Side
		entry     float64
		stopLoss  float64
		wantError bool
	}{
		{
			name:      "LONG valid - SL below entry",
			side:      types.SideLong,
			entry:     45000.0,
			stopLoss:  44500.0,
			wantError: false,
		},
		{
			name:      "LONG invalid - SL above entry",
			side:      types.SideLong,
			entry:     45000.0,
			stopLoss:  45500.0,
			wantError: true,
		},
		{
			name:      "SHORT valid - SL above entry",
			side:      types.SideShort,
			entry:     45000.0,
			stopLoss:  45500.0,
			wantError: false,
		},
		{
			name:      "SHORT invalid - SL below entry",
			side:      types.SideShort,
			entry:     45000.0,
			stopLoss:  44500.0,
			wantError: true,
		},
		{
			name:      "LONG invalid - SL equals entry",
			side:      types.SideLong,
			entry:     45000.0,
			stopLoss:  45000.0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := calc.ValidateStopLoss(tt.side, tt.entry, tt.stopLoss)
			if (err != nil) != tt.wantError {
				t.Errorf("ValidateStopLoss() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestCalculatePnLPercent(t *testing.T) {
	calc := New(125)

	tests := []struct {
		name       string
		side       types.Side
		entryPrice float64
		markPrice  float64
		want       float64
	}{
		{
			name:       "LONG profit",
			side:       types.SideLong,
			entryPrice: 45000.0,
			markPrice:  46000.0,
			want:       2.222, // (46000 - 45000) / 45000 * 100 ≈ 2.222%
		},
		{
			name:       "LONG loss",
			side:       types.SideLong,
			entryPrice: 45000.0,
			markPrice:  44000.0,
			want:       -2.222, // (44000 - 45000) / 45000 * 100 ≈ -2.222%
		},
		{
			name:       "SHORT profit",
			side:       types.SideShort,
			entryPrice: 45000.0,
			markPrice:  44000.0,
			want:       2.222, // (45000 - 44000) / 45000 * 100 ≈ 2.222%
		},
		{
			name:       "SHORT loss",
			side:       types.SideShort,
			entryPrice: 45000.0,
			markPrice:  46000.0,
			want:       -2.222, // (45000 - 46000) / 45000 * 100 ≈ -2.222%
		},
		{
			name:       "Break even",
			side:       types.SideLong,
			entryPrice: 45000.0,
			markPrice:  45000.0,
			want:       0.0,
		},
		{
			name:       "Zero entry price",
			side:       types.SideLong,
			entryPrice: 0.0,
			markPrice:  45000.0,
			want:       0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.CalculatePnLPercent(tt.side, tt.entryPrice, tt.markPrice)
			if math.Abs(got-tt.want) > 0.01 {
				t.Errorf("CalculatePnLPercent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateDistanceToPrice(t *testing.T) {
	calc := New(125)

	tests := []struct {
		name         string
		side         types.Side
		currentPrice float64
		targetPrice  float64
		want         float64
	}{
		{
			name:         "LONG distance to TP above",
			side:         types.SideLong,
			currentPrice: 45000.0,
			targetPrice:  46000.0,
			want:         2.222, // (46000 - 45000) / 45000 * 100 ≈ 2.222%
		},
		{
			name:         "LONG distance to SL below",
			side:         types.SideLong,
			currentPrice: 45000.0,
			targetPrice:  44000.0,
			want:         -2.222, // (44000 - 45000) / 45000 * 100 ≈ -2.222%
		},
		{
			name:         "SHORT distance to TP below",
			side:         types.SideShort,
			currentPrice: 45000.0,
			targetPrice:  44000.0,
			want:         2.222, // (45000 - 44000) / 45000 * 100 ≈ 2.222%
		},
		{
			name:         "SHORT distance to SL above",
			side:         types.SideShort,
			currentPrice: 45000.0,
			targetPrice:  46000.0,
			want:         -2.222, // (45000 - 46000) / 45000 * 100 ≈ -2.222%
		},
		{
			name:         "Zero distance",
			side:         types.SideLong,
			currentPrice: 45000.0,
			targetPrice:  45000.0,
			want:         0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.CalculateDistanceToPrice(tt.side, tt.currentPrice, tt.targetPrice)
			if math.Abs(got-tt.want) > 0.01 {
				t.Errorf("CalculateDistanceToPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateExpectedPnL(t *testing.T) {
	calc := New(125)

	tests := []struct {
		name           string
		side           types.Side
		entryPrice     float64
		exitPrice      float64
		size           float64
		wantNominal    float64
		wantPercentage float64
	}{
		{
			name:           "LONG profit",
			side:           types.SideLong,
			entryPrice:     45000.0,
			exitPrice:      46000.0,
			size:           0.5,
			wantNominal:    500.0, // (46000 - 45000) * 0.5 = 500
			wantPercentage: 2.222, // (46000 - 45000) / 45000 * 100
		},
		{
			name:           "LONG loss",
			side:           types.SideLong,
			entryPrice:     45000.0,
			exitPrice:      44000.0,
			size:           0.5,
			wantNominal:    -500.0, // (44000 - 45000) * 0.5 = -500
			wantPercentage: -2.222,
		},
		{
			name:           "SHORT profit",
			side:           types.SideShort,
			entryPrice:     45000.0,
			exitPrice:      44000.0,
			size:           0.5,
			wantNominal:    500.0, // (45000 - 44000) * 0.5 = 500
			wantPercentage: 2.222,
		},
		{
			name:           "SHORT loss",
			side:           types.SideShort,
			entryPrice:     45000.0,
			exitPrice:      46000.0,
			size:           0.5,
			wantNominal:    -500.0, // (45000 - 46000) * 0.5 = -500
			wantPercentage: -2.222,
		},
		{
			name:           "Break even",
			side:           types.SideLong,
			entryPrice:     45000.0,
			exitPrice:      45000.0,
			size:           1.0,
			wantNominal:    0.0,
			wantPercentage: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNominal, gotPercentage := calc.CalculateExpectedPnL(tt.side, tt.entryPrice, tt.exitPrice, tt.size)
			if math.Abs(gotNominal-tt.wantNominal) > 0.01 {
				t.Errorf("CalculateExpectedPnL() nominal = %v, want %v", gotNominal, tt.wantNominal)
			}
			if math.Abs(gotPercentage-tt.wantPercentage) > 0.01 {
				t.Errorf("CalculateExpectedPnL() percentage = %v, want %v", gotPercentage, tt.wantPercentage)
			}
		})
	}
}

func TestValidateInputs(t *testing.T) {
	calc := New(125)

	tests := []struct {
		name          string
		side          types.Side
		entryPrice    float64
		stopLoss      float64
		riskPercent   float64
		accountEquity float64
		wantError     bool
	}{
		{
			name:          "Valid LONG inputs",
			side:          types.SideLong,
			entryPrice:    45000.0,
			stopLoss:      44500.0,
			riskPercent:   2.0,
			accountEquity: 1000.0,
			wantError:     false,
		},
		{
			name:          "Invalid - negative entry price",
			side:          types.SideLong,
			entryPrice:    -45000.0,
			stopLoss:      44500.0,
			riskPercent:   2.0,
			accountEquity: 1000.0,
			wantError:     true,
		},
		{
			name:          "Invalid - negative stop loss",
			side:          types.SideLong,
			entryPrice:    45000.0,
			stopLoss:      -44500.0,
			riskPercent:   2.0,
			accountEquity: 1000.0,
			wantError:     true,
		},
		{
			name:          "Invalid - risk percent too high",
			side:          types.SideLong,
			entryPrice:    45000.0,
			stopLoss:      44500.0,
			riskPercent:   150.0,
			accountEquity: 1000.0,
			wantError:     true,
		},
		{
			name:          "Invalid - zero equity",
			side:          types.SideLong,
			entryPrice:    45000.0,
			stopLoss:      44500.0,
			riskPercent:   2.0,
			accountEquity: 0.0,
			wantError:     true,
		},
		{
			name:          "Invalid - LONG SL above entry",
			side:          types.SideLong,
			entryPrice:    45000.0,
			stopLoss:      45500.0,
			riskPercent:   2.0,
			accountEquity: 1000.0,
			wantError:     true,
		},
		{
			name:          "Valid SHORT inputs",
			side:          types.SideShort,
			entryPrice:    45000.0,
			stopLoss:      45500.0,
			riskPercent:   2.0,
			accountEquity: 1000.0,
			wantError:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := calc.ValidateInputs(tt.side, tt.entryPrice, tt.stopLoss, tt.riskPercent, tt.accountEquity)
			if (err != nil) != tt.wantError {
				t.Errorf("ValidateInputs() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		maxLeverage int
	}{
		{
			name:        "Standard leverage",
			maxLeverage: 125,
		},
		{
			name:        "Low leverage",
			maxLeverage: 10,
		},
		{
			name:        "High leverage",
			maxLeverage: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := New(tt.maxLeverage)
			if calc == nil {
				t.Error("New() returned nil")
			}
		})
	}
}
