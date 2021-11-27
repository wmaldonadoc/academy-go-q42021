package schemas

type BatchSearchingSchema struct {
	Type           string `form:"type" binding:"required,oneof=odd even"`
	Items          int    `form:"items" binding:"required"`
	ItemsPerWorker int    `form:"items_per_worker" binding:"required"`
}
