package syncpool

import (
	"testing"
)

type Entity struct {
	Name string
	Age  int
}

func TestPool(t *testing.T) {
	entityPool := NewPool[Entity]()

	a := entityPool.Get()
	a.Age = 123
	a.Name = "OldMan"
	entityPool.Put(a) // a should be reset here, before release

	got := entityPool.Get()
	want := Entity{} //zero
	if *got != want {
		t.Errorf("got %v, want %v", *got, want)
	}
}
func BenchmarkPool(b *testing.B) {
	entityPool := NewPool[Entity]()
	for i := 0; i < b.N; i++ {
		a := entityPool.Get()
		a.Age = 123
		a.Name = "OldMan"
		entityPool.Put(a)
	}
}

func BenchmarkNoPool(b *testing.B) {
	var a *Entity
	for i := 0; i < b.N; i++ {
		a = new(Entity)
		a.Age = 123
		a.Name = "OldMan"
	}
}
