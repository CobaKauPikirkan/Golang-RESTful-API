package web

type CategoryCreateRequest struct{
	Name string `validate:"required,min=1" json:"name"`
}
