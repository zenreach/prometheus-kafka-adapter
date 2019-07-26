// Copyright 2018 Telefónica
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "github.com/prometheus/client_golang/prometheus"

var (
	queueSize = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "kafka_queue_size",
			Help: "Queue size for metrics sent to Kafka",
		})
	httpRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Count of all http requests",
		})
	metricsBlocked = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "prometheus_kafka_adapter_metrics_blocked",
			Help: "Count of metrics blocked due to whitelist",
		})
	metricsProduced = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "prometheus_kafka_adapter_metrics_produced",
			Help: "Count of metrics forwarded to kafka",
		})
	produceErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "prometheus_kafka_adapter_produce_errors",
			Help: "Count of errors producing to kafka",
		})

)

func init() {
	prometheus.MustRegister(queueSize)
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(metricsProduced)
	prometheus.MustRegister(metricsBlocked)
	prometheus.MustRegister(produceErrors)
}
