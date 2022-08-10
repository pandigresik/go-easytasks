package models

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Pic   string `json:"pic"`
	Deadline   string `json:"deadline"`
	Status int    `json:"status"`
}

func GetTasks() []Task {
	sql := "SELECT id, name, pic, deadline, status FROM tasks"
	rows, err := dbCon.Query(sql)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := []Task{}	
	for rows.Next() {
		var status, id int
		var name, pic, deadline string
		task := Task{}
		err2 := rows.Scan(&id, &name, &pic, &deadline, &status)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}	
		task.ID = id
		task.Name = name
		task.Pic = pic
		task.Deadline = deadline
		task.Status = status
		
		result = append(result, task)
	}
	return result
}

func GetTask(id int) Task {
	sql := "SELECT id, name, pic, deadline, status FROM tasks where id = ?"
	rows, err := dbCon.Query(sql, id)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := Task{}	
	for rows.Next() {
		var status, id int
		var name, pic, deadline string
		task := Task{}
		err2 := rows.Scan(&id, &name, &pic, &deadline, &status)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}	
		task.ID = id
		task.Name = name
		task.Pic = pic
		task.Deadline = deadline
		task.Status = status
		
		result = task
	}
	return result
}

func PutTask(name string, pic string, deadline string, status int) (int64, error) {
	sql := "INSERT INTO tasks(name, pic, deadline, status) VALUES(?,?,?,?)"

	// Create a prepared SQL statement
	stmt, err := dbCon.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(name, pic, deadline, status)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func EditTask(taskId int, name string, pic string, deadline string, status int) (int64, error) {
	sql := "UPDATE tasks set name = ?, pic = ?, deadline = ?, status = ? WHERE id = ?"

	stmt, err := dbCon.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(name, pic, deadline, status, taskId)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

func DeleteTask(id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	// Create a prepared SQL statement
	stmt, err := dbCon.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}

	// Replace the '?' in our prepared statement with 'id'
	result, err2 := stmt.Exec(id)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
