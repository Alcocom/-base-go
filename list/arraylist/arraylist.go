package arraylist

type ArrayList[T any] struct {
	slice []T
}

func New[T any](data ...T) *ArrayList[T] {
	arr := make([]T, 10)

	if data != nil {
		arr = append(arr, data...)
	}

	return &ArrayList[T]{
		slice: arr,
	}
}

func (al *ArrayList[T]) Add(data ...T) {
	al.slice = append(al.slice, data...)
}

// TODO
/*
[ Add() ] Java ArrayList add()
inserts the element to the arraylist
[ Add() ] Java ArrayList addAll()
adds all elements of a collection to arraylist

Java ArrayList clear()
removes all the elements from arraylist

Java ArrayList clone()
makes a copy of the array list

Java ArrayList contains()
checks if the element is present in the arraylist

Java ArrayList get()
returns the element present in the specified index

Java ArrayList indexOf()
returns the position of the specified element

Java ArrayList removeAll()
removes multiple elements from the arraylist

Java ArrayList remove()
removes the single element from the arraylist

Java ArrayList size()
returns the length of an arraylist

Java ArrayList isEmpty()
checks if the arraylist is empty

Java ArrayList subList()
returns a portion of the arraylist

Java ArrayList set()
replace the single element from an arraylist

Java ArrayList sort()
sorts the arraylist according to specified order

Java ArrayList toArray()
converts an arraylist to an array

Java ArrayList toString()
converts the arraylist into a String

Java ArrayList ensureCapacity()
set the size of an araylist

Java ArrayList lastIndexOf()
returns position of last occurrence of the element

Java ArrayList retainAll()
retains only the common elements

Java ArrayList containsAll()
checks if a collection is a subset of arraylist

Java ArrayList trimToSize()
trims the capacity of arraylist equal to the size

Java ArrayList removeRange()
removes a portion of the arraylist

Java ArrayList replaceAll()
replace all elements from the arraylist

Java ArrayList removeIf()
removes element that satisfy the condition

Java ArrayList forEach()
performs an action to all elements of arraylist

Java ArrayList iterator()
returns an iterate to loop through the ArrayList
*/
