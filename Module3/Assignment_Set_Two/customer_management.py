class Customer:
    def __init__(self, name, email, phone):
        self.name = name
        self.email = email
        self.phone = phone

    def display_customer(self):
        return f"Name: {self.name}, Email: {self.email}, Phone: {self.phone}"

# Add a new customer
def add_customer(customers, name, email, phone):
    customer = Customer(name, email, phone)
    customers.append(customer)
    print(f"Customer '{name}' added successfully!")

# View all customers
def view_customers(customers):
    for customer in customers:
        print(customer.display_customer())
