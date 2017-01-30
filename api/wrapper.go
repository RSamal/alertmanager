package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/prometheus/alertmanager/silence"
	"github.com/prometheus/alertmanager/types"
)

// SilenceRequest is a request message for create and update Silence Batch API
type SilenceRequest struct {

	// Information about the hosts which needs to be silenced
	Hosts Hosts `json:"hosts"`

	// Time range of the silences.
	//
	// * StartsAt must not be before creation time
	// * EndsAt must be after StartsAt
	// * Deleting a silence means to set EndsAt to now
	// * Time range must not be modified in different ways
	//
	StartsAt time.Time `json:"startsAt"`
	EndsAt   time.Time `json:"endsAt"`

	// Information about who created the silence for which reason.
	CreatedBy string `json:"createdBy"`
	Comment   string `json:"comment,omitempty"`
}

// Hosts contains list of host, that is received in the input request.
// This is request message for delete Silence Batch API
type Hosts []Host

// Host contains the host specific information
type Host struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	IP   string `json:"ip"`
}

type CommonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var re *regexp.Regexp

func init() {
	re, _ = regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
}

func (api *API) addSilenceBatch(w http.ResponseWriter, r *http.Request) {
	var reqs SilenceRequest
	if err := receive(r, &reqs); err != nil {
		respondError(w, apiError{
			typ: errorBadData,
			err: err,
		}, nil)
		return
	}

	if err := validData(reqs.Hosts); err != nil {
		respondError(w, apiError{
			typ: errorBadData,
			err: err,
		}, nil)
		return
	}

	if err := api.createSilencer(&reqs); err != nil {
		api.deleteSilencer(reqs.Hosts)
		respondError(w, apiError{
			typ: errorInternal,
			err: err,
		}, nil)
	}

	json.NewEncoder(w).Encode(reqs.Hosts)
}

func validData(hosts Hosts) error {

	for _, host := range hosts {
		if host.Name == "" {
			return fmt.Errorf("Host name is empty for %s", host.IP)
		}
		if !re.MatchString(host.IP) {
			return fmt.Errorf("Invalid IP address %s", host.IP)
		}
	}
	return nil
}

func (api *API) delSilenceBatch(w http.ResponseWriter, r *http.Request) {
	var reqs Hosts

	if err := receive(r, &reqs); err != nil {
		respondError(w, apiError{
			typ: errorBadData,
			err: err,
		}, nil)
		return
	}

	if err := api.validDeleteData(reqs); err != nil {
		respondError(w, apiError{
			typ: errorBadData,
			err: err,
		}, nil)
		return
	}

	if err := api.deleteSilencer(reqs); err != nil {
		respondError(w, apiError{
			typ: errorInternal,
			err: err,
		}, nil)
		return
	}

	json.NewEncoder(w).Encode(&CommonResponse{
		Status:  "Success",
		Message: "Deleted Silencers",
	})
}

func (api *API) validDeleteData(hosts Hosts) error {

	for _, host := range hosts {
		sils, err := api.silences.Query(silence.QIDs(host.ID))
		if err != nil || len(sils) == 0 {
			fmt.Println(err)
			return fmt.Errorf("%v %v", host.ID, err.Error())

		}
	}
	return nil
}

func (api *API) createSilencer(req *SilenceRequest) error {
	var matchers types.Matchers

	for i, host := range req.Hosts {

		// Add the hostname as a matchers with wildcard "*" for Regular expression
		matchers = append(matchers, &types.Matcher{
			Name:    "host",
			Value:   host.Name + "*",
			IsRegex: true,
		})

		// Add IP as a matchers
		matchers = append(matchers, &types.Matcher{
			Name:    "ip",
			Value:   host.IP,
			IsRegex: false,
		})

		psil, _ := silenceToProto(&types.Silence{
			Matchers:  matchers,
			StartsAt:  req.StartsAt,
			EndsAt:    req.EndsAt,
			CreatedBy: req.CreatedBy,
			Comment:   req.Comment,
		})

		// Drop start time for new silences so we default to now.
		if host.ID == "" && req.StartsAt.Before(time.Now()) {
			psil.StartsAt = nil
		}

		sid, err := api.silences.Create(psil)
		if err != nil {
			return err
		}
		req.Hosts[i].ID = sid
		matchers = nil
	}
	return nil
}

func (api *API) deleteSilencer(hosts Hosts) error {
	for _, host := range hosts {
		if host.ID != "" {
			if err := api.silences.Expire(host.ID); err != nil {
				return err
			}
		}
	}
	return nil
}
