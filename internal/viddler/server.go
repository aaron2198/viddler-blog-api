package viddler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/aiservice"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/generator"
)

func (viddler *Viddler) Server(port int) {
	kill := make(chan os.Signal, 1)
	signal.Notify(kill, os.Interrupt)
	mux := http.NewServeMux()
	mux.HandleFunc("/api/generate", corsMiddleware(viddler.generateHandler))
	mux.HandleFunc("/api/phaseoptions", corsMiddleware(phaseOptionsHandler))
	mux.HandleFunc("/api/clientoptions", corsMiddleware(viddler.clientOptionsHandler))

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	go func() {
		fmt.Println("Starting server on port", port)
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	<-kill
	fmt.Println("Shutting down...")
	server.Shutdown(context.Background())
}

func (viddler *Viddler) generateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var options generator.UserProvidedOptions
	err := json.NewDecoder(r.Body).Decode(&options)
	if err != nil {
		http.Error(w, "Failed to decode options", http.StatusBadRequest)
		return
	}

	fmt.Println(options)
	params := generator.ArticleGeneratorParams{
		Config:      viddler.Config.ArticleGenerator,
		BucketStore: viddler.BucketStore,
		Options:     &options,
	}
	article, err := generator.New(&params).GenerateArticle()
	if err != nil {
		msg := fmt.Sprintf("Failed to generate article: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"article": article})
}

type PhaseOptions struct {
	Phases []PhaseOption
}

type PhaseOption struct {
	Name        string
	Description string
	Clients     map[string][]string
}

func phaseOptionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	phaseOptions := PhaseOptions{}

	for _, availablePhase := range aiservice.PhaseOrder {
		phaseOptions.Phases = append(phaseOptions.Phases, PhaseOption{
			Name:        availablePhase,
			Description: aiservice.AvailablePhases[availablePhase],
			Clients:     aiservice.ModelOptions(aiservice.StructuredOutputs, aiservice.PsuedoStructuredOutputs),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(phaseOptions)
}

func (viddler *Viddler) clientOptionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(aiservice.ModelOptions())
}
