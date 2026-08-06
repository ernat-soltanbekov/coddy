package main

import "fmt"

func calculateCalories(activity string, duration int, intensity float64) float64 {
    var baseCalories float64
    switch activity {
        case "running": baseCalories = 10
        case "swimming": baseCalories = 8
        case "cycling": baseCalories = 7
        default: baseCalories = 5
    }
    timeDuration := float64(duration)
    return baseCalories * timeDuration * intensity
}

func main() {
    fmt.Println("Running for 30 minutes at intensity 1.2:", calculateCalories("running", 30, 1.2))
    fmt.Println("Swimming for 45 minutes at intensity 1.0:", calculateCalories("swimming", 45, 1.0))
    fmt.Println("Cycling for 60 minutes at intensity 1.5:", calculateCalories("cycling", 60, 1.5))
    fmt.Println("Yoga for 60 minutes at intensity 0.8:", calculateCalories("yoga", 60, 0.8))
}
