package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/dev", DevHandler)
	r.HandleFunc("/session", SessionCreateHandler)
	r.HandleFunc("/listen", ListenHandler)
	r.HandleFunc("/test", TestHandler)

	// interplay tests
	r.HandleFunc("/evil", EvilHandler)
	r.HandleFunc("/bridge", BridgeHandler)
	r.HandleFunc("/victim", VictimHandler)

	// multiple-stage CSP
	//r.HandleFunc("/multi_csp/{file:[a-z_]+}", MultiCSPHandler)
	r.HandleFunc("/multi_csp/{file}", MultiCSPHandler)
	//r.HandleFunc("/multi_csp_js/{file:[a-z]+}", MultiCSPScriptHandler)
	r.HandleFunc("/multi_csp_js/{file}", MultiCSPScriptHandler)

	// postmessage
	r.HandleFunc("/postmsg/{file}", PostmsgHandler)

	// shared resource bypassing
	r.HandleFunc("/share/{file}", ShareHandler)

	// title test page
	r.HandleFunc("/titletest/{file}", TitleTestHandler)

	// clickjack - csp test
	r.HandleFunc("/clickjackcsp/{file}", ClickjackCSPHandler)

	// third-party cookie blocking
	r.HandleFunc("/thirdpartycookie/{file}", ThirdPartyCookieHandler)
	r.HandleFunc("/setthirdpartycookie/{file}", SetThirdPartyCookieHandler)

	// tracker A
	r.HandleFunc("/trackera/{file}", TrackerAHandler)

	// tracker B
	r.HandleFunc("/trackerb/{file}", TrackerBHandler)

	// tracker C
	r.HandleFunc("/trackerc/{file}", TrackerCHandler)

	// tracker E
	r.HandleFunc("/trackere/{file}", TrackerEHandler)

	// Start listening on the given IP address and port
	http.Handle("/", r)
	var httpListenAddr = fmt.Sprintf("%s:%d",
		"127.0.0.1",
		8001)
	if err := http.ListenAndServe(httpListenAddr, nil); err != nil {
		log.Fatalf("Could not start HTTP server listening: %s\n", err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func DevHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./dev.html")
}

func SessionCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	r.ParseForm()
	fmt.Println("username:", r.Form["uname"])
	fmt.Println("password:", r.Form["passwd"])
}

func ListenHandler(w http.ResponseWriter, r *http.Request) {
	// http://mylab.com/listen?login=param_value
	//	httpHeader := http.Header(r.Header)
	//	log.Println(httpHeader)
	param1 := r.URL.Query().Get("login")
	fmt.Println(param1)
	http.ServeFile(w, r, "./static/1.png")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./test.html")
}

func EvilHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./interplay/evil.html")
}

func BridgeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./interplay/bridge.html")
}

func VictimHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./interplay/victim.html")
}

func MultiCSPHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./week2/"+mux.Vars(r)["file"])
	fmt.Println("./week2/" + mux.Vars(r)["file"])
}

func MultiCSPScriptHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./week2/js/"+mux.Vars(r)["file"])
}

func PostmsgHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./week2/postmsg/"+mux.Vars(r)["file"])
}

func ShareHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./week2/sharedpage/"+mux.Vars(r)["file"])
}

func TitleTestHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./week3/title_test/"+mux.Vars(r)["file"])
}

func ClickjackCSPHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./week3/clickjackcsp/"+mux.Vars(r)["file"])
}

func GifCSPBypassHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./week3/gifbypass/"+mux.Vars(r)["file"])
}

func ThirdPartyCookieHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./week3/thirdparty/"+mux.Vars(r)["file"])
}

func SetThirdPartyCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "setcookie",
		Value: "42",
	}
	http.SetCookie(w, cookie)
	http.ServeFile(w, r, "./week3/thirdparty/"+mux.Vars(r)["file"])
}

func TrackerAHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./week4/trackera/"+mux.Vars(r)["file"])
}

func TrackerBHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./week4/trackerb/"+mux.Vars(r)["file"])
}

func TrackerCHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./week4/trackerc/"+mux.Vars(r)["file"])
}

func TrackerEHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./week4/trackere/"+mux.Vars(r)["file"])
}
