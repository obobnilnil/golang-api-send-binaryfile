package main

import (
	"fmt"
	"io"
	"net/http"
	"returnFileMany/database"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func main() {

	db := database.Postgresql()
	defer db.Close()
	fmt.Println("connected db")

	r := gin.Default()

	// create simple route with gin-framework
	r.POST("/upload", func(c *gin.Context) {
		// using gin-framework to sent binary file
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		files := form.File["files"]

		var fileBytesArray [][]byte
		for _, file := range files {
			openedFile, err := file.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			defer openedFile.Close()
			fileBytes, err := io.ReadAll(openedFile)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			fileBytesArray = append(fileBytesArray, fileBytes)
		}

		// _, err = db.Exec(`INSERT INTO "testbinary" (binaryFiles) VALUES ($1)`, pq.Array(fileBytesArray)) // pgAdmin
		_, err = db.Exec(`INSERT INTO "testdb" (files) VALUES ($1)`, pq.Array(fileBytesArray)) // dbReaver

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfully"})
	})

	r.POST("/upload2", func(c *gin.Context) {
		file, err := c.FormFile("file") //
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		openedFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer openedFile.Close()

		fileBytes, err := io.ReadAll(openedFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		_, err = db.Exec(`INSERT INTO "testbinary" (binaryFiles) VALUES ($1)`, pq.Array(fileBytes)) // ใช้ `fileBytes` โดยตรงแทน `pq.Array(fileBytesArray)`

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
	})

	// r.POST("/upload/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	var uploadedFiles [][]byte

	// 	// แก้ไข query เพื่อรับชื่อไฟล์ด้วย
	// 	err := db.QueryRow("SELECT binaryFiles, filename FROM testbinary WHERE id = $1", id).Scan(pq.Array(&uploadedFiles))
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file from the database"})
	// 		return
	// 	}

	// 	// ตรวจสอบว่าไฟล์มีข้อมูลหรือไม่
	// 	if len(uploadedFiles) == 0 || len(uploadedFiles[0]) == 0 {
	// 		c.JSON(http.StatusNotFound, gin.H{"error": "No file found"})
	// 		return
	// 	}

	// 	// ตั้งค่า Content-Type และ Content-Disposition
	// 	fileType := http.DetectContentType(uploadedFiles[0])
	// 	c.Header("Content-Type", fileType)
	// 	c.Header("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", filename))

	// 	// ส่ง binary data กลับไปที่ client
	// 	c.Data(http.StatusOK, fileType, uploadedFiles[0])
	// })
	// r.GET("/getfiles/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	var uploadedFiles [][]byte

	// 	// แก้ไข query เพื่อรับไฟล์ binary
	// 	err := db.QueryRow("SELECT binaryFiles FROM testbinary WHERE id = $1", id).Scan(pq.Array(&uploadedFiles))
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file from the database"})
	// 		return
	// 	}

	// 	// ตรวจสอบว่าไฟล์มีข้อมูลหรือไม่
	// 	if len(uploadedFiles) == 0 || len(uploadedFiles[0]) == 0 {
	// 		c.JSON(http.StatusNotFound, gin.H{"error": "No file found"})
	// 		return
	// 	}

	// 	// ตั้งค่า Content-Type และส่ง binary data กลับไปที่ client
	// 	fileType := http.DetectContentType(uploadedFiles[0])
	// 	c.Header("Content-Type", fileType)
	// 	c.Data(http.StatusOK, fileType, uploadedFiles[0])
	// })
	r.GET("/getfiles/:id", func(c *gin.Context) {
		id := c.Param("id")
		var uploadedFiles [][]byte

		// err := db.QueryRow("SELECT binaryFiles FROM testbinary WHERE id = $1", id).Scan(pq.Array(&uploadedFiles)) // pgadmin
		err := db.QueryRow("SELECT files FROM testdb WHERE id = $1", id).Scan(pq.Array(&uploadedFiles)) // dbReaver
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get files from the database"})
			return
		}

		if len(uploadedFiles) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No files found"})
			return
		}

		// loop all the files
		for i, fileBytes := range uploadedFiles {
			fileType := http.DetectContentType(fileBytes)

			fmt.Printf("File %d MIME type: %s\n", i, fileType)
		}

		fileType := http.DetectContentType(uploadedFiles[0])
		c.Header("Content-Type", fileType)
		c.Data(http.StatusOK, fileType, uploadedFiles[0])
	})
	r.Run(":7777")
}

// func GetUploadedFiles(c *gin.Context) { /// ดีกว่าแต่ต้องแก้ logic ที่ frontend ด้วย อาจจะได้ใช้ในอนาคต
// 	id := c.Param("organizeID")
// 	var uploadedFiles [][]byte

// 	err := db.QueryRow("SELECT org_logo_binary FROM organize_master WHERE org_id = $1", id).Scan(pq.Array(&uploadedFiles))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file from the database"})
// 		return
// 	}

// 	// ตรวจสอบว่าไฟล์มีข้อมูลหรือไม่
// 	if len(uploadedFiles) == 0 || len(uploadedFiles[0]) == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "No file found"})
// 		return
// 	}

// 	// ตั้งค่า Content-Type และ Content-Disposition
// 	fileType := http.DetectContentType(uploadedFiles[0])
// 	c.Header("Content-Type", fileType)
// 	c.Header("Content-Disposition", "inline")

// 	// ส่ง binary data กลับไปที่ client
// 	c.Data(http.StatusOK, fileType, uploadedFiles[0])
// }
