# StudentAPIs

Students Apis for fetching data from in-memory database developed in go

## Steps to run:
1. Either clone the repository in your local system then run the getStudentAPIs.exe file or simply run go run main.go if you have go installed
2. Or open the hosted [URL](https://student-terrific-elephant-ot.cfapps.us10-001.hana.ondemand.com) and simple add the routes to get desired results

## Routes:
1. **/getStudents**
     - Returns all the students in database
     - no params
     - Return type : JSON/Student{s}
3. **/getStudent**
     - Returns particular student from DB
     - param : id/RollNo
     - Return type: JSON/Student
4. **/addStudent**
     - Adds student into db
     - body : Student object
     - No returns
5. **/deleteStudent**
     - Deletes student from db
     - params : id/RollNo
     - No returns
