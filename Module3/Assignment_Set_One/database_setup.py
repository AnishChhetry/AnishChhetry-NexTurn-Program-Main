import sqlite3

def create_database():
    # Connect to SQLite database (or create it if it doesn't exist)
    conn = sqlite3.connect('books.db')
    cursor = conn.cursor()

    # Create the books table
    cursor.execute('''
    CREATE TABLE IF NOT EXISTS books (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        author TEXT NOT NULL,
        published_year INTEGER NOT NULL,
        genre TEXT NOT NULL
    )
    ''')

    # Sample books data
    books = [
        ("The Great Gatsby", "F. Scott Fitzgerald", 1925, "Fiction"),
        ("To Kill a Mockingbird", "Harper Lee", 1960, "Fiction"),
        ("Alice in Wonderland", "Lewis Caroll", 1865, "Fiction"),
        ("Hello World", "abc", 2025, "Thriller")
    ]

    # Insert sample data
    cursor.executemany('''
    INSERT INTO books (title, author, published_year, genre)
    VALUES (?, ?, ?, ?)
    ''', books)

    conn.commit()
    conn.close()

if __name__ == '__main__':
    create_database()
    print("Database created and populated with sample data.")
