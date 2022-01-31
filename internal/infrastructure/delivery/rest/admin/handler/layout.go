package handler

import (
	"net/http"
	"strconv"

	"github.com/nozgurozturk/usher/internal/application"
	"github.com/nozgurozturk/usher/internal/application/command"
	"github.com/nozgurozturk/usher/internal/application/query"
	"github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest"
	. "github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest/admin"
)

type LayoutHandler interface {
	// (GET /layouts)
	GetLayouts(w http.ResponseWriter, r *http.Request)

	// (POST /layouts)
	CreateLayout(w http.ResponseWriter, r *http.Request)

	// (GET /layouts/{layoutID})
	GetLayout(w http.ResponseWriter, r *http.Request, layoutID string)
}

type layoutHandler struct {
	app *application.Application
}

// NewLayoutHandler creates a new layoutHandler
func NewLayoutHandler(app *application.Application) LayoutHandler {
	return layoutHandler{
		app: app,
	}
}

func (h layoutHandler) GetLayouts(w http.ResponseWriter, r *http.Request) {
	layouts, err := h.app.Queries.GetLayouts.Handle(query.GetLayoutsQuery{})
	if err != nil {
		rest.JSON(w, http.StatusInternalServerError, ErrorResponse{
			Error:   "Error getting layouts",
			Message: err.Error(),
		})
		return
	}

	respLayouts := make([]Hall, len(layouts))
	for i, layout := range layouts {
		respLayouts[i] = Hall{
			Name: layout.Name(),
			Id:   layout.ID(),
		}
	}

	rest.JSON(w, http.StatusOK, respLayouts)
}

func (h layoutHandler) GetLayout(w http.ResponseWriter, r *http.Request, layoutID string) {
	l, err := h.app.Queries.GetLayout.Handle(query.GetLayoutQuery{
		LayoutID: layoutID,
	})
	if err != nil {
		rest.JSON(w, http.StatusInternalServerError, ErrorResponse{
			Error:   "Error getting layout",
			Message: err.Error(),
		})
		return
	}

	sections := make([]Section, len(l.Sections()))
	for i, section := range l.Sections() {
		sections[i].Name = section.Name()
		sections[i].Rows = make([]Row, len(section.Rows()))
		for j, row := range section.Rows() {
			sections[i].Rows[j].Name = row.Name()
			sections[i].Rows[j].Seats = make([]Seat, len(row.Seats()))
			for k, seat := range row.Seats() {
				f := int(seat.Feature())
				row, col := seat.Position()
				sections[i].Rows[j].Seats[k] = Seat{
					Id:        seat.ID(),
					Available: seat.Available(),
					Features:  &f,
					Number:    seat.Number(),
					Position: struct {
						Col int "json:\"col\""
						Row int "json:\"row\""
					}{
						Col: col,
						Row: row,
					},
					Rank: seat.Rank(),
				}
			}
		}
	}

	rest.JSON(w, http.StatusOK, GetLayoutResponse{
		Id:       l.ID(),
		Name:     l.Name(),
		Sections: sections,
	})
}

func (h layoutHandler) CreateLayout(w http.ResponseWriter, r *http.Request) {
	var body CreateLayoutRequest
	if err := rest.ParseBody(r, &body); err != nil {
		rest.JSON(w, http.StatusUnprocessableEntity, ErrorResponse{
			Error:   "Error parsing body",
			Message: err.Error(),
		})
		return
	}

	sections := make([]command.Section, len(body.Sections))
	for i, section := range body.Sections {
		sections[i] = command.Section{
			Name:    section.Name,
			Rows:    make([]command.Row, len(section.Rows)),
			Feature: string(*section.Feature),
		}
		for j, row := range section.Rows {
			sections[i].Rows[j] = command.Row{
				Seats: make([]command.Seat, len(row.Seats)),
			}
			for k, seat := range row.Seats {
				sections[i].Rows[j].Seats[k] = command.Seat(strconv.Itoa(seat.Rank))
			}
		}
	}

	l, err := h.app.Commands.CreateLayout.Handle(command.CreateLayoutCommand{
		Name:      "",
		Numbering: "",
		Sections:  sections,
	})

	if err != nil {
		rest.JSON(w, http.StatusInternalServerError, ErrorResponse{
			Error:   "Error creating layout",
			Message: err.Error(),
		})
		return
	}

	respSections := make([]Section, len(l.Sections()))
	for i, section := range l.Sections() {
		respSections[i].Name = section.Name()
		respSections[i].Rows = make([]Row, len(section.Rows()))
		for j, row := range section.Rows() {
			respSections[i].Rows[j].Name = row.Name()
			respSections[i].Rows[j].Seats = make([]Seat, len(row.Seats()))
			for k, seat := range row.Seats() {
				f := int(seat.Feature())
				row, col := seat.Position()
				respSections[i].Rows[j].Seats[k] = Seat{
					Available: seat.Available(),
					Features:  &f,
					Number:    seat.Number(),
					Position: struct {
						Col int "json:\"col\""
						Row int "json:\"row\""
					}{
						Col: col,
						Row: row,
					},
					Rank: seat.Rank(),
				}
			}
		}
	}

	rest.JSON(w, http.StatusOK, CreateLayoutResponse{
		Name:     l.Name(),
		Sections: respSections,
	})

}
