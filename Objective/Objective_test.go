package Objective

import "testing"

func TestObjective(t *testing.T) {
	testAckley := GetObjective("Ackley")
	testAckley.Min()
	testAckley.Max()
	testAckley.EvaluateFitness([]float64{1.0, 2.0, 3.0})

	testSchwefel := GetObjective("Schwefel")
	testSchwefel.Min()
	testSchwefel.Max()
	testSchwefel.EvaluateFitness([]float64{1.0, 2.0, 3.0})

	testBohachevsky := GetObjective("Bohachevksy")
	testBohachevsky.Min()
	testBohachevsky.Max()
	testBohachevsky.EvaluateFitness([]float64{1.0, 2.0, 3.0})
}
