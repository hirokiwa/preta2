package query

import "hackz.com/m/v2/graph/model"

type DiaryQuery interface {
	FindAll() ([]*model.Diary)
}