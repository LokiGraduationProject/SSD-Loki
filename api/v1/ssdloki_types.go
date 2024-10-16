/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BloomBuild 설정 구조체
type BloomBuild struct {
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`

	// +optional
	// +kubebuilder:validation:Optional
	Builder *BloomBuildBuilder `json:"builder,omitempty"`
}

type BloomBuildBuilder struct {
	// +kubebuilder:validation:Required
	PlannerAddress string `json:"plannerAddress"`
}

// BloomGateway 설정 구조체
type BloomGateway struct {
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`

	// +optional
	// +kubebuilder:validation:Optional
	Client *BloomGatewayClient `json:"client,omitempty"`
}

type BloomGatewayClient struct {
	// +kubebuilder:validation:Required
	Addresses string `json:"addresses"`
}

// ChunkStoreConfig 설정 구조체
type ChunkStoreConfig struct {
	// +optional
	// +kubebuilder:validation:Optional
	ChunkCacheConfig *ChunkCacheConfig `json:"chunkCacheConfig,omitempty"`
}

type ChunkCacheConfig struct {
	// +optional
	// +kubebuilder:validation:Optional
	Background *CacheBackgroundConfig `json:"background,omitempty"`

	// +kubebuilder:validation:Required
	DefaultValidity string `json:"defaultValidity"`

	// +optional
	// +kubebuilder:validation:Optional
	Memcached *MemcachedConfig `json:"memcached,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	MemcachedClient *MemcachedClientConfig `json:"memcachedClient,omitempty"`
}

type CacheBackgroundConfig struct {
	// +kubebuilder:validation:Required
	WritebackBuffer int `json:"writebackBuffer"`

	// +kubebuilder:validation:Required
	WritebackGoroutines int `json:"writebackGoroutines"`

	// +kubebuilder:validation:Required
	WritebackSizeLimit string `json:"writebackSizeLimit"`
}

type MemcachedConfig struct {
	// +kubebuilder:validation:Required
	BatchSize int `json:"batchSize"`

	// +kubebuilder:validation:Required
	Parallelism int `json:"parallelism"`
}

type MemcachedClientConfig struct {
	// +kubebuilder:validation:Required
	Addresses string `json:"addresses"`

	// +kubebuilder:validation:Required
	ConsistentHash bool `json:"consistentHash"`

	// +optional
	// +kubebuilder:validation:Optional
	MaxIdleConns int `json:"maxIdleConns,omitempty"`

	// +kubebuilder:validation:Required
	Timeout string `json:"timeout"`

	// +optional
	// +kubebuilder:validation:Optional
	UpdateInterval string `json:"updateInterval,omitempty"`
}

// Common 설정 구조체
type CommonConfig struct {
	// +kubebuilder:validation:Required
	CompactorAddress string `json:"compactorAddress"`

	// +kubebuilder:validation:Required
	PathPrefix string `json:"pathPrefix"`

	// +kubebuilder:validation:Required
	ReplicationFactor int `json:"replicationFactor"`

	// +optional
	// +kubebuilder:validation:Optional
	Storage *CommonStorage `json:"storage,omitempty"`
}

type CommonStorage struct {
	// +optional
	// +kubebuilder:validation:Optional
	S3 *S3Config `json:"s3,omitempty"`
}

type S3Config struct {
	// +kubebuilder:validation:Required
	AccessKeyID string `json:"accessKeyId"`

	// +kubebuilder:validation:Required
	SecretAccessKey string `json:"secretAccessKey"`

	// +kubebuilder:validation:Required
	BucketNames string `json:"bucketnames"`

	// +kubebuilder:validation:Required
	Endpoint string `json:"endpoint"`

	// +kubebuilder:validation:Required
	Insecure bool `json:"insecure"`

	// +kubebuilder:validation:Required
	S3ForcePathStyle bool `json:"s3ForcePathStyle"`
}

// Frontend 설정 구조체
type FrontendConfig struct {
	// +kubebuilder:validation:Required
	SchedulerAddress string `json:"schedulerAddress"`

	// +kubebuilder:validation:Required
	TailProxyURL string `json:"tailProxyUrl"`
}

type FrontendWorkerConfig struct {
	// +kubebuilder:validation:Required
	SchedulerAddress string `json:"schedulerAddress"`
}

// IndexGateway 설정 구조체
type IndexGatewayConfig struct {
	// +kubebuilder:validation:Required
	Mode string `json:"mode"`
}

// Ingester 설정 구조체
type IngesterConfig struct {
	// +kubebuilder:validation:Required
	ChunkEncoding string `json:"chunkEncoding"`
}

// LimitsConfig 설정 구조체
type LimitsConfig struct {
	// +kubebuilder:validation:Required
	MaxCacheFreshnessPerQuery string `json:"maxCacheFreshnessPerQuery"`

	// +kubebuilder:validation:Required
	QueryTimeout string `json:"queryTimeout"`

	// +kubebuilder:validation:Required
	RejectOldSamples bool `json:"rejectOldSamples"`

	// +kubebuilder:validation:Required
	RejectOldSamplesMaxAge string `json:"rejectOldSamplesMaxAge"`

	// +kubebuilder:validation:Required
	SplitQueriesByInterval string `json:"splitQueriesByInterval"`

	// +kubebuilder:validation:Required
	VolumeEnabled bool `json:"volumeEnabled"`
}

// Memberlist 설정 구조체
type MemberlistConfig struct {
	// +kubebuilder:validation:Required
	JoinMembers []string `json:"joinMembers"`
}

// PatternIngester 설정 구조체
type PatternIngesterConfig struct {
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

// Querier 설정 구조체
type QuerierConfig struct {
	// +kubebuilder:validation:Required
	MaxConcurrent int `json:"maxConcurrent"`
}

// QueryRange 설정 구조체
type QueryRangeConfig struct {
	// +kubebuilder:validation:Required
	AlignQueriesWithStep bool `json:"alignQueriesWithStep"`

	// +kubebuilder:validation:Required
	CacheResults bool `json:"cacheResults"`

	// +optional
	// +kubebuilder:validation:Optional
	ResultsCache *ResultsCacheConfig `json:"resultsCache,omitempty"`
}

type ResultsCacheConfig struct {
	// +optional
	// +kubebuilder:validation:Optional
	Cache *CacheConfig `json:"cache,omitempty"`
}

type CacheConfig struct {
	// +optional
	// +kubebuilder:validation:Optional
	Background *CacheBackgroundConfig `json:"background,omitempty"`

	// +kubebuilder:validation:Required
	DefaultValidity string `json:"defaultValidity"`

	// +optional
	// +kubebuilder:validation:Optional
	MemcachedClient *MemcachedClientConfig `json:"memcachedClient,omitempty"`
}

// Ruler 설정 구조체
type RulerConfig struct {
	// +optional
	// +kubebuilder:validation:Optional
	Storage *RulerStorageConfig `json:"storage,omitempty"`
}

type RulerStorageConfig struct {
	// +kubebuilder:validation:Required
	Type string `json:"type"`

	// +optional
	// +kubebuilder:validation:Optional
	S3 *RulerS3Config `json:"s3,omitempty"`
}

type RulerS3Config struct {
	// +kubebuilder:validation:Required
	BucketNames string `json:"bucketnames"`
}

// RuntimeConfig 설정 구조체
type RuntimeConfig struct {
	// +kubebuilder:validation:Required
	File string `json:"file"`
}

// SchemaConfig 설정 구조체
type SchemaConfig struct {
	// +kubebuilder:validation:Required
	Configs []SchemaConfigEntry `json:"configs"`
}

type SchemaConfigEntry struct {
	// +kubebuilder:validation:Required
	From string `json:"from"`

	// +optional
	// +kubebuilder:validation:Optional
	Index *SchemaConfigIndex `json:"index,omitempty"`

	// +kubebuilder:validation:Required
	ObjectStore string `json:"objectStore"`

	// +kubebuilder:validation:Required
	Schema string `json:"schema"`

	// +kubebuilder:validation:Required
	Store string `json:"store"`
}

type SchemaConfigIndex struct {
	// +kubebuilder:validation:Required
	Period string `json:"period"`

	// +kubebuilder:validation:Required
	Prefix string `json:"prefix"`
}

// Server 설정 구조체
type ServerConfig struct {
	// +kubebuilder:validation:Required
	GRPCListenPort int `json:"grpcListenPort"`

	// +kubebuilder:validation:Required
	HTTPListenPort int `json:"httpListenPort"`

	// +kubebuilder:validation:Required
	HTTPServerReadTimeout string `json:"httpServerReadTimeout"`

	// +kubebuilder:validation:Required
	HTTPServerWriteTimeout string `json:"httpServerWriteTimeout"`
}

// StorageConfig 설정 구조체
type StorageConfig struct {
	// +optional
	// +kubebuilder:validation:Optional
	BloomShipper *BloomShipperConfig `json:"bloomShipper,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	BoltDBShipper *BoltDBShipperConfig `json:"boltdbShipper,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	Hedging *HedgingConfig `json:"hedging,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	TSDBShipper *TSDBShipperConfig `json:"tsdbShipper,omitempty"`
}

type BloomShipperConfig struct {
	// +kubebuilder:validation:Required
	WorkingDirectory string `json:"workingDirectory"`
}

type BoltDBShipperConfig struct {
	// +optional
	// +kubebuilder:validation:Optional
	IndexGatewayClient *IndexGatewayClientConfig `json:"indexGatewayClient,omitempty"`
}

type IndexGatewayClientConfig struct {
	// +kubebuilder:validation:Required
	ServerAddress string `json:"serverAddress"`
}

type HedgingConfig struct {
	// +kubebuilder:validation:Required
	At string `json:"at"`

	// +kubebuilder:validation:Required
	MaxPerSecond int `json:"maxPerSecond"`

	// +kubebuilder:validation:Required
	UpTo int `json:"upTo"`
}

type TSDBShipperConfig struct {
	// +optional
	// +kubebuilder:validation:Optional
	IndexGatewayClient *IndexGatewayClientConfig `json:"indexGatewayClient,omitempty"`
}

// Tracing 설정 구조체
type TracingConfig struct {
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`
}

