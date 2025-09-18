package routes

import (
	"github.com/gin-gonic/gin"
	"hufeng-code/db"
	"hufeng-code/models"
	"net/http"
	"time"
)

// 之后再拆分

func BindEmployeeSetting(router *gin.Engine) {
	emp := router.Group("/employee")
	emp.POST("/add", EmployeeEntry)
	emp.DELETE("/dimission", EmployeeDiMission)
	emp.PUT("/change", EmployeeUpdate)
	emp.GET("/list", EmployeeList)
}

func BindEmployeeJob(router *gin.Engine) {
	job := router.Group("/job")
	job.POST("/record", JobRecord)
	//	//job.GET("/records", JobRecordsList)
	//	//job.PUT("/checkout", JobCheckOut)
}

func BindEmployeeCompany(router *gin.Engine) {
	company := router.Group("/company")
	company.POST("/department", DepartmentCreate)
	//company.GET("/departments", DepartmentList)
	//company.PUT("/department", DepartmentUpdate)
	//company.DELETE("/department", DepartmentDelete)
}

func EmployeeDiMission(c *gin.Context) {
}

func EmployeeUpdate(c *gin.Context) {
}

func EmployeeEntry(c *gin.Context) {
	var employee models.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee created successfully", "employee": employee})
}
func EmployeeList(c *gin.Context) {
	
}

// JobRecord 添加工作记录
func JobRecord(c *gin.Context) {
	var record models.Records

	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 设置打卡时间
	record.ClickOn = time.Now()

	if err := db.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job record created successfully", "record": record})
}

func DepartmentCreate(c *gin.Context) {
	var department models.Department

	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create department"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Department created successfully", "department": department})
}
