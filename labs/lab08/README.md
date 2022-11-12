# LAB 08

Build a product/ordering system API in the language of your choice (the solution includes Python using FastAPI). Include the following models within that API:

* Product Structure
    - Internal ID (auto-increment)
    - Product number (e.g., "ABC-123")
    - Product description
    - Unit cost
* Order Structure
    - Internal ID (auto-increment)
    - Product ID (foreign key)
    - Order number (e.g., "12345")
    - Quantity
    - Total
* Operations to Include
    - List all products in the catalog
    - List a specific product in the catalog by product number
    - Accept a new order for a single product
    - Retrieve a specific order's details by order number

Use a database (or in memory data structure) of your choice for persisting your catalog and order information. The starting point for the lab includes sample data for a set of products.

You'll be seeking to build the components of your API in a loosely-coupled manner, layer components in modules correctly, etc. such that the components of your API are clean and well architected. For example, whatever storage/persistence strategy you use, the rest of the API should not need to know or care. Ensure that functionality is sufficiently isolated to modules and that each module is insulated from the others (to prevent brittle coupling).

## Part 1

Build the set of features outlined above using TDD. Use mocking where possible to mock the components of your application so your tests remain true unit tests (isolated).

## Part 2

Add additional operations to your API:

* Add a new product to the catalog
* Edit an existing product

As with the previous, use TDD to build a corresponding set of automated tests to accompany the production code. Practice Red/Green/Refactor as you move through the coding of this new set of operations. Ensure you use mocking with the new operations to maintain isolation.
