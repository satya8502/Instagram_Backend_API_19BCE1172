package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User struct
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Post struct
type Post struct {
	ID               string `json:"id"`
	UID              string `json:"uid"`
	Caption          string `json:"caption"`
	Image_URL        string `json:"image_url"`
	Posted_Timestamp string `json:"Posted_Timestamp"`
}

//user_handle function will handle all the operations associated with the user
//If it's called with the get method, the user details will be fetched
//If it's the called with the post method, a new user will be created with the passed details
func user_handle(w http.ResponseWriter, r *http.Request) {
	//Establishing connection with MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	conn, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(conn)
	if err != nil {
		log.Fatal(err)
	}
	//Fetching the user details corresponding to the given ID
	if r.Method == "GET" {
		id := strings.TrimPrefix(r.URL.Path, "/users/")
		post_collection := client.Database("instadb").Collection("users")
		filterCursor, err := post_collection.Find(conn, bson.M{"id": id})
		if err != nil {
			log.Fatal(err)
		}
		var users []bson.M
		if err = filterCursor.All(conn, &users); err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(users)
	}
	//Creating a new user with details passed in the JSON body
	if r.Method == "POST" {
		defer client.Disconnect(conn)
		database_conn := client.Database("instadb")
		user_collection := database_conn.Collection("users")
		r.Body = http.MaxBytesReader(w, r.Body, 1048576)
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var p User
		_ = dec.Decode(&p)
		s := p.Password
		h := sha1.New()
		h.Write([]byte(s))
		temp := hex.EncodeToString(h.Sum(nil))
		var encrypted_p = User{
			ID:       p.ID,
			Name:     p.Name,
			Email:    p.Email,
			Password: temp,
		}

		_, insertErr := user_collection.InsertOne(conn, encrypted_p)
		if insertErr != nil {
			fmt.Println(insertErr)
		}
		json.NewEncoder(w).Encode(encrypted_p)
	}
}

//post_handle function will handle all the operations associated with the post
//If it's called with the get method, the post details will be fetched
//If it's the called with the post method, a new post will be created with the passed details
func post_handle(w http.ResponseWriter, r *http.Request) {
	//Establishing connection with MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	conn, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(conn)
	if err != nil {
		log.Fatal(err)
	}
	//Fetching the post details corresponding to the given ID
	if r.Method == "GET" {
		id := strings.TrimPrefix(r.URL.Path, "/posts/")
		post_collection := client.Database("instadb").Collection("posts")
		filterCursor, err := post_collection.Find(conn, bson.M{"id": id})
		if err != nil {
			log.Fatal(err)
		}
		var posts []bson.M
		if err = filterCursor.All(conn, &posts); err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(posts)
	}
	//Creating a new post with details passed in the JSON body
	if r.Method == "POST" {
		defer client.Disconnect(conn)
		database_conn := client.Database("instadb")
		posts_collection := database_conn.Collection("posts")
		r.Body = http.MaxBytesReader(w, r.Body, 1048576)
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var p Post
		_ = dec.Decode(&p)
		_, insertErr := posts_collection.InsertOne(conn, p)
		if insertErr != nil {
			fmt.Println(insertErr)
		}
		json.NewEncoder(w).Encode(p)
	}
}
func user_posts_handle(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	conn, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(conn)
	if err != nil {
		log.Fatal(err)
	}
	if r.Method == "GET" {
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatal(err)
		}
		conn, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(conn)
		if err != nil {
			log.Fatal(err)
		}
		id := strings.TrimPrefix(r.URL.Path, "/posts/users/")
		post_collection := client.Database("instadb").Collection("posts")
		filterCursor, err := post_collection.Find(conn, bson.M{"uid": id})
		if err != nil {
			log.Fatal(err)
		}
		var user_posts []bson.M
		if err = filterCursor.All(conn, &user_posts); err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(user_posts)
	}
}
func main() {
	http.HandleFunc("/users/", user_handle)
	http.HandleFunc("/posts/", post_handle)
	http.HandleFunc("/posts/users/", user_posts_handle)
	http.ListenAndServe(":8080", nil)
}
