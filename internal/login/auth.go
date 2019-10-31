package login

import (
	"github.com/brianbroderick/agora"
)

type response struct {
	Users []User `json:"users,omitempty"`
}

type User struct {
	UID   string `json:"uid,omitempty"`
	Type  string `json:"dgraph.type,omitempty"`
	Name  string `json:"userName,omitempty"`
	Pass  string `json:"pass,omitempty"`
	Email string `json:"email,omitempty"`
	Auth  bool   `json:"auth,omitempty"`
}

func Auth(name string, pass string) (bool, error) {
	query := `query User($name: string, $pass: string){
				users(first:1, func: eq(userName, $name)) {
      				auth : checkpwd(pass, $pass)
				}
	  		}`
	variables := make(map[string]string)
	variables["$name"] = name
	variables["$pass"] = pass

	var r response
	err := agora.ResolveQueryWithVars(&r, query, variables)
	if err != nil {
		return false, err
	}

	users := r.Users

	if len(users) > 0 && users[0].Auth {
		return true, nil
	}
	return false, nil
}
