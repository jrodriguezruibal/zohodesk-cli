package models

type Product struct {
	ID            string `json:"id" yaml:"id"`
	Name          string `json:"name" yaml:"name"`
	Description   string `json:"description,omitempty" yaml:"description,omitempty"`
	ProductCode   string `json:"productCode,omitempty" yaml:"productCode,omitempty"`
	Category      string `json:"category,omitempty" yaml:"category,omitempty"`
	IsActive      bool   `json:"isActive" yaml:"isActive"`
	CreatedTime   string `json:"createdTime,omitempty" yaml:"createdTime,omitempty"`
	ModifiedTime  string `json:"modifiedTime,omitempty" yaml:"modifiedTime,omitempty"`
}

type ProductListResponse struct {
	Data []Product `json:"data"`
}

type Account struct {
	ID            string `json:"id" yaml:"id"`
	Name          string `json:"name" yaml:"name"`
	Email         string `json:"email,omitempty" yaml:"email,omitempty"`
	Phone         string `json:"phone,omitempty" yaml:"phone,omitempty"`
	Website       string `json:"website,omitempty" yaml:"website,omitempty"`
	Type          string `json:"type,omitempty" yaml:"type,omitempty"`
	Industry      string `json:"industry,omitempty" yaml:"industry,omitempty"`
	OwnerID       string `json:"ownerId,omitempty" yaml:"ownerId,omitempty"`
	OwnerName     string `json:"ownerName,omitempty" yaml:"ownerName,omitempty"`
	IsActive      bool   `json:"isActive" yaml:"isActive"`
	CreatedTime   string `json:"createdTime,omitempty" yaml:"createdTime,omitempty"`
	ModifiedTime  string `json:"modifiedTime,omitempty" yaml:"modifiedTime,omitempty"`
}

type AccountListResponse struct {
	Data []Account `json:"data"`
}

type AccountCreateRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Website   string `json:"website,omitempty"`
	Type      string `json:"type,omitempty"`
	Industry  string `json:"industry,omitempty"`
	OwnerID   string `json:"ownerId,omitempty"`
}

type AccountUpdateRequest struct {
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Website   string `json:"website,omitempty"`
	Type      string `json:"type,omitempty"`
	Industry  string `json:"industry,omitempty"`
	OwnerID   string `json:"ownerId,omitempty"`
	IsActive  bool   `json:"isActive,omitempty"`
}