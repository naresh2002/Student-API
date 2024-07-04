# Student-Management-System
REST API in Go that performs basic CRUD (Create, Read, Update, Delete) operations on a list of students  
Using **Ollama** with model **llama3** to generate summary for given student  
Added **validations** to check proper request data for our struct using **middleware**  

# How to run this project
1. Clone the repository using the command :  
` git clone https://github.com/naresh2002/Student-Management-System.git `  
2. Open terminal in the cloned directory and run :
` go mod tidy `  
3. To run GetStudentSummary endpoint we need to first install
   a. ([**Ollama**](https://www.ollama.com/download))  
   b. then install ([**llama3**](https://www.ollama.com/library/llama3)) model of Ollama using  
   ` ollama run llama3 `  
   c. Once it's done open a new terminal and start Ollama using the commands  
   ` sudo systemctl enable ollama `  
   ` sudo systemctl start ollama `  
   ` ollama serve `  
4. Then you are all ready to test the project, just start the project using  
` go run main.go `  
Then you can test endpoints given below using either Postman or running curl commands through terminal  

# Endpoints
1. GetAllStudents [GET]  
   ``` http://127.0.0.1:8000/student/all ```  
   curl command :  
   curl -X GET http://localhost:8000/student/all | jq  

2. GetStudentById [GET]  
   ``` http://127.0.0.1:8000/student/1 ```  
   curl command :  
   curl -X GET http://localhost:8000/student/1 | jq  
   
3. CreateStudent [POST]  
   ``` http://127.0.0.1:8000/student/add ```  
   JSON input :  
   {  
    "name": "name2",  
    "age": 23,  
    "email": "name3@gmail.com"  
    }  
   curl command :  
   curl -X POST http://localhost:8000/student/add -d '{  
      "name": "name2",  
      "age": 23,  
      "email": "name3@gmail.com"  
    }'

4. UpdateStudent [PUT]  
   ``` http://127.0.0.1:8000/student/update/2 ```  
   JSON input :  
   {  
    "name": "name2",  
    "age": 23,  
    "email": "name2@gmail.com"  
    }  
   curl command :  
   curl -X PUT http://localhost:8000/student/update/2 -d '{  
      "name": "name2",  
      "age": 23,  
      "email": "name2@gmail.com"  
    }'

5. DeleteStudent [DELETE]  
   ``` http://127.0.0.1:8000/student/delete/1 ```  
   curl command :  
   curl -X DELETE http://localhost:8000/student/delete/1

7. GetStudentSummary [GET]  
   ``` http://127.0.0.1:8000/student/summary/1 ```  
   curl command :  
   curl -X GET http://localhost:8000/student/summary/1 | jq
