package main

import (
	"fmt"
	"github.com/MaibornWolff/elcep/main/config"
	"github.com/olivere/elastic"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloName(t *testing.T) {
	response := `{
  "took": 2,
  "timed_out": false,
  "_shards": {
    "total": 10,
    "successful": 10,
    "skipped": 0,
    "failed": 0
  },
  "hits": {
    "total": 690,
    "max_score": 0.0021378943,
    "hits": [
      {
        "_index": "myindex",
        "_type": "mydoc",
        "_id": "UMg-gHwBvceD-auiSvii",
        "_score": 0.0021378943,
        "_source": {
          "@timestamp": "2021-10-14T19:18:58.454Z",
          "log": "Exception this is a sample log message",
          "bucket": true,
          "kubernetes.app/name": "testapp",
          "key2": 3,
          "somefield": "somevalue"
        }
      },
      {
        "_index": "myindex",
        "_type": "mydoc",
        "_id": "Y8g-gHwBvceD-auicPiH",
        "_score": 0.0021378943,
        "_source": {
          "@timestamp": "2021-10-14T19:19:08.152Z",
          "log": "Exception this is a sample log message",
          "bucket": true,
          "kubernetes.app/name": "testapp2",
          "key2": 1,
          "somefield": "somevalue"
        }
      },
      {
        "_index": "myindex",
        "_type": "mydoc",
        "_id": "mMg_gHwBvceD-auiHPjW",
        "_score": 0.0021378943,
        "_source": {
          "@timestamp": "2021-10-14T19:19:52.260Z",
          "log": "Exception this is a sample log message",
          "bucket": true,
          "kubernetes.app/name": "testapp",
          "key2": 3,
          "somefield": "somevalue"
        }
      },
      {
        "_index": "myindex",
        "_type": "mydoc",
        "_id": "s8g_gHwBvceD-auiXfid",
        "_score": 0.0021378943,
        "_source": {
          "@timestamp": "2021-10-14T19:20:08.847Z",
          "log": "Exception this is a sample log message",
          "bucket": true,
          "kubernetes.app/name": "testapp",
          "key2": 2,
          "somefield": "somevalue"
        }
      },
      {
        "_index": "myindex",
        "_type": "mydoc",
        "_id": "wMg_gHwBvceD-auifvgU",
        "_score": 0.0021378943,
        "_source": {
          "@timestamp": "2021-10-14T19:20:17.121Z",
          "log": "Exception this is a sample log message",
          "bucket": true,
          "kubernetes.app/name": "testapp",
          "key2": 3,
          "somefield": "somevalue"
        }
      },
      {
        "_index": "myindex",
        "_type": "mydoc",
        "_id": "zcg_gHwBvceD-auir_ih",
        "_score": 0.0021378943,
        "_source": {
          "@timestamp": "2021-10-14T19:20:29.843Z",
          "log": "Exception this is a sample log message",
          "bucket": true,
          "kubernetes.app/name": "testapp2",
          "key2": 3,
          "somefield": "somevalue"
        }
      },
      {
        "_index": "myindex",
        "_type": "mydoc",
        "_id": "R8hBgHwBvceD-auiJ_kE",
        "_score": 0.0021378943,
        "_source": {
          "@timestamp": "2021-10-14T19:22:05.943Z",
          "log": "Exception this is a sample log message",
          "bucket": true,
          "kubernetes.app/name": "testapp",
          "key2": 3,
          "somefield": "somevalue"
        }
      },
      {
        "_index": "myindex",
        "_type": "mydoc",
        "_id": "SMhBgHwBvceD-auiJ_km",
        "_score": 0.0021378943,
        "_source": {
          "@timestamp": "2021-10-14T19:22:05.978Z",
          "log": "Exception this is a sample log message",
          "bucket": true,
          "kubernetes.app/name": "testapp2",
          "key2": 2,
          "somefield": "somevalue"
        }
      },
      {
        "_index": "myindex",
        "_type": "mydoc",
        "_id": "9shDgHwBvceD-auiEPmu",
        "_score": 0.0021378943,
        "_source": {
          "@timestamp": "2021-10-14T19:24:11.298Z",
          "log": "Exception this is a sample log message",
          "bucket": true,
          "kubernetes.app/name": "testapp",
          "key2": 1,
          "somefield": "somevalue"
        }
      },
      {
        "_index": "myindex",
        "_type": "mydoc",
        "_id": "98hDgHwBvceD-auiEfny",
        "_score": 0.0021378943,
        "_source": {
          "@timestamp": "2021-10-14T19:24:11.622Z",
          "log": "Exception this is a sample log message",
          "bucket": true,
          "kubernetes.app/name": "testapp",
          "key2": 2,
          "somefield": "somevalue"
        }
      }
    ]
  },
  "aggregations": {
    "kubernetes.app/name": {
      "doc_count_error_upper_bound": 0,
      "sum_other_doc_count": 0,
      "buckets": [
        {
          "key": "testapp",
          "doc_count": 191,
          "key2": {
            "doc_count_error_upper_bound": 0,
            "sum_other_doc_count": 0,
            "buckets": [
              {
                "key": 2,
                "doc_count": 52
              },
              {
                "key": 0,
                "doc_count": 50
              },
              {
                "key": 3,
                "doc_count": 47
              },
              {
                "key": 1,
                "doc_count": 42
              }
            ]
          }
        },
        {
          "key": "testapp2",
          "doc_count": 173,
          "key2": {
            "doc_count_error_upper_bound": 0,
            "sum_other_doc_count": 0,
            "buckets": [
              {
                "key": 3,
                "doc_count": 48
              },
              {
                "key": 2,
                "doc_count": 47
              },
              {
                "key": 1,
                "doc_count": 46
              },
              {
                "key": 0,
                "doc_count": 32
              }
            ]
          }
        }
      ]
    }
  }
}`

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(response))
	}))

	dummy, err := createElasticSearchClient(testServer.URL)
	if err != nil {
		fmt.Println("Did not work to create the ES Client", err)
	}

	query := config.CreateQuery("myName", "log:exception AND bucket:true")
	addAggregationsToQuery(query)

	// create plugin, Build Prometheus collectors and
	aggregationPlugin := NewPlugin(config.Options{TimeKey: "@timestamp"}, nil)
	collectors := aggregationPlugin.BuildMetrics([]config.Query{query})
	aggregationPlugin.Perform(dummy)

	for _, collector := range collectors {
		countervec, _ := collector.(*prometheus.CounterVec)
		value := testutil.ToFloat64(countervec.With(prometheus.Labels{"kubernetes_app_name": "testapp", "key2": "0"}))

		log.Print("tihs is our real structure: ", countervec, " with value ", value)
	}

	var examplecountervec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "elcep_logs_matched_myName_buckets",
		Help: "Aggregates logs matching log:exception AND bucket:true to buckets",
	}, []string{"key1", "key2"})

	prometheus.MustRegister(examplecountervec)

	examplecountervec.With(prometheus.Labels{"key1": "0", "key2": "2"}).Add(34)
	value := testutil.ToFloat64(examplecountervec.With(prometheus.Labels{"key1": "0", "key2": "2"}))
	log.Println("this is our fake structure ", examplecountervec, "with value ", value)

}

func addAggregationsToQuery(query config.Query) {
	var aggregations [2]string
	aggregations[0] = "kubernetes.app/name"
	aggregations[1] = "key2"
	b := make([]interface{}, len(aggregations))
	for i := range aggregations {
		b[i] = aggregations[i]
	}
	query["aggregations"] = b
}

func createElasticSearchClient(endpoint string) (*elastic.Client, error) {
	var client, err = elastic.NewClient(
		elastic.SetURL(endpoint),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	return client, err
}
