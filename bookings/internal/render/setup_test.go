package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/techieasif/TheGo/bookings/internal/config"
	"github.com/techieasif/TheGo/bookings/models"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	//What I am going to put in session
	gob.Register(models.Reservation{})

	testApp.InProduction = false
	session = scs.New()

	//24 hour session
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = testApp.InProduction

	testApp.Session = session

	app = &testApp
	os.Exit(m.Run())
}
