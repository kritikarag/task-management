package controllers

import (
	//"context"
	//"fmt"
	"net/http"
	"task-management/models"
	"time"
	"github.com/gin-gonic/gin"
)

type CreateTaskInput struct {
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
}

type UpdateTaskInput struct {
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
}

func GetTasks(c *gin.Context){
	var tasks []models.Task
	//db := c.MustGet("db").(*gorm.DB)
	models.DB.Find(&tasks)
	c.JSON(http.StatusOK,gin.H{"data": tasks})

}

func CreateTask(c *gin.Context){
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currTime := time.Now();
	task:= models.Task{TaskName: input.TaskName, TaskDetail: input.TaskDetail, Date: currTime.Format("2006-01-02 15:04:05")}
	models.DB.Create(&task)
}

func GetTaskById(c *gin.Context){
	var task models.Task
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	//id := c.Request.URL.Query().Get("id")
	//db := c.MustGet("db").(*gorm.DB) 
	if err:= models.DB.Where("id = ?",id).First(&task).Error;err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Record not found!"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"data":task})
}

func UpdateTask(c *gin.Context){

	var task models.Task
	
	c.Header("Content-Type", "application/json")
	id := c.Param("id")

	if err:= models.DB.Where("id = ?",id).First(&task).Error;err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Record not found!"})
		return
	}

	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedTask models.Task

	updatedTask.TaskName = input.TaskName
	updatedTask.TaskDetail = input.TaskDetail
	updatedTask.Date = time.Now().Format("2006-01-02 15:04:05")

	models.DB.Model(&task).Updates(updatedTask)
	c.JSON(http.StatusOK,gin.H{"data":task})
}

func DeleteTask(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	var task models.Task
	if err := models.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func UpdateTaskName(c *gin.Context) {
	taskID := c.Param("id")

	type OneTask struct {
		TaskName *string `json:"task_name"`
	}

	var task OneTask;

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) 
		return
	}

	stmt, err := models.DB.DB().Prepare("UPDATE tasks SET task_name = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) 
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.TaskName, taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) 
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Waiter updated successfully"})
}

func GetTasksByName(c *gin.Context) {
	taskName := c.Params.ByName("task_name") 

	var task_list []models.Task

	var tasks []models.Task

	models.DB.Find(&tasks)

	for _,task:= range(tasks){
		if (task.TaskName==taskName){
			task_list = append(task_list, task)
		}
	} 
	c.JSON(http.StatusOK,gin.H{"data": task_list})
}