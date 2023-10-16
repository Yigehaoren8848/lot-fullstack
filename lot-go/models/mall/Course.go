// models/course.go

package models

import (
	"github.com/astaxie/beego/orm"
	"meilian/lib"
)

type Courses struct {
	Id          int    `orm:"auto;pk"`
	ApiCourseid int    `orm:"column(api_courseid);null"`
	CourseName  string `orm:"column(course_name);size(255);null"`
	WatchTimes  int    `orm:"column(watch_times);null"`
	Videourl    string `orm:"column(videourl);type(text);null"`
	AddTime     string `orm:"column(add_time);size(255);null"`
	Imgurl      string `orm:"column(imgurl);size(255);null"`

	Time    lib.Time `orm:"column(time);type(timestamp);auto_now_add"`
	Deleted string   `orm:"column(deleted);size(255);default(0)"`
	Teacher string   `orm:"column(teacher);size(255);null"`

	ApiCoursetype     string  `orm:"column(api_coursetype);size(255);null"`
	Coursetype        string  `orm:"column(coursetype);size(255);null"`
	Name              string  `orm:"column(name);size(255);null"`
	Articleid         string  `orm:"column(articleid);size(255);null"`
	ImageIntroduction string  `orm:"column(image_introduction);size(255);null"`
	VideoIntroduction string  `orm:"column(video_introduction);size(255);null"`
	TextIntroduction  string  `orm:"column(text_introduction);type(text);null"`
	Price             float64 `orm:"column(price);null"`
	Free              string  `orm:"column(free);size(10);default(0)"`
	Orginal           string  `orm:"column(orginal);size(255);default(course)"`
	Imglink           string  `orm:"column(imglink);size(255);null"`
	ArticleContent    string  `orm:"column(article_content);type(text);null"`
	CourseType        string  `orm:"column(course_type);size(255);null"`
	Title             string  `orm:"column(title);size(255);null"`
	Subtitle          string  `orm:"column(subtitle);size(255);null"`
	CategoryList      string  `orm:"column(catagory_list);size(255);null"`
	Belong            string  `orm:"column(belong);size(255);default('')"`
}

func init() {
	// 注册模型
	orm.RegisterModel(new(Courses))
}
