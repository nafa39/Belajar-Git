package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	//"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println("JWT secret:", os.Getenv("JWT_SECRET"))

	router := httprouter.New()

	// router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 	w.Write([]byte("Hello World!"))
	// })

	// router.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 	w.Write([]byte("Hello " + ps.ByName("name") + "!"))
	// })

	// router.GET("/", middleware(welcome))

	newrouter := middleware(router)

	router.GET("/", welcome)

	router.POST("/login", login)

	router.GET("/dashboard", authenticate(dashboad))

	router.POST("/user/all", authorization(authenticate(getUser)))

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", newrouter))
}

func welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Welcome"))
}

func middleware(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware")
		router.ServeHTTP(w, r)
	})
}

func authenticate(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tokenstring := r.Header.Get("Authorization")
		if tokenstring == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			log.Println("Token not found")
			return
		}

		tokenstring = strings.Replace(tokenstring, "Bearer ", "", 1)

		log.Print(os.Getenv("JWT_SECRET"))

		token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				log.Println("Invalid token wrong")
				return nil, fmt.Errorf("invalid token")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err == nil && token.Valid {
			next(w, r, ps)
		} else {
			log.Print(err)
			log.Print(token.Valid)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			log.Print("Invalid token")
			return
		}
		log.Println("Authenticate")
		//next(w, r, ps)
	}
}

func authorization(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Println("Authorization")
		next(w, r, ps)
	}
}

func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	log.Println("Username:", username)
	log.Println("Password:", password)

	// Check username and password
	if username == "admin" && password == "admin" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

		log.Print(os.Getenv("JWT_SECRET"))

		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte(tokenString))
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Wrong username or password"))
}

func dashboad(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Dashboard"))
}

func getUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := map[string]string{
		"username": "admin",
		"email":    "apaaja@mail.com",
	}

	w.Write([]byte(fmt.Sprintf("Get User: %v", user)))
	//w.Write([]byte("Get User"))
}
