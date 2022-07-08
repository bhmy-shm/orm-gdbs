package examples

type ProdModels struct {
	Id         int32   `json:"id" mapstructure:"id,omitempty"`                //商品ID
	ProdName   string  `json:"prodName" mapstructure:"prod_name,omitempty"`   //商品名
	ProdPrice  float32 `json:"prodPrice" mapstructure:"prod_price,omitempty"` //价格
	ProdOnsale bool    `json:"prodOnsale" mapstructure:"prod_onsale,omitempty"`
}

