This is a Golang-based code that demonstrates how you can use the Gin framework to send multiple binary files using Postman. To do this, use the POST method and select 'form-data' as the body type, then store the binaries in a PostgreSQL database as an array of binaries. 
This code shows that you can send multiple files with a single API request and store them in the same array. However, when using the GET method to retrieve the binary files, you can only retrieve one file at a time, which is selected by an index parameter that specifies the file in the array.

Golang/Gin-framework database postgresql+pgadmin/postgresql+Dbeaver
