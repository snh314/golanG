package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("Can't find the product")
	ErrCantDecodeProducts = errors.New("Can't find the product")
	ErrUserIdISNotValid   = errors.New("This user is not Valid")
	ErrCantUpdateUser     = errors.New("Can't add this product to the cart")
	ErrCantRemoveItemCart = errors.New("Can't remove this product to the cart")
	ErrCantGetItem        = errors.New("was unable to get item from the cart")
	ErrcantBuycartitem    = errors.New("Can't update the purchase")
)

func AddProductToCart() {

}
func RemoveCartItem() {

}
func BuyItemFromCart() {

}
func InstantBuyer() {

}
