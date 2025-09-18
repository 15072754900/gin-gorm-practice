package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hufeng-code/db"
	"hufeng-code/models"
	"log"
	"net/http"
)

// 之后再拆分

func BindEmployeeSetting(router *gin.Engine) {
	emp := router.Group("/employee")
	emp.POST("/add", EmployeeEntry)
	emp.DELETE("/dimission", EmployeeDiMission)
	emp.PUT("/change", EmployeeUpdate)
	emp.PATCH("/change/:name", EmployeeUpdateByEmail)
	emp.GET("/list", EmployeeList)
	emp.GET("/:name", EmployeeMsgByName)
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
	employeeId := c.Request.URL.Query()["employee_id"]
	emp := models.Employee{}
	_ = db.DB.Where("id = ?", employeeId).First(&emp)
	_ = db.DB.Delete(emp)
	c.JSON(http.StatusOK, fmt.Sprintf("删除成功：%s", emp.Name))
}

func EmployeeUpdate(c *gin.Context) {
	emp := models.Employee{}
	if err := c.Bind(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("输入有误：与原始结构体不符，%v", err)
	}
	if rows := db.DB.Preload("Department").Save(&emp); rows.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据修改失败"})
		log.Printf("数据修改错误：%v", rows.Error)
	}
	c.JSON(http.StatusOK, gin.H{"修改字段成功": emp})
}

func EmployeeUpdateByEmail(c *gin.Context) {
	emp := models.Employee{}
	name := c.Param("name")
	db.DB.Where("name=?", name).Find(&emp)
	data := struct {
		Email string
	}{}
	_ = c.Bind(&data)
	if rows := db.DB.Model(&emp).Where("name=?", name).Update("email", data.Email); rows.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据修改失败"})
		log.Printf("数据修改错误：%v", rows.Error)
	}
	db.DB.Preload("Department").Where("name=?", name).First(&emp)
	c.JSON(http.StatusOK, gin.H{"修改字段后": emp})
}

func EmployeeEntry(c *gin.Context) {
	var employee models.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "职员插入失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "职员导入成功", "employee": employee})
}
func EmployeeList(c *gin.Context) {
	var empList []models.Employee
	rows := db.DB.Preload("Department").Find(&empList)
	if rows.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据查看失败"})
		log.Printf("查询数据错误：%v", rows.Error)
	}
	c.JSON(http.StatusOK, gin.H{"用户数据：": empList})
}

func EmployeeMsgByName(c *gin.Context) {
	empName := c.Param("name")
	emp := models.Employee{}
	rows := db.DB.Where("name=?", empName).Find(&emp)
	if rows.Error != nil {
		log.Printf("查询数据错误：%v", rows.Error)
	}
	db.DB.Preload("Department").Find(&emp, emp.ID)
	c.JSON(http.StatusOK, gin.H{"用户数据：": emp})
}

// JobRecord 添加工作记录
func JobRecord(c *gin.Context) {

}

func DepartmentCreate(c *gin.Context) {
	var department models.Department

	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建部门失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "公司部门创建成功", "department": department})
}
