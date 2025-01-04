import book_management
import customer_management
import sales_management

def menu():
    # Initialize books and customers list in main.py
    books = []
    customers = []
    
    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        choice = input("Enter your choice: ")

        if choice == "1":
            book_management_menu(books)
        elif choice == "2":
            customer_management_menu(customers)
        elif choice == "3":
            sales_management_menu(customers, books)
        elif choice == "4":
            print("Exiting program.")
            break
        else:
            print("Invalid choice. Please try again.")

def book_management_menu(books):
    while True:
        print("\nBook Management")
        print("1. Add Book")
        print("2. View Books")
        print("3. Search Book")
        print("4. Back to Main Menu")
        choice = input("Enter your choice: ")

        if choice == "1":
            title = input("Enter book title: ")
            author = input("Enter book author: ")
            try:
                price = float(input("Enter book price: "))
                quantity = int(input("Enter book quantity: "))
                if price <= 0 or quantity <= 0:
                    print("Price and quantity must be positive numbers.")
                else:
                    book_management.add_book(books, title, author, price, quantity)
            except ValueError:
                print("Invalid input! Price must be a number and quantity must be an integer.")
        elif choice == "2":
            book_management.view_books(books)
        elif choice == "3":
            search_term = input("Enter book title or author to search: ")
            book_management.search_book(books, search_term)
        elif choice == "4":
            break
        else:
            print("Invalid choice. Please try again.")

def customer_management_menu(customers):
    while True:
        print("\nCustomer Management")
        print("1. Add Customer")
        print("2. View Customers")
        print("3. Back to Main Menu")
        choice = input("Enter your choice: ")

        if choice == "1":
            name = input("Enter customer name: ")
            email = input("Enter customer email: ")
            phone = input("Enter customer phone: ")
            customer_management.add_customer(customers, name, email, phone)
        elif choice == "2":
            customer_management.view_customers(customers)
        elif choice == "3":
            break
        else:
            print("Invalid choice. Please try again.")

def sales_management_menu(customers, books):
    while True:
        print("\nSales Management")
        print("1. Sell Book")
        print("2. View Sales")
        print("3. Back to Main Menu")
        choice = input("Enter your choice: ")

        if choice == "1":
            customer_name = input("Enter customer name: ")
            book_title = input("Enter book title: ")
            try:
                quantity = int(input("Enter quantity: "))
                if quantity <= 0:
                    print("Quantity must be a positive integer.")
                else:
                    sales_management.sell_book(customers, books, customer_name, book_title, quantity)
            except ValueError:
                print("Invalid input! Quantity must be an integer.")
        elif choice == "2":
            sales_management.view_sales()
        elif choice == "3":
            break
        else:
            print("Invalid choice. Please try again.")

# Run the program
if __name__ == "__main__":
    menu()
