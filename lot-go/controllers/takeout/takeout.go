package takeout

import (
	"meilian/constants"
	controller "meilian/controllers/base"
	"meilian/models/takeout"
	"meilian/utils/pagination"
	"meilian/vo"
)

type TakeoutController struct {
	controller.BaseController
}

/**
*店内搜索
 */
func (c *TakeoutController) SearchTakeoutInshop() {
	page, err := c.GetInt("page", 1)
	if err != nil {
		c.ResponseJSON(constants.ParamErrorCode, "！", "")
		return
	}
	pageSize, err := c.GetInt("pageSize", 10)
	if err != nil {
		c.ResponseJSON(constants.ParamErrorCode, "！", "")
		return
	}
	keyword := c.GetString("keyword")
	if keyword == "" {
		c.ResponseJSON(constants.ParamErrorCode, "请输入搜索关键字！", "")
		return
	}
	//searchSql := "" +
	//	"SELECT tk.* FROM takeout as tk  WHERE tk.shop_id = ? AND tk.takeout_name like  AND tk.deleted = 0 ";
	var takeouts []takeout.Takeout
	qs := c.O.QueryTable("takeout").
		Filter("ShopId", 46501).
		Filter("TakeoutName__icontains", keyword).
		Filter("Deleted", 0)
	total, err := qs.Count()
	if err != nil {
		c.ResponseJSON(constants.QueryError, "数据获取异常,请重试或联系开发者！", "")
		return
	}
	total, err = pagination.Paginate(qs, page, pageSize, &takeouts)
	//_,err = qs.Limit(10,0).All(&takeouts)
	dataMap := make(map[string]interface{})
	if err != nil {
		c.ResponseJSON(constants.QueryError, "数据获取异常,请重试或联系开发者！", "")
		return
	}
	dataMap["total"] = total
	dataMap["data"] = takeouts
	c.ResponseJSON(constants.SuccessCode, "成功！", dataMap)
}

/*
*店内外卖列表
 */
func (c *TakeoutController) List() {
	page, err := c.GetInt("page", 1)
	if err != nil {
		c.ResponseJSON(constants.ParamErrorCode, "！", "")
		return
	}
	pageSize, err := c.GetInt("page_size", 10)
	if err != nil {
		c.ResponseJSON(constants.ParamErrorCode, "！", "")
		return
	}
	tkCategoryID, err := c.GetInt("category_id", 10)
	if err != nil {
		c.ResponseJSON(constants.ParamErrorCode, "！", "")
		return
	}
	shopID, err := c.GetInt("shop_id", 10)
	if err != nil {
		c.ResponseJSON(constants.ParamErrorCode, "！", "")
		return
	}
	queryForList := `
        SELECT tk.*, sp.id, sp.shop_name, sp.business, sc.category_id
        FROM takeout tk 
        INNER JOIN shop sp 
        ON tk.shop_id = sp.id 
		AND sp.deleted = 0
        INNER JOIN shop_category sc
        ON tk.category_id = sc.category_id 
		AND sc.deleted = 0 
		AND tk.category_id = ?
        WHERE sp.id = ? 
		AND tk.deleted = 0 
        LIMIT ? OFFSET ?
    `

	var results []vo.TakeoutShopCategoryVO
	_, err = c.O.Raw(queryForList, tkCategoryID, shopID, pageSize, page).QueryRows(&results)
	c.ResponseJSON(constants.SuccessCode, "OK", results)

}

/**
*规格
 */
func (c *TakeoutController) Specifications() {

}

/**
*下单
 */
func (c *TakeoutController) Order() {

}
