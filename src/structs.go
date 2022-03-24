package grpt

var (
	Artists     = "https://groupietrackers.herokuapp.com/api/artists"
	Locations   = "https://groupietrackers.herokuapp.com/api/locations"
	Dates       = "https://groupietrackers.herokuapp.com/api/dates"
	RelationURL = "https://groupietrackers.herokuapp.com/api/relation"
)

var API []Artist

type Artist struct {
	Id             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	DatesLocations map[string][]string
}

type Relation struct {
	Index []FoundRelation `json:"index"`
}

type FoundRelation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Error struct {
	Code int
	Msg  string
}

type Search struct {
	Input  string
	Option string
	Result []Artist
}

type Filter struct {
	CD      FilterParams
	Members FilterParams
	FAlbum  FilterParams
	Loc     LocationParams
	Result  []Artist
}
type FilterParams struct {
	isSelected string
	From       string
	To         string
}
type LocationParams struct {
	isSelected string
	Location   string
}
