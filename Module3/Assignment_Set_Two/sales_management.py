from customer_management import Customer
from book_management import Book

class Transaction(Customer):
    def __init__(self, name, email, phone, book_title, quantity_sold):
        super().__init__(name, email, phone)
        self.book_title = book_title
        self.quantity_sold = quantity_sold

    def display_transaction(self):
        return f"Customer: {self.name}, Book: {self.book_title}, Quantity Sold: {self.quantity_sold}"

# List to hold sales records
sales_records = []

# Sell a book
def sell_book(customers, books, customer_name, book_title, quantity):
    # Find customer
    customer = next((customer for customer in customers if customer.name == customer_name), None)
    if not customer:
        print("Error: Customer not found.")
        return

    # Find book
    book = next((book for book in books if book.title == book_title), None)
    if not book:
        print("Error: Book not found.")
        return

    # Check if book has sufficient quantity
    if book.quantity < quantity:
        print(f"Error: Only {book.quantity} copies available. Sale cannot be completed.")
        return

    # Process the sale
    book.quantity -= quantity
    transaction = Transaction(customer_name, customer.email, customer.phone, book_title, quantity)
    sales_records.append(transaction)
    print(f"Sale successful! Remaining quantity of '{book_title}': {book.quantity}")
    
# View all sales
def view_sales():
    for transaction in sales_records:
        print(transaction.display_transaction())
