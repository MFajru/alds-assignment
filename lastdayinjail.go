package main

type Person struct {
	Name          string
	CriminalScore int
}

func mergeSort(criminals []Person) []Person {
	if len(criminals) < 2 {
		return criminals
	}

	middle := len(criminals) / 2
	left := mergeSort(criminals[:middle])
	right := mergeSort(criminals[middle:])

	return merge(left, right)
}

func merge(left, right []Person) []Person {
	result := make([]Person, len(left)+len(right))
	lIndex, rIndex := 0, 0

	for i := 0; i < len(result); i++ {
		if lIndex >= len(left) {
			result[i] = right[rIndex]
			rIndex++
		} else if rIndex >= len(right) {
			result[i] = left[lIndex]
			lIndex++
		} else if left[lIndex].CriminalScore < right[rIndex].CriminalScore {
			result[i] = left[lIndex]
			lIndex++
		} else if left[lIndex].CriminalScore > right[rIndex].CriminalScore {
			result[i] = right[rIndex]
			rIndex++
		} else if left[lIndex].CriminalScore == right[rIndex].CriminalScore && left[lIndex].Name < right[rIndex].Name {
			result[i] = left[lIndex]
			lIndex++
		} else {
			result[i] = right[rIndex]
			rIndex++
		}
	}
	return result
}

func findChosenPerson(criminals []Person, chosenPerson string) Person {
	var chosen Person
	for _, person := range criminals {
		if chosenPerson == person.Name {
			chosen = person
			break
		}
	}
	return chosen
}

// Task 1.a
func LastDayInJail(criminals []Person, chosenPerson string) (onTransport []Person, waiting []Person) {
	sorted := mergeSort(criminals)
	var releasedCriminals []Person

	if len(criminals) < 5 {
		releasedCriminals = sorted
		return releasedCriminals, nil
	} else {
		releasedCriminals = sorted[:5]
	}

	chosen := findChosenPerson(criminals, chosenPerson)

	if chosenPerson != "" {
		releasedCriminals = append(releasedCriminals, chosen)
	}

	return releasedCriminals[:3], releasedCriminals[3:]
}
