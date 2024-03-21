package handler

import (
	"encoding/json"
	"net/http"
	logger "testTask_retail/logs"
)

type TempOrderDetails struct {
	OrderId           int      `json:"заказ"`
	ProductId         int      `json:"id"`
	ProductName       string   `json:"Товар"`
	OrderCount        int      `json:"Количество"`
	AdditionalShelves []string `json:"Доп. стеллаж,omitempty"`
}

// @Summary Get Order Details
// @Tags Orders
// @Description Get List of All Order Details
// @Accept json
// @Produce json
// @Success 200 {object} TempOrderDetails
// @Failure 400 {object} Err
// @Failure 404 {object} Err
// @Failure 500 {object} Err
// @Router /api/orderdetails [get]
func (h *Handler) handleGetOrderDetails(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("Handling Get Order Details")

	orderDetails, err := h.service.OrderDetails.GetOrderDetails()
	if err != nil {
		logger.Log.Error("Failed to Get Order Details: ", err.Error())
		NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	groupedOrders := make(map[string][]TempOrderDetails)
	for _, detail := range orderDetails {
		tempDetail := TempOrderDetails{
			OrderId:           detail.OrderId,
			ProductId:         detail.ProductId,
			ProductName:       detail.ProductName,
			OrderCount:        detail.OrderCount,
			AdditionalShelves: detail.AdditionalShelves,
		}
		groupedOrders[detail.ShelfLocation] = append(groupedOrders[detail.ShelfLocation], tempDetail)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(groupedOrders); err != nil {
		logger.Log.Error("Failed to encode response", err.Error())
		NewErrorResponse(w, http.StatusInternalServerError, "failed to encode response")
	}
}
