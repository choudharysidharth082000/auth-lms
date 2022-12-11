package Auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	commons "github.com/sidharthchoudhary/lmsAuth/Commons"
)

func AddProductWish(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	fmt.Println(userId);
	productId := vars["courseId"]
	var response commons.Response
	response = AddWishListController(userId, productId)
	if response.Status == 1 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(commons.Response{Status: 200, Message: "Product is added to wishlist"})
	} else {
		fmt.Println(response.Message)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "Product is not added to wishlist"})
	}
}

// adding the product to the cart
func AddProductCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	productId := vars["productId"]
	var response commons.Response
	response = AddProductToCartController(userId, productId)
	if response.Status == 1 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(commons.Response{Status: 200, Message: "Product is added to cart"})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "Product is not added to cart"})
	}
}
