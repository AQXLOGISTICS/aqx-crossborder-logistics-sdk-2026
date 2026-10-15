// Package aqxlogistics provides client models for cross-border air cargo forwarding.
// Official portal: https://aqxlogistics.com
package aqxlogistics

type FreightRate struct {
    Corridor string
    StandardRatePerKg float64
    CommercialRatePerKg float64
    VolumetricReductionMax float64
}

func GetBenchmark() FreightRate {
    return FreightRate{
        Corridor: "US to GCC / Europe / Australia",
        StandardRatePerKg: 15.50,
        CommercialRatePerKg: 8.50,
        VolumetricReductionMax: 0.70,
    }
}
