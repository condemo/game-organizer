package store

import (
	"os"
	"testing"
	"time"

	"github.com/condemo/game-organizer/services/common/types"
	"github.com/stretchr/testify/require"
)

var mockupDB *Storage

var gameMockup = types.Game{
	IgdbID:      1,
	Title:       "Juego",
	Screenshot:  "link fake",
	ReleaseDate: time.Now(),
	Genres:      "acción,movidas",
	Developer:   "Noentiendo",
	Cover:       "https://fakecovers.com",
	Publisher:   "Capcom",
	Platforms:   "PC,Nintendo Switch",
	Rating:      78,
	Url:         "https://random.com",
	Played:      false,
	Pending:     false,
}

func TestMain(m *testing.M) {
	db := NewPosgresqlStore()
	mockupDB = NewStorage(db)
	os.Exit(m.Run())
}

func TestCreateGame(t *testing.T) {
	err := mockupDB.CreateGame(&gameMockup)
	require.NoError(t, err)
}

func TestGetGame(t *testing.T) {
	games, err := mockupDB.GetGamesPoltrait()
	require.NoError(t, err)
	require.NotEmpty(t, games)
}

func TestGetGamePoltrait(t *testing.T) {
	games, err := mockupDB.GetGamesPoltrait()
	require.NoError(t, err)
	require.NotEmpty(t, games)
}

func TestGetOneGame(t *testing.T) {
	var game types.Game
	err := mockupDB.GetOneGame(gameMockup.ID, &game)
	require.NoError(t, err)
	require.NotEmpty(t, game)
}

func TestDeleteGame(t *testing.T) {
	err := mockupDB.DeleteGame(gameMockup.ID)
	require.NoError(t, err)
}
