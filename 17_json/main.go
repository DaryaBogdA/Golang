package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	ID       int    `json:"id"`                        // поле в JSON -> "id"
	Name     string `json:"name,omitempty"`            // если пустая строка, поле не выводится
	Email    string `json:"-"`                         // поле игнорируется при сериализации
	Password string `json:"password,omitempty,string"` // преобразовать в строку
}

func main() {
	u := User{ID: 1, Name: "", Email: "test@test.com", Password: "secret"}
	data, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println(string(data))
	//{
	//	"id": 1,
	//	"password": "\"secret\""
	//}

	// Потоки
	// Запись в поток
	file, _ := os.Create("users.json")
	encoder := json.NewEncoder(file)
	for _, user := range users {
		if err := encoder.Encode(user); err != nil {
			// обработка ошибки
		}
	}

	// Чтение из потока
	file, _ := os.Open("users.json")
	decoder := json.NewDecoder(file)
	var users []User
	for decoder.More() {
		var u User
		if err := decoder.Decode(&u); err != nil {
			break
		}
		users = append(users, u)
	}
}

type MyTime time.Time

// Кастомная сериализация
func (mt MyTime) MarshalJSON() ([]byte, error) {
	// форматируем в строку RFC3339
	return []byte(fmt.Sprintf(`"%s"`, time.Time(mt).Format(time.RFC3339))), nil
}

func (mt *MyTime) UnmarshalJSON(data []byte) error {
	// убираем кавычки и парсим
	s := strings.Trim(string(data), `"`)
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	*mt = MyTime(t)
	return nil
}
