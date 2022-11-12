# LAB 06

Using the language of your choice, create a Movie class that includes attributes for name, year released, and genre (science fiction, romance, drama, other). Prompt the user for inputs for each attribute, create a new instance of the Movie class using user input values, and output a summary using the format of “<movie name> released in <year> is a <genre> film”.

Use TDD and include unit tests for each property and method defined on the Movie class. You will have a Movie class and a separate runner that prompts the user and creates an instance of the Movie class, passing input values for properties and calling methods on the Movie object. You will be writing unit tests against the Movie class only.

In the Movie class, raise an Exception in each of the following cases:

Name is less than 5 characters or more than 35 characters
Year released is not a whole number
Year released is less than 1950 or greater than 2022
Genre is not in the list of accepted genre values

Include unit tests that verify the expected Exceptions are raised when invalid data is provided.