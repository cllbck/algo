package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMap(t *testing.T) {
	h := NewHashMap(5)
	const count = 100
	for i := range [count]struct{}{} {
		h.Set(strconv.FormatInt(int64(i), 10), strconv.FormatInt(int64(i), 10))
	}
	assert.Equal(t, count, h.len)
	assert.Equal(t, 80, h.capacity)
	for i := range [count]struct{}{} {
		v, ok := h.Get(strconv.FormatInt(int64(i), 10))
		assert.Equal(t, v, strconv.FormatInt(int64(i), 10))
		assert.True(t, ok)
	}
	assert.Equal(t, count, h.len)
	h.Print()
	for i := range [count]struct{}{} {
		err := h.Delete(strconv.FormatInt(int64(i), 10))
		assert.NoError(t, err)
	}
}

func TestHashMapDeleteMiddleElementFromChain(t *testing.T) {
	h := NewHashMap(5)
	h.Set("123", "123")
	h.Set("321", "321")
	h.Set("312", "312")

	assert.Equal(t, h.len, 3)

	err := h.Delete("321")

	assert.NoError(t, err)
	assert.Equal(t, h.len, 2)

	_, ok312 := h.Get("321")
	assert.False(t, ok312)

	value312, ok312 := h.Get("312")
	assert.True(t, ok312)
	assert.Equal(t, value312, "312")

	value123, ok123 := h.Get("123")
	assert.True(t, ok123)
	assert.Equal(t, value123, "123")

}

func TestHashMapRehash(t *testing.T) {
	h := NewHashMap(10)
	h.Set("123", "123")
	h.Set("123", "test")
	h.Set("312", "321")
	assert.Equal(t, h.len, 2)
	assert.Equal(t, 10, h.capacity)
	h.rehash()
	assert.Equal(t, h.len, 2)
	assert.Equal(t, 20, h.capacity)
}