// SsdLokiSpec 정의
type SsdLokiSpec struct {
	// +kubebuilder:validation:Required
	AuthEnabled bool `json:"authEnabled"`

	// +optional
	// +kubebuilder:validation:Optional
	BloomBuild *BloomBuild `json:"bloomBuild,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	BloomGateway *BloomGateway `json:"bloomGateway,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	ChunkStoreConfig *ChunkStoreConfig `json:"chunkStoreConfig,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	Common *CommonConfig `json:"common,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	Frontend *FrontendConfig `json:"frontend,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	FrontendWorker *FrontendWorkerConfig `json:"frontendWorker,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	IndexGateway *IndexGatewayConfig `json:"indexGateway,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	Ingester *IngesterConfig `json:"ingester,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	LimitsConfig *LimitsConfig `json:"limitsConfig,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	Memberlist *MemberlistConfig `json:"memberlist,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	PatternIngester *PatternIngesterConfig `json:"patternIngester,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	Querier *QuerierConfig `json:"querier,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	QueryRange *QueryRangeConfig `json:"queryRange,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	Ruler *RulerConfig `json:"ruler,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	RuntimeConfig *RuntimeConfig `json:"runtimeConfig,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	SchemaConfig *SchemaConfig `json:"schemaConfig,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	Server *ServerConfig `json:"server,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	StorageConfig *StorageConfig `json:"storageConfig,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	Tracing *TracingConfig `json:"tracing,omitempty"`
}

