package main

import "net/http"

func main() {
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("deneme"))
		})
	*/
	//----------------------------
	//http.HandleFunc("/",indexHandler)
	//http.HandleFunc("/about",aboutHandler)
	//http.ListenAndServe(":8080", nil)
	//-----------------------------------
	/*
		var i ironman
		var w wolverine
		mux := http.NewServeMux() //http hanle func yerina arık mux kullanacagız
		mux.Handle("/iron",i)
		mux.Handle("/wolwer",w)


		http.ListenAndServe(":8080", mux)

	*/

	/* //coustumHandler
	mux := http.NewServeMux()
	x1 := &messageHandler{"bu bir mesaj"}
	x2 := &messageHandler{"bu da bir mesaj"}

	mux.Handle("/bir",x1)
	mux.Handle("/iki",x2)

	log.Println("listening")

	http.ListenAndServe(":8080",mux)

	*/

	//queryString
	http.HandleFunc("/search", search)
	http.ListenAndServe(":8080", nil)

}
