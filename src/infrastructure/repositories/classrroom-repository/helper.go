package classrroomrepository

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/classroom/v1"
	"google.golang.org/api/option"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// handle the auth-code returned from google
func handleCode(authCode chan string, wg *sync.WaitGroup) {
	fmt.Println("Waiting for the user confirmation...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		authCode <- r.FormValue("code")
		fmt.Fprintf(w, "You may close this window now!")
		wg.Done()
	})

	err := http.ListenAndServe(":5003", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Allow the application accessing the link: \n%v\n", authURL)

	// channel for the authCode
	authCode := make(chan string)

	// creates wait group
	var wg sync.WaitGroup
	wg.Add(1)

	// runs code handler concurrently
	go handleCode(authCode, &wg)

	// exchanges the authorization token
	tok, err := config.Exchange(context.TODO(), <-authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}

	wg.Wait()

	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func GetClassroomService() *classroom.Service {
	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read credentials file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(
		b,
		classroom.ClassroomCoursesReadonlyScope,
		classroom.ClassroomStudentSubmissionsStudentsReadonlyScope,
		classroom.ClassroomRostersReadonlyScope,
		classroom.ClassroomProfileEmailsScope,
	)

	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := classroom.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create classroom Client %v", err)
	}

	return srv
}
