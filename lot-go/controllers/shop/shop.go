package shop

import (
	"encoding/json"
	"meilian/constants"
	controller "meilian/controllers/base"
	"meilian/models/shop"
)

type ShopController struct {
	controller.BaseController
}

/*
*主页查询商家列表并根据距离排序
 */
func (c *ShopController) ShopList() {

	// 构建原始SQL语句
	sql := `
		SELECT
  		id,
  		distance,
		shop_name,
		address,
		lng,
		lat,
		bg_image_id,
		module,
		user_id
		FROM
  		(
    	SELECT
     	 id,
				shop_name,
				address,
				user_id,
				bg_image_id,
				lng,
				module,
				lat,
      ST_Distance_Sphere(
        POINT(114.643700, 37.850790),
        location
      ) AS distance
    FROM
      shop limit 10000
  ) AS distances
WHERE
  distance <= 40000
ORDER BY
  distance limit 10 offset 0;

	`

	// 设置固定的最大距离（以米为单位）

	// 执行查询并获取结果
	var shops []shop.Shop
	_, err := c.O.Raw(sql).QueryRows(&shops)
	if err != nil {
		c.ResponseJSON(constants.QueryError, "系统异常", "")
		return
	} else {
		c.ResponseJSON(constants.SuccessCode, "ok", shops)
	}

}

/**
*成为商家
 */
func (c *ShopController) AddShop() {
	shop := shop.Shop{}
	shopBody := c.Ctx.Input.RequestBody
	json.Unmarshal(shopBody, &shop)
	/**
	*校验必要参数
	 */
	if shop.Address == "" {
		c.ResponseJSON(constants.ParamErrorCode, "地址不可为空,请返回重试。", "")
		return
	}
	if shop.ShopName == "" {
		c.ResponseJSON(constants.ParamErrorCode, "商家名称不可为空,请返回重试。", "")
		return
	}
	if shop.Lat == "" || shop.Lng == "" {
		c.ResponseJSON(constants.ParamErrorCode, "位置信息获取失败,请返回重试。", "")
		return
	}

	shop.Business = 1
	_, err := c.O.Insert(shop)
	if err != nil {
		c.ResponseJSON(constants.InsertError, "网络异常，请联系管理员！", "")
		return
	}
	c.ResponseJSON(constants.SuccessCode, "成功", shop)
	return
}

/**
*店内搜索
 */

func (c *ShopController) SearchInshop() {
	keywords := c.GetString("keyword")
	if keywords == "" {
		c.ResponseJSON(constants.ParamErrorCode, "请输入搜索关键字！", "")
		return
	}
	//searchSql := "" +
	//	"SELECT * FROM shop ";
}

//func (c *ShopController) Order
