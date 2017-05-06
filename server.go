package main

import (
	"net/http"
	"log"
	"io"
	"bytes"
	"strconv"
	"errors"
	"k-means/k_means"
	"encoding/json"
	"k-means/templates"
)

const ALLOWED_FILE_TYPE = "text/plain; charset=utf-8"

const (
	clusterNumberParam = "clusters"
	iterationNumberParam = "iterations"
	distanceMethodParam = "distance"
)

func paramToInt(param string) (int, error) {
	if param == "" {
		return 0, errors.New(" param is empty")
	}

	return strconv.Atoi(param)
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseMultipartForm(32 << 20)

		// Fetching process params
		clustersParam, err := paramToInt(r.Form.Get(clusterNumberParam))
		if err != nil {
			http.Error(w,
				clusterNumberParam + err.Error(),
				http.StatusBadRequest)
			return
		} else if clustersParam < 0 {
			http.Error(w,
				clusterNumberParam + " is negative or zero ",
				http.StatusBadRequest)
			return
		}

		iterationsParam, err := paramToInt(r.Form.Get(iterationNumberParam))
		if err != nil {
			http.Error(w,
				iterationNumberParam + err.Error(),
				http.StatusBadRequest)
			return
		} else if iterationsParam < 0 {
			http.Error(w,
				iterationNumberParam + " is negative or zero ",
				http.StatusBadRequest)
			return
		}

		var distanceParam = r.Form.Get(distanceMethodParam)
		var distanceMethod k_means.DistanceFunc

		switch distanceParam {
		case k_means.EuclidDistance:
			distanceMethod = k_means.EuclidDistanceFunc
		default:
			http.Error(w,
				distanceMethodParam + " is unknown",
				http.StatusBadRequest)
			return
		}

		// Fetching file
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w,
				err.Error(),
				http.StatusInternalServerError)

			return
		}

		defer file.Close()

		var Buf bytes.Buffer
		io.Copy(&Buf, file)

		// Checking it's content type
		if http.DetectContentType(Buf.Bytes()) != ALLOWED_FILE_TYPE {
			http.Error(w,
				http.StatusText(http.StatusUnsupportedMediaType),
				http.StatusUnsupportedMediaType)
			return
		}

		// Processing file
		var claims Claims
		err = json.Unmarshal(Buf.Bytes(), &claims)
		if err != nil {
			http.Error(w,
				err.Error(),
				http.StatusInternalServerError)

			return
		}

		Buf.Reset()

		points, err := ClaimsToPoints(claims)
		if err != nil {
			http.Error(w,
				err.Error(),
				http.StatusInternalServerError)

			return
		}

		log.Println("new data request accepted...")


		cls, err := k_means.Calc(points, int32(clustersParam), int32(iterationsParam), distanceMethod)
		if err != nil {
			http.Error(w,
				err.Error(),
				http.StatusInternalServerError)
			return
		}

		data, err := json.MarshalIndent(&cls, "  ", "    ")
		if err != nil {
			http.Error(w,
				err.Error(),
				http.StatusInternalServerError)
			return
		}

		template, err := templates.GetTemplateWithData("map.html", struct{ Data string }{Data: string(data)})
		if err != nil {
			http.Error(w,
				err.Error(),
				http.StatusInternalServerError)
			return
		}

		w.Write(template)
	} else {
		http.Error(w,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed)
	}
}

func Start() {
	log.Println("starting server on port localhost:9090")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/form.html")
	})

	http.HandleFunc("/process", processHandler)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("FATAL:", err.Error())
	}
}