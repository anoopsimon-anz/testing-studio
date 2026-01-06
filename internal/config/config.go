package config

import (
	"encoding/json"
	"os"
	"sync"
)

const configFile = "configs.json"

type PubSubConfig struct {
	Name           string `json:"name"`
	EmulatorHost   string `json:"emulatorHost"`
	ProjectID      string `json:"projectId"`
	SubscriptionID string `json:"subscriptionId"`
}

type KafkaConfig struct {
	Name           string `json:"name"`
	Brokers        string `json:"brokers"`
	Topic          string `json:"topic"`
	ConsumerGroup  string `json:"consumerGroup"`
	SchemaRegistry string `json:"schemaRegistry"`
}

type SpannerConfig struct {
	Name         string `json:"name"`
	EmulatorHost string `json:"emulatorHost"`
	ProjectID    string `json:"projectId"`
	InstanceID   string `json:"instanceId"`
	DatabaseID   string `json:"databaseId"`
}

type GCSConfig struct {
	Name         string `json:"name"`
	EmulatorHost string `json:"emulatorHost"`
	ProjectID    string `json:"projectId"`
}

type SavedRequest struct {
	Name       string                 `json:"name"`
	Method     string                 `json:"method"`
	URL        string                 `json:"url"`
	Headers    map[string]string      `json:"headers"`
	Parameters map[string]string      `json:"parameters"`
	Body       string                 `json:"body"`
	TLSCert    string                 `json:"tlsCert,omitempty"`
	TLSKey     string                 `json:"tlsKey,omitempty"`
}

type RequestCollection struct {
	Name     string          `json:"name"`
	Requests []SavedRequest  `json:"requests"`
}

type Config struct {
	PubSubConfigs       []PubSubConfig       `json:"pubsubConfigs"`
	KafkaConfigs        []KafkaConfig        `json:"kafkaConfigs"`
	SpannerConfigs      []SpannerConfig      `json:"spannerConfigs"`
	GCSConfigs          []GCSConfig          `json:"gcsConfigs,omitempty"`
	RequestCollections  []RequestCollection  `json:"requestCollections,omitempty"`
}

var (
	mu     sync.RWMutex
	config Config
)

func Load() error {
	mu.Lock()
	defer mu.Unlock()

	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			config = Config{
				PubSubConfigs: []PubSubConfig{
					{
						Name:           "TMS Local",
						EmulatorHost:   "localhost:8086",
						ProjectID:      "tms-suncorp-local",
						SubscriptionID: "cloudevents.subscription",
					},
				},
				KafkaConfigs: []KafkaConfig{
					{
						Name:           "TMS Unica Local",
						Brokers:        "localhost:19092",
						Topic:          "unica.marketing.response.events",
						ConsumerGroup:  "cloudevents-explorer",
						SchemaRegistry: "http://localhost:18081",
					},
				},
				SpannerConfigs: []SpannerConfig{
					{
						Name:         "TMS Local",
						EmulatorHost: "localhost:9010",
						ProjectID:    "tms-suncorp-local",
						InstanceID:   "tms-suncorp-local",
						DatabaseID:   "tms-suncorp-db",
					},
				},
			}
			return saveLocked()
		}
		return err
	}

	return json.Unmarshal(data, &config)
}

func saveLocked() error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configFile, data, 0644)
}

func Save() error {
	mu.Lock()
	defer mu.Unlock()
	return saveLocked()
}

func Get() Config {
	mu.RLock()
	defer mu.RUnlock()
	return config
}

func AddOrUpdatePubSubConfig(newConfig PubSubConfig) error {
	mu.Lock()
	found := false
	for i, cfg := range config.PubSubConfigs {
		if cfg.Name == newConfig.Name {
			config.PubSubConfigs[i] = newConfig
			found = true
			break
		}
	}
	if !found {
		config.PubSubConfigs = append(config.PubSubConfigs, newConfig)
	}
	mu.Unlock()

	return Save()
}

func AddOrUpdateKafkaConfig(newConfig KafkaConfig) error {
	mu.Lock()
	found := false
	for i, cfg := range config.KafkaConfigs {
		if cfg.Name == newConfig.Name {
			config.KafkaConfigs[i] = newConfig
			found = true
			break
		}
	}
	if !found {
		config.KafkaConfigs = append(config.KafkaConfigs, newConfig)
	}
	mu.Unlock()

	return Save()
}

func AddOrUpdateSpannerConfig(newConfig SpannerConfig) error {
	mu.Lock()
	found := false
	for i, cfg := range config.SpannerConfigs {
		if cfg.Name == newConfig.Name {
			config.SpannerConfigs[i] = newConfig
			found = true
			break
		}
	}
	if !found {
		config.SpannerConfigs = append(config.SpannerConfigs, newConfig)
	}
	mu.Unlock()

	return Save()
}

func SaveRequestToCollection(collectionName string, req SavedRequest) error {
	mu.Lock()
	defer mu.Unlock()

	// Find or create collection
	var collection *RequestCollection
	for i := range config.RequestCollections {
		if config.RequestCollections[i].Name == collectionName {
			collection = &config.RequestCollections[i]
			break
		}
	}

	if collection == nil {
		// Create new collection
		config.RequestCollections = append(config.RequestCollections, RequestCollection{
			Name:     collectionName,
			Requests: []SavedRequest{},
		})
		collection = &config.RequestCollections[len(config.RequestCollections)-1]
	}

	// Add or update request in collection
	found := false
	for i, r := range collection.Requests {
		if r.Name == req.Name {
			collection.Requests[i] = req
			found = true
			break
		}
	}
	if !found {
		collection.Requests = append(collection.Requests, req)
	}

	return saveLocked()
}

func GetRequestCollections() []RequestCollection {
	mu.RLock()
	defer mu.RUnlock()
	return config.RequestCollections
}

func DeleteRequestFromCollection(collectionName, requestName string) error {
	mu.Lock()
	defer mu.Unlock()

	for i, coll := range config.RequestCollections {
		if coll.Name == collectionName {
			for j, req := range coll.Requests {
				if req.Name == requestName {
					config.RequestCollections[i].Requests = append(
						coll.Requests[:j],
						coll.Requests[j+1:]...,
					)
					return saveLocked()
				}
			}
		}
	}

	return saveLocked()
}

func DeleteCollection(name string) error {
	mu.Lock()
	defer mu.Unlock()

	for i, coll := range config.RequestCollections {
		if coll.Name == name {
			config.RequestCollections = append(
				config.RequestCollections[:i],
				config.RequestCollections[i+1:]...,
			)
			return saveLocked()
		}
	}

	return saveLocked()
}
