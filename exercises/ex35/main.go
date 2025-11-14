package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// formHandler handles requests to the "/form" endpoint.
// It serves a form on GET requests and processes the submitted data on POST requests.
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Switch on the request method.
	switch r.Method {
	case "GET":
		// If the method is GET, just serve the HTML form.
		tmpl, err := template.ParseFiles("templates/form.html")
		if err != nil {
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)

	case "POST":
		// If the method is POST, we need to parse the form data.
		// It's crucial to call ParseForm() before accessing form values.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		// After parsing, the form values are available in r.PostForm.
		// We can retrieve them using r.PostFormValue, which is a convenient shortcut.
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")

		// We can now use the extracted data.
		// Here, we'll just display it back to the user.
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<h1>Submission Successful</h1>")
		fmt.Fprintf(w, "<p>Thank you, <strong>%s</strong>!</p>", name)
		fmt.Fprintf(w, "<p>We have received your email: <strong>%s</strong></p>", email)
		fmt.Fprintf(w, `<a href="/form">Go Back</a>`)

	default:
		// For any other method, we return a "Method Not Allowed" error.
		http.Error(w, "Method is not supported.", http.StatusNotFound)
	}
}

// helloHandler is a simple handler for a different route to show a basic GET request.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello from the /hello endpoint!")
}

func main() {
	// Register our handler functions for the given routes.
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting server on port 8080...")
	fmt.Println("Visit http://localhost:8080/form to see the form.")
	fmt.Println("Visit http://localhost:8080/hello for a simple GET response.")

	// Start the server on port 8080.
	// log.Fatal will print any error and exit the program if the server fails to start.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
