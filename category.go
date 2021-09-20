package category_nurzhas_store

type Category struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	SmallDescription string `json:"small_description"`
	BigDescription   string `json:"big_description"`
	ImageUrl         string `json:"image_url"`
}

type CategoryUpdate struct {
	Id               string  `json:"id"`
	Name             *string `json:"name"`
	SmallDescription *string `json:"small_description"`
	BigDescription   *string `json:"big_description"`
	ImageUrl         *string `json:"image_url"`
}
