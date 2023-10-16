package controllers

import (
	controller "meilian/controllers/base"
	models "meilian/models/mall"
)

type CourseController struct {
	controller.BaseController
}

// Get 方法用于处理查询所有课程数据的请求
func (c *CourseController) GetAllCourses() {
	var courses []models.Courses

	_, err := c.O.QueryTable("courses").All(&courses, "Id", "Title", "Subtitle", "Time")
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = courses
	}
	c.ServeJSON()
}
