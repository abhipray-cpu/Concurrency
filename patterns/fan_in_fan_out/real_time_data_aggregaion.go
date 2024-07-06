package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Data struct {
	SourceId int
	Value    int
}

// simulateDataSource simulates real-time data generation from a source
func simulateDataResource(sourceId int, dataStream chan<- Data) {
	for {
		// simulate data generation
		data := Data{
			SourceId: sourceId,
			Value:    rand.Intn(100),
		}

		// send data to channel
		dataStream <- data

		// wait for a bit before generating more data
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

// fanIn aggregates multiple data streams into a single stream

func fanIn(dataStreams ...<-chan Data) <-chan Data {
	var wg sync.WaitGroup
	aggregatedSum := make(chan Data)

	ouput := func(ds <-chan Data) {
		for data := range ds {
			aggregatedSum <- data
		}
		wg.Done()
	}
	wg.Add(len(dataStreams))
	for _, ds := range dataStreams {
		go ouput(ds)
	}

	// close the aggregatedSum channel once all the goroutines have finished

	go func() {
		wg.Wait()
		close(aggregatedSum)
	}()
	return aggregatedSum
}

func AggregateSum() {
	// Number of data sources
	numSources := 3
	dataStreams := make([]<-chan Data, numSources)

	// Create and start data source simulations
	for i := 0; i < numSources; i++ {
		ds := make(chan Data)
		dataStreams[i] = ds
		go simulateDataResource(i+1, ds)
	}

	// Aggregate data from all sources
	aggregatedStream := fanIn(dataStreams...)

	// Process the aggregated data stream
	for data := range aggregatedStream {
		fmt.Printf("Received data: SourceID=%d, Value=%d\n", data.SourceId, data.Value)
	}
}
