package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var db *gorm.DB

func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:joleif@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&noteModel{})
}

func main() {

	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins: "http://localhost:8080",
		ValidateHeaders: false,
		Credentials: true,
	}))

	v1 := router.Group("/api/v1/notes")
	{
		v1.POST("/", createNote)
		v1.GET("/", fetchAllNotes)
		v1.GET("/:id", fetchSingleNote)
		v1.PUT("/:id", updateNote)
		v1.DELETE("/:id", deleteNote)
	}
	router.Run(":3000")

}

type (
	//noteModel represents a Note
	noteModel struct {
		gorm.Model
		Title string `json:"title"`
		Text string `json:"text"`
	}

	//noteModel represents a Note
	publicNote struct {
		ID uint `json:"id"`
		Title string `json:"title"`
		Text string `json:"text"`
	}

)

// create new Note
func createNote(c *gin.Context)  {
	note := noteModel{Title: c.PostForm("title"), Text: c.PostForm("text")}
	db.Save(&note)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Note successfully created!", "resourceId": note.ID})

}

//list all notes
func fetchAllNotes(c *gin.Context)  {
	var notes []noteModel
	var _notes []publicNote

	db.Find(&notes)

	if len(notes) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Note found!"})
		return
	}

	for _, item := range notes{
		_notes = append(_notes, publicNote{ID: item.ID, Title: item.Title, Text: item.Text})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _notes})
}

//fetch single note
func fetchSingleNote(c *gin.Context) {
	var note noteModel
	noteID := c.Param("id")

	db.First(&note, noteID)

	if note.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Note not found!"})
		return
	}

	_note := publicNote{ID: note.ID, Title: note.Title}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _note})
}

//update note
func updateNote(c *gin.Context) {
	var note noteModel
	noteID := c.Param("id")

	db.First(&note, noteID)

	if note.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No note found!"})
		return
	}

	db.Model(&note).Update("title", c.PostForm("title"))
	db.Model(&note).Update("text", c.PostForm("text"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Note updated successfully!"})
}

//delete note
func deleteNote(c *gin.Context) {
	var note noteModel
	noteID := c.Param("id")

	db.First(&note, noteID)

	if note.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No note found!"})
	}

	db.Delete(&note)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Note deleted successfully!"})

}