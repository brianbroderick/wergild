package login

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/brianbroderick/agora"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

func CreateUser(name string, pass string) (User, error) {
	u := User{Name: name, Pass: pass, Type: "User"}

	c := agora.NewDgraphTxn()
	defer c.DiscardTxn()

	exists, u, err := UserExists(c, u)
	if err != nil {
		return u, err
	}

	if exists {
		return u, fmt.Errorf("User already exists.")
	}

	u, err = saveUser(c, u)
	if err != nil {
		return u, err
	}

	return u, nil
}

func UserExistsTxn(u User) (bool, User, error) {
	c := agora.NewDgraphTxn()
	defer c.DiscardTxn()

	return UserExists(c, u)
}

func UserExists(c *agora.DgraphConn, u User) (bool, User, error) {
	query := `query User($name: string){
		users(first:1, func: eq(userName, $name)) {
		  uid 
		}
	  }`
	variables := make(map[string]string)
	variables["$name"] = u.Name

	resp, err := c.Txn.QueryWithVars(c.Ctx, query, variables)
	if err != nil {
		return false, User{}, err
	}

	var r response
	j := resp.Json
	err = json.Unmarshal(j, &r)
	if err != nil {
		return false, User{}, err
	}

	if len(r.Users) > 0 {
		u.UID = r.Users[0].UID
		return true, u, nil
	}

	return false, u, nil
}

func saveUser(c *agora.DgraphConn, u User) (User, error) {
	// returned in the assigned var
	u.UID = "_:user"

	jsonStr, err := json.Marshal(u)
	if err != nil {
		return User{}, err
	}

	mu := &api.Mutation{
		SetJson: jsonStr,
	}

	assigned, err := c.Txn.Mutate(c.Ctx, mu)
	if err != nil {
		return User{}, err
	}
	err = c.Txn.Commit(c.Ctx)
	if err != nil {
		return User{}, err
	}

	if assigned.Uids["user"] != "" {
		u.UID = assigned.Uids["user"]
	}
	return u, nil
}

func UpdateUser(u User) {
	if u.UID == "" {
		fmt.Println("Didn't find UID to update User: ", u.Name)
		return
	}

	j, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	agora.MutateDgraph(j)
}

// Upsert isn't working yet. Seems to not find the variables

// func UpsertEmail(u User) {
// 	query := `query user($name: string){
// 		user as var(func: eq(userName, $name))
// 	  }`
// 	variables := make(map[string]string)
// 	variables["$name"] = u.Name

// 	// cond := "@if(eq(len(user), 1))"
// 	cond := ""
// 	upd := User{Email: u.Email}
// 	j, err := json.Marshal(upd)
// 	if err != nil {
// 		fmt.Println("HERE")
// 		log.Fatal(err)
// 	}

// 	agora.UpsertDgraph(query, variables, cond, j)
// }
