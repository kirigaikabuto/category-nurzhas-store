package category_nurzhas_store

import setdata_common "github.com/kirigaikabuto/setdata-common"

type Category struct {
	Id               string                      `json:"id"`
	Name             string                      `json:"name"`
	SmallDescription string                      `json:"small_description"`
	BigDescription   string                      `json:"big_description"`
	ImageUrl         string                      `json:"image_url"`
	CategoryType     setdata_common.CategoryType `json:"category_type"`
}

type CategoryUpdate struct {
	Id               string                       `json:"id"`
	Name             *string                      `json:"name"`
	SmallDescription *string                      `json:"small_description"`
	BigDescription   *string                      `json:"big_description"`
	ImageUrl         *string                      `json:"image_url"`
	CategoryType     *setdata_common.CategoryType `json:"category_type"`
}
