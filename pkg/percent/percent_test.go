// SPDX-License-Identifier: Apache-2.0

package percent_test

import (
	"errors"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/sentenz/percent/internal/pkg/resource"
	"github.com/sentenz/percent/pkg/percent"
)

func TestPercent(t *testing.T) {
	t.Parallel()

	type in struct {
		percent float64
		value   float64
	}

	type want struct {
		value float64
		err   error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "valid positive input",
			in: in{
				percent: 25,
				value:   100,
			},
			want: want{
				value: 25,
				err:   nil,
			},
		},
		{
			name: "valid negative input",
			in: in{
				percent: 50,
				value:   -200,
			},
			want: want{
				value: -100,
				err:   nil,
			},
		},
		{
			name: "zero percent",
			in: in{
				percent: 0,
				value:   100,
			},
			want: want{
				value: 0,
				err:   nil,
			},
		},
		{
			name: "hundred percent",
			in: in{
				percent: 100,
				value:   50,
			},
			want: want{
				value: 50,
				err:   nil,
			},
		},
		{
			name: "negative percent",
			in: in{
				percent: -10,
				value:   100,
			},
			want: want{
				value: 0,
				err:   resource.ErrOutOfRange,
			},
		},
		{
			name: "invalid percentage over 100",
			in: in{
				percent: 150,
				value:   100,
			},
			want: want{
				value: 0.0,
				err:   resource.ErrOutOfRange,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange

			// Act
			got, err := percent.Percent(tt.in.percent, tt.in.value)

			// Assert
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Percent() error = %v, want err %v", err, tt.want.err)
			}
			if !cmp.Equal(got, tt.want.value) {
				t.Errorf("Percent(%+v) = %v, want %v", tt.in, got, tt.want.value)
			}
		})
	}
}

func TestOf(t *testing.T) {
	t.Parallel()

	type in struct {
		part  float64
		total float64
	}

	type want struct {
		value float64
		err   error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "valid positive input",
			in: in{
				part:  25,
				total: 100,
			},
			want: want{
				value: 25,
				err:   nil,
			},
		},
		{
			name: "valid negative input",
			in: in{
				part:  -200,
				total: -50,
			},
			want: want{
				value: 400,
				err:   nil,
			},
		},
		{
			name: "part equals total",
			in: in{
				part:  100,
				total: 100,
			},
			want: want{
				value: 100,
				err:   nil,
			},
		},
		{
			name: "zero part",
			in: in{
				part:  0,
				total: 100,
			},
			want: want{
				value: 0,
				err:   nil,
			},
		},
		{
			name: "part greater than total",
			in: in{
				part:  150,
				total: 100,
			},
			want: want{
				value: 0.0,
				err:   resource.ErrPartGreaterThanTotal,
			},
		},
		{
			name: "zero total",
			in: in{
				part:  150,
				total: 0,
			},
			want: want{
				value: 0.0,
				err:   resource.ErrDivideByZero,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange

			// Act
			got, err := percent.Of(tt.in.part, tt.in.total)

			// Assert
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Of() error = %v, want err %v", err, tt.want.err)
			}
			if !cmp.Equal(got, tt.want.value) {
				t.Errorf("Of(%+v) = %v, want %v", tt.in, got, tt.want.value)
			}
		})
	}
}

func TestChange(t *testing.T) {
	t.Parallel()

	type in struct {
		oldValue float64
		newValue float64
	}

	type want struct {
		value float64
		err   error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "valid increase input",
			in: in{
				oldValue: 25,
				newValue: 100,
			},
			want: want{
				value: 300,
				err:   nil,
			},
		},
		{
			name: "valid decrease input",
			in: in{
				oldValue: -50,
				newValue: -200,
			},
			want: want{
				value: -300,
				err:   nil,
			},
		},
		{
			name: "no change",
			in: in{
				oldValue: 100,
				newValue: 100,
			},
			want: want{
				value: 0,
				err:   nil,
			},
		},
		{
			name: "zero old value",
			in: in{
				oldValue: 0,
				newValue: 100,
			},
			want: want{
				value: 0,
				err:   resource.ErrDivideByZero,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange

			// Act
			got, err := percent.Change(tt.in.oldValue, tt.in.newValue)

			// Assert
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Change() error = %v, want err %v", err, tt.want.err)
			}
			if !cmp.Equal(got, tt.want.value) {
				t.Errorf("Change(%+v) = %v, want %v", tt.in, got, tt.want.value)
			}
		})
	}
}

