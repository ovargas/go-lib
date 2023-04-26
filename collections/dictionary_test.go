package collections

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestDictionary_Get(t *testing.T) {
	tests := []struct {
		name       string
		dictionary Dictionary[string, string]
		key        string
		want       string
		wantErr    error
	}{
		{
			name:       "Get existing key",
			dictionary: Dictionary[string, string]{"key": "value"},
			key:        "key",
			want:       "value",
		},
		{
			name:       "Get non-existing key",
			dictionary: Dictionary[string, string]{"key": "value"},
			key:        "non-existing-key",
			want:       "",
			wantErr:    ErrKeyNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.dictionary.Get(tt.key)

			if tt.wantErr == nil {
				assert.NoError(t, err)
				assert.Equalf(t, tt.want, actual, "Get(%v)", tt.key)
			} else {
				assert.EqualError(t, err, tt.wantErr.Error())
				assert.Equalf(t, tt.want, actual, "Get(%v)", tt.key)
			}
		})
	}
}

func TestDictionary_GetOrDefault(t *testing.T) {
	tests := []struct {
		name         string
		dictionary   Dictionary[string, string]
		key          string
		want         string
		defaultValue string
	}{
		{
			name:         "Get existing key",
			dictionary:   Dictionary[string, string]{"key": "value"},
			key:          "key",
			want:         "value",
			defaultValue: "default-value",
		},
		{
			name:         "Get non-existing key",
			dictionary:   Dictionary[string, string]{"key": "value"},
			key:          "non-existing-key",
			defaultValue: "default-value",
			want:         "default-value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.dictionary.GetOrDefault(tt.key, tt.defaultValue), "GetOrDefault(%v, %v)", tt.key, tt.defaultValue)
		})
	}
}

func TestDictionary_Has(t *testing.T) {
	tests := []struct {
		name       string
		dictionary Dictionary[string, string]
		key        string
		want       bool
	}{
		{
			name:       "Get existing key",
			dictionary: Dictionary[string, string]{"key": "value"},
			key:        "key",
			want:       true,
		},
		{
			name:       "Get non-existing key",
			dictionary: Dictionary[string, string]{"key": "value"},
			key:        "non-existing-key",
			want:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.dictionary.Has(tt.key), "Has(%v)", tt.key)
		})
	}
}

func TestDictionary_Keys(t *testing.T) {
	tests := []struct {
		name       string
		dictionary Dictionary[string, string]
		want       []string
	}{
		{
			name:       "Get non empty dictionary",
			dictionary: Dictionary[string, string]{"key": "value", "key2": "value2"},
			want:       []string{"key", "key2"},
		},
		{
			name:       "Get empty dictionary",
			dictionary: Dictionary[string, string]{},
			want:       []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual := tt.dictionary.Keys()

			// Sort the slices to make sure they are equal, the keys from the
			// dictionary are not ensured to be in order.
			sort.Strings(actual)

			assert.Equalf(t, tt.want, actual, "Keys()")
		})
	}
}

func TestDictionary_Remove(t *testing.T) {
	tests := []struct {
		name       string
		dictionary Dictionary[string, string]
		key        string
		want       bool
		wantLength int
	}{
		{
			name:       "Delete existing key",
			dictionary: Dictionary[string, string]{"key": "value"},
			key:        "key",
			wantLength: 0,
		},
		{
			name:       "Try deleting a non-existing key",
			dictionary: Dictionary[string, string]{"key": "value"},
			key:        "non-existing-key",
			wantLength: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dictionary.Remove(tt.key)

			assert.Equalf(t, tt.wantLength, len(tt.dictionary), "Remove(%v)", tt.key)
			assert.Falsef(t, tt.dictionary.Has(tt.key), "Has(%v)", tt.key)
		})
	}
}

func TestDictionary_Set(t *testing.T) {
	tests := []struct {
		name       string
		dictionary Dictionary[string, string]
		key        string
		value      string
		want       string
		wantLength int
	}{
		{
			name:       "Set existing key",
			dictionary: Dictionary[string, string]{"key": "value"},
			key:        "key",
			value:      "new-value",
			want:       "new-value",
			wantLength: 1,
		},
		{
			name:       "Set non-existing key",
			dictionary: Dictionary[string, string]{"key": "value"},
			key:        "non-existing-key",
			value:      "the-value",
			want:       "the-value",
			wantLength: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			oldValue, _ := tt.dictionary.Get(tt.key)

			tt.dictionary.Set(tt.key, tt.value)

			newValue, _ := tt.dictionary.Get(tt.key)

			assert.Equalf(t, tt.want, newValue, "Set(%v, %v)", tt.key, tt.value)
			assert.NotEqualf(t, oldValue, newValue, "Set(%v, %v)", tt.key, tt.value)
			assert.Equalf(t, tt.wantLength, len(tt.dictionary), "Set(%v, %v)", tt.key, tt.value)
		})
	}
}

func TestDictionary_Values(t *testing.T) {
	tests := []struct {
		name       string
		dictionary Dictionary[string, string]
		want       []string
	}{
		{
			name:       "Get non empty dictionary",
			dictionary: Dictionary[string, string]{"key": "value", "key2": "value2"},
			want:       []string{"value", "value2"},
		},
		{
			name:       "Get empty dictionary",
			dictionary: Dictionary[string, string]{},
			want:       []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual := tt.dictionary.Values()

			// Sort the slices to make sure they are equal, the values from the
			// dictionary are not ensured to be in order.
			sort.Strings(actual)

			assert.Equalf(t, tt.want, actual, "Values()")
		})
	}
}

func TestDictionary_Merge(t *testing.T) {
	tests := []struct {
		name       string
		dictionary Dictionary[string, string]
		other      Dictionary[string, string]
		want       Dictionary[string, string]
		wantLength int
	}{
		{
			name:       "Adding new keys",
			dictionary: Dictionary[string, string]{"key": "value"},
			other:      Dictionary[string, string]{"key2": "value2"},
			want:       Dictionary[string, string]{"key": "value", "key2": "value2"},
			wantLength: 2,
		},
		{
			name:       "Overwriting existing keys",
			dictionary: Dictionary[string, string]{"key": "value"},
			other:      Dictionary[string, string]{"key": "new-value"},
			want:       Dictionary[string, string]{"key": "new-value"},
			wantLength: 1,
		},
		{
			name:       "Overwriting existing keys and adding new ones",
			dictionary: Dictionary[string, string]{"key": "value"},
			other:      Dictionary[string, string]{"key": "new-value", "key2": "value2"},
			want:       Dictionary[string, string]{"key": "new-value", "key2": "value2"},
			wantLength: 2,
		},
		{
			name:       "Do nothing",
			dictionary: Dictionary[string, string]{"key": "value"},
			other:      nil,
			want:       Dictionary[string, string]{"key": "value"},
			wantLength: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dictionary.Merge(tt.other)

			assert.Equalf(t, tt.wantLength, len(tt.dictionary), "Merge(%v)", tt.other)
			assert.Equalf(t, tt.want, tt.dictionary, "Merge(%v)", tt.other)
		})
	}
}

func TestDictionary_Iterator(t *testing.T) {
	dictionary := Dictionary[string, string]{"key": "value", "key2": "value2"}

	var result []string
	for it := dictionary.Iterator(); it.HasNext(); {
		n := it.Next()
		result = append(result, fmt.Sprintf("%s+%s", n.Key(), n.Value()))
	}

	// Sort the slices to make sure they are equal, the values from the
	// dictionary are not ensured to be in order.
	sort.Strings(result)

	assert.Equalf(t, []string{"key+value", "key2+value2"}, result, "Iterator()")
}
