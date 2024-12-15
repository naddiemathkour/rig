package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	recipes "test/recipes"

	"github.com/gosimple/slug"
)

var (
	RecipeRe       = regexp.MustCompile(`^/recipes/*$`)
	RecipeReWithID = regexp.MustCompile(`^/recipes/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)

func main() {
	// For demo purposes, a memory store will be used as a storage solution
	store := recipes.NewMemStore()
	recipesHandler := NewRecipeHandler(store)

	// Create a multiplexer to interface w/ http requests and match them with
	// URI patterns
	mux := http.NewServeMux()

	// Register URI paths and define which HTTP Handler to execute for the URI
	mux.Handle("/", &homeHandler{})
	mux.Handle("/recipes", recipesHandler)
	mux.Handle("/recipes/", recipesHandler)

	// Call the ListenAndServe function to start your sever, assigning the mux
	// created above to handle requests
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

// Create URI path structs and implement ServeHTTP
type homeHandler struct{}
type recipeHandler struct {
	store recipeStore
}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page\n"))
}

// recipeStore defines an interface for interacting with your storage solution.
// This abstracts the storage implementation, allowing the server to work with
// any storage backend that implements this interface.
type recipeStore interface {
	// Add a new recipe (params: name, Recipe) returns: error
	Add(name string, recipe recipes.Recipe) error
	// Retrieve a stored recipe (params: name) returns: Recipe, error
	Get(name string) (recipes.Recipe, error)
	// Update an existing recipe (params: name, Recipe) returns: error
	Update(name string, recipe recipes.Recipe) error
	// List all recipes in the store (params: none) returns: Recipe map, error
	List() (map[string]recipes.Recipe, error)
	// Delete a recipe from storage (params: name) returns: error
	Remove(name string) error
}

// Handler function that routes requests to their handlers via a switch
func (h *recipeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
		h.CreateRecipe(w, r)
		return

	case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
		h.ListRecipes(w, r)
		return

	case r.Method == http.MethodGet && RecipeReWithID.MatchString(r.URL.Path):
		h.GetRecipe(w, r)
		return

	case r.Method == http.MethodPut && RecipeReWithID.MatchString(r.URL.Path):
		h.UpdateRecipe(w, r)
		return

	case r.Method == http.MethodDelete && RecipeReWithID.MatchString(r.URL.Path):
		h.DeleteRecipe(w, r)
		return

	default:
		return
	}
}

// Correct implementation of handler when using an interface
func NewRecipeHandler(s recipeStore) *recipeHandler {
	return &recipeHandler{
		store: s,
	}
}

// Handler functions to process CRUD operations and business logic
func (h *recipeHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe recipes.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	resourceID := slug.Make(recipe.Name)
	if err := h.store.Add(resourceID, recipe); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func (h *recipeHandler) ListRecipes(w http.ResponseWriter, r *http.Request)  {}
func (h *recipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request)    {}
func (h *recipeHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {}
func (h *recipeHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {}

// Generic error handlers
func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}
