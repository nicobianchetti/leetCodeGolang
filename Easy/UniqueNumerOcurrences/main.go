package main

func uniqueOccurrences(arr []int) bool {
	mapHelper := make(map[int]int)

	for _, v := range arr {
		//Verifico si el valor existe en el mapa , sino , lo creo
		_, found := mapHelper[v]
		if found {
			mapHelper[v] = mapHelper[v] + 1 //si ya existe le sumo 1
		} else {
			mapHelper[v] = 1 // Creo en el mapa una key nueva (el valor que estoy recorriendo del array) y lo inicialo con 1 xq es la primera vez que lo encuentro
		}
	}

	// fmt.Print(mapHelper)

	var arrayHelper []int // creo nuevo array vacio para ir adjuntandole valores bajo demanda segun los valores que tenga en el hashmap. Es dinamico

	for i := range mapHelper {
		arrayHelper = append(arrayHelper, mapHelper[i]) // a append le digo: appendeama en arrayHelper , el valor que tenga mapHelper en el indice i que es el que estoy recorriendo
	}

	// fmt.Print(arrayHelper)

	var auxBool bool = true //true es porque comienzo considerando que todos los valores son únicos

	for i := 0; i < len(arrayHelper); i++ {
		valor := arrayHelper[i] // tomo el valor referencia sobre el cual me paro y ahora voy a recorrer de aca en adelante para ver si este valor se repite
		for j := (i + 1); j < len(arrayHelper); j++ {
			if valor == arrayHelper[j] {
				auxBool = false // defino como false porque ya lo encontre repido al menos una vez
				break
			}
		}
		if auxBool == false {
			break
		}
	}

	return auxBool
}
