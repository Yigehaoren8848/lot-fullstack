// pagination.go

package pagination

import (
	"github.com/astaxie/beego/orm"
)

func Paginate(query orm.QuerySeter, page int, pageSize int, data interface{}) (int64, error) {
	total, err := query.Count()
	if err != nil {
		return 0, err
	}

	_, err = query.Limit(pageSize, (page-1)*pageSize).All(data)
	if err != nil {
		return 0, err
	}

	return total, nil
}
