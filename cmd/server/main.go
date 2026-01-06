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
	http.HandleFunc("/spanner", handlers.HandleSpanner)
	http.HandleFunc("/config-editor", handlers.HandleConfigEditor)
	http.HandleFunc("/flimflam-explorer", handlers.FlimFlamExplorerHandler)

	// Register API handlers
	http.HandleFunc("/api/configs", handlers.HandleGetConfigs)
	http.HandleFunc("/api/configs/save", handlers.HandleSaveAllConfigs)
	http.HandleFunc("/api/pubsub/configs", handlers.HandleSavePubSubConfig)
	http.HandleFunc("/api/kafka/configs", handlers.HandleSaveKafkaConfig)
	http.HandleFunc("/api/pubsub/pull", handlers.HandlePullPubSub)
	http.HandleFunc("/api/kafka/pull", handlers.HandlePullKafka)
	http.HandleFunc("/api/kafka/publish", handlers.HandlePublishKafka)
	http.HandleFunc("/api/rest/send", handlers.HandleRestSend)
	http.HandleFunc("/api/rest/save", handlers.HandleSaveRequest)
	http.HandleFunc("/api/rest/collections", handlers.HandleGetCollections)
	http.HandleFunc("/api/rest/delete", handlers.HandleDeleteRequest)
	http.HandleFunc("/api/rest/collection/delete", handlers.HandleDeleteCollection)
	http.HandleFunc("/api/docker/status", handlers.HandleDockerStatus)
	http.HandleFunc("/api/gcloud/status", handlers.HandleGCloudStatus)
	http.HandleFunc("/api/gcs/buckets", handlers.HandleListBuckets)
	http.HandleFunc("/api/gcs/objects", handlers.HandleListObjects)
	http.HandleFunc("/api/gcs/object/content", handlers.HandleGetObjectContent)
	http.HandleFunc("/api/gcs/object/download", handlers.HandleDownloadObject)
	http.HandleFunc("/api/trace/search", handlers.HandleTraceSearch)
	http.HandleFunc("/api/spanner/connect", handlers.HandleSpannerConnect)
	http.HandleFunc("/api/spanner/tables", handlers.HandleSpannerTables)
	http.HandleFunc("/api/spanner/query", handlers.HandleSpannerQuery)
	http.HandleFunc("/api/spanner/configs", handlers.HandleSaveSpannerConfig)
	http.HandleFunc("/api/spanner/schema", handlers.HandleSpannerSchema)
	http.HandleFunc("/api/flimflam/apis", handlers.FlimFlamAPIsHandler)
	http.HandleFunc("/api/flimflam/send", handlers.FlimFlamProxyHandler)
	http.HandleFunc("/api/flimflam/status", handlers.FlimFlamStatusHandler)

	port := "8888"
	log.Printf("🚀 Testing Studio starting on http://localhost:%s", port)
	log.Printf("📝 Configuration file: configs.json")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
