package garbage

import (
	"fmt"
)

// Define object colors
const (
	White = iota // Unreachable (Garbage)
	Gray         // In Progress (Scanned but Not Fully Processed)
	Black        // Reachable (Live Object)
)

// Object represents a node in memory
type Object struct {
	id       int
	color    int
	children []*Object
}

// Mark and Sweep Garbage Collector
type GC struct {
	objects []*Object // List of all objects in memory
	roots   []*Object // List of root objects
}

// Mark Phase: Implements the tri-color marking algorithm
func (gc *GC) mark() {
	// Step 1: Move all root objects to gray
	grayQueue := []*Object{}
	for _, root := range gc.roots {
		root.color = Gray
		grayQueue = append(grayQueue, root)
	}

	// Step 2: Process each gray object
	for len(grayQueue) > 0 {
		obj := grayQueue[0]
		grayQueue = grayQueue[1:]

		// Scan and move its children to gray if they are white
		for _, child := range obj.children {
			if child.color == White {
				child.color = Gray
				grayQueue = append(grayQueue, child)
			}
		}

		// Once scanned, move the object to black
		obj.color = Black
	}
}

// Sweep Phase: Removes white objects (garbage)
func (gc *GC) sweep() {
	newObjects := []*Object{}
	for _, obj := range gc.objects {
		if obj.color == White {
			// Object is unreachable and will be collected
			fmt.Println("Collecting garbage:", obj.id)
		} else {
			// Retain reachable objects
			obj.color = White // Reset for next cycle
			newObjects = append(newObjects, obj)
		}
	}
	gc.objects = newObjects
}

// Run the Garbage Collector
func (gc *GC) runGC() {
	fmt.Println("Starting Mark Phase...")
	gc.mark()
	fmt.Println("Starting Sweep Phase...")
	gc.sweep()
}

// Create an object and track it
func (gc *GC) createObject(id int) *Object {
	obj := &Object{id: id, color: White}
	gc.objects = append(gc.objects, obj)
	return obj
}

// Set root objects
func (gc *GC) addRoot(obj *Object) {
	gc.roots = append(gc.roots, obj)
}

func tricolorAlgorithm() {
	gc := &GC{}

	// Create objects
	obj1 := gc.createObject(1)
	obj2 := gc.createObject(2)
	obj3 := gc.createObject(3)
	obj4 := gc.createObject(4)

	// Set relationships (object graph)
	obj1.children = []*Object{obj2, obj3}
	obj3.children = []*Object{obj4}

	// Add root objects
	gc.addRoot(obj1)

	// Run garbage collection
	gc.runGC()
}
