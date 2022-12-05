package main

type Caretaker struct {
	memetoArray []*Memento
}

func (c *Caretaker) addMemento(m *Memento) {
	c.memetoArray = append(c.memetoArray, m)
}

func (c *Caretaker) getMemento(index int) *Memento {
	return c.memetoArray[index]
}
