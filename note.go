package main

type Note struct {
	ID       int
	Title    string
	Category string
}

func fetchNotes() ([]Note, error) {
	var notes []Note

	rows, err := DB.Query("select id, title, category from notes;")
	if err != nil {
		return []Note{}, err
	}
	defer rows.Close()

	for rows.Next() {
		note := Note{}
		err := rows.Scan(&note.ID, &note.Title, &note.Category)
		if err != nil {
			return []Note{}, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func fetchNote(ID int) (Note, error) {
	var note Note
	err := DB.QueryRow("select id, title, category from notes where id = (?)", ID).Scan(&note.ID, &note.Title, &note.Category)
	if err != nil {
		return Note{}, err
	}
	return note, nil
}

func fetchCount() (int, error) {
	var count int
	err := DB.QueryRow("select count(*) from notes;").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func insertNote(title string, category string) (Note, error) {
	var id int
	err := DB.QueryRow("insert into notes(title, category) values (?, ?) returning id", title, category).Scan(&id)
	if err != nil {
		return Note{}, err
	}
	note := Note{ID: id, Title: title, Category: category}
	return note, nil
}
