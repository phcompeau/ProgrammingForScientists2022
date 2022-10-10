package main

import "strconv"

//UPGMA takes a distance matrix and a collection of species names as input.
//It returns a Tree (an array of nodes) resulting from applying
//UPGMA to this dataset.
func UPGMA(mtx DistanceMatrix, speciesNames []string) Tree {
	if len(mtx) != len(speciesNames) {
		panic("Different sized matrix than number of species names.")
	}
	t := InitializeTree(speciesNames)
	clusters := t.InitializeClusters()
	numLeaves := len(speciesNames)

	//range over all internal nodes of the tree and set appropriate ages and connect edges and update clusters and update matrix
	for k := numLeaves; k < 2*numLeaves-1; k++ {
		row, col, val := FindMinElement(mtx)
		t[k].Age = val / 2.0 // set age
		//now, we need to connect edges
		t[k].Child1 = clusters[row]
		t[k].Child2 = clusters[col]

		//now we update matrix
		mtx = AddRowCol(mtx, clusters, row, col)
		mtx = DeleteRowCol(mtx, row, col)

		//now we update clusters by adding a pointer to the new node
		clusters = append(clusters, t[k])
		clusters = DeleteClusters(clusters, row, col)

	}

	return t
}

/*
AddRowCol(D, clusters, row, col)
	newRow  decimal array of length NumCols(D) + 1
	for every integer r between 0 and |newRow| – 2
		size1  CountLeaves(clusters[row])
		size2  CountLeaves(clusters[col])
		newRow[r]  (size1 * D[row][r] + size2 * D[r][col] ) 				/ (size1 + size2)
	D  append(D, newRow)
	for every integer c between 0 and |newRow| – 2
		D[c]  append(D[c], newRow[c])
	return D

*/

func AddRowCol(mtx DistanceMatrix, clusters []*Node, row, col int) DistanceMatrix {
	numCols := len(mtx) + 1
	newRow := make([]float64, numCols)
	for c := 0; c < len(newRow)-1; c++ {
		if c != row && c != col {
			sizeCluster1 := CountLeaves(clusters[row])
			sizeCluster2 := CountLeaves(clusters[col])
			//now set new distance value equal to weighted average of distance
			//between c-th element and the two clusters at clusters[row]
			//and clusters[col]
			newRow[c] = (float64(sizeCluster1)*mtx[row][c] + float64(sizeCluster2)*mtx[col][c]) / float64(sizeCluster1+sizeCluster2)
		}
	}

	//append our new row to mtx!
	mtx = append(mtx, newRow)

	//our matrix has n+1 rows, the first n having n cols and the last one having n+1 cols
	//use symmetry of distance matrix to fill in missing values in final column
	for i := 0; i < len(mtx)-1; i++ {
		mtx[i] = append(mtx[i], newRow[i])
	}

	return mtx
}

func DeleteRowCol(mtx DistanceMatrix, row, col int) DistanceMatrix {
	//col > row by assumption
	//let's delete rows first
	mtx = append(mtx[:col], mtx[col+1:]...)
	mtx = append(mtx[:row], mtx[row+1:]...)

	//so now I have an (n-2) x n matrix and I need to delete columns at the given indices
	//for every row that is left, delete the (row)-th and (col)-th element
	for i := range mtx {
		mtx[i] = append(mtx[i][:col], mtx[i][col+1:]...)
		mtx[i] = append(mtx[i][:row], mtx[i][row+1:]...)
	}

	return mtx
}

func DeleteClusters(clusters []*Node, row, col int) []*Node {
	//delete clusters at indices row and col
	//col > row by assumption
	clusters = append(clusters[:col], clusters[col+1:]...)
	clusters = append(clusters[:row], clusters[row+1:]...)

	return clusters
}

//BIG NOTE: col will always be > row
func FindMinElement(mtx DistanceMatrix) (int, int, float64) {
	//range over values of matrix and find smallest one
	if len(mtx) <= 1 || len(mtx[0]) <= 1 {
		panic("bad")
	}
	row := 0
	col := 1
	minVal := mtx[row][col]

	for i := 0; i < len(mtx)-1; i++ { //row indexing
		for j := i + 1; j < len(mtx[i]); j++ {
			// is it better?
			if mtx[i][j] < mtx[row][col] {
				//update!
				row = i
				col = j
				minVal = mtx[i][j]
			}
		}
	}

	return row, col, minVal
}

func (t Tree) InitializeClusters() []*Node {
	//slice of nodes point to leaves in order
	//how many nodes are there then?
	totalNodes := len(t)
	numLeaves := (totalNodes + 1) / 2

	clusters := make([]*Node, numLeaves)
	for i := range clusters {
		clusters[i] = t[i]
	}

	return clusters
}

func InitializeTree(speciesNames []string) Tree {
	numLeaves := len(speciesNames)

	var t Tree // slice of node pointers
	//make the slice
	t = make([]*Node, 2*numLeaves-1)

	//how many total nodes do I need to create? 2n-1 total
	for i := 0; i < 2*numLeaves-1; i++ {
		//create a node
		var vx Node
		//set its fields
		vx.Num = i
		//don't need to set children or ages
		//do need to set label
		//leaves get labels from speciesNames
		if i < numLeaves { //leaf!
			vx.Label = speciesNames[i]
		} else {
			vx.Label = "Ancestor Species: " + strconv.Itoa(i)
		}

		//add this node into tree ...
		t[i] = &vx // i-th node pointer points to reference of current node

	}

	return t

}

func CountLeaves(v *Node) int {
	//base case: both children are nil
	if v.Child1 == nil && v.Child2 == nil {
		return 1
	}
	//if we make it here, we know that one child is not nil (but not necessarily both)
	//what if only one child is nil?
	if v.Child1 == nil {
		return CountLeaves(v.Child2)
	}
	if v.Child2 == nil {
		return CountLeaves(v.Child1)
	}
	// If I make it here, both are not nil (i.e., two children)
	return CountLeaves(v.Child1) + CountLeaves(v.Child2)

}
