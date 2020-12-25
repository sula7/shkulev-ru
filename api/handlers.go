package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"shkulev-ru/utils"
)

type (
	Handlers interface {
		ping(w http.ResponseWriter, req *http.Request)

		validate(w http.ResponseWriter, req *http.Request)
		fix(w http.ResponseWriter, req *http.Request)
	}

	Response struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}

	Validity struct {
		Value   string `json:"value"`
		IsValid bool   `json:"is_valid"`
	}
)

func (api *API) ping(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "OK")
	if err != nil {
		log.Println("error sending response: ", err)
		return
	}
}

func (api *API) validate(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		err := json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "only POST method is supported",
		})
		if err != nil {
			log.Println("can not send response:", err)
			return
		}
		return
	}

	for _, headerValue := range req.Header["Content-Type"] {
		if headerValue != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)

			err := json.NewEncoder(w).Encode(Response{
				Success: false,
				Message: "Content-Type should be an application/json",
			})
			if err != nil {
				log.Println("can not send response:", err)
				return
			}
			return
		}
	}

	values := []string{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("can not read the request body:", err)

		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: err.Error(),
		})
		if err != nil {
			log.Println("can not send response:", err)
			return
		}
		return
	}

	err = json.Unmarshal(body, &values)
	if err != nil {
		log.Println("can not unmarshal the request body:", err)

		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: err.Error(),
		})
		if err != nil {
			log.Println("can not send response:", err)
			return
		}
		return
	}

	if len(values) > 19 {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "more than 20 values",
		})
		if err != nil {
			log.Println("can not send response:", err)
			return
		}
		return
	}

	validities := []Validity{}

	for i := 0; i < len(values); i++ {
		validity := Validity{
			Value:   values[i],
			IsValid: utils.ValidateBrackets(values[i]),
		}
		validities = append(validities, validity)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    validities,
	})
	if err != nil {
		log.Println("can not send response:", err)
		return
	}
}

func (api *API) fix(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		err := json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "only POST method is supported",
		})
		if err != nil {
			log.Println("can not send response:", err)
			return
		}
		return
	}

	for _, headerValue := range req.Header["Content-Type"] {
		if headerValue != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)

			err := json.NewEncoder(w).Encode(Response{
				Success: false,
				Message: "Content-Type should be an application/json",
			})
			if err != nil {
				log.Println("can not send response:", err)
				return
			}
			return
		}
	}

	values := []string{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("can not read the request body:", err)

		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: err.Error(),
		})
		if err != nil {
			log.Println("can not send response:", err)
			return
		}
		return
	}

	err = json.Unmarshal(body, &values)
	if err != nil {
		log.Println("can not unmarshal the request body:", err)

		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: err.Error(),
		})
		if err != nil {
			log.Println("can not send response:", err)
			return
		}
		return
	}

	if len(values) > 19 {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "more than 20 values",
		})
		if err != nil {
			log.Println("can not send response:", err)
			return
		}
		return
	}

	fixes := []map[string]string{}

	for i := 0; i < len(values); i++ {
		fix := map[string]string{
			"value":       values[i],
			"fixed_value": utils.Fix(values[i]),
		}
		fixes = append(fixes, fix)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    fixes,
	})
	if err != nil {
		log.Println("can not send response:", err)
		return
	}
}
