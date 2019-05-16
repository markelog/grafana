package ldap

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/ldap.v3"

	"github.com/grafana/grafana/pkg/infra/log"
)

func TestLdapLogin(t *testing.T) {
	Convey("Login using ldap", t, func() {
		authScenario("When login with invalid credentials", func(scenario *scenarioContext) {
			connection := &mockConnection{}
			entry := ldap.Entry{}
			result := ldap.SearchResult{Entries: []*ldap.Entry{&entry}}
			connection.setSearchResult(&result)

			connection.bindProvider = func(username, password string) error {
				return &ldap.Error{
					ResultCode: 49,
				}
			}
			auth := &Server{
				config: &ServerConfig{
					Attr: AttributeMap{
						Username: "username",
						Name:     "name",
						MemberOf: "memberof",
					},
					SearchBaseDNs: []string{"BaseDNHere"},
				},
				connection: connection,
				log:        log.New("test-logger"),
			}

			_, err := auth.Login(scenario.loginUserQuery)

			Convey("it should return invalid credentials error", func() {
				So(err, ShouldEqual, ErrInvalidCredentials)
			})
		})

		authScenario("When login with valid credentials", func(scenario *scenarioContext) {
			connection := &mockConnection{}
			entry := ldap.Entry{
				DN: "dn", Attributes: []*ldap.EntryAttribute{
					{Name: "username", Values: []string{"markelog"}},
					{Name: "surname", Values: []string{"Gaidarenko"}},
					{Name: "email", Values: []string{"markelog@gmail.com"}},
					{Name: "name", Values: []string{"Oleg"}},
					{Name: "memberof", Values: []string{"admins"}},
				},
			}
			result := ldap.SearchResult{Entries: []*ldap.Entry{&entry}}
			connection.setSearchResult(&result)

			connection.bindProvider = func(username, password string) error {
				return nil
			}
			auth := &Server{
				config: &ServerConfig{
					Attr: AttributeMap{
						Username: "username",
						Name:     "name",
						MemberOf: "memberof",
					},
					SearchBaseDNs: []string{"BaseDNHere"},
				},
				connection: connection,
				log:        log.New("test-logger"),
			}

			resp, err := auth.Login(scenario.loginUserQuery)

			Convey("it should not return error", func() {
				So(err, ShouldBeNil)
			})

			Convey("it should get user", func() {
				So(resp.Login, ShouldEqual, "markelog")
			})
		})
	})
}
