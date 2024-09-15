package main

import "conducting-tenders/internal/app"

func main() {
	/*
		http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "SERVER_ADDRESS"+" --"+os.Getenv("SERVER_ADDRESS"))
		})

		fmt.Println("Server is listening on port 8080...")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("Error starting server: %s\n", err)
		}
	*/
	app.Run()
}
