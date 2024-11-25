package models

import "github.com/jackc/pgx/v5/pgtype"

type CreateUserReq struct {
	UserName     *string `json:"user_name" binding:"required"`
	UserPassword *string `json:"user_password" binding:"required"`
}

type UserResponse struct {
	UserID       int32               `json:"user_id"`
	UserName     *string             `json:"user_name"`
	UserPassword *string             `json:"user_password"`
	UserPhone    *string             `json:"user_phone"`
	UserToken    *string             `json:"user_token"`
	Roles        []*UserRoleResponse `json:"roles"`
}

type UserRoleResponse struct {
	RoleID   int32   `json:"role_id"`
	RoleName *string `json:"role_name"`
}

type CategoryPostReq struct {
	CategoryName string  `json:"category_name" binding:"required"`
	Description  *string `json:"description"`
}

type CategoryUpdateReq struct {
	CategoryName string  `json:"category_name"`
	Description  *string `json:"description"`
}

type CreateCartParams struct {
	CustomerID string   `json:"customer_id"`
	ProductID  int32    `json:"product_id"`
	UnitPrice  *float32 `json:"unit_price"`
	Qty        *int32   `json:"qty"`
}

type CartResponse struct {
	CartId      int                    `json:"cart_id"`
	CustomerID  string                 `json:"customer_id"`
	CompanyName string                 `json:"company_name"`
	Products    []*CartProductResponse `json:"products"`
}

type CartProductResponse struct {
	ProductID   *int16         `json:"product_id"`
	ProductName *string        `json:"product_name"`
	UnitPrice   *float32       `json:"unit_price"`
	Qty         *int32         `json:"qty"`
	Price       pgtype.Numeric `json:"price"`
}

type CreateOrderRequest struct {
	CustomerID     *string     `json:"customer_id" binding:"required"`
	EmployeeID     *int16      `json:"employee_id" binding:"required"`
	OrderDate      pgtype.Date `json:"order_date"`
	RequiredDate   pgtype.Date `json:"required_date"`
	ShippedDate    pgtype.Date `json:"shipped_date"`
	ShipVia        *int16      `json:"ship_via" binding:"required"`
	Freight        *float32    `json:"freight"`
	ShipName       *string     `json:"ship_name"`
	ShipAddress    *string     `json:"ship_address"`
	ShipCity       *string     `json:"ship_city"`
	ShipRegion     *string     `json:"ship_region"`
	ShipPostalCode *string     `json:"ship_postal_code"`
	ShipCountry    *string     `json:"ship_country"`
}
