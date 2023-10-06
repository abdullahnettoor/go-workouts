//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

func (m *MemoryUsage) AverageMemoryUsage() float64 {
	var sum float64
	for _, v := range m.amount {
		sum += float64(v)
	}
	avg := sum / float64(len(m.amount))
	return avg
}

func (c *CpuTemp) AverageCpuTemp() float64 {
	var sum float64
	for _, v := range c.temp {
		sum += float64(v)
	}
	avg := sum / float64(len(c.temp))
	return avg
}

func (b *BandwidthUsage) AverageBandwidthUsage() float64 {
	var sum float64
	for _, v := range b.amount {
		sum += float64(v)
	}
	avg := sum / float64(len(b.amount))
	return avg
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dash := Dashboard{BandwidthUsage: bandwidth, CpuTemp: temp, MemoryUsage: memory}
	a, b, c := dash.AverageBandwidthUsage(), dash.AverageCpuTemp(), dash.AverageMemoryUsage()
	fmt.Printf("Avg Bandwidth: %v	Avg Cpu Temp: %v	Avg Memory: %v", a, b, c)
}
