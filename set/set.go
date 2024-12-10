package set

// Set represents a collection of unique elements
type Set[T comparable] struct {
	elements map[T]struct{}
}

// NewSet creates and initializes a new set
//
// Returns:
//   - a new set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

// Add inserts an element into the set
//
// Parameters:
//   - value: the element to insert
func (s *Set[T]) Add(value T) {
	s.elements[value] = struct{}{}
}

// Remove deletes an element from the set
//
// Parameters:
//   - value: the element to delete
func (s *Set[T]) Remove(value T) {
	delete(s.elements, value)
}

// Contains checks if the set contains the specified element
//
// Parameters:
//   - value: the element to check
func (s *Set[T]) Contains(value T) bool {
	_, exists := s.elements[value]
	return exists
}

// Size returns the number of elements in the set
func (s *Set[T]) Size() int {
	return len(s.elements)
}

// Clear removes all elements from the set
//
// Parameters:
//   - s: the set
func (s *Set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

// Elements returns a slice of all elements in the set
//
// Returns:
//   - a slice of all elements in the set
func (s *Set[T]) Elements() []T {
	result := make([]T, 0, len(s.elements))
	for key := range s.elements {
		result = append(result, key)
	}
	return result
}
