# School Management System

This project is a simple School Management System that allows school administrators to manage some activities and processes of the school. The system provides features such as student management and lectures/courses management.

This project is my submission to a Altschool assignment for semester 3.

## Installation

To install and run the School Management System, follow these steps:

1. Clone this repository to your local machine:
   `git clone https://github.com/<your-username>/school-management-system.git`
2. Change into the project directory:
   `cd school-management-system`
3. Run the following command to start the project using Docker:
   `docker-compose up`

This will start the project and create a MySQL database container and a RESTful API server container.

## Usage

To use the School Management System, you can send HTTP requests to the RESTful API using tools such as cURL or Postman.

The following endpoints are available:

- `GET /students`: Get a list of all students.
- `POST /students`: Add a new student.
- `GET /students/:id/courses`: Get a list of courses taken by a student.
- `PUT /students/:id/courses`: Update the list of courses taken by a student.
- `GET /courses/:id/students`: Get a list of students taking a particular course.

Refer to the OpenAPI documentation for more details on the available endpoints.

## Contributing

If you would like to contribute to this project, please open an issue or a pull request.

## License

This project is licensed under the MIT License.
