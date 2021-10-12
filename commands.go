package category_nurzhas_store

import "bytes"

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

type UploadPricesFileCommand struct {
	Name string        `json:"name"`
	File *bytes.Buffer `json:"file" form:"file"`
}

func (cmd *UploadPricesFileCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(CategoryService).UploadPricesFile(cmd)
}

type UploadPricesFileResponse struct {
	Name    string `json:"name"`
	FileUrl string `json:"file_url"`
}

type GetPricesFileCommand struct {
	Name string `json:"name"`
}

func (cmd *GetPricesFileCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(CategoryService).GetPricesFile(cmd)
}

type GetPricesFileResponse struct {
	FileUrl string `json:"file_url"`
}

type UploadCategoryImageCommand struct {
	File *bytes.Buffer
	Id   string `json:"id"`
}

func (cmd *UploadCategoryImageCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(CategoryService).UploadCategoryImage(cmd)
}

type CreateUserCommand struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (cmd *CreateUserCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(UserService).CreateUser(cmd)
}

type UpdateUserCommand struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (cmd *UpdateUserCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(UserService).UpdateUser(cmd)
}

type DeleteUserCommand struct {
	Id string `json:"id"`
}

func (cmd *DeleteUserCommand) Exec(svc interface{}) (interface{}, error) {
	return nil, svc.(UserService).DeleteUser(cmd)
}

type GetUserCommand struct {
	Id string `json:"id"`
}

func (cmd *GetUserCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(UserService).GetUser(cmd)
}

type ListUserCommand struct {
}

func (cmd *ListUserCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(UserService).ListUser(cmd)
}

type GetUserByUsernameAndPassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cmd *GetUserByUsernameAndPassword) Exec(svc interface{}) (interface{}, error) {
	return svc.(UserService).GetUserByUsernameAndPassword(cmd)
}