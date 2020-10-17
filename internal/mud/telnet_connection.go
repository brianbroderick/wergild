package mud

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/brianbroderick/wergild/internal/login"
)

type ConnState int8

const (
	STATE_WELCOME ConnState = iota
	STATE_LOGIN_USERNAME
	STATE_LOGIN_PASSWORD
	STATE_SHOULD_CREATE_NEW
	STATE_USER_CREATION
	STATE_PASSWORD_CREATION
	STATE_EMAIL_CREATION
	STATE_CHARACTER_CREATION
	STATE_PLAYING
	STATE_STATUE
)

type Connection struct {
	conn          net.Conn
	timeConnected time.Time
	state         ConnState
	user          login.User
}

func (connection *Connection) Write(message string) {
	if connection != nil {
		connection.conn.Write([]byte(message))
	}
}

// listen goes through the login/registration process.
// once logged in, it passes messages to ServerInstance.onCommandReceived
func (connection *Connection) listen() {
	reader := bufio.NewReader(connection.conn)
	// connection.Write("bobby")

	// connection.Write(motd())
	connection.Write("What is your name: ")
	connection.state = STATE_LOGIN_USERNAME

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("Connection error for %s: %v\n", connection.user.Name, err)
			connection.state = STATE_STATUE

			connection.conn.Close()
			ServerInstance.onClientConnectionClosed(connection, err)
			return
		}

		message = strings.TrimSpace(message)
		switch connection.state {
		case STATE_LOGIN_USERNAME:
			connection.user.Name = message

			exists, _, err := login.UserExistsTxn(connection.user)
			if err != nil {
				connection.hasError(err)
				return
			}

			if exists {
				connection.state = STATE_LOGIN_PASSWORD
				connection.Write("Password: ")
			} else {
				connection.state = STATE_USER_CREATION
				connection.Write("That username was not found. Create new: y/N? ")
			}
		case STATE_USER_CREATION:
			if strings.ToLower(message) == "y" {
				connection.state = STATE_PASSWORD_CREATION
				connection.Write("Choose a password of at least 6 characters: ")
			} else {
				connection.Write("OK, have a great day.\n")
				connection.conn.Close()
			}
		case STATE_PASSWORD_CREATION:
			ok := true
			// If password is only numbers
			if _, err := strconv.Atoi(message); err == nil {
				ok = false
			}
			// If password contains less than 6 characters
			if len(message) < 6 {
				ok = false
			}
			if ok {
				user, err := login.CreateUser(connection.user.Name, message)
				if err != nil {
					connection.hasError(err)
				}
				CreateUserMob(user)
				connection.user = user
				connection.state = STATE_EMAIL_CREATION
				connection.Write("Choose an email address: ")
			} else {
				connection.Write("You don't know too much about security, do you?  Try again.\nPassword: ")
			}
		case STATE_EMAIL_CREATION:
			// TODO: validate email
			if message == "" {
				connection.Write("Choose an email address: ")
			} else {
				connection.user.Email = message
				login.UpdateUser(connection.user)
				ServerInstance.onPlayerAuthenticated(connection)
			}
		case STATE_CHARACTER_CREATION:

		case STATE_LOGIN_PASSWORD:
			auth, err := login.Auth(connection.user.Name, message)
			if err != nil {
				connection.hasError(err)
				return
			}
			if auth {
				ServerInstance.onPlayerAuthenticated(connection)
			} else {
				connection.Write("Username or password was incorrect.\n")
				connection.conn.Close()
				return
			}

		default:
			ServerInstance.onCommandReceived(connection, message)
		}
	}
}

func (connection *Connection) hasError(err error) {
	connection.Write("There was an unexpected error.")
	connection.conn.Close()
	ServerInstance.onClientConnectionClosed(connection, err)
}
