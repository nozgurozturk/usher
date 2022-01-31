package test

import (
	"context"
	"net/http/httptest"
	"testing"

	"entgo.io/ent/dialect"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nozgurozturk/usher/internal/application"
	public "github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest/public/handler"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/migrate"
	"github.com/stretchr/testify/suite"
)

type E2ETestSuite struct {
	suite.Suite
	app    *application.Application
	db     *ent.Client
	server *httptest.Server
}

func (s *E2ETestSuite) SetupSuite() {

	db, err := ent.Open(dialect.SQLite, "file:test.db?cache=shared&mode=memory&_fk=1")

	s.Require().NoError(err)
	s.db = db

	app := application.NewApplication(db)
	s.app = app

	h := public.NewAPIHandler(app)

	server := httptest.NewServer(h)
	s.server = server
}

func (s *E2ETestSuite) TearDownSuite() {
	s.server.Close()
	s.db.Close()
}

func (s *E2ETestSuite) SetupTest() {
	s.db.Schema.Create(
		context.Background(),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	)
}

func (s *E2ETestSuite) TearDownTest() {
	ResetDB(s.T(), s.db)
}

func TestE2ETestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test suite")
	}
	suite.Run(t, new(E2ETestSuite))
}

func ResetDB(t *testing.T, client *ent.Client) {
	t.Helper()
	t.Log("drop data from database")
	ctx := context.Background()
	client.Event.Delete().ExecX(ctx)
	client.Reservation.Delete().ExecX(ctx)
	client.Seat.Delete().ExecX(ctx)
	client.Row.Delete().ExecX(ctx)
	client.Section.Delete().ExecX(ctx)
	client.Layout.Delete().ExecX(ctx)
	client.Ticket.Delete().ExecX(ctx)
	client.User.Delete().ExecX(ctx)
}
