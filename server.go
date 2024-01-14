package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/history", historyhandle)
	http.HandleFunc("/download", downloadhandle)
	http.HandleFunc("/auth", authhandle)
	http.HandleFunc("/enter", enterhandle)
	http.ListenAndServe(":8080", nil)
}

func historyhandle(w http.ResponseWriter, r *http.Request) {
	request := r.URL.Query().Get("token")
	user := decodeValid(request)
	login := fmt.Sprintf("%v", user["login"])
	password := fmt.Sprintf("%v", user["password"])
	fmt.Fprint(w, HistoryToJWT(GetHistory(login, password)))
}

func downloadhandle(w http.ResponseWriter, r *http.Request) {
	request := r.FormValue("token")
	user := decodeValid(request)
	url := fmt.Sprintf("%v", user["url"])
	login := fmt.Sprintf("%v", user["login"])
	password := fmt.Sprintf("%v", user["password"])
	link := fmt.Sprintf("%v", user["link"])
	flag := 0
	Download(url)
	videoFile, err := os.Open("video.mp4")
	fileInfo, err := videoFile.Stat()
	if err != nil {
		// Выводим ошибку в консоль и в ответ
		fmt.Println(err)
		http.Error(w, "Ошибка при получении информации о файле", http.StatusInternalServerError)
		flag += 1
	}
	fileSize := fileInfo.Size()

	// Устанавливаем заголовки ответа
	w.Header().Set("Content-Type", "video/mp4")                             // Тип содержимого
	w.Header().Set("Content-Length", strconv.FormatInt(fileSize, 10))       // Длина содержимого в байтах
	w.Header().Set("Content-Disposition", "attachment; filename=video.mp4") // Предложение сохранить файл под определенным именем

	// Копируем данные из видео файла в ответ
	_, err = io.Copy(w, videoFile)
	if err != nil {
		// Выводим ошибку в консоль
		fmt.Println(err)
		flag += 1
		// Не отправляем ошибку в ответ, так как данные уже начали передаваться
	}
	if flag == 0 {
		if link != "" {
			SetHistory(login, password, link)
		}
	}

}

func authhandle(w http.ResponseWriter, r *http.Request) {
	request := r.URL.Query().Get("token")
	user := decodeValid(request)
	login := fmt.Sprintf("%v", user["login"])
	password := fmt.Sprintf("%v", user["password"])
	if IsInDB(login, password) {
		fmt.Fprint(w, ZacodeStr("Вы уже зарегистрированы"))
	} else {
		AddUserToDB(login, password)
		fmt.Fprint(w, ZacodeStr("Регистрация прошла успешно"))
		fmt.Println("Успешная регистрация")
	}
}
func enterhandle(w http.ResponseWriter, r *http.Request) {
	request := r.URL.Query().Get("token")
	user := decodeValid(request)
	login := fmt.Sprintf("%v", user["login"])
	password := fmt.Sprintf("%v", user["password"])
	if IsInDB(login, password) {
		fmt.Fprint(w, ZacodeStr("Успешный вход"))
		fmt.Println("Успешный вход")
	} else {
		fmt.Fprint(w, ZacodeStr("Зарегистрируйтесь"))
		fmt.Println("Нет в БД")
	}
}
func decodeValid(tokenString string) jwt.MapClaims {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("123456789"), nil
	})
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
	}
	return claims
}

func ZacodeStr(str string) string {
	tokeExpiresAt := time.Now().Add(time.Minute * time.Duration(1))
	user := jwt.MapClaims{
		"Massage":    str,
		"Expires_at": tokeExpiresAt.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	tokenString, err := token.SignedString([]byte(JWTCODE))
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
		return ""
	}
	return tokenString
}
