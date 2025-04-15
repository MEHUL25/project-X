package leetcode

import (
	"fmt"
	"sync"
)

// channel-based solution for the H₂O problem in Go, where we use channels to coordinate hydrogen and oxygen threads so that exactly 2 hydrogen and 1 oxygen print together per water molecule.
// You're simulating the creation of water molecules, where:

// Each water molecule needs 2 Hydrogen (H) and 1 Oxygen (O).
// Hydrogen threads call releaseHydrogen() to print "H".
// Oxygen threads call releaseOxygen() to print "O".
// The output must group HHO together for every molecule.
// No thread can move forward until a valid molecule (2H + 1O) can be formed.

type H2O struct {
	hChan chan struct{}
	oChan chan struct{}
	bond  chan struct{} // for barrier synchronization
	done  chan struct{} // after printing
}

func NewH2O() *H2O {
	return &H2O{
		hChan: make(chan struct{}, 2), // buffer for 2 hydrogens
		oChan: make(chan struct{}, 1), // buffer for 1 oxygen
		bond:  make(chan struct{}, 3), // wait until 3 have printed
		done:  make(chan struct{}, 3), // used to signal completion of one molecule
	}
}

func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	h2o.hChan <- struct{}{} // signal hydrogen is ready
	<-h2o.bond              // wait until it’s this hydrogen’s turn to print
	releaseHydrogen()
	h2o.done <- struct{}{}
}

func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.oChan <- struct{}{} // signal oxygen is ready

	// wait until 2 H and 1 O are ready
	<-h2o.hChan
	<-h2o.hChan
	<-h2o.oChan

	// allow 3 threads to print (2 H + 1 O)
	for i := 0; i < 2; i++ {
		h2o.bond <- struct{}{}
	}
	releaseOxygen()
	h2o.done <- struct{}{}

	// wait until all 3 are done printing before allowing next group
	for i := 0; i < 3; i++ {
		<-h2o.done
	}
}
func solution2() {
	input := "OOHHHH" // Must contain twice as many H as O
	h2o := NewH2O()
	var wg sync.WaitGroup

	printHydrogen := func() {
		fmt.Print("H")
	}
	printOxygen := func() {
		fmt.Print("O")
	}

	for _, atom := range input {
		wg.Add(1)
		if atom == 'H' {
			go func() {
				defer wg.Done()
				h2o.hydrogen(printHydrogen)
			}()
		} else {
			go func() {
				defer wg.Done()
				h2o.oxygen(printOxygen)
			}()
		}
	}

	wg.Wait()
	fmt.Println()
}
