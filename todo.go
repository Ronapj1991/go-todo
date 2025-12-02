package main

type Todo struct {
	ID          int
	Description string
	Completed   bool
}

func (t *Todo) MarkCompleted() {
	t.Completed = true
}

func (t *Todo) SetDescription(desc string) {
	t.Description = desc
}
