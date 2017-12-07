package Objective

import a "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Objective/Ackley"
import b "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Objective/Bohachevsky"
import s "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Objective/Schwefel"

//Objective Interface for universal use of Objectives
type Objective interface {
	EvaluateFitness([]float64) float64
	Min() float64
	Max() float64
}

//GetObjective Generates Objective By Name
func GetObjective(name string) Objective {
	switch name {
	case "Ackley":
		return a.Ackley{}
	case "Schwefel":
		return s.Schwefel{}
	default:
		return b.Bohachevsky{}

	}
}
