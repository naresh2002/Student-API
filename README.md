# Student Management System

A REST API in Go that performs basic CRUD (Create, Read, Update, Delete) operations on a list of students.  
Utilizes **Ollama** with the **llama3** model to generate summaries for given students.  
Includes **validations** to ensure proper request data for the struct using **middleware**.  
Ensures **concurrency** safety using **RWMutex** to protect shared data.  

The project is designed to manage student data, providing endpoints to create, read, update, and delete student records. Additionally, it integrates with Ollama's AI model to generate concise summaries of student details, enhancing the system's functionality with intelligent data processing.

## How to Run This Project

1. Clone the repository using the command:  
    ```bash
    git clone https://github.com/naresh2002/Student-Management-System.git
    ```

2. Open a terminal in the cloned directory and run:  
    ```bash
    go mod tidy
    ```

3. **OPTIONAL:** <a name="optional"></a>  
To run the **Get Student Summary** endpoint, you need to first install:  
    a. [**Ollama**](https://www.ollama.com/download)  
    b. Then install the [**llama3**](https://www.ollama.com/library/llama3) model of Ollama using:  
     ```bash
     ollama run llama3
     ```  
    c. Once installation is complete, open a new terminal and start Ollama using the commands:  
     ```bash
     sudo systemctl enable ollama  
     sudo systemctl start ollama  
     ollama serve
     ```

4. Now you are all set to test the project. Start the project using:  
    ```bash
    go run main.go
    ```  
   You can test the endpoints given below using either Postman or by running curl commands through the terminal.

## Endpoints

1. **Get All Students** [GET]  
    ```http://127.0.0.1:8000/student/all```  
    curl command:  
    ```bash
    curl -X GET http://localhost:8000/student/all | jq
    ```

2. **Get Student By ID** [GET]  
    ```http://127.0.0.1:8000/student/{id}```  
    curl command:  
    ```bash
    curl -X GET http://localhost:8000/student/1 | jq
    ```

3. **Create Student** [POST]  
    ```http://127.0.0.1:8000/student/add```  
    JSON input:  
    ```json
    {  
        "name": "ABC DEF",  
        "age": 23,  
        "email": "abcdef@gmail.com"  
    }
    ```  
    curl command:  
    ```bash
    curl -X POST http://localhost:8000/student/add -d '{  
        "name": "ABC DEF",  
        "age": 23,  
        "email": "abcdef@gmail.com"  
    }'
    ```

4. **Update Student** [PUT]  
    ```http://127.0.0.1:8000/student/update/{id}```  
    JSON input:  
    ```json
    {  
        "name": "Abc Def",  
        "age": 23,  
        "email": "abcdef@gmail.com"  
    }
    ```  
    curl command:  
    ```bash
    curl -X PUT http://localhost:8000/student/update/2 -d '{  
        "name": "Abc Def",  
        "age": 23,  
        "email": "abcdef@gmail.com"  
    }'
    ```

5. **Delete Student** [DELETE]  
    ```http://127.0.0.1:8000/student/delete/{id}```  
    curl command:  
    ```bash
    curl -X DELETE http://localhost:8000/student/delete/1
    ```

6. **Get Student Summary** [GET]  
    Mandatory to have Ollama along with model llama3 installed as described in How to Run This Project [step 3](#optional).  
    ```http://127.0.0.1:8000/student/summary/{id}```  
    curl command:  
    ```bash
    curl -X GET http://localhost:8000/student/summary/1 | jq
    ```

## How to Remove Ollama

If you no longer require Ollama and wish to uninstall it, follow [**this guide**](https://collabnix.com/how-to-uninstall-ollama/) to completely remove Ollama from your system.
