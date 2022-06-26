package main

import "fmt"

type HashMapEntity struct {
	key   string
	value string
	next  *HashMapEntity
	prev  *HashMapEntity
}

func (e *HashMapEntity) SetNext(key, value string) {
	if e.next == nil {
		e.next = &HashMapEntity{
			key:   key,
			value: value,
			next:  nil,
			prev:  e,
		}
		return
	}
	e.next.SetNext(key, value)
}

type HashMap struct {
	len         int
	capacity    int
	rehashIndex float64
	data        []*HashMapEntity
}

func (h *HashMap) Get(key string) (string, bool) {
	e, _, ok := h.get(key)
	if ok {
		return e.value, true
	}
	return "", false
}

func (h *HashMap) Set(key, value string) {
	if float64(h.len)*h.rehashIndex > float64(h.capacity) {
		h.rehash()
	}
	e, _, ok := h.get(key)
	if !ok {
		h.set(key, value)
		h.len += 1
		return
	}
	e.value = value
}

func (h *HashMap) Delete(key string) error {
	e, index, ok := h.get(key)
	if !ok {
		return fmt.Errorf("%s key not found", key)
	}
	if e.prev != nil {
		if e.next != nil {
			e.prev.next = e.next
		} else {
			e.prev.next = nil
		}
		h.len -= 1
		return nil
	}
	if e.next != nil {
		h.data[index] = e.next
	} else {
		h.data[index] = nil
	}

	h.len -= 1
	return nil
}

func (h *HashMap) rehash() {
	newHashMap := NewHashMap(h.capacity * 2)
	for _, e := range h.data {
		if e == nil {
			continue
		}
		newHashMap.set(e.key, e.value)
		if e.next != nil {
			next := e.next
			for next != nil {
				newHashMap.Set(next.key, next.value)
				next = next.next
			}
		}
	}
	h.data = newHashMap.data
	h.capacity = newHashMap.capacity
}

func (h *HashMap) hashCode(key string) int {
	var hashcode int32
	for _, r := range key { // r is rune - int32 unicode codepoint
		hashcode += r
	}
	return int(hashcode)
}

func (h *HashMap) hash(key string) int {
	return h.hashCode(key) % len(h.data)
}

func (h *HashMap) get(key string) (*HashMapEntity, int, bool) {
	hash := h.hash(key)
	e := h.data[hash]
	if e != nil {
		if e.key == key {
			return e, hash, true
		}
		next := e.next
		for next != nil {
			if next.key == key {
				return next, hash, true
			}
			next = next.next
		}
	}
	return nil, 0, false
}

func (h *HashMap) set(key, value string) {
	hash := h.hash(key)
	e := h.data[hash]
	if e != nil {
		if e.key == key {
			e.value = value
			return
		}
		e.SetNext(key, value)
		return
	}
	h.data[hash] = &HashMapEntity{
		key:   key,
		value: value,
		next:  nil,
		prev:  nil,
	}
}

func (h HashMap) Print() {
	fmt.Println("=====================================================================")
	for i, e := range h.data {
		line := fmt.Sprintf("[%v] ", i)
		if e != nil {
			line += fmt.Sprintf(" -> HashMapEntity{%s:%s}", e.key, e.value)
		} else {
			fmt.Println(line)
			continue
		}
		if e.next != nil {
			next := e.next
			for next != nil {
				line += fmt.Sprintf(" -> HashMapEntity{%s:%s}", next.key, next.value)
				next = next.next
			}
		}
		fmt.Println(line)
	}
	fmt.Println("=====================================================================")
}

func NewHashMap(capacity int) HashMap {
	return HashMap{
		len:         0,
		capacity:    capacity,
		rehashIndex: 0.75,
		data:        make([]*HashMapEntity, capacity),
	}
}

func main() {
	h := NewHashMap(5)
	h.Set("27", "test27")
	h.Set("72", "test72")
	h.Set("d", "testD")
	err := h.Delete("321")
	if err != nil {
		fmt.Println(err)
	}
	h.Print()
	h.rehash()
	h.Print()
}
