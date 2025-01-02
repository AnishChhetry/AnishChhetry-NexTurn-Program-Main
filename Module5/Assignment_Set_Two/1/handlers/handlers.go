package handlers

import (
	"Blog_Management_System/dbconfig"
	"Blog_Management_System/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var blog models.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	if blog.Title == "" || blog.Content == "" || blog.Author == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		return
	}

	db := dbconfig.GetDB()

	statement, err := db.Prepare("INSERT INTO Blog (title, content, author, timestamp) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Error preparing statement: ", err)
		http.Error(w, "Error preparing statement", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	_, err = statement.Exec(blog.Title, blog.Content, blog.Author, time.Now())
	if err != nil {
		log.Fatal("Error inserting blog: ", err)
		http.Error(w, "Error inserting blog", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(blog)
}

func GetBlogByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	db := dbconfig.GetDB()

	row := db.QueryRow("SELECT id, title, content, author, timestamp FROM Blog WHERE id = ?", id)

	var blog models.Blog
	err = row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
	if err != nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(blog)
}

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := dbconfig.GetDB()

	rows, err := db.Query("SELECT id, title, content, author, timestamp FROM Blog")
	if err != nil {
		log.Fatal("Error fetching blogs: ", err)
		http.Error(w, "Error fetching blogs", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var blogs []models.Blog
	for rows.Next() {
		var blog models.Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp); err != nil {
			http.Error(w, "Error scanning blog data", http.StatusInternalServerError)
			return
		}
		blogs = append(blogs, blog)
	}

	json.NewEncoder(w).Encode(blogs)
}

func UpdateBlogByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var blog models.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	if blog.Title == "" || blog.Content == "" || blog.Author == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		return
	}

	db := dbconfig.GetDB()

	statement, err := db.Prepare("UPDATE Blog SET title = ?, content = ?, author = ?, timestamp = ? WHERE id = ?")
	if err != nil {
		log.Fatal("Error preparing statement: ", err)
		http.Error(w, "Error preparing statement", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	_, err = statement.Exec(blog.Title, blog.Content, blog.Author, time.Now(), id)
	if err != nil {
		log.Fatal("Error updating blog: ", err)
		http.Error(w, "Error updating blog", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(blog)
}

func DeleteBlogByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	db := dbconfig.GetDB()

	statement, err := db.Prepare("DELETE FROM Blog WHERE id = ?")
	if err != nil {
		log.Fatal("Error preparing statement: ", err)
		http.Error(w, "Error preparing statement", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		log.Fatal("Error deleting blog: ", err)
		http.Error(w, "Error deleting blog", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
