package types

// Money предстовляет собой денежную сумму в минимальных еденицах (центы, копейки, дирамы и е.д.).
type Money int64

// Category  предстовляет собой категорию, в которой был совершён платёж (авто, аптеки, рестораны и т.д.).
type Category string

// Payment предстовляет информацию о платеже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
