# Bookstore API

This is a simple REST API for a bookstore, enabling users to explore books, add them to a cart, and proceed with a purchase simulation (without actual payment processing).

## Scope of Work and Expectations

There are three user permission levels:

- Anonymous users: can view and filter books by category. Multiple categories can be selected at once, which will display all books matching any of the chosen categories.
- Authenticated users: have the same access as anonymous users, but can also add books to their cart, proceed to checkout, and purchase books.
- Administrators: have permission to create, read, update, and delete (CRUD) categories and books.

## Functional Requirements

- Users should be able to register and authenticate with an email and password via the API.
- Only database modifications can assign admin status to users.
- Admins can create, update, and delete categories. Each category has a unique name and is associated with books. Categories are non-hierarchical, meaning they cannot be nested.
- Admins can also manage books. Each book has a title, publication year, author, price in USD, and category. Books are required to belong to a category and have an inventory count. Books that are out of stock should not appear in the listing and cannot be purchased. Stock is set when a book is created and cannot be modified later.
- Visitors (including those who are not logged in) should be able to view and filter the list of books.
- Authenticated users can add books to their cart. Users can buy multiple books at once, but only one copy of each title (no quantity adjustments needed).
- A checkout endpoint should finalize the purchase for items in the cart. This endpoint simulates a payment process without requiring any payment details. It clears the cart and deducts the purchased books from stock.
- Handle cases where two users attempt to buy the last copy of a book simultaneously; only one should succeed.
- If a user adds a book to their cart and does not complete the purchase within 30 minutes, the book should automatically become available to others again.

## Running the Application

- `make dc` runs the application using Docker Compose, with the app container exposed on port 8080.
- `make test` runs tests.
- `make run` launches the app locally on port 8080 without Docker.
- `make lint` runs the linter.


## Solution Details

- Clean architecture setup (handler → service → repository).
- Standard Go project structure (or close to it 😊).
- Includes Docker Compose and a Makefile.
- PostgreSQL migrations are provided.
- Includes a Postman collection.
- Race conditions are managed with SQL transactions and `SELECT ... FOR UPDATE` queries to prevent issues.