// SsdLokiStatus defines the observed state of SsdLoki
type SsdLokiStatus struct {
	// +optional
	// +kubebuilder:validation:Optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	Phase string `json:"phase,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	ComponentStatuses *ComponentStatuses `json:"componentStatuses,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	Message string `json:"message,omitempty"`
}

type ComponentStatuses struct {
	// +optional
	// +kubebuilder:validation:Optional
	Ingester ComponentStatus `json:"ingester,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	Querier ComponentStatus `json:"querier,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	Distributor ComponentStatus `json:"distributor,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	QueryFrontend ComponentStatus `json:"queryFrontend,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	Ruler ComponentStatus `json:"ruler,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	Compactor ComponentStatus `json:"compactor,omitempty"`
}

type ComponentStatus struct {
	// +optional
	// +kubebuilder:validation:Optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	UpdatedReplicas int32 `json:"updatedReplicas,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	DesiredReplicas int32 `json:"desiredReplicas,omitempty"`

	// +optional
	// +kubebuilder:validation:Optional
	AvailableReplicas int32 `json:"availableReplicas,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SsdLoki는 ssdlokis API의 스키마입니다
type SsdLoki struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec SsdLokiSpec `json:"spec"`

	// +optional
	// +kubebuilder:validation:Optional
	Status SsdLokiStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SsdLokiList는 SsdLoki의 목록을 포함합니다
type SsdLokiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsdLoki `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SsdLoki{}, &SsdLokiList{})
}
