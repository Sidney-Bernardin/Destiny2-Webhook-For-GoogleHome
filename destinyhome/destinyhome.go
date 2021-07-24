package destinyhome

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"root/destinyhome/repository"

	"cloud.google.com/go/logging"
	"github.com/pkg/errors"
)

func DestinyHome(w http.ResponseWriter, r *http.Request) {

	var (
		req webhookRequest
		res webhookResponse
	)

	// Decode request body.
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// Lookup the operation.
	operation, ok := operations[req.Handler.Name]
	if !ok {
		http.Error(w, "operation not found", http.StatusNotFound)
		return
	}

	// Put the resolved params into a map.
	var params map[string]string
	for k, v := range req.Intent.Params {
		params[k] = v.Resolved
	}

	// Do the operation.
	var err error
	res.Prompt.FirstSimple.Speech, err = operation(params)
	if err != nil {

		errMap := map[error]int{
			repository.ErrUserNotFound: http.StatusNotFound,
		}

		if code, ok := errMap[err]; ok {

			// Setup the response.
			res.Prompt.FirstSimple.Speech = errors.Cause(err).Error()
			w.WriteHeader(code)
		} else {

			// Encode the error.
			b, err := json.Marshal(err)
			if err != nil {

				// Log a warning.
				log.Println(logging.Entry{
					Severity: logging.Warning,
					Payload:  errors.Wrap(err, "cannot marshal error").Error(),
				})

				// Log the error as a string.
				log.Println(logging.Entry{
					Severity: logging.Critical,
					Payload:  fmt.Sprintf("%#v", err),
				})
			}

			// Log the error.
			log.Println(logging.Entry{
				Severity: logging.Critical,
				Payload:  b,
			})

			// Setup the response.
			res.Prompt.FirstSimple.Speech = "My backend systems are not working right now, try again later."
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	// Setup the response.
	res.Session.ID = req.Session.ID
	res.Scene.Name = req.Scene.Name

	// Encode the result.
	b, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Respond.
	r.Header.Set("Content-Type", "application/json")
	_, err = w.Write(b)
	if err != nil {
		log.Println(err.Error())
	}
}
