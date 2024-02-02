package initializers

import (
	"log"
	"os"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/salehWeb/chat-app/server/src/helpers"
)

func InitClient() {
	helpers.CopyDir("./client/public", "./dist/assets")
	helpers.CopyFile("./client/index.html", "./dist/index.html")

	context, err := api.Context(api.BuildOptions{
		EntryPoints: []string{"./client/src/main.tsx"},
		Splitting:   true,
		KeepNames:   true,
		AssetNames:  "[name]",
		ChunkNames:  "[name]",
		Loader: map[string]api.Loader{
			".png": api.LoaderFile,
			".svg": api.LoaderFile,
		},
		Format:            api.FormatESModule,
		LogLevel:          api.LogLevelInfo,
		Outdir:            "./dist/assets",
		Bundle:            true,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Write:             true,
	})

	if err != nil && len(err.Errors) != 0 {
		for i := 0; i < len(err.Errors); i++ {
			log.Println(err.Errors[i])
		}
		os.Exit(1)
	}

	e := context.Watch(api.WatchOptions{})

	if e != nil {
		log.Fatal(e)
	}
}
