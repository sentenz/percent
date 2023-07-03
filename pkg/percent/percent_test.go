// MIT License
//
// Copyright (c) 2023 sentenz
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package percent_test

import (
	"testing"

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
			name: "valid nagative input",
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
			name: "invalid percentage",
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
			got, err := percent.Percent(tt.in.percent, tt.in.value)
			if err != tt.want.err {
				t.Errorf("Percent() error = %v, want err %v", err, tt.want.err)
				return
			}
			if got != tt.want.value {
				t.Errorf("Percent() = %v, want %v", got, tt.want.value)
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
			name: "valid nagative input",
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
			name: "invalid part",
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
			name: "invalid total",
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
			got, err := percent.Of(tt.in.part, tt.in.total)
			if err != tt.want.err {
				t.Errorf("Of() error = %v, want err %v", err, tt.want.err)
				return
			}
			if got != tt.want.value {
				t.Errorf("Of() = %v, want %v", got, tt.want.value)
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
			name: "invalid zero value",
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
			got, err := percent.Change(tt.in.oldValue, tt.in.newValue)
			if err != tt.want.err {
				t.Errorf("Change() error = %v, want err %v", err, tt.want.err)
				return
			}
			if got != tt.want.value {
				t.Errorf("Change() = %v, want %v", got, tt.want.value)
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
			name: "valid nagative input",
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
			name: "invalid percentage input",
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
			got, err := percent.Remain(tt.in.percent, tt.in.value)
			if err != tt.want.err {
				t.Errorf("Remain() error = %v, want err %v", err, tt.want.err)
				return
			}
			if got != tt.want.value {
				t.Errorf("Remain() = %v, want %v", got, tt.want.value)
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
			name: "invalid input",
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
			got, err := percent.FromRatio(tt.in.ratio)
			if err != tt.want.err {
				t.Errorf("FromRatio() error = %v, want err %v", err, tt.want.err)
				return
			}
			if got != tt.want.value {
				t.Errorf("FromRatio() = %v, want %v", got, tt.want.value)
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
			name: "invalid input",
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
			got, err := percent.ToRatio(tt.in.percent)
			if err != tt.want.err {
				t.Errorf("ToRatio() error = %v, want err %v", err, tt.want.err)
				return
			}
			if got != tt.want.value {
				t.Errorf("ToRatio() = %v, want %v", got, tt.want.value)
			}
		})
	}
}
