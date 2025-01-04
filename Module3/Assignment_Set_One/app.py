from flask import Flask, request, jsonify
import sqlite3

app = Flask(__name__)

# Database connection function
def get_db_connection():
    conn = sqlite3.connect('books.db')
    conn.row_factory = sqlite3.Row
    return conn

# POST /books - Add a new book
@app.route('/books', methods=['POST'])
def add_book():
    data = request.get_json()
    title = data.get('title')
    author = data.get('author')
    published_year = data.get('published_year')
    genre = data.get('genre')

    # Input validation
    if not title or not author or not published_year or not genre:
        return jsonify({"error": "Invalid input", "message": "All fields are required"}), 400

    # Genre validation (Optional: Add predefined genres)
    valid_genres = ["Fiction", "Non-Fiction", "Mystery", "Sci-Fi"]
    if genre not in valid_genres:
        return jsonify({"error": "Invalid genre", "message": f"Valid genres are: {', '.join(valid_genres)}"}), 400

    try:
        conn = get_db_connection()
        cursor = conn.cursor()
        cursor.execute('''
        INSERT INTO books (title, author, published_year, genre)
        VALUES (?, ?, ?, ?)
        ''', (title, author, published_year, genre))
        conn.commit()

        # Get the ID of the newly inserted book
        book_id = cursor.lastrowid
        conn.close()

        return jsonify({"message": "Book added successfully", "book_id": book_id}), 201

    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

# GET /books - Retrieve all books
@app.route('/books', methods=['GET'])
def get_books():
    try:
        conn = get_db_connection()
        cursor = conn.cursor()
        cursor.execute('SELECT * FROM books')
        books = cursor.fetchall()
        conn.close()

        # Convert the query result to a list of dictionaries
        books_list = [dict(book) for book in books]
        return jsonify(books_list), 200

    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

# GET /books/<id> - Retrieve a specific book by id
@app.route('/books/<int:id>', methods=['GET'])
def get_book(id):
    try:
        conn = get_db_connection()
        cursor = conn.cursor()
        cursor.execute('SELECT * FROM books WHERE id = ?', (id,))
        book = cursor.fetchone()
        conn.close()

        if book is None:
            return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

        return jsonify(dict(book)), 200

    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

# PUT /books/<id> - Update an existing book by id
@app.route('/books/<int:id>', methods=['PUT'])
def update_book(id):
    data = request.get_json()
    title = data.get('title')
    author = data.get('author')
    published_year = data.get('published_year')
    genre = data.get('genre')

    # Input validation
    if not title or not author or not published_year or not genre:
        return jsonify({"error": "Invalid input", "message": "All fields are required"}), 400

    valid_genres = ["Fiction", "Non-Fiction", "Mystery", "Sci-Fi"]
    if genre not in valid_genres:
        return jsonify({"error": "Invalid genre", "message": f"Valid genres are: {', '.join(valid_genres)}"}), 400

    try:
        conn = get_db_connection()
        cursor = conn.cursor()
        cursor.execute('''
        UPDATE books SET title = ?, author = ?, published_year = ?, genre = ? WHERE id = ?
        ''', (title, author, published_year, genre, id))
        conn.commit()

        if cursor.rowcount == 0:
            return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

        conn.close()
        return jsonify({"message": "Book updated successfully"}), 200

    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

# DELETE /books/<id> - Delete a book by id
@app.route('/books/<int:id>', methods=['DELETE'])
def delete_book(id):
    try:
        conn = get_db_connection()
        cursor = conn.cursor()
        cursor.execute('DELETE FROM books WHERE id = ?', (id,))
        conn.commit()

        if cursor.rowcount == 0:
            return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

        conn.close()
        return jsonify({"message": "Book deleted successfully"}), 200

    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

if __name__ == '__main__':
    app.run(debug=True)
