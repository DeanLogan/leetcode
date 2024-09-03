package main

import (
	"fmt"
)

func main(){
	// fmt.Println(minNumberOfHours(5, 3, []int{1,4,3,2}, []int{2,6,3,1}))
	// fmt.Println(minNumberOfHours(2, 4, []int{1}, []int{3}))
	fmt.Println(minNumberOfHours(1, 1, []int{1,1,1,1}, []int{1,1,1,50}))
}

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
    trainingHours := 0
    for i := 0; i < len(energy); i++ {
        if initialEnergy <= energy[i] {
            trainingHours += energy[i] - initialEnergy + 1
            initialEnergy = energy[i] + 1
        }
        if initialExperience <= experience[i] {
            trainingHours += experience[i] - initialExperience + 1
            initialExperience = experience[i] + 1
        }
        initialEnergy -= energy[i]
        initialExperience += experience[i]
    }
    return trainingHours
}