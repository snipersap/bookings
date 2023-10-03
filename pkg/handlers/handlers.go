package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/snipersap/bookings/pkg/config"
	"github.com/snipersap/bookings/pkg/models"
	"github.com/snipersap/bookings/pkg/render"
)

const (
	homeTpl                = "home.page.tmpl"
	aboutTpl               = "about.page.tmpl"
	contactTpl             = "contact.page.tmpl"
	makeReservationTpl     = "make-reservation.page.tmpl"
	roomFamilyFTpl         = "room.familyfirst.page.tmpl"
	roomRomanticGetawayTpl = "room.romanticgetaway.page.tmpl"
	roomRoyalGardenTpl     = "room.royalgarden.page.tmpl"
	roomYourLadyTpl        = "room.yourlady.page.tmpl"
	roomYourMajestyTpl     = "room.yourmajesty.page.tmpl"
	SearchAvailabilityTpl  = "search-availability.page.tmpl"
)

// Repo stores the pointer to a Repository
var Repo *Repository

// Repository defines the structure to store the app config
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository with the app config and returns a pointer to the new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{a}
}

// SetRepo initializes the reference to a Repository
func SetRepo(r *Repository) {
	Repo = r
}

// Home handles the home page URL or /
func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	rp.setIPToSession(r)

	if !isURLPathHome(r) {
		http.NotFound(w, r)
		return
	}
	err := render.RenderTemplate(w, r, homeTpl, &models.TemplateData{})
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		log.Println("Rendered Home")
	}
}

// setIPToSession add remote IP address into session
func (rp *Repository) setIPToSession(r *http.Request) {
	rp.App.Session.Put(r.Context(), "remote_Ip", r.RemoteAddr)
}

// isURLPathHome checks for the path '/' in the URL and returns true if path is '/'
// This is used to avoid duplicate requests because of favicon
func isURLPathHome(r *http.Request) bool {
	return r.URL.Path == "/"
}

// About handles the about page with URL /about
func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {

	td := rp.dynContentAbout(r)
	err := render.RenderTemplate(w, r, aboutTpl, td) //Using .html notation to use Emmet abbreviations in VS Code
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		log.Println("Rendered About")
	}
}

// dynContentAbout returns the templated data filled with dynamic content for the about page
func (rp *Repository) dynContentAbout(r *http.Request) *models.TemplateData {
	sMap := make(map[string]string)
	sMap["About me"] = "Software Engineering Manager from Germany"
	sMap["remote_Ip"] = rp.App.Session.GetString(r.Context(), "remote_Ip")
	td := models.TemplateData{StringMap: sMap}
	return &td
}

// Contact handles the about page with URL /about
func (rp *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}
	err := render.RenderTemplate(w, r, contactTpl, td)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		log.Println("Rendered Contact")
	}
}

// MakeReservation handles the about page with URL /about
func (rp *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}
	err := render.RenderTemplate(w, r, makeReservationTpl, td)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		log.Println("Rendered Make Reservation")
	}
}

// RoomFamilyFirst handles the about page with URL /about
func (rp *Repository) RoomFamilyFirst(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}
	err := render.RenderTemplate(w, r, roomFamilyFTpl, td)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		log.Println("Rendered Room Family First")
	}
}

// RoomRomanticGetaway handles the about page with URL /about
func (rp *Repository) RoomRomanticGetaway(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}
	err := render.RenderTemplate(w, r, roomRomanticGetawayTpl, td)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		log.Println("Rendered Room Romantic Getaway")
	}
}

// RoomRoyalGarden handles the about page with URL /about
func (rp *Repository) RoomRoyalGarden(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}
	err := render.RenderTemplate(w, r, roomRoyalGardenTpl, td)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		log.Println("Rendered Room Royal Garden")
	}
}

// RoomYourLady handles the about page with URL /about
func (rp *Repository) RoomYourLady(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}
	err := render.RenderTemplate(w, r, roomYourLadyTpl, td)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		log.Println("Rendered Room Your Lady")
	}
}

// RoomYourMajesty handles the about page with URL /about
func (rp *Repository) RoomYourMajesty(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}
	err := render.RenderTemplate(w, r, roomYourMajestyTpl, td)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		log.Println("Rendered Room Your Majesty")
	}
}

// SearchAvailability handles the about page with URL /about
func (rp *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}
	err := render.RenderTemplate(w, r, SearchAvailabilityTpl, td)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		log.Println("Rendered Search Availability")
	}
}

// PostSearchAvailability handles the about page with URL /about
func (rp *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	checkinDate := r.Form.Get("checkInDate")
	checkoutDate := r.Form.Get("checkOutDate")
	// w.Write([]byte("From " + checkinDate + " to " + checkoutDate))
	w.Write([]byte(fmt.Sprintf("Check-in date is %s and Check-out date is %s", checkinDate, checkoutDate)))
}
