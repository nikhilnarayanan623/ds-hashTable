package interfaces

type HashMapInterface interface {
	//public
	Put(key string, value string)  // to insert a new value to hashMap
	Get(key string) (string, bool) // to get value with its status is found or not
	Contains(key string) bool      // same as Get but only return its the key contain or not as boo
	Display()                      // dispalay all values in the hashMap
}
