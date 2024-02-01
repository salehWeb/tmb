package main

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/salehWeb/chat-app/server/src/initializers"
	"github.com/salehWeb/chat-app/server/src/routes"
	"github.com/salehWeb/chat-app/server/src/socket"
)

type gzipResponseWriter struct {
	gw *gzip.Writer
	http.ResponseWriter
}

func (grw gzipResponseWriter) Write(b []byte) (int, error) {
	return grw.gw.Write(b)
}

func gzipHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Content-Encoding", "gzip")

			gzw := gzip.NewWriter(w)
			defer gzw.Close()

			w = gzipResponseWriter{gzw, w}
		}

		next.ServeHTTP(w, r)
	})
}

var clientPaths = map[string]bool{
	"/login":   true,
	"/sign-up": true,
	"/chat":    true,
	"/":        true,
}

func readFile(filename string) ([]byte, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content
	content := make([]byte, 0)
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if n == 0 || err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		content = append(content, buffer[:n]...)
	}

	return content, nil
}

func handelClient(next *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n hello pro i run and path is %s \n", r.URL.Path)
		if clientPaths[r.URL.Path] {
			htmlContent, err := readFile("./dist/index.html")
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "text/html")
			w.Header().Set("Custom-Header", "Custom Value")
			w.WriteHeader(http.StatusOK)
			w.Write(htmlContent)

		} else {
			log.Println("i should run the fuck up")
			next.ServeHTTP(w, r)
		}
	})
}

func main() {
	// copyDir("../client/public", "./dist/assets")
	// copyFile("../client/index.html", "./dist/index.html")

	// context, err := api.Context(api.BuildOptions{
	// 	EntryPoints: []string{"../client/src/main.tsx"},
	// 	Splitting:   true,
	// 	KeepNames:   true,
	// 	AssetNames:  "[name]",
	// 	ChunkNames:  "[name]",
	// 	Loader: map[string]api.Loader{
	// 		".png": api.LoaderFile,
	// 		".svg": api.LoaderFile,
	// 	},
	// 	Format:            api.FormatESModule,
	// 	LogLevel:          api.LogLevelInfo,
	// 	Outdir:            "./dist/assets",
	// 	Bundle:            true,
	// 	MinifyWhitespace:  true,
	// 	MinifyIdentifiers: true,
	// 	MinifySyntax:      true,
	// 	Write:             true,
	// })

	// if err != nil && len(err.Errors) != 0 {
	// 	for i := 0; i < len(err.Errors); i++ {
	// 		log.Println(err.Errors[i])
	// 	}
	// 	os.Exit(1)
	// }

	// e := context.Watch(api.WatchOptions{})

	// if e != nil {
	// 	log.Fatal(e)
	// }

	initializers.GetENV()

	mux := http.NewServeMux()

	// mux1 := http.NewServeMux()
	// mux1.Handle("/", gzipHandler(mux))

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./dist"))))
	mux.Handle("/api", routes.HandelRoutes())
	mux.Handle("/ws", socket.UseSocket())

	http.Handle("/", handelClient(mux))

	initializers.Listen(nil)
}

func copyDir(src, dest string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dest, 0755); err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, destPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
