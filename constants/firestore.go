package constants

type FirestoreCollection struct {
	Clients   string
	Tokens    string
	AuthCodes string
}

var FirestoreCollections = FirestoreCollection{
	Clients:   "clients",
	Tokens:    "tokens",
	AuthCodes: "auth_codes",
}
