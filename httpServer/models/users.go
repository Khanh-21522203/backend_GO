package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Profile  string `json:"profileImage"`
}

// func (u *User) NewUser()

func (u *User) IsExist() (bool, error) {
	userList, err := u.readFromFile()
	if err != nil {
		return false, err
	}
	for _, user := range userList {
		if user.Username == u.Username {
			return true, nil
		}
	}

	return false, nil
}

func (u *User) CheckPassword() (bool, error) {
	userList, err := u.readFromFile()
	if err != nil {
		return false, err
	}
	for _, user := range userList {
		if user.Username == u.Username {
			if user.Password == u.Password {
				return true, nil
			} else {
				return false, errors.New("wrong password")
			}
		}
	}

	return false, errors.New("username does not exist")
}

func (u *User) AddNew() error {
	userList, err := u.readFromFile()
	if err != nil {
		return err
	}
	checkExist, err := u.IsExist()
	if err != nil {
		return err
	}
	if checkExist {
		return errors.New("username have been already used")
	}

	userList = append(userList, *u)
	err = u.writeToFile(userList)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) UpdateInfo() error {
	userList, err := u.readFromFile()
	if err != nil {
		return err
	}
	checkExist, err := u.IsExist()
	if err != nil {
		return err
	}
	if !checkExist {
		return errors.New("the username is not exist")
	}

	for i, user := range userList {
		if user.Username == u.Username {
			userList[i] = *u
			break
		}
	}
	err = u.writeToFile(userList)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) readFromFile() ([]User, error) {
	jsonFile, err := os.Open("user.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var users []User
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return users, nil
}

func (u *User) writeToFile(userList []User) error {
	data, err := json.Marshal(userList)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return err
	}

	err = os.WriteFile("user.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	return nil
}

// func main() {
// 	user := User{
// 		Username: "khan11h",
// 		Password: "123456",
// 	}
// 	user.AddNew()
// 	list, _ := user.readFromFile()
// 	fmt.Print(list)
// }
