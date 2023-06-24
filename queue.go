package main

type (
	Queue struct {
		persons []Person
	}

	Person struct {
		name string
		age  int
	}
)

func (q *Queue) Enqueue(name string, age int) {
	person := &Person{name, age}
	q.persons = append(q.persons, *person)
}

func (q *Queue) Dequeue() Person {
	if len(q.persons) == 0 {
		return q.persons[0]
	}
	person := q.persons[0]
	q.persons = q.persons[1:]
	return person
}
