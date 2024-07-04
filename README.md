# Student-Management-System
REST API in Go that performs basic CRUD (Create, Read, Update, Delete) operations on a list of students  
Using **Ollama** with model **llama3** to generate summary for given student  
Added **validations** to check proper request data for our struct using **middleware**  


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

7. StudentSummary [GET]  
   ``` http://127.0.0.1:8000/student/summary/1 ```  
   curl command :  
   curl -X GET http://localhost:8000/student/summary/1 | jq
