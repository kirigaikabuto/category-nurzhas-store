package category_nurzhas_store

import (
	"errors"
	"github.com/google/uuid"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type CategoryService interface {
	CreateCategory(cmd *CreateCategoryCommand) (*Category, error)
	UpdateCategory(cmd *UpdateCategoryCommand) (*Category, error)
	ListCategory(cmd *ListCategoryCommand) ([]Category, error)
	GetCategory(cmd *GetCategoryCommand) (*Category, error)
	DeleteCategory(cmd *DeleteCategoryCommand) error
}

type categoryService struct {
	categoryStore CategoryStore
}

func NewCategoryService(cStore CategoryStore) CategoryService {
	return &categoryService{categoryStore: cStore}
}

func (c *categoryService) CreateCategory(cmd *CreateCategoryCommand) (*Category, error) {
	if !setdata_common.IsCategoryExist(cmd.CategoryType) {
		return nil, errors.New("Incorrect category type")
	}
	category := &Category{
		Id:               uuid.New().String(),
		Name:             cmd.Name,
		SmallDescription: cmd.SmallDescription,
		BigDescription:   cmd.BigDescription,
		ImageUrl:         cmd.ImageUrl,
		CategoryType:     setdata_common.ToCategoryType(cmd.CategoryType),
	}
	return c.categoryStore.CreateCategory(category)
}

func (c *categoryService) UpdateCategory(cmd *UpdateCategoryCommand) (*Category, error) {
	categoryUpdate := &CategoryUpdate{Id: cmd.Id}
	oldCategory, err := c.categoryStore.GetCategory(cmd.Id)
	if err != nil {
		return nil, err
	}
	if cmd.Name != "" && oldCategory.Name != cmd.Name {
		categoryUpdate.Name = &cmd.Name
	}
	if cmd.SmallDescription != "" && oldCategory.SmallDescription != cmd.SmallDescription {
		categoryUpdate.SmallDescription = &cmd.SmallDescription
	}
	if cmd.BigDescription != "" && oldCategory.BigDescription != cmd.BigDescription {
		categoryUpdate.BigDescription = &cmd.BigDescription
	}
	if cmd.ImageUrl != "" && oldCategory.ImageUrl != cmd.ImageUrl {
		categoryUpdate.ImageUrl = &cmd.ImageUrl
	}
	if cmd.CategoryType != "" && oldCategory.CategoryType.ToString() != cmd.CategoryType {
		if !setdata_common.IsCategoryExist(cmd.CategoryType) {
			return nil, errors.New("Incorrect category type")
		}
		catType := setdata_common.ToCategoryType(cmd.CategoryType)
		categoryUpdate.CategoryType = &catType
	}
	return c.categoryStore.UpdateCategory(categoryUpdate)
}

func (c *categoryService) ListCategory(cmd *ListCategoryCommand) ([]Category, error) {
	return c.categoryStore.ListCategory()
}

func (c *categoryService) GetCategory(cmd *GetCategoryCommand) (*Category, error) {
	return c.categoryStore.GetCategory(cmd.Id)
}

func (c *categoryService) DeleteCategory(cmd *DeleteCategoryCommand) error {
	return c.categoryStore.DeleteCategory(cmd.Id)
}
