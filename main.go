package main

// import "fmt"
import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

var RE_TZ = regexp.MustCompile(`(?s)BEGIN:VTIMEZONE.+END:VTIMEZONE`)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// log.Println("new req", r.RemoteAddr, r.Header.Get("X-Forwarded-For"))
	calendarURL := os.Getenv("ICAL_URL")

	resp, err := http.Get(calendarURL)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	replaced := bytes.ReplaceAll(body, []byte("Central Europe Standard Time"), []byte("Europe/Zagreb"))
	replaced = bytes.ReplaceAll(replaced, []byte("Central European Standard Time"), []byte("Europe/Zagreb"))
	replaced = bytes.ReplaceAll(replaced, []byte(";TZID=Europe/Zagreb"), []byte{})
	replaced = RE_TZ.ReplaceAll(replaced, []byte("X-WR-TIMEZONE:Europe/Zagreb"))

	r.Header.Add("Content-type", resp.Header.Get("Content-type"))
	w.WriteHeader(http.StatusOK)
	w.Write(replaced)
}

func main() {
	godotenv.Load()
	if os.Getenv("ICAL_URL") == "" {
		log.Println("The ICAL_URL env var must be present either in the environment or .env")
		os.Exit(1)
	}
	path := os.Getenv("SERVER_PATH")
	if path == "" {
		path = "/ical"
	}
	listenAddress := os.Getenv("SERVER_LISTEN_ADDRESS")
	if listenAddress == "" {
		listenAddress = "0.0.0.0:80"
	}
	http.HandleFunc(path, handleRequest)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
