package recipes

// Represents a recipe
type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

// Represents an individual ingredient
type Ingredient struct {
	Name string `json:"name"`
}

// Note:	the 'json' tag in structs will be used to encode and decode structs
//				into JSON when the api is shipping data. Adding the 'json' struct tag
//				defines the name of the field in JSON representation.
