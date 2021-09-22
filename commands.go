package category_nurzhas_store

type CreateCategoryCommand struct {
	Name             string `json:"name"`
	SmallDescription string `json:"small_description"`
	BigDescription   string `json:"big_description"`
	ImageUrl         string `json:"image_url"`
	CategoryType     string `json:"category_type"`
}

func (cmd *CreateCategoryCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(CategoryService).CreateCategory(cmd)
}

type UpdateCategoryCommand struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	SmallDescription string `json:"small_description"`
	BigDescription   string `json:"big_description"`
	ImageUrl         string `json:"image_url"`
	CategoryType     string `json:"category_type"`
}

func (cmd *UpdateCategoryCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(CategoryService).UpdateCategory(cmd)
}

type GetCategoryCommand struct {
	Id string `json:"id"`
}

func (cmd *GetCategoryCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(CategoryService).GetCategory(cmd)
}

type ListCategoryCommand struct {
}

func (cmd *ListCategoryCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(CategoryService).ListCategory(cmd)
}

type DeleteCategoryCommand struct {
	Id string `json:"id"`
}

func (cmd *DeleteCategoryCommand) Exec(svc interface{}) (interface{}, error) {
	return nil, svc.(CategoryService).DeleteCategory(cmd)
}
