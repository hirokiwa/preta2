package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"hackz.com/m/v2/graph"
	"hackz.com/m/v2/graph/generated"
	"hackz.com/m/v2/infrastructure"
)

func main() {
	v2_db := infrastructure.Init()
	db,err := v2_db.DB()
	if err != nil {
		
	}
	defer db.Close()

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000","https://alexture-diary-7ak4y4niw-mochimochi.vercel.app","http://20.222.244.179:8443","https://20.222.244.179:8443","https://20.63.152.225:8443","https://alexturediary.vercel.app"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
    srv.AddTransport(&transport.Websocket{
        Upgrader: websocket.Upgrader{
            CheckOrigin: func(r *http.Request) bool {
                // Check against your desired domains here
                 return r.Host == "example.org"
            },
            ReadBufferSize:  1024,
            WriteBufferSize: 1024,
        },
    })

	router.Handle("/", playground.Handler("Starwars", "/query"))
	router.Handle("/query", srv)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

