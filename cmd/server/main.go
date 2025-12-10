package main

import (
	"log"
	"net/http"

	"cloudevents-explorer/internal/config"
	"cloudevents-explorer/internal/handlers"
)

func main() {
	// Load configuration
	if err := config.Load(); err != nil {
		log.Printf("Warning: Failed to load config: %v", err)
	}

	// Register page handlers
	http.HandleFunc("/", handlers.HandleIndex)
	http.HandleFunc("/pubsub", handlers.HandlePubSub)
	http.HandleFunc("/kafka", handlers.HandleKafka)
	http.HandleFunc("/rest-client", handlers.HandleRestClient)
	http.HandleFunc("/flow-diagram", handlers.HandleFlowDiagram)
	http.HandleFunc("/gcs", handlers.HandleGCS)
	http.HandleFunc("/trace-journey", handlers.HandleTraceJourney)

	// Register API handlers
	http.HandleFunc("/api/configs", handlers.HandleGetConfigs)
	http.HandleFunc("/api/pubsub/configs", handlers.HandleSavePubSubConfig)
	http.HandleFunc("/api/kafka/configs", handlers.HandleSaveKafkaConfig)
	http.HandleFunc("/api/pubsub/pull", handlers.HandlePullPubSub)
	http.HandleFunc("/api/kafka/pull", handlers.HandlePullKafka)
	http.HandleFunc("/api/kafka/publish", handlers.HandlePublishKafka)
	http.HandleFunc("/api/rest/send", handlers.HandleRestSend)
	http.HandleFunc("/api/docker/status", handlers.HandleDockerStatus)
	http.HandleFunc("/api/gcloud/status", handlers.HandleGCloudStatus)
	http.HandleFunc("/api/gcs/buckets", handlers.HandleListBuckets)
	http.HandleFunc("/api/gcs/objects", handlers.HandleListObjects)
	http.HandleFunc("/api/gcs/object/content", handlers.HandleGetObjectContent)
	http.HandleFunc("/api/gcs/object/download", handlers.HandleDownloadObject)
	http.HandleFunc("/api/trace/search", handlers.HandleTraceSearch)

	port := "8888"
	log.Printf("🚀 Testing Studio starting on http://localhost:%s", port)
	log.Printf("📝 Configuration file: configs.json")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
