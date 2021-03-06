package main

import( 
  "github.com/gin-gonic/gin"
  "net/http"
)
func index(c *gin.Context){
  c.HTML(http.StatusOK, "index.html", gin.H{})
}

func projects(c *gin.Context){
  c.HTML(http.StatusOK,"projects.html",gin.H{})
}

func about(c *gin.Context){
  c.HTML(http.StatusOK,"about.html",gin.H{})
}

func main(){
  r := gin.Default()
  r.Static("/assets", "./assets")
  r.Static("/styles","./styles")
  r.LoadHTMLGlob("templates/*")
  r.GET("/",index)
  r.GET("/projects",projects)
  r.GET("/about",about)
  r.Run()
}
