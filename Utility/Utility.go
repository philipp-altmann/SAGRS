package Utility

import (
	"encoding/csv"
	"fmt"
	"os"

	a "github.com/philipp-altmann/SAGRS/Approximation"
	e "github.com/philipp-altmann/SAGRS/Environment"
	o "github.com/philipp-altmann/SAGRS/Objective"
	p "github.com/philipp-altmann/SAGRS/Population"
)

var objective o.Objective
var converging int
var convergenceCycle int
var writer *csv.Writer
var file *os.File

//Setup initialises CSV file
func Setup(e e.Environment) {

	//Setup Convergence Meassure
	converging = 0
	convergenceCycle = e.Cycles

	if e.WriteProgress {
		objective = o.GetObjective(e.Objective)
		file, _ = os.Create(e.ProgressFileName + ".csv")
		writer = csv.NewWriter(file)
		writer.Write([]string{"Cycle", "Best in Pool", "Approximation Success", "Approximation Error"})
		writer.Flush()
	}
}

//WriteProgress writes progress to csv file
func WriteProgress(e e.Environment, cycle int, evaluationPool, suggestions p.Population, approximator a.Approximator) {
	if !e.Verbose && !e.WriteProgress {
		return
	}
	better, worse, _ := computeSuccess(suggestions, evaluationPool)
	bestFitness := evaluationPool[0].Fitness
	updateConvergence(e, worse, cycle, evaluationPool)
	if e.WriteProgress {
		writeCSV(e, cycle, worse, bestFitness, suggestions, approximator)
	}

	if e.Verbose {
		writeConsole(e, cycle, better, worse, evaluationPool, approximator)
	}
}

func computeSuccess(suggestions, evaluationPool p.Population) (better, worse int, bestFitness float64) {
	if len(suggestions) == 0 {
		return 0, 0, evaluationPool[0].Fitness
	}
	bestFitness = evaluationPool[0].Fitness
	worstFitness := evaluationPool[len(evaluationPool)-1].Fitness
	better = 0
	worse = 0
	for _, individual := range suggestions {
		fmt.Println(objective.EvaluateFitness(individual.Value))
		if objective.EvaluateFitness(individual.Value) > worstFitness {
			worse++
		}
		if objective.EvaluateFitness(individual.Value) < bestFitness {
			better++
		}
	}
	if better == 0 {
		for _, individual := range suggestions {
			if objective.EvaluateFitness(individual.Value) > bestFitness {
				bestFitness = objective.EvaluateFitness(individual.Value)
			}
		}
	}
	return
}

func updateConvergence(e e.Environment, worse, cycle int, evaluationPool p.Population) {
	if worse == e.SuggestToEvaluation || evaluationPool[0].Fitness == 0.0 {
		if converging > 0 {
			converging++
		} else {
			converging = 1
			convergenceCycle = cycle
		}
	} else {
		converging = 0
		convergenceCycle = cycle
	}
}

//GetConvergenceCycle Getter for privately maintained Cycle of Convergence
func GetConvergenceCycle() int {
	file.Close()
	return convergenceCycle
}
