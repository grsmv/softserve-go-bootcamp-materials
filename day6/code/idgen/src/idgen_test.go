package main

import (
	"testing"
)

func testIdGenerator(idgen idGenerator, t *testing.T) {
	for i := 1; i <= 100; i++ {
		if id := idgen.Generate(); id != int32(i) {
			t.Errorf("Wrong generated id!!! %d != %d", id, i)
		}
		if id := idgen.Current(); id != int32(i) {
			t.Errorf("Wrong current id!!! %d != %d", id, i)
		}
	}
}

func TestIdGeneratorAtomic(t *testing.T) {
	testIdGenerator(newIdGeneratorAtomic(), t)
}

func TestIdGeneratorMutex(t *testing.T) {
	testIdGenerator(newIdGeneratorMutex(), t)
}

func TestIdGeneratorChan(t *testing.T) {
	testIdGenerator(newIdGeneratorChan(), t)
}

func benchmarkIdGenerator_Generate(idgen idGenerator, b *testing.B) {
	for i := 0; i < b.N; i++ {
		idgen.Generate()
	}
}

func benchmarkIdGenerator_Current(idgen idGenerator, b *testing.B) {
	for i := 0; i < b.N; i++ {
		idgen.Current()
	}
}

func BenchmarkIdGeneratorAtomic_Current(b *testing.B) {
	benchmarkIdGenerator_Current(newIdGeneratorAtomic(), b)
}

func BenchmarkIdGeneratorAtomic_Generate(b *testing.B) {
	benchmarkIdGenerator_Generate(newIdGeneratorAtomic(), b)
}

func BenchmarkIdGeneratorMutex_Current(b *testing.B) {
	benchmarkIdGenerator_Current(newIdGeneratorMutex(), b)
}

func BenchmarkIdGeneratorMutex_Generate(b *testing.B) {
	benchmarkIdGenerator_Generate(newIdGeneratorMutex(), b)
}

func BenchmarkIdGeneratorChan_Current(b *testing.B) {
	benchmarkIdGenerator_Current(newIdGeneratorChan(), b)
}

func BenchmarkIdGeneratorChan_Generate(b *testing.B) {
	benchmarkIdGenerator_Generate(newIdGeneratorChan(), b)
}
