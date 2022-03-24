package grpt

import (
	"errors"
	"html/template"
	"math"
	"net/http"
	"strconv"
	"strings"
)

var templates, _ = template.ParseGlob("templates/*.html")

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {

	if err := templates.ExecuteTemplate(w, tmpl, data); err != nil {
		return errors.New("Can not execute " + tmpl + " template")
	}
	return nil
}

func (search *Search) SearchArtists() error {

	if i := strings.Index(search.Input, " -->"); i != -1 {
		switch search.Input[i+5:] {
		case "Artist/Band":
			search.Option = "Name"
		case "Member":
			search.Option = "Members"
		case "Creation Date":
			search.Option = "CreationDate"
		case "First Album":
			search.Option = "FirstAlbum"
		case "Location":
			search.Option = "Locations"
		case "Concert date":
			search.Option = "Dates"
		}
		search.Input = search.Input[:i]
	}

	if search.Input == "" {
		return errors.New("400:Empty search input")
	}

	for i := range API {

		switch search.Option {
		case "Name":
			if strings.Contains(API[i].Name, search.Input) {
				search.Result = append(search.Result, API[i])
			}
		case "Members":
			for _, member := range API[i].Members {
				if strings.Contains(member, search.Input) {
					search.Result = append(search.Result, API[i])
					break
				}
			}
		case "CreationDate":
			if strings.Contains(strconv.Itoa(API[i].CreationDate), search.Input) {
				search.Result = append(search.Result, API[i])
			}
		case "FirstAlbum":
			if strings.Contains(API[i].FirstAlbum, search.Input) {
				search.Result = append(search.Result, API[i])
			}
		case "Locations":
			for key := range API[i].DatesLocations {
				if strings.Contains(key, search.Input) {
					search.Result = append(search.Result, API[i])
					break
				}
			}
		case "Dates":
			var flag bool
			for _, value := range API[i].DatesLocations {
				for _, dates := range value {
					if strings.Contains(dates, search.Input) && !flag {
						search.Result = append(search.Result, API[i])
						flag = true
						break
					}
				}
			}
		default:
			return errors.New("400:Error search option")
		}
	}
	return nil
}

func (fp *FilterParams) convertParams() (int, int, error) {

	if fp.From == "" {
		fp.From = "0"
	}
	from, err := strconv.Atoi(fp.From)
	if err != nil || from < 0 {
		return -1, -1, err
	}

	if fp.To == "" {
		fp.To = "0"
	}
	to, err := strconv.Atoi(fp.To)
	if err != nil {
		return -1, -1, err
	}
	if to == 0 {
		to = math.MaxInt32
	}
	return from, to, nil
}

func (f *Filter) FilterArtists() error {

	for index := range API {

		from, to, err := f.CD.convertParams()
		if err != nil || from < 0 || to < 0 {
			return errors.New("atoiError")
		}

		if (API[index].CreationDate < from || API[index].CreationDate > to) && f.CD.isSelected == "on" {
			continue
		}

		from, to, err = f.Members.convertParams()
		if err != nil || from < 0 || to < 0 {
			return errors.New("atoiError")
		}
		if (len(API[index].Members) < from || len(API[index].Members) > to) && f.Members.isSelected == "on" {
			continue
		}

		from, to, err = f.FAlbum.convertParams()
		if err != nil || from < 0 || to < 0 {
			return errors.New("atoiError")
		}
		falbum, _ := strconv.Atoi(API[index].FirstAlbum[len(API[index].FirstAlbum)-4:])
		if (falbum < from || falbum > to) && f.FAlbum.isSelected == "on" {
			continue
		}

		if f.Loc.isSelected == "on" {
			ctr := 0
			for key := range API[index].DatesLocations {
				newLoc := strings.Replace(f.Loc.Location, ", ", "-", 1)
				newLoc = strings.Replace(newLoc, " ", "_", -1)
				if strings.Contains(key, strings.ToLower(newLoc)) {
					ctr++
				}
			}
			if ctr == 0 {
				continue
			}
		}
		f.Result = append(f.Result, API[index])
	}
	return nil
}
