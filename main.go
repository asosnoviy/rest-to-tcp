package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	tcpsender "github.com/asosnoviy/rest-to-tcp/internal"
	"github.com/kardianos/service"
)

type program struct{}

func (p program) Start(s service.Service) error {
	fmt.Println(s.String() + " started")
	go p.run()
	return nil
}

func (p program) Stop(s service.Service) error {
	fmt.Println(s.String() + " stopped")
	return nil
}

func (p program) run() {

	port := flag.String("port", "8050", "port")
	flag.Parse()

	http.Handle("/simpletcp", http.HandlerFunc(handleRequest))

	fmt.Println("Server started at port ", *port)
	http.ListenAndServe(":"+*port, nil)

}

func main() {

	serviceConfig := &service.Config{
		Name:        "simple tcp http proxy service",
		DisplayName: "simple tcp http proxy service",
		Description: "Web service for simple tcp query by http",
	}
	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		fmt.Println("Cannot create the service: " + err.Error())
	}
	err = s.Run()
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}

}

type answer struct {
	Error      string
	Data       []byte
	StringData string
}

type request struct {
	Server string `json:"server"`
	Data   []byte `json:"data"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	var request request

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)

	if err != nil {
		b, _ := json.Marshal(answer{Error: err.Error()})
		w.Write(b)
		return
	}

	answerdata, err := tcpsender.Send(request.Server, []byte(request.Data))
	w.Header().Set("Content-Type", "application/json")

	answer := answer{}

	if err != nil {
		answer.Error = err.Error()
	}
	answer.Data = answerdata
	answer.StringData = string(answerdata)

	b, _ := json.Marshal(answer)

	w.Write(b)

}
