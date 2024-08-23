package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Import pprof for profiling
	"os"
)

// BioWrapper maps the top-level "bio" key in the JSON.
type BioWrapper struct {
	Bio Bio `json:"bio"`
}

// Bio maps the structure within the "bio" key.
type Bio struct {
	PersonalInfo struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Location string `json:"location"`
	} `json:"personal_info"`
	Education struct {
		HighSchool struct {
			SchoolName     string `json:"school_name"`
			GraduationYear int    `json:"graduation_year"`
		} `json:"high_school"`
		University struct {
			SchoolName     string   `json:"school_name"`
			GraduationYear int      `json:"graduation_year"`
			Degree         string   `json:"degree"`
			Courses        []string `json:"courses"`
		} `json:"university"`
	} `json:"education"`
	WorkExperience struct {
		Job1 struct {
			Company string `json:"company"`
			Role    string `json:"role"`
			Years   string `json:"years"`
		} `json:"job1"`
		Job2 struct {
			Company string `json:"company"`
			Role    string `json:"role"`
			Years   string `json:"years"`
		} `json:"job2"`
	} `json:"work_experience"`
}

// extractNames is a leaf function that processes course names.
func extractNames(data []string) []string {
	names := make([]string, 0, len(data))
	for _, name := range data {
		names = append(names, name)
	}
	return names
}

func main() {
	// Set up the HTTP server with pprof enabled
	http.HandleFunc("/", processor)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// processor handles incoming HTTP requests, processes the JSON, and returns the JSON data.
func processor(w http.ResponseWriter, r *http.Request) {
	// Read the JSON file
	file, err := os.ReadFile("./bio.json")
	if err != nil {
		http.Error(w, "Error reading bio.json", http.StatusInternalServerError)
		return
	}

	// Unmarshal the JSON into the BioWrapper struct
	var bioWrapper BioWrapper
	err = json.Unmarshal(file, &bioWrapper)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Access the Bio data
	bio := bioWrapper.Bio

	// Process the Bio data (e.g., print to console)
	fmt.Println("Name:", bio.PersonalInfo.Name)
	fmt.Println("University:", bio.Education.University.SchoolName)
	fmt.Println("Current Job:", bio.WorkExperience.Job2.Role)

	// Use the leaf function to extract course names
	courseNames := extractNames(bio.Education.University.Courses)
	fmt.Println("Courses:", courseNames)

	// Set the response header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")

	// Using json.MarshalIndent for pretty-printing
	responseData, err := json.MarshalIndent(bioWrapper, "", "    ")
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Write(responseData)
}
