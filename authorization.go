package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const JWTCODE = "123456789"

type User struct {
	Login   string
	Pswrd   string
	History []string
}

var NoHistory []string

func AddUserToDB(login, pswrd string) string {
	// создаём дэфолтного клиента
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://asankhalilov2005:UenvNf3MEQ4W6zfV@cluster0.mb8ybk1.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
		return "Ошибка регистрации"
	}
	// создаём соединение
	err = client.Connect(context.TODO())
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
		return "Ошибка регистрации"
	}
	// проверяем соединение
	err = client.Ping(context.TODO(), nil)
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
		return "Ошибка регистрации"
	}
	// обращаемся к коллекции Users из базы TubeSaver
	collection := client.Database("TubeSaver").Collection("Users")
	// создаём переменную в виде структуры User
	CurUser := User{login, pswrd, NoHistory}
	if CurUser.Login == "<nil>" || CurUser.Pswrd == "<nil>" {
		return "Ошибка регистрации"
	}
	// добавляем одиночный документ в коллекцию
	insertResult, err := collection.InsertOne(context.TODO(), CurUser)
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
		return "Ошибка регистрации"
	}
	filter := bson.D{{"login", login}, {"pswrd", pswrd}}
	update := bson.D{
		{"$set", bson.D{
			{"history", bson.A{}},
		}},
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return "Ошибка регистрации"
	}
	// выводим внутренний ID добавленного документа
	log.Println("Inserted a single document: ", insertResult.InsertedID)
	return "Успешная регистрация"
} // Функция добавления данных нового пользователя по умолчанию в бд

func IsInDB(login, pswrd string) bool {
	// создаём дэфолтного клиента
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://asankhalilov2005:UenvNf3MEQ4W6zfV@cluster0.mb8ybk1.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
		return false
	}
	// создаём соединение
	err = client.Connect(context.TODO())
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
		return false
	}
	// проверяем соединение
	err = client.Ping(context.TODO(), nil)
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
		return false
	}
	// обращаемся к коллекции Users из базы TubeSaver
	collection := client.Database("TubeSaver").Collection("Users")
	// создаём фильтр по которму мы будем искать клиента. был взят именно ID потому что они не повторяются
	filter := bson.D{{"login", login}, {"pswrd", pswrd}}
	// создаём переменную в которую будем записывать полученного клиента в результате поиска
	var result User
	// собственно ищем
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil { // проверяем ошибку если она есть то возвращаем пустую структуру вида User
		log.Println(err)
		return false
	}
	log.Println("User is in DB")
	return true
} // Функция проверяющаяя был ли пользоваетль зарегестрирован ранее

func GetHistory(login, pswrd string) []string {
	// создаём дэфолтного клиента
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://asankhalilov2005:UenvNf3MEQ4W6zfV@cluster0.mb8ybk1.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
		return NoHistory
	}
	// создаём соединение
	err = client.Connect(context.TODO())
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
		return NoHistory
	}
	// проверяем соединение
	err = client.Ping(context.TODO(), nil)
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
		return NoHistory
	}
	// обращаемся к коллекции Users из базы TubeSaver
	collection := client.Database("TubeSaver").Collection("Users")
	// создаём фильтр по которму мы будем искать клиента. был взят именно ID потому что они не повторяются
	filter := bson.D{{"login", login}, {"pswrd", pswrd}}
	// создаём переменную в которую будем записывать полученного клиента в результате поиска
	var result User
	// собственно ищем
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil { // проверяем ошибку если она есть то возвращаем пустую структуру вида User
		log.Println(err)
		return NoHistory
	}
	log.Println("History was found")
	return result.History
} // Функция возвращающая историю загрузок пользователя

func SetHistory(login, pswrd, link string) {
	// создаём дэфолтного клиента
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://asankhalilov2005:UenvNf3MEQ4W6zfV@cluster0.mb8ybk1.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
	}
	// создаём соединение
	err = client.Connect(context.TODO())
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
	}
	// проверяем соединение
	err = client.Ping(context.TODO(), nil)
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
	}
	// обращаемся к коллекции Users из базы TubeSaver
	collection := client.Database("TubeSaver").Collection("Users")
	filter := bson.D{{"login", login}, {"pswrd", pswrd}}
	update := bson.D{
		{"$push", bson.D{
			{"history", link},
		}},
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	log.Println("History was updated")
} // Функция добавляющая в историю загрузок новую ссылку