func TestRemain(t *testing.T) {
	t.Parallel()

	type in struct {
		percent float64
		value   float64
	}

	type want struct {
		value float64
		err   error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "valid positive input",
			in: in{
				percent: 25,
				value:   100,
			},
			want: want{
				value: 75,
				err:   nil,
			},
		},
		{
			name: "valid negative input",
			in: in{
				percent: 50,
				value:   -200,
			},
			want: want{
				value: -100,
				err:   nil,
			},
		},
		{
			name: "zero percent",
			in: in{
				percent: 0,
				value:   100,
			},
			want: want{
				value: 100,
				err:   nil,
			},
		},
		{
			name: "hundred percent",
			in: in{
				percent: 100,
				value:   50,
			},
			want: want{
				value: 0,
				err:   nil,
			},
		},
		{
			name: "negative percent",
			in: in{
				percent: -10,
				value:   100,
			},
			want: want{
				value: 0,
				err:   resource.ErrOutOfRange,
			},
		},
		{
			name: "invalid percentage over 100",
			in: in{
				percent: 150,
				value:   100,
			},
			want: want{
				value: 0,
				err:   resource.ErrOutOfRange,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange

			// Act
			got, err := percent.Remain(tt.in.percent, tt.in.value)

			// Assert
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Remain() error = %v, want err %v", err, tt.want.err)
			}
			if !cmp.Equal(got, tt.want.value) {
				t.Errorf("Remain(%+v) = %v, want %v", tt.in, got, tt.want.value)
			}
		})
	}
}

func TestFromRatio(t *testing.T) {
	t.Parallel()

	type in struct {
		ratio float64
	}

	type want struct {
		value float64
		err   error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "valid input",
			in: in{
				ratio: 0.25,
			},
			want: want{
				value: 25,
				err:   nil,
			},
		},
		{
			name: "zero ratio",
			in: in{
				ratio: 0,
			},
			want: want{
				value: 0,
				err:   nil,
			},
		},
		{
			name: "one ratio",
			in: in{
				ratio: 1,
			},
			want: want{
				value: 100,
				err:   nil,
			},
		},
		{
			name: "negative ratio",
			in: in{
				ratio: -0.1,
			},
			want: want{
				value: 0,
				err:   resource.ErrOutOfRange,
			},
		},
		{
			name: "ratio over 1",
			in: in{
				ratio: 2,
			},
			want: want{
				value: 0,
				err:   resource.ErrOutOfRange,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange

			// Act
			got, err := percent.FromRatio(tt.in.ratio)

			// Assert
			if !errors.Is(err, tt.want.err) {
				t.Errorf("FromRatio() error = %v, want err %v", err, tt.want.err)
			}
			if !cmp.Equal(got, tt.want.value) {
				t.Errorf("FromRatio(%+v) = %v, want %v", tt.in, got, tt.want.value)
			}
		})
	}
}

func TestToRatio(t *testing.T) {
	t.Parallel()

	type in struct {
		percent float64
	}

	type want struct {
		value float64
		err   error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "valid positive input",
			in: in{
				percent: 50,
			},
			want: want{
				value: 0.5,
				err:   nil,
			},
		},
		{
			name: "zero percent",
			in: in{
				percent: 0,
			},
			want: want{
				value: 0,
				err:   nil,
			},
		},
		{
			name: "hundred percent",
			in: in{
				percent: 100,
			},
			want: want{
				value: 1,
				err:   nil,
			},
		},
		{
			name: "negative percent",
			in: in{
				percent: -10,
			},
			want: want{
				value: 0,
				err:   resource.ErrOutOfRange,
			},
		},
		{
			name: "percent over 100",
			in: in{
				percent: 150,
			},
			want: want{
				value: 0,
				err:   resource.ErrOutOfRange,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange

			// Act
			got, err := percent.ToRatio(tt.in.percent)

			// Assert
			if !errors.Is(err, tt.want.err) {
				t.Errorf("ToRatio() error = %v, want err %v", err, tt.want.err)
			}
			if !cmp.Equal(got, tt.want.value) {
				t.Errorf("ToRatio(%+v) = %v, want %v", tt.in, got, tt.want.value)
			}
		})
	}
}

