package main

import "testing"

// тестировать при пустой бд, либо после каждого использования удалять тех пользователей, что были созданы во время теста

// чтобы начать тест нужно через терминал находясь в папке с проектом(желательно чтобы в папке были только файл который проверяется и файл тестер) ввести команду go test -cover

// может срабатывать через раз или два, поэтому если при запуске возникает ошибка то попробуй запустить еще пару раз. Если ошибка не исчезла надо попытатся разобраться из-за чего она. Возможно забыл почистить бд

func TestJWTCODE(t *testing.T) {
	result := JWTCODE
	expected := "123456789"
	if result != expected {
		t.Errorf("Incorrect result. Expected %s, got %s", expected, result)
	}
}

func TestAddUserToDB(t *testing.T) {
	login := "Testing"
	pswrd := "TestPswrd"
	User := User{login, pswrd, NoHistory}
	AddUserToDB(User.Login, User.Pswrd)
}

func TestAddUserToDB2(t *testing.T) {
	login := "Testing2"
	pswrd := "TestPswrd2"
	User := User{login, pswrd, NoHistory}
	AddUserToDB(User.Login, User.Pswrd)
}

func TestIsInDB(t *testing.T) {
	login := "Testing"
	pswrd := "TestPswrd"
	result := IsInDB(login, pswrd)
	expected := true
	if result != expected {
		t.Error("Incorrect result. Expected", expected, ", got", result)
	}
}

func TestIsInDB2(t *testing.T) {
	login := "Testing2"
	pswrd := "TestPswrd2"
	result := IsInDB(login, pswrd)
	expected := true
	if result != expected {
		t.Error("Incorrect result. Expected", expected, ", got", result)
	}
}

func TestSetHistory(t *testing.T) {
	nembers := "123321"
	login := "Testing"
	pswrd := "TestPswrd"
	result := "123321"
	SetHistory(login, pswrd, nembers)
	expected := GetHistory(login, pswrd)[0]
	if result != expected {
		t.Errorf("Incorrect result. Expected %s, got %s", expected, result)
	}
}

func TestGetHistory(t *testing.T) {
	login := "Testing"
	pswrd := "TestPswrd"
	result := "123321"
	expected := GetHistory(login, pswrd)[0]
	if result != expected {
		t.Errorf("Incorrect result. Expected %s, got %s", expected, result)
	}
}
