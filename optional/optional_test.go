package optional

import (
	"github.com/ovargas/go-lib/constant"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Of_Nullable(t *testing.T) {

	tests := []struct {
		name            string
		value           *string
		expectedPresent bool
		expectedValue   *string
	}{
		{
			name:            "Present value",
			value:           constant.AsPointer("Hello, World!"),
			expectedPresent: true,
			expectedValue:   constant.AsPointer("Hello, World!"),
		},
		{
			name:            "Nil Value",
			value:           nil,
			expectedPresent: false,
			expectedValue:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Of(tt.value)

			assert.Equal(t, tt.expectedPresent, o.IsPresent())
			assert.Equal(t, tt.expectedValue, o.Get())
		})
	}
}

func Test_Of_Non_Nullable(t *testing.T) {

	tests := []struct {
		name            string
		value           string
		expectedPresent bool
		expectedValue   string
	}{
		{
			name:            "Present value",
			value:           "Hello, World!",
			expectedPresent: true,
			expectedValue:   "Hello, World!",
		},
		{
			name:            "Zero Value",
			expectedPresent: true,
			expectedValue:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Of(tt.value)

			assert.Equal(t, tt.expectedPresent, o.IsPresent())
			assert.Equal(t, tt.expectedValue, o.Get())
		})
	}
}

func TestOptional_OrElse(t *testing.T) {

	tests := []struct {
		name          string
		value         *string
		orElseValue   *string
		expectedValue *string
	}{
		{
			name:          "Return value",
			value:         constant.AsPointer("Hello, World!"),
			orElseValue:   constant.AsPointer("Goodbye, World!"),
			expectedValue: constant.AsPointer("Hello, World!"),
		},
		{
			name:          "Return orElseValue",
			value:         nil,
			orElseValue:   constant.AsPointer("Goodbye, World!"),
			expectedValue: constant.AsPointer("Goodbye, World!"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			actual := Of(tt.value).OrElse(tt.orElseValue)
			assert.Equal(t, tt.expectedValue, actual)
		})

	}
}

func TestOptional_OrElseGet(t *testing.T) {

	tests := []struct {
		name          string
		value         *string
		orElseFunc    func() *string
		expectedValue *string
	}{
		{
			name:          "Return value",
			value:         constant.AsPointer("Hello, World!"),
			orElseFunc:    func() *string { return constant.AsPointer("Goodbye, World!") },
			expectedValue: constant.AsPointer("Hello, World!"),
		},
		{
			name:          "Return orElseValue",
			value:         nil,
			orElseFunc:    func() *string { return constant.AsPointer("Goodbye, World!") },
			expectedValue: constant.AsPointer("Goodbye, World!"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			actual := Of(tt.value).OrElseGet(tt.orElseFunc)
			assert.Equal(t, tt.expectedValue, actual)
		})

	}
}

func TestOptional_IfPresent(t *testing.T) {

	tests := []struct {
		name     string
		value    *string
		executed bool
	}{
		{
			name:     "Present value",
			value:    constant.AsPointer("Hello, World!"),
			executed: true,
		},
		{
			name:     "Nil value",
			value:    nil,
			executed: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := false
			Of(tt.value).IfPresent(func(s *string) {
				actual = true
				assert.Equal(t, tt.value, s)
			})

			assert.Equal(t, tt.executed, actual)
		})
	}
}

func TestOptional_IfPresentOrElse(t *testing.T) {

	tests := []struct {
		name   string
		value  *string
		orElse bool
	}{
		{
			name:   "Present value",
			value:  constant.AsPointer("Hello, World!"),
			orElse: false,
		},
		{
			name:   "Nil value",
			value:  nil,
			orElse: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual *bool
			Of(tt.value).IfPresentOrElse(func(s *string) {
				assert.Equal(t, tt.value, s)
				actual = constant.AsPointer(false)
			}, func() {
				actual = constant.AsPointer(true)
			})

			assert.NotNil(t, actual)
			assert.Equal(t, tt.orElse, *actual)
		})
	}
}

func Test_OfNullable(t *testing.T) {

	tests := []struct {
		name     string
		value    *string
		expected *Optional[*string]
	}{
		{
			name:     "With value",
			value:    constant.AsPointer("Hello, World!"),
			expected: Of(constant.AsPointer("Hello, World!")),
		},
		{
			name:     "With nil",
			value:    nil,
			expected: Empty[*string](),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := OfNullable(tt.value)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_Map(t *testing.T) {
	type From struct {
		name string
	}

	type To struct {
		title string
	}

	tests := []struct {
		name     string
		from     *From
		expected *Optional[*To]
	}{
		{
			name:     "Map Value",
			from:     &From{name: "Hello, World!"},
			expected: Of(&To{title: "Hello, World!"}),
		},
		{
			name:     "Map Nil",
			from:     nil,
			expected: Empty[*To](),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionalFrom := Of(tt.from)
			actual := Map(optionalFrom, func(f *From) *To {
				return &To{title: f.name}
			})
			assert.Equal(t, tt.expected, actual)
		})
	}
}
