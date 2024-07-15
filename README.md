<div align="center">
    <img src="https://miro.medium.com/v2/resize:fit:921/1*HtCjHzGwf6iWNqXu5Cndsg.png" alt="Gin" width="600"/>
</div>
<H1>ReadME</H1>
This is a Golang-based code that demonstrates how you can use the Gin framework to send multiple binary files using Postman. To do this,
use the POST method and select 'form-data' as the body type, then store the binaries in a PostgreSQL database as an array of binaries.

<div align="center">
    <img src="https://github.com/user-attachments/assets/5bcbac5e-2642-4344-92cf-4c7af91949b8" alt="binaryFile" width="150"/>
    <img src="https://github.com/user-attachments/assets/a743dbc3-1398-44ad-a788-a096a7c0ecab" alt="some_arrow" width="150"/>
    <img src="https://www.somkiat.cc/wp-content/uploads/2024/04/postgresql-data-01.png" alt="postgresql" width="350"/>
</div>


This code shows that you can send multiple files with a single API request and store them in the same array. However, 
when using the GET method to retrieve the binary files, you can only retrieve one file at a time, which is selected by an index parameter that specifies the file in the array.

Golang/Gin-framework database postgresql+pgadmin/postgresql+Dbeaver