func FuzzPercent(f *testing.F) {
	// Seed corpus with edge cases using testcases array
	testcases := []struct {
		percent float64
		value   float64
	}{
		{0.0, 100.0},   // zero percent
		{100.0, 100.0}, // hundred percent
		{50.0, 200.0},  // typical case
		{25.0, -100.0}, // negative value
		{-10.0, 100.0}, // negative percent (should error)
		{150.0, 100.0}, // over 100 percent (should error)
	}
	for _, tc := range testcases {
		f.Add(tc.percent, tc.value) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, pct float64, value float64) {
		// Arrange
		// No special arrangement needed

		// Act
		got, err := percent.Percent(pct, value)

		// Assert
		// Property 1: Function should never panic
		// Property 2: If percent is out of range [0, 100], should return error
		// Property 3: If no error, result should be mathematically correct

		if pct < 0 || pct > 100 {
			// Invalid range - should return error
			if err == nil {
				t.Errorf("Percent(%v, %v) should return error for out of range percent", pct, value)
			}
			// Result should be zero on error
			if got != 0 {
				t.Errorf("Percent(%v, %v) = %v, want 0 on error", pct, value, got)
			}
		} else {
			// Valid range - should not return error
			if err != nil {
				t.Errorf("Percent(%v, %v) returned unexpected error: %v", pct, value, err)
			}
			// Verify mathematical correctness (within floating point precision)
			expected := value * (pct / 100.0)
			if math.Abs(got-expected) > 1e-10 {
				t.Errorf("Percent(%v, %v) = %v, want %v", pct, value, got, expected)
			}
		}
	})
}

func FuzzOf(f *testing.F) {
	// Seed corpus with edge cases using testcases array
	testcases := []struct {
		part  float64
		total float64
	}{
		{25.0, 100.0},   // typical case
		{100.0, 100.0},  // part equals total
		{0.0, 100.0},    // zero part
		{150.0, 100.0},  // part greater than total (should error)
		{150.0, 0.0},    // zero total (should error)
		{-200.0, -50.0}, // negative values
	}
	for _, tc := range testcases {
		f.Add(tc.part, tc.total) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, part float64, total float64) {
		// Arrange
		// No special arrangement needed

		// Act
		got, err := percent.Of(part, total)

		// Assert
		// Property 1: Function should never panic
		// Property 2: If total is zero, should return error
		// Property 3: If part > total, should return error
		// Property 4: If no error, result should be mathematically correct

		if total == 0 {
			// Zero total - should return error
			if err == nil {
				t.Errorf("Of(%v, %v) should return error for zero total", part, total)
			}
			// Result should be zero on error
			if got != 0 {
				t.Errorf("Of(%v, %v) = %v, want 0 on error", part, total, got)
			}
		} else if part > total {
			// Part greater than total - should return error
			if err == nil {
				t.Errorf("Of(%v, %v) should return error for part > total", part, total)
			}
			// Result should be zero on error
			if got != 0 {
				t.Errorf("Of(%v, %v) = %v, want 0 on error", part, total, got)
			}
		} else {
			// Valid input - should not return error
			if err != nil {
				t.Errorf("Of(%v, %v) returned unexpected error: %v", part, total, err)
			}
			// Verify result is mathematically correct
			expected := (part / total) * 100.0
			if math.Abs(got-expected) > 1e-10 {
				t.Errorf("Of(%v, %v) = %v, want %v", part, total, got, expected)
			}
		}
	})
}

func FuzzChange(f *testing.F) {
	// Seed corpus with edge cases using testcases array
	testcases := []struct {
		oldValue float64
		newValue float64
	}{
		{25.0, 100.0},   // increase
		{-50.0, -200.0}, // decrease with negative values
		{100.0, 100.0},  // no change
		{0.0, 100.0},    // zero old value (should error)
		{50.0, 75.0},    // typical increase
	}
	for _, tc := range testcases {
		f.Add(tc.oldValue, tc.newValue) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, oldValue float64, newValue float64) {
		// Arrange
		// No special arrangement needed

		// Act
		got, err := percent.Change(oldValue, newValue)

		// Assert
		// Property 1: Function should never panic
		// Property 2: If oldValue is zero, should return error
		// Property 3: If no error, result should be mathematically correct

		if oldValue == 0 {
			// Zero old value - should return error
			if err == nil {
				t.Errorf("Change(%v, %v) should return error for zero old value", oldValue, newValue)
			}
			// Result should be zero on error
			if got != 0 {
				t.Errorf("Change(%v, %v) = %v, want 0 on error", oldValue, newValue, got)
			}
		} else {
			// Valid input - should not return error
			if err != nil {
				t.Errorf("Change(%v, %v) returned unexpected error: %v", oldValue, newValue, err)
			}
			// Verify mathematical correctness
			expected := ((newValue - oldValue) / math.Abs(oldValue)) * 100.0
			if math.Abs(got-expected) > 1e-10 {
				t.Errorf("Change(%v, %v) = %v, want %v", oldValue, newValue, got, expected)
			}
		}
	})
}

