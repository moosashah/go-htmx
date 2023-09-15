package database

import "fmt"

type Todo struct {
	Content   string
	Completed bool
	Id        int
}

type ErrorMap = map[string]string

func (t *Todo) validate() ErrorMap {
	var errors ErrorMap = make(ErrorMap)

	if t.Content == "" {
		errors["content"] = "Task required"
	}

	return errors
}

func DeleteTodo(id int) error {
	_, err := Db.Exec("DELETE FROM todos WHERE Id = $1", id)
	if err != nil {
		return fmt.Errorf("unable to detele contact: %+v", err)
	}
	return nil
}

func GetTodos() ([]Todo, error) {
	res, err := Db.Query("select * from todos")
	if err != nil {
		return nil, fmt.Errorf("unable to query db: %+v", err)
	}

	defer res.Close()

	var todos []Todo = make([]Todo, 0)
	for res.Next() {
		i := Todo{}
		err := res.Scan(&i.Content, &i.Completed, &i.Id)
		if err != nil {
			return nil, fmt.Errorf("unable to scan db row: %+v", err)
		}

		todos = append(todos, i)
	}
	return todos, nil
}

func (t *Todo) Save() (ErrorMap, error) {
	errors := t.validate()
	if len(errors) > 0 {
		return errors, nil
	}
	var err error
	if t.Id == -1 {
		_, err = Db.Exec(`INSERT INTO todos (content,completed) VALUES (?,?)`, t.Content, t.Completed)
	} else {
		_, err = Db.Exec(`UPDATE TODOS SET content = ?, completed = ?`, t.Content, t.Completed, t.Id)
	}
	return errors, err
}
