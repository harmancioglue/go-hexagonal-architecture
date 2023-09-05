package entity

import "time"

type Todo struct {
	Id          int32     `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"-"`
	Title       string    `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Description string    `gorm:"column:description;type:varchar(255);not null" json:"description"`
	DueDate     time.Time `gorm:"column:due_date;type:datetime" json:"dueDate"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime" json:"createdAt"`
	Completed   bool      `gorm:"column:completed;type:tinyint(1);not null" json:"completed"`
}

func (*Todo) TableName() string {
	return "todo_task"
}