func FuzzRemain(f *testing.F) {
	// Seed corpus with edge cases using testcases array
	testcases := []struct {
		percent float64
		value   float64
	}{
		{25.0, 100.0},  // typical case
		{0.0, 100.0},   // zero percent
		{100.0, 50.0},  // hundred percent
		{50.0, -200.0}, // negative value
		{-10.0, 100.0}, // negative percent (should error)
		{150.0, 100.0}, // over 100 percent (should error)
	}
	for _, tc := range testcases {
		f.Add(tc.percent, tc.value) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, pct float64, value float64) {
		// Arrange
		// No special arrangement needed

		// Act
		got, err := percent.Remain(pct, value)

		// Assert
		// Property 1: Function should never panic
		// Property 2: If percent is out of range [0, 100], should return error
		// Property 3: If no error, result should be mathematically correct

		if pct < 0 || pct > 100 {
			// Invalid range - should return error
			if err == nil {
				t.Errorf("Remain(%v, %v) should return error for out of range percent", pct, value)
			}
			// Result should be zero on error
			if got != 0 {
				t.Errorf("Remain(%v, %v) = %v, want 0 on error", pct, value, got)
			}
		} else {
			// Valid range - should not return error
			if err != nil {
				t.Errorf("Remain(%v, %v) returned unexpected error: %v", pct, value, err)
			}
			// Verify mathematical correctness
			expected := value * ((100.0 - pct) / 100.0)
			if math.Abs(got-expected) > 1e-10 {
				t.Errorf("Remain(%v, %v) = %v, want %v", pct, value, got, expected)
			}
		}
	})
}

func FuzzFromRatio(f *testing.F) {
	// Seed corpus with edge cases using testcases array
	testcases := []float64{
		0.25, // typical ratio
		0.0,  // zero ratio
		1.0,  // one ratio
		-0.1, // negative ratio (should error)
		2.0,  // ratio over 1 (should error)
	}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, ratio float64) {
		// Arrange
		// No special arrangement needed

		// Act
		got, err := percent.FromRatio(ratio)

		// Assert
		// Property 1: Function should never panic
		// Property 2: If ratio is out of range [0, 1], should return error
		// Property 3: If no error, result should be mathematically correct

		if ratio < 0 || ratio > 1 {
			// Invalid range - should return error
			if err == nil {
				t.Errorf("FromRatio(%v) should return error for out of range ratio", ratio)
			}
			// Result should be zero on error
			if got != 0 {
				t.Errorf("FromRatio(%v) = %v, want 0 on error", ratio, got)
			}
		} else {
			// Valid range - should not return error
			if err != nil {
				t.Errorf("FromRatio(%v) returned unexpected error: %v", ratio, err)
			}
			// Verify mathematical correctness
			expected := ratio * 100.0
			if math.Abs(got-expected) > 1e-10 {
				t.Errorf("FromRatio(%v) = %v, want %v", ratio, got, expected)
			}
		}
	})
}

func FuzzToRatio(f *testing.F) {
	// Seed corpus with edge cases using testcases array
	testcases := []float64{
		50.0,  // typical percent
		0.0,   // zero percent
		100.0, // hundred percent
		-10.0, // negative percent (should error)
		150.0, // percent over 100 (should error)
	}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, pct float64) {
		// Arrange
		// No special arrangement needed

		// Act
		got, err := percent.ToRatio(pct)

		// Assert
		// Property 1: Function should never panic
		// Property 2: If percent is out of range [0, 100], should return error
		// Property 3: If no error, result should be mathematically correct

		if pct < 0 || pct > 100 {
			// Invalid range - should return error
			if err == nil {
				t.Errorf("ToRatio(%v) should return error for out of range percent", pct)
			}
			// Result should be zero on error
			if got != 0 {
				t.Errorf("ToRatio(%v) = %v, want 0 on error", pct, got)
			}
		} else {
			// Valid range - should not return error
			if err != nil {
				t.Errorf("ToRatio(%v) returned unexpected error: %v", pct, err)
			}
			// Verify mathematical correctness
			expected := pct / 100.0
			if math.Abs(got-expected) > 1e-10 {
				t.Errorf("ToRatio(%v) = %v, want %v", pct, got, expected)
			}
		}
	})
}

