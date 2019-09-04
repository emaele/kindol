package types

//OffersURL è il link all'offerta lampo per kindle store
const OffersURL = "https://www.amazon.it/Offerta-Lampo-Kindle/b?ie=UTF8&node=5689487031"

//Deal è la struttura del libro in offerta lampo
type Deal struct {
	Title  string
	Author string
	Cover  string
	Price  string
	Link   string
}
