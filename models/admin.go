package models

import "time"

// 不建立登录模块
// 建立员工相关三张表：员工信息，员工部门和工作记录

type Employee struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	Name         string     `gorm:"not null;size:20" json:"name"`
	Email        string     `gorm:"unique;size:30" json:"email"`
	DepartmentID uint       `gorm:"not null" json:"department_id"` // 外键约束
	Department   Department `gorm:"foreignKey:DepartmentID" json:"department"`
}

// 员工的部门，使用外键关联

type Department struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Company    string `gorm:"not null" json:"company"`
	Level1Dept string `gorm:"not null" json:"level_1_dept"`
	Level2Dept string `gorm:"not null" json:"level_2_dept"`
}

type Records struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	EmployeeID uint      `gorm:"not null" json:"employee_id"`
	Employee   Employee  `gorm:"foreignKey:EmployeeID" json:"employee"`
	Date       time.Time `gorm:"size:20;not null;index" json:"date"`
	ClickOn    time.Time `gorm:"size:10" json:"click_on"`
	ClickOut   time.Time `gorm:"size:10" json:"click_out"`
	OverTime   time.Time `gorm:"size:10" json:"over_time"`
}
