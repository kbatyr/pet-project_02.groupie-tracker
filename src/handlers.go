package grpt

import (
	"fmt"
	"net/http"
	"strconv"
)

// The main page handler
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if err := MethodChecker(w, r); err != nil {
		fmt.Println(err)
		ErrPrint(w, &Error{
			Code: http.StatusMethodNotAllowed,
			Msg:  http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	if r.URL.Path != "/" {
		ErrPrint(w, &Error{
			Code: http.StatusNotFound,
			Msg:  http.StatusText(http.StatusNotFound),
		})
		return
	}

	if err := DecodeAPI(GetRequestApi(Artists), &API); err != nil {
		ErrPrint(w, &Error{
			Code: http.StatusInternalServerError,
			Msg:  http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	tempRel := &Relation{}
	if err := DecodeAPI(GetRequestApi(RelationURL), &tempRel); err != nil {
		ErrPrint(w, &Error{
			Code: http.StatusInternalServerError,
			Msg:  http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	for i, rel := range tempRel.Index {
		API[i].DatesLocations = rel.DatesLocations
	}

	if err := RenderTemplate(w, "index", API); err != nil {
		fmt.Println(err)
		ErrPrint(w, &Error{
			Code: http.StatusInternalServerError,
			Msg:  http.StatusText(http.StatusInternalServerError),
		})
		return
	}
}

// Artist's personal page handler
func ArtistInfoHandler(w http.ResponseWriter, r *http.Request) {

	if err := MethodChecker(w, r); err != nil {
		fmt.Println(err)
		ErrPrint(w, &Error{
			Code: http.StatusMethodNotAllowed,
			Msg:  http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	if r.URL.Path[9:] == "" {
		ErrPrint(w, &Error{
			Code: http.StatusBadRequest,
			Msg:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	ArtistID, _ := strconv.Atoi(r.URL.Path[9:])

	//checking for validation of URL address
	if ArtistID < 1 || ArtistID > 52 {
		ErrPrint(w, &Error{
			Code: http.StatusNotFound,
			Msg:  http.StatusText(http.StatusNotFound),
		})
		return
	}

	// Executing of personal artist's info
	for _, id := range API {
		if ArtistID == int(id.Id) {

			if err := RenderTemplate(w, "artist", API[ArtistID-1]); err != nil {
				fmt.Println(err)
				ErrPrint(w, &Error{
					Code: http.StatusInternalServerError,
					Msg:  http.StatusText(http.StatusInternalServerError),
				})
				return
			}
			break
		}
	}
}

// Search bar handler
func SearchHandler(w http.ResponseWriter, r *http.Request) {

	if err := MethodChecker(w, r); err != nil {
		fmt.Println(err)
		ErrPrint(w, &Error{
			Code: http.StatusMethodNotAllowed,
			Msg:  http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	search := &Search{Input: r.FormValue("sinput"), Option: r.FormValue("soption")}

	if err := search.SearchArtists(); err != nil {
		fmt.Println(err)
		ErrPrint(w, &Error{
			Code: http.StatusBadRequest,
			Msg:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if err := RenderTemplate(w, "index", search.Result); err != nil {
		fmt.Println(err)
		ErrPrint(w, &Error{
			Code: http.StatusInternalServerError,
			Msg:  http.StatusText(http.StatusInternalServerError),
		})
		return
	}
}

func FilterHandler(w http.ResponseWriter, r *http.Request) {

	if err := MethodChecker(w, r); err != nil {
		fmt.Println(err)
		ErrPrint(w, &Error{
			Code: http.StatusMethodNotAllowed,
			Msg:  http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	filter := &Filter{
		CD: FilterParams{
			isSelected: r.FormValue("cdate"),
			From:       r.FormValue("cdate[from]"),
			To:         r.FormValue("cdate[to]"),
		},
		Members: FilterParams{
			isSelected: r.FormValue("members"),
			From:       r.FormValue("members[from]"),
			To:         r.FormValue("members[to]"),
		},
		FAlbum: FilterParams{
			isSelected: r.FormValue("falbum"),
			From:       r.FormValue("falbum[from]"),
			To:         r.FormValue("falbum[to]"),
		},
		Loc: LocationParams{
			isSelected: r.FormValue("locations"),
			Location:   r.FormValue("location"),
		},
	}

	if err := filter.FilterArtists(); err != nil {
		fmt.Println(err)
		ErrPrint(w, &Error{
			Code: http.StatusBadRequest,
			Msg:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if err := RenderTemplate(w, "index", filter.Result); err != nil {
		fmt.Println(err)
		ErrPrint(w, &Error{
			Code: http.StatusInternalServerError,
			Msg:  http.StatusText(http.StatusInternalServerError),
		})
		return
	}
}
