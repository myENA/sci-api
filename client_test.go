package api_test

import (
	"context"
	"crypto/tls"
	"flag"
	"github.com/myENA/sci-api"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	host     string
	port     int
	username string
	password string
)

func init() {
	flag.StringVar(&host, "host", "", "Hostname of SCI instance")
	flag.IntVar(&port, "port", 443, "HTTP port to use")
	flag.StringVar(&username, "username", "", "API Username")
	flag.StringVar(&password, "password", "", "API Password")
	flag.Parse()

	if host == "" {
		log.Println("host cannot be empty")
		os.Exit(1)
	}
	if 0 >= port {
		log.Println("port must be >= 0")
		os.Exit(1)
	}
	if username == "" {
		log.Println("username cannot be empty")
		os.Exit(1)
	}
	if password == "" {
		log.Println("password cannot be empty")
		os.Exit(1)
	}

	api.EnableDebug()
}

func testClient(t *testing.T) *api.Client {
	client, err := api.NewClient(
		api.DefaultConfig(host),
		api.NewPasswordAuthenticator(username, password),
		&http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				DisableKeepAlives:     true,
				MaxIdleConnsPerHost:   -1,
			},
		},
	)
	if err != nil {
		t.Logf("Error creating client: %s", err)
		t.FailNow()
	}
	return client
}

func TestClient(t *testing.T) {
	var ctx context.Context
	var cancel context.CancelFunc

	client := testClient(t)

	jan1 := time.Date(2018, 01, 01, 0, 0, 0, 0, time.Local)
	jan3 := time.Date(2018, 01, 03, 0, 0, 0, 0, time.Local)

	t.Run("Synchronous", func(t *testing.T) {
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		_, _, err := client.Reports().GetAPMACFacetData(ctx, 4, &api.Query{
			Start: jan1,
			End:   jan3,
		})
		cancel()
		if err != nil {
			t.Logf("Unable to execute query: %s", err)
			t.FailNow()
		}

		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		_, _, err = client.Users().Logout(ctx)
		cancel()
		if err != nil {
			t.Logf("Unable to logout: %s", err)
			t.FailNow()
		}

		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		_, _, err = client.Reports().GetAPMACFacetData(ctx, 4, &api.Query{
			Start: jan1,
			End:   jan3,
		})
		cancel()
		if err != nil {
			t.Logf("Unable to execute query: %s", err)
			t.FailNow()
		}

		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		_, _, err = client.Users().Logout(ctx)
		cancel()
		if err != nil {
			t.Logf("Unable to logout: %s", err)
			t.FailNow()
		}
	})

	// This is mostly here to find issues traceable with -race, and to test the robustness of your vsz setup :)
	t.Run("Asynchronous", func(t *testing.T) {
		start := make(chan struct{})

		wg := new(sync.WaitGroup)
		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func(routine int, client *api.Client, start <-chan struct{}) {
				var ctx context.Context
				var cancel context.CancelFunc
				var defs api.ReportsListDefinitions200ResponseSlice
				var err error
				<-start
				for i := 0; i < 5; i++ {
					if t.Failed() {
						break
					}
					ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
					_, defs, err = client.Reports().ListDefinitions(ctx, nil)
					cancel()
					if err != nil {
						t.Logf("Routine %d loop %d query 1 saw error: %+v", routine, i, err)
					} else if len(defs) == 0 {
						t.Logf("Routine %d loop %d returned an empty zone list, expected at least 1: %+v", routine, i, defs)
					}

					if routine == 0 {
						ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
						_, _, err = client.Users().Logout(ctx)
						cancel()
						if err != nil {
							t.Logf("Routine %d loop %d query 2 saw error: %+v", routine, i, err)
						}
					}

					ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
					_, _, err = client.Systems().Count(ctx)
					cancel()
					if err != nil {
						t.Logf("Routine %d loop %d query 3 saw error: %+v", routine, i, err)
					}
				}
				wg.Done()
			}(i, client, start)
		}

		close(start)
		wg.Wait()
	})
}
