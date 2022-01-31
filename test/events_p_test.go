package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/nozgurozturk/usher/internal/infrastructure/store/seeder"
)

func (s *E2ETestSuite) TestE2E_GetEvents() {

	ls := seeder.CreateLayouts(s.db, seeder.Layout{
		Name:      "Test Layout",
		Numbering: "sequential",
		Sections: []seeder.Section{
			{
				Name:    "A",
				Feature: "default",
				Rows: []seeder.Row{
					{
						Seats: []seeder.Seat{"1", "1", "1"},
					},
				},
			},
		},
	})

	event1 := seeder.Event{
		Name:     "Test Event-1",
		LayoutID: ls[0].ID,
	}

	event2 := seeder.Event{
		Name:     "Test Event-2",
		LayoutID: ls[0].ID,
	}

	events := []seeder.Event{event1, event2}
	_ = seeder.CreateEvents(s.db, events...)

	req, err := http.NewRequest("GET", s.server.URL+"/events", nil)
	s.Require().NoError(err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.server.Client().Do(req)

	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.StatusCode)

	var jsonEvents []interface{}

	err = json.NewDecoder(resp.Body).Decode(&jsonEvents)
	s.Require().NoError(err)

	s.Require().Len(jsonEvents, 2)

}

func (s *E2ETestSuite) TestE2E_GetEvent() {
	ls := seeder.CreateLayouts(s.db, seeder.Layout{
		Name:      "Test Layout",
		Numbering: "sequential",
		Sections: []seeder.Section{
			{
				Name:    "A",
				Feature: "default",
				Rows: []seeder.Row{
					{

						Seats: []seeder.Seat{"1"},
					},
				},
			},
		},
	})

	event := seeder.Event{
		Name:     "Test Event-1",
		LayoutID: ls[0].ID,
	}

	es := seeder.CreateEvents(s.db, event)

	req, err := http.NewRequest("GET", s.server.URL+"/events/"+es[0].ID.String(), nil)
	s.Require().NoError(err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.server.Client().Do(req)

	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.StatusCode)

	byteBody, err := ioutil.ReadAll(resp.Body)
	s.Require().NoError(err)

	timeLayout := "2006-01-02T15:04:05.000000Z"
	expected := fmt.Sprintf(`
	{
		"id": "%s",
		"name": "Test Event-1",
		"description": "",
		"startDate": "%s",
		"endDate": "%s",
		"location": {
		  "name": "Test Layout",
		  "sections": [
			{
			  "name": "A",
			  "rows": [
				{
				  "name": "A1",
				  "order": 0,
				  "seats": [
					{
					  "position": {
						"row": 0,
						"col": 0
					  },
					  "number": 1,
					  "features": 5,
					  "rank": 1,
					  "available": true
					}
				  ]
				}
			  ]
			}
		  ]
		}
	  }
	`, es[0].ID.String(), es[0].StartAt.Format(timeLayout), es[0].EndAt.Format(timeLayout))

	s.Require().JSONEq(expected, string(byteBody))
}

func (s *E2ETestSuite) TestE2E_CheckEventSeats() {
	/*
		Features are:
		"default": 1,
		"balcony": 2,
		"front":   4,
	*/
	ls := seeder.CreateLayouts(s.db, seeder.Layout{
		Name:      "Test Layout",
		Numbering: "sequential",
		Sections: []seeder.Section{
			{
				Name:    "A",
				Feature: "default",
				Rows: []seeder.Row{
					{
						Seats: []seeder.Seat{"1", "1", "2", "2"},
					},
					{
						Seats: []seeder.Seat{"2"},
					},
				},
			},
			{
				Name:    "B",
				Feature: "balcony",
				Rows: []seeder.Row{
					{
						Seats: []seeder.Seat{"1", "1", "1", "2", "2", "2", "2"},
					},
				},
			},
		},
	})

	event := seeder.Event{
		Name:     "Test Event-1",
		LayoutID: ls[0].ID,
	}

	es := seeder.CreateEvents(s.db, event)

	s.Run("Seats available with rank and feature", func() {
		req, err := http.NewRequest("POST", s.server.URL+"/events/"+es[0].ID.String()+"/check", strings.NewReader(
			`
			{
				"count": 3,
				"rank": 1,
				"features": 4
			  }
			`,
		))

		s.Require().NoError(err)

		req.Header.Set("Content-Type", "application/json")

		resp, err := s.server.Client().Do(req)

		s.Require().NoError(err)
		s.Require().Equal(http.StatusOK, resp.StatusCode)

		byteBody, err := ioutil.ReadAll(resp.Body)
		s.Require().NoError(err)

		expected := `
		{
			"remaining": 3
		}`

		s.Require().JSONEq(expected, string(byteBody))

	})

	s.Run("Seats available with rank", func() {
		req, err := http.NewRequest("POST", s.server.URL+"/events/"+es[0].ID.String()+"/check", strings.NewReader(
			`
			{
				"count": 3,
				"rank": 2
			  }
			`,
		))

		s.Require().NoError(err)

		req.Header.Set("Content-Type", "application/json")

		resp, err := s.server.Client().Do(req)

		s.Require().NoError(err)
		s.Require().Equal(http.StatusOK, resp.StatusCode)

		byteBody, err := ioutil.ReadAll(resp.Body)
		s.Require().NoError(err)

		expected := `
		{
			"remaining": 4
		}`

		s.Require().JSONEq(expected, string(byteBody))

	})

	s.Run("Seats available with feature", func() {
		req, err := http.NewRequest("POST", s.server.URL+"/events/"+es[0].ID.String()+"/check", strings.NewReader(
			`
			{
				"count": 3,
				"features": 4
			  }
			`,
		))

		s.Require().NoError(err)

		req.Header.Set("Content-Type", "application/json")

		resp, err := s.server.Client().Do(req)

		s.Require().NoError(err)
		s.Require().Equal(http.StatusOK, resp.StatusCode)

		byteBody, err := ioutil.ReadAll(resp.Body)
		s.Require().NoError(err)

		expected := `
		{
			"remaining": 7
		}`

		s.Require().JSONEq(expected, string(byteBody))

	})

}
