package hashmap

import (
	"fmt"
	"testing"
)

func TestMapPutGet(t *testing.T) {
	m := New()
	m.Put(1, "a")
	m.Put(2, "b")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(5, "e")
	m.Put(3, "f") //overwrite

	if actualSize := m.Size(); actualSize != 5 {
		t.Errorf("Size() - Got %v expected %v", actualSize, 5)
	}
	if actualKeys, expectedKeys := m.Keys(), []interface{}{1, 2, 3, 4, 5}; !Equals(actualKeys, expectedKeys) {
		t.Errorf("Keys() - Got %v expected %v", actualKeys, expectedKeys)
	}
	if actualValues, expectedValues := m.Values(), []interface{}{"a", "b", "f", "d", "e"}; !Equals(actualValues, expectedValues) {
		t.Errorf("Values() - Got %v expected %v", actualValues, expectedValues)
	}

	// Testing Get Method
	// key,expectedValue,expectedFound
	expectedMap := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{8, nil, false},
		{3, "f", true},
		{4, "d", true},
		{5, "e", true},
	}

	for _, expectedArray := range expectedMap {
		actualValue, actualFound := m.Get(expectedArray[0])
		if actualValue != expectedArray[1] || actualFound != expectedArray[2] {
			t.Errorf("Get() - Got %v expected %v", actualValue, expectedArray[1])
		}
	}
}

func TestMapPutGetRemove(t *testing.T) {

	// Testing Put
	m := New()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	// Testing Remove
	m.Remove(5)
	m.Remove(6)
	m.Remove(5)
	m.Remove(7)
	m.Remove(8)

	if actualKeys, expectedKeys := m.Keys(), []interface{}{1, 2, 3, 4}; !Equals(actualKeys, expectedKeys) {
		t.Errorf("Keys() - Got %v expected %v", actualKeys, expectedKeys)
	}

	if actualValues, expectedValues := m.Values(), []interface{}{"a", "b", "c", "d"}; !Equals(actualValues, expectedValues) {
		t.Errorf("Values() - Got %v expected %v", actualValues, expectedValues)
	}
	if actualSize := m.Size(); actualSize != 4 {
		t.Errorf("Size() - Got %v expected %v", actualSize, 4)
	}

	// Testing Get
	expectedMap := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, nil, false},
	}

	for _, expectedArray := range expectedMap {
		actualValue, actualFound := m.Get(expectedArray[0])
		if actualValue != expectedArray[1] || actualFound != expectedArray[2] {
			t.Errorf("Get() - Got %v expected %v", actualValue, expectedArray[1])
		}
	}

	// Remove Everything else
	m.Remove(1)
	m.Remove(2)
	m.Remove(4)
	m.Remove(2)
	m.Remove(3)
	m.Remove(2)

	if actualKeys, expectedKeys := fmt.Sprintf("%s", m.Keys()), "[]"; actualKeys != expectedKeys {
		t.Errorf("Keys() - Got %v expected %v", actualKeys, expectedKeys)
	}
	if actualValues, expectedValues := fmt.Sprintf("%s", m.Values()), "[]"; actualValues != expectedValues {
		t.Errorf("Values() - Got %v expected %v", actualValues, expectedValues)
	}
	if actualSize := m.Size(); actualSize != 0 {
		t.Errorf("Size() - Got %v expected %v", actualSize, 0)
	}
	if isEmpty := m.IsEmpty(); !isEmpty {
		t.Errorf("IsEmpty() - Got %v expected %v", isEmpty, true)
	}
}

func TestMapContainsClear(t *testing.T) {
	m := New()
	m.Put(1, "a")
	m.Put(2, "b")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(5, "e")

	// Testing contains method
	if contains := m.Contains(1, 2, 3, 4); !contains {
		t.Errorf("Contains() - Got %v expected %v", contains, true)
	}

	if contains := m.Contains(1, 2, 7, 4); contains {
		t.Errorf("!Contains() - Got %v expected %v", contains, false)
	}

	// Testing Clear method
	m.Clear()

	if actualKeys, expectedKeys := fmt.Sprintf("%s", m.Keys()), "[]"; actualKeys != expectedKeys {
		t.Errorf("Keys() - Got %v expected %v", actualKeys, expectedKeys)
	}
	if actualValues, expectedValues := fmt.Sprintf("%s", m.Values()), "[]"; actualValues != expectedValues {
		t.Errorf("Values() - Got %v expected %v", actualValues, expectedValues)
	}
	if actualSize := m.Size(); actualSize != 0 {
		t.Errorf("Size() - Got %v expected %v", actualSize, 0)
	}
	if isEmpty := m.IsEmpty(); !isEmpty {
		t.Errorf("IsEmpty() - Got %v expected %v", isEmpty, true)
	}

}

func Equals(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for _, av := range a {
		found := false
		for _, bv := range b {
			if av == bv {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func benchmarkGet(b *testing.B, m *Map, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, m *Map, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Put(n, struct{}{})
		}
	}
}

func benchmarkRemove(b *testing.B, m *Map, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Remove(n)
		}
	}
}

func BenchmarkHashMapGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashMapGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashMapGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashMapGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashMapPut100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New()
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashMapPut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashMapPut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashMapPut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashMapRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkHashMapRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkHashMapRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkHashMapRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}
