// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START gae_go111_app]

// Sample helloworld is an App Engine app.
package main

// [START import]
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"cloud.google.com/go/datastore"
	"context"
)

type Ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

// [END import]
// [START main_func]

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/apollo", apolloHandler)
	http.HandleFunc("/data", dataHandler)
	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	// [END setting_port]
}

// [END main_func]

// [START indexHandler]

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}
func apolloHandler(w http.ResponseWriter, r *http.Request){
	
	ping := Ping{Status: 200, Result: "ok"}

     res, err := json.Marshal(ping)

     if err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
         return
     }

     w.Header().Set("Content-Type", "application/json")
     w.Write(res)

}
type Task struct{
	Description string
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Set your Google Cloud Platform project ID.
	projectID := "my-first-go-project-309402"

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the kind for the new entity.
	kind := "Task"
	// Sets the name/ID for the new entity.
	name := "sampletask1"
	// Creates a Key instance.
	taskKey := datastore.NameKey(kind, name, nil)

	// Creates a Task instance.
	task := Task{
		Description: "Buy milk",
	}

	// Saves the new entity.
	if _, err := client.Put(ctx, taskKey, &task); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}

	fmt.Printf("Saved %v: %v\n", taskKey, task.Description)
	fmt.Fprint(w, "Hello, World!")
}

// [END indexHandler]
// [END gae_go111_app]
