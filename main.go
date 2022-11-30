package main

import (
	"fmt"
	"github.com/aidenwang9867/ShorlServer/app"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting ShorlServer on port %s...\n", port)
	router := mux.NewRouter().StrictSlash(true)

	// API Index.
	router.HandleFunc("/", app.Index)

	// Generate the short link given the original.
	// GET query to generate the short link for an input long link.
	router.HandleFunc("/generate", app.GenerateGetHandler).Methods(http.MethodGet)
	// POST query to generate the short link for an input long link, bulk access is supported by POST.
	router.HandleFunc("/generate", app.GeneratePostHandler).Methods(http.MethodPost)

	// Redirect short to original.
	router.HandleFunc("/r/{short_link}", app.RedirectGetHandler).Methods(http.MethodGet)

	http.Handle("/", router)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}

	//ctx, client, err := utils.InitDatabase()
	//if err != nil {
	//	log.Fatalf("Error initializing the database: %v", err)
	//}

	//var users []utils.User
	//for i := 0; i < 30000; i++ {
	//	user := utils.User{
	//		UserID:       fmt.Sprintf("%d", i),
	//		PasswordHash: fmt.Sprintf("%d", crc32.ChecksumIEEE([]byte(fmt.Sprintf("%d", i)))),
	//		PasswordSalt: fmt.Sprintf("%d", crc32.ChecksumIEEE([]byte(fmt.Sprintf("%d", i+1)))),
	//	}
	//	fmt.Printf("userID: %s, passwordHash: %s, salt:%s \n", user.UserID, user.PasswordHash, user.PasswordSalt)
	//	users = append(users, user)
	//}
	//err = utils.WriteToUsers(ctx, client, users)
	//if err != nil {
	//	log.Fatalf("Error writing to the users table: %v", err)
	//}
	//
	//var tokens []utils.Token
	//for i := 0; i < 50000; i++ {
	//	token := utils.Token{
	//		Token:             utils.EncodeLink(fmt.Sprintf("%d", i)),
	//		AccessLimitPerMin: fmt.Sprintf("%d", i+1),
	//	}
	//	fmt.Printf("Token: %s, accessLimit: %s\n", token.Token, token.AccessLimitPerMin)
	//	tokens = append(tokens, token)
	//}
	//err = utils.WriteToTokens(ctx, client, tokens)
	//if err != nil {
	//	log.Fatalf("Error writing to the users table: %v", err)
	//}

	//fmt.Println(utils.GetLongLinkByShortLink(ctx, client, "5kThM"))
}
