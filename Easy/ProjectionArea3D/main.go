package main

/*
https://leetcode.com/problems/projection-area-of-3d-shapes/

Runtime: 4 ms, faster than 100.00% .
Memory Usage: 3.7 MB, less than 100.00% .

Descrition: On a N * N grid, we place some 1 * 1 * 1 cubes that are axis-aligned with the x, y, and z axes.

Each value v = grid[i][j] represents a tower of v cubes placed on top of grid cell (i, j).

Now we view the projection of these cubes onto the xy, yz, and zx planes.

A projection is like a shadow, that maps our 3 dimensional figure to a 2 dimensional plane.

Here, we are viewing the "shadow" when looking at the cubes from the top, the front, and the side.

Return the total area of all three projections.

Description solution:

Obtengo por separadas las tres proyecciones : XY, YZ, YX
a) XY es la vista de arriba,dado v= grid[i][j], es la suma de cantidad de elementos [j] distintos de 0 (0 representaria un cubo sin volumen)
b) YZ es igual a sumatoria de los v = [j] máximos de cada [i]
c) XZ es igual a la sumatoria de los valores máximos para cada índice j en cada [i]

*/

func projectionArea(grid [][]int) int {
	var axisXY, axisXZ, axisYZ int
	var mapXz = make(map[int]int)

	for _, i := range grid {
		var axisYZAux int
		for index, j := range i {
			if j > axisYZAux {
				axisYZAux = j
			}
			if j != 0 {
				axisXY++
			}
			if mapXz[index] < j {
				mapXz[index] = j
			}
		}
		axisYZ += axisYZAux
	}

	for _, value := range mapXz {
		axisXZ += value
	}

	return (axisYZ + axisXY + axisXZ)
}

/*
Error, no se toman los i valores mas altos, sino los valores mas altos para c/Indice de i
func getAxisXZ(arrayXZ []int, length int) int {
	sort.Sort(sort.Reverse((sort.IntSlice(arrayXZ))))
	ArrayAux := arrayXZ[:length]

	var SumArray int

	for _, i := range ArrayAux {
		SumArray += i
	}
	return SumArray
}
*/

/*
intSlice := []int{3,5,6,4,2,293,-34}
sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
*/
