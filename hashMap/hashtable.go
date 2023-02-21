package hashMap

import (
	"fmt"
	"hash/fnv"

	"github.com/nikhilnarayanan623/ds-hashTable/hashMap/interfaces"
)

type node struct {
	key   string
	value string
	next  *node
}

type hashMap struct {
	array       []*node // to store key value pair
	arrayLength int
	nodeCount   int
	threshold   int
}

// to get a new instance of hashMap
func GetHashMap() interfaces.HashMapInterface {
	return &hashMap{array: make([]*node, 5), arrayLength: 5, threshold: 0}
}

func (h *hashMap) Put(key string, value string) {
	fmt.Println("map lenght ", h.arrayLength)
	//check the hashTable need to create a new array and re hash
	h.checkReHash()

	// convert key into its correspoding index
	index := h.hashTheKey(key)

	// if there is no node the assign first node
	if h.array[index] == nil {
		h.array[index] = &node{key: key, value: value}
		h.nodeCount++
		return
	}

	// loop throgh the linked list and find the node if it present then change its value
	currentNode := h.array[index]

	for currentNode != nil {

		if currentNode.key == key { // if the key found at any node then replace its value
			currentNode.value = value
			return
		} else if currentNode.next == nil { // not found the key until at the last positoin then create a new node at the end
			currentNode.next = &node{key: key, value: value}
		}
		currentNode = currentNode.next
	}
}

func (h *hashMap) Get(key string) (string, bool) {

	// get the corresponding index of key
	index := h.hashTheKey(key)

	//loop throgh the linked list if its formed a linked list and
	currentNode := h.array[index]
	for currentNode != nil {
		if currentNode.key == key { // key found in list then return value with true
			return currentNode.value, true
		}
	}

	// if key not found then return empty string with false
	return "", false
}

func (h *hashMap) Contains(key string) bool {

	// get the corresponding index of key
	index := h.hashTheKey(key)

	//loop throgh the linked list if its formed a linked list and
	currentNode := h.array[index]
	for currentNode != nil {
		if currentNode.key == key { // key found in list then return value with true
			return true
		}
		currentNode = currentNode.next
	}

	// if key not found then return empty string with false
	return false

}

func (h *hashMap) Display() {

	fmt.Print("\nValues in the HashMap\n\n")
	for _, val := range h.array { //range through the array
		if val != nil { // if index have node present
			for val != nil { //loop throght the linked and print values if its form a linked list
				fmt.Printf("key: %s , value: %s \n", val.key, val.value)
				val = val.next
			}
		}
	}
	fmt.Println()
}

// function to hash a string to integer and return index according to array length
func (h *hashMap) hashTheKey(key string) int {

	// create an instance of hash
	hashObj := fnv.New64()
	//write string to hash
	hashObj.Write([]byte(key))

	//get int value of hash from obj
	hashValue := hashObj.Sum64()

	// calculate the index according to the array length
	index := hashValue % uint64(h.arrayLength)

	return int(index) //return the uint index as int
}

// function to check the hashMap array need to extend
func (h *hashMap) checkReHash() {

	// if nodes count / array lenght is not over than threshold then simply return
	if h.nodeCount/h.arrayLength <= h.threshold {
		return
	} // node count to array lenght is over threshold value create a new array and re hash all valued

	//if nodes count / array lenght is over threshold then create a new array and re hash
	oldArray := h.array // hold old array

	//set new array on hashMap
	newsize := h.arrayLength * 2
	h.array = make([]*node, newsize)

	//reset the count and assign new array size
	h.nodeCount = 0
	h.arrayLength = newsize

	//loop throgh old array and get nodes and set it on new array on the hashMap
	for _, node := range oldArray {
		if node != nil { // if node is not nil then set it
			//loop through the linked list
			for node != nil {
				h.Put(node.key, node.value)
				node = node.next
			}
		}
	}
}
