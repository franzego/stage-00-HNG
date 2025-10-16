package internal

import (
	"net/http"

	"github.com/franzego/stage00/pkg"
	"github.com/gin-gonic/gin"
)

type Profile struct {
	Status    string `json:"status"`
	User      `json:"user"`
	Timestamp string `json:"timestamp"`
	Fact      string `json:"fact"`
}

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Stack string `json:"stack"`
}

func GetProfile(c *gin.Context) {
	// var profile Profile
	timeNow := pkg.Timestamp()
	newFact, err := pkg.CatApi()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, &Profile{
		Status: "success",
		User: User{
			Email: "davidenenama@gmail.com",
			Name:  "Enenamah David",
			Stack: "Go/gin",
		},
		Timestamp: timeNow,
		Fact:      newFact,
	})

}