// Benchmark tests

var (
	benchResult float64
	benchError  error
)

func BenchmarkPercent(b *testing.B) {
	benchmarks := []struct {
		name    string
		percent float64
		value   float64
	}{
		{
			name:    "typical case",
			percent: 25.0,
			value:   100.0,
		},
		{
			name:    "zero percent",
			percent: 0.0,
			value:   100.0,
		},
		{
			name:    "hundred percent",
			percent: 100.0,
			value:   100.0,
		},
		{
			name:    "negative value",
			percent: 50.0,
			value:   -200.0,
		},
		{
			name:    "large value",
			percent: 15.5,
			value:   1000000.0,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				benchResult, benchError = percent.Percent(bm.percent, bm.value)
			}
		})
	}
}

func BenchmarkOf(b *testing.B) {
	benchmarks := []struct {
		name  string
		part  float64
		total float64
	}{
		{
			name:  "typical case",
			part:  25.0,
			total: 100.0,
		},
		{
			name:  "part equals total",
			part:  100.0,
			total: 100.0,
		},
		{
			name:  "zero part",
			part:  0.0,
			total: 100.0,
		},
		{
			name:  "negative values",
			part:  -200.0,
			total: -50.0,
		},
		{
			name:  "large values",
			part:  500000.0,
			total: 1000000.0,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				benchResult, benchError = percent.Of(bm.part, bm.total)
			}
		})
	}
}

func BenchmarkChange(b *testing.B) {
	benchmarks := []struct {
		name     string
		oldValue float64
		newValue float64
	}{
		{
			name:     "typical increase",
			oldValue: 25.0,
			newValue: 100.0,
		},
		{
			name:     "typical decrease",
			oldValue: 100.0,
			newValue: 25.0,
		},
		{
			name:     "no change",
			oldValue: 100.0,
			newValue: 100.0,
		},
		{
			name:     "negative values decrease",
			oldValue: -50.0,
			newValue: -200.0,
		},
		{
			name:     "large values",
			oldValue: 500000.0,
			newValue: 1000000.0,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				benchResult, benchError = percent.Change(bm.oldValue, bm.newValue)
			}
		})
	}
}

func BenchmarkRemain(b *testing.B) {
	benchmarks := []struct {
		name    string
		percent float64
		value   float64
	}{
		{
			name:    "typical case",
			percent: 25.0,
			value:   100.0,
		},
		{
			name:    "zero percent",
			percent: 0.0,
			value:   100.0,
		},
		{
			name:    "hundred percent",
			percent: 100.0,
			value:   50.0,
		},
		{
			name:    "negative value",
			percent: 50.0,
			value:   -200.0,
		},
		{
			name:    "large value",
			percent: 33.3,
			value:   1000000.0,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				benchResult, benchError = percent.Remain(bm.percent, bm.value)
			}
		})
	}
}

func BenchmarkFromRatio(b *testing.B) {
	benchmarks := []struct {
		name  string
		ratio float64
	}{
		{
			name:  "typical ratio",
			ratio: 0.25,
		},
		{
			name:  "zero ratio",
			ratio: 0.0,
		},
		{
			name:  "one ratio",
			ratio: 1.0,
		},
		{
			name:  "half ratio",
			ratio: 0.5,
		},
		{
			name:  "small ratio",
			ratio: 0.001,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				benchResult, benchError = percent.FromRatio(bm.ratio)
			}
		})
	}
}

func BenchmarkToRatio(b *testing.B) {
	benchmarks := []struct {
		name    string
		percent float64
	}{
		{
			name:    "typical percent",
			percent: 50.0,
		},
		{
			name:    "zero percent",
			percent: 0.0,
		},
		{
			name:    "hundred percent",
			percent: 100.0,
		},
		{
			name:    "quarter percent",
			percent: 25.0,
		},
		{
			name:    "small percent",
			percent: 0.1,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				benchResult, benchError = percent.ToRatio(bm.percent)
			}
		})
	}
}
