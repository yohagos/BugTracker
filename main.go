package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type person struct {
	ID        primitive.ObjectID `bson:"_id"`
	name      string             `bson:"name"`
	city      string             `bson:"city"`
	createdAt time.Time          `bson:"createdAt"`
	updatedAt time.Time          `bson:"updatedAt"`
}

var (
	templates *template.Template
	msgList   []string

	collection  *mongo.Collection
	mongoClient *mongo.Client
)

var ctx = context.TODO()

func main() {
	initMongoDB()
	collection = mongoClient.Database("mydb").Collection("person")
	yosef := person{
		ID:        primitive.NewObjectID(),
		name:      "Yosef",
		city:      "Darmstadt",
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	_, err := collection.InsertOne(ctx, bson.M{
		"name": yosef.name,
	})
	ifError(err)

	loadTemplate("static/*.html")

	r := newRouter()
	log.Println("Server starting..")
	if err := http.ListenAndServe(":8888", r); err != nil {
		log.Fatal("Server error! Message : ", err)
	}
}

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexGETHandler).Methods("GET")
	router.HandleFunc("/", indexPOSTHandler).Methods("POST")

	fs := http.FileServer(http.Dir("static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return router
}

func indexGETHandler(w http.ResponseWriter, r *http.Request) {
	var displayList bool
	displayList = false
	if len(msgList) > 0 {
		displayList = true
	}

	executeTemplate(w, "index.html", struct {
		Display  bool
		Messages []string
	}{
		Display:  displayList,
		Messages: msgList,
	})
}

func indexPOSTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	message := r.PostForm.Get("txtArea")
	msgList = append(msgList, message)
	http.Redirect(w, r, "/", 302)
}

func loadTemplate(pattern string) {
	templates = template.Must(template.ParseGlob(pattern))
}

func executeTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

func initMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	ifError(err)

	err = client.Ping(context.TODO(), nil)
	ifError(err)
	mongoClient = client
	fmt.Println("Connected to MongoDB!")
}

func ifError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createDocument(person person) error {
	_, err := collection.InsertOne(ctx, person)
	return err
}
