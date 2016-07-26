package main

import (
	"encoding/json"
	"strings"

	"github.com/snikch/goodman/hooks"
	trans "github.com/snikch/goodman/transaction"
)

var response_stash = make(map[string]string)

func main() {
	h := hooks.NewHooks()
	server := hooks.NewServer(h)
	h.After("Users > users > Create user", func(t *trans.Transaction) {
		response_stash[t.Name] = t.Real.Body
	})
	h.Before("Users > user details > Get user", func(t *trans.Transaction) {
		setID(t, "Users > users > Create user")
	})
	h.Before("Users > user details > Update user", func(t *trans.Transaction) {
		setID(t, "Users > users > Create user")
	})
	h.Before("Users > user details > Delete user", func(t *trans.Transaction) {
		setID(t, "Users > users > Create user")
	})

	h.After("Profiles > profiles > Create profile", func(t *trans.Transaction) {
		response_stash[t.Name] = t.Real.Body
	})
	h.Before("Profiles > profile details > Get profile", func(t *trans.Transaction) {
		setID(t, "Profiles > profiles > Create profile")
	})
	h.Before("Profiles > profile details > Update profile", func(t *trans.Transaction) {
		setID(t, "Profiles > profiles > Create profile")
	})
	h.Before("Profiles > profile details > Delete profile", func(t *trans.Transaction) {
		setID(t, "Profiles > profiles > Create profile")
	})
	server.Serve()
	defer server.Listener.Close()
}

func setID(t *trans.Transaction, event string) {
	// reusing data from previous response here
	var parseMap map[string]*json.RawMessage
	_ = json.Unmarshal([]byte(response_stash[event]), &parseMap)
	replaceId := parseMap["id"]
	// #replacing id in URL with stashed id from previous response
	t.FullPath = strings.Replace(t.FullPath, "1", string(*replaceId), 1)
}
