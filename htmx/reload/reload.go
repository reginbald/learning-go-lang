package reload

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

const script = `<script>
	function reload(retry) {
		if (typeof(EventSource) !== "undefined") {
			const es = new EventSource("/__reload");
			es.onmessage = function(event) {
				if (retry && event.data === "ping") {
					location.reload();
				}
			}
			es.onerror = function(err) {
				console.error("ES:", err, err.target);
				es.close();
				setTimeout(function() { reload(true); }, 1000);
			};
		}
	}
	reload(false);
</script>`

// HTTP middleware setting a value on the request context
func ReloadMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/__reload" {
			reload(w, r)
			return
		}

		// set headers first so that they're sent with the initial response
		// disable caching, because http.FileServer will
		// send Last-Modified headers, prompting the browser to cache it
		w.Header().Set("Cache-Control", "no-cache")

		body := &bytes.Buffer{}
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		// copy body so that we can sniff the content type
		ww.Tee(body)

		next.ServeHTTP(ww, r)

		contentType := w.Header().Get("Content-Type")

		if contentType == "" {
			contentType = http.DetectContentType(body.Bytes())
		}

		if strings.HasPrefix(contentType, "text/html") {
			// just append the script to the end of the document
			// this is invalid HTML, but browsers will accept it anyways
			// should be fine for development purposes
			w.Write([]byte(script))
		}
	})
}

func reload(w http.ResponseWriter, r *http.Request) {
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	for {
		io.WriteString(w, "data: ping\n\n")
		f.Flush()
		time.Sleep(2 * time.Second)
	}
}
