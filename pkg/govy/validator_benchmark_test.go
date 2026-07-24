package govy

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkValidatorRemovePropertiesByID(b *testing.B) {
	runRemovePropertiesByIDBenchmark(b, removePropertiesByIDBenchmark{
		name:          "no_ids/props=32/ids=0/retained=32",
		propertyCount: 32,
		retainedCount: 32,
	})
	runRemovePropertiesByIDBenchmark(b, removePropertiesByIDBenchmark{
		name:          "empty_validator/props=0/ids=1/retained=0",
		propertyCount: 0,
		retainedCount: 0,
		missIDCount:   1,
	})
	runRemovePropertiesByIDBenchmark(b, removePropertiesByIDBenchmark{
		name:          "single_miss/props=32/ids=1/retained=32",
		propertyCount: 32,
		retainedCount: 32,
		missIDCount:   1,
	})
	runRemovePropertiesByIDBenchmark(b, removePropertiesByIDBenchmark{
		name:          "single_hit_first/props=32/ids=1/retained=31",
		propertyCount: 32,
		retainedCount: 31,
		hitIndexes:    []int{0},
	})
	runRemovePropertiesByIDBenchmark(b, removePropertiesByIDBenchmark{
		name:          "single_hit_middle/props=32/ids=1/retained=31",
		propertyCount: 32,
		retainedCount: 31,
		hitIndexes:    []int{16},
	})
	runRemovePropertiesByIDBenchmark(b, removePropertiesByIDBenchmark{
		name:          "single_hit_last/props=32/ids=1/retained=31",
		propertyCount: 32,
		retainedCount: 31,
		hitIndexes:    []int{31},
	})
	runRemovePropertiesByIDBenchmark(b, removePropertiesByIDBenchmark{
		name:          "duplicate_hit/props=32/ids=2/retained=31",
		propertyCount: 32,
		retainedCount: 31,
		hitIndexes:    []int{16, 16},
	})
	runRemovePropertiesByIDBenchmark(b, removePropertiesByIDBenchmark{
		name:          "mixed/props=32/ids=4/retained=30",
		propertyCount: 32,
		retainedCount: 30,
		hitIndexes:    []int{8, 24},
		missIDCount:   2,
	})

	for _, propertyCount := range []int{8, 32, 128} {
		for _, idCount := range []int{2, 4, 8, 16, 32} {
			if idCount > propertyCount {
				continue
			}
			hitIndexes := make([]int, idCount)
			for i := range idCount {
				hitIndexes[i] = i * propertyCount / idCount
			}
			runRemovePropertiesByIDBenchmark(b, removePropertiesByIDBenchmark{
				name: fmt.Sprintf(
					"crossover_all_hits/props=%d/ids=%d/retained=%d",
					propertyCount,
					idCount,
					propertyCount-idCount,
				),
				propertyCount: propertyCount,
				retainedCount: propertyCount - idCount,
				hitIndexes:    hitIndexes,
			})
			runRemovePropertiesByIDBenchmark(b, removePropertiesByIDBenchmark{
				name: fmt.Sprintf(
					"crossover_all_misses/props=%d/ids=%d/retained=%d",
					propertyCount,
					idCount,
					propertyCount,
				),
				propertyCount: propertyCount,
				retainedCount: propertyCount,
				missIDCount:   idCount,
			})
		}
	}
}

type removePropertiesByIDBenchmark struct {
	name          string
	propertyCount int
	retainedCount int
	hitIndexes    []int
	missIDCount   int
}

func runRemovePropertiesByIDBenchmark(b *testing.B, benchmark removePropertiesByIDBenchmark) {
	b.Helper()
	b.Run(benchmark.name, func(b *testing.B) {
		validator, propertyIDs := newRemovePropertiesByIDBenchmarkValidator(benchmark.propertyCount)
		ids := make([]string, 0, len(benchmark.hitIndexes)+benchmark.missIDCount)
		for _, index := range benchmark.hitIndexes {
			ids = append(ids, propertyIDs[index])
		}
		for i := range benchmark.missIDCount {
			ids = append(ids, "missing-"+strconv.Itoa(i))
		}
		result := validator.RemovePropertiesByID(ids...)
		if len(result.props) != benchmark.retainedCount {
			b.Fatalf("retained %d properties, expected %d", len(result.props), benchmark.retainedCount)
		}

		for b.Loop() {
			result = validator.RemovePropertiesByID(ids...)
		}
		b.ReportMetric(float64(benchmark.propertyCount), "props/op")
		b.ReportMetric(float64(len(ids)), "ids/op")
		b.ReportMetric(float64(len(benchmark.hitIndexes)), "hit-ids/op")
		b.ReportMetric(float64(benchmark.missIDCount), "miss-ids/op")
		b.ReportMetric(float64(benchmark.propertyCount-benchmark.retainedCount), "removed/op")
		b.ReportMetric(float64(benchmark.retainedCount), "retained/op")
		benchmarkValidatorSink = result
	})
}

func newRemovePropertiesByIDBenchmarkValidator(
	propertyCount int,
) (validator Validator[struct{}], ids []string) {
	properties := make([]PropertyRulesInterface[struct{}], propertyCount)
	ids = make([]string, propertyCount)
	getter := func(struct{}) string { return "" }
	for i := range propertyCount {
		id := "property-" + strconv.Itoa(i)
		ids[i] = id
		properties[i] = For(getter).WithID(id)
	}
	return New(properties...), ids
}

var benchmarkValidatorSink Validator[struct{}]
