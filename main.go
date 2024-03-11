package main

import (
	"log"
	"os"

	"github.com/eduzgun/api-gateway-footy/models"
	"github.com/eduzgun/api-gateway-footy/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	/*

				url := "https://api-football-v1.p.rapidapi.com/v3/timezone"

				req, _ := http.NewRequest("GET", url, nil)

				req.Header.Add("X-RapidAPI-Key", "76a5c68254msh15c95aa5f37d156p1a408ajsn7bc8944039ed")
				req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")

				res, _ := http.DefaultClient.Do(req)

				defer res.Body.Close()
				body, _ := io.ReadAll(res.Body)

				fmt.Println(res)
				fmt.Println(string(body))



			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello World! \n COYG"))
			})

			log.Println("Starting an API Gateway on port 8080")
			log.Fatal(http.ListenAndServe(":8080", nil))



			http.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
				if req.Method != "POST" {
					http.NotFound(res, req)
					return
				}


					params := req.URL.Query()
					body, err := ioutil.ReadAll(req.Body)
					if err != nil {
						http.Error(res, http.StatusText(http.StatusBadRequest), 400)
						return
					}

					login logic here
					token, err := myAuth.DoLogin(body, params)
					if err == nil {
						res.WriteHeader(http.StatusOK)
					} else {
						res.WriteHeader(http.StatusUnauthorized)
					}

			})

			http.HandleFunc("/logout", func(res http.ResponseWriter, req *http.Request) {
				if req.Method != "POST" {
					http.NotFound(res, req)
					return
				}
					if !myAuth.authenticate(req.Header.Get("Authorization")) {
						res.WriteHeader(http.StatusUnauthorized)
						return
					}
			})

		// end session and remove any state related to user login account

		http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
			log.Printf("incoming request: %s %", req.Host, req.URL.String())

				if !myAuth.authenticate(req.Header.Get("Authorization")) {
					res.WriteHeader(http.StatusUnauthorized)
					return
				}
		})

		http.HandleFunc("/coyg", func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte("YEO HAHAHAHA"))
			log.Printf("incoming request: %s %", req.Host, req.URL.String())

				if !myAuth.authenticate(req.Header.Get("Authorization")) {
					res.WriteHeader(http.StatusUnauthorized)
					return
				}
		})

		//implement routing and forwarding

		log.Println("API gateway listening on port 8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Panic(err)
		}
	*/

	// Create a new gin instance
	r := gin.Default()

	// Load .env file and Create a new connection to the database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Initialize DB
	models.InitDB(config)

	// Load the routes
	routes.AuthRoutes(r)

	// Run the server
	r.Run(":8080")

}
