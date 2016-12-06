package server

import (
  "encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"io/ioutil"

	"vscale-metadata/utils"
	"github.com/gorilla/mux"
	sjson "github.com/tidwall/sjson"
	vscale_api_go "github.com/vscale/go-vscale"
)

type Server struct {
	config   *Config
}

// Interface Server interface
type Interface interface {
	// start Server
	Serve() error
}

// NewServer return new instance of Server
func NewServer(c *Config) Interface {
	return &Server{
		config: c,
	}
}

// Serve start Server serve service
func (s *Server) Serve() error {
	mx := mux.NewRouter()
	mx.HandleFunc("/metadata/v1/", s.getMetaDataHandler).Methods("GET")
	utils.Log().Infof("[Metadata service] start on:", s.config.getAddr())
	utils.Log().Fatal(http.ListenAndServe(s.config.getAddr(), mx))
	return nil
}

func (s *Server) getMetaDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json_result := scalet_info(s.config.ApiKey)
  json_result, _ = sjson.Set(json_result, "user-data", s.config.getUserData())
	w.Write([]byte(json_result))
}

func scalet_info(api_key string) string{
	if api_key != "" {
		content, err := ioutil.ReadFile("/proc/sys/kernel/hostname")
		if err != nil {
			return ""
		}
		scalet_id := regexp.MustCompile(`[0-9]+`).FindString(string(content))
		if scalet_id != "" {
			id, err := strconv.ParseInt(scalet_id, 10, 64)
			if err != nil {
				return ""
			}
			vscale_client := vscale_api_go.NewClient(api_key)
			scalet, _, err := vscale_client.Scalet.Get(id)
			if err != nil {
				return ""
			}else{
				b, err := json.Marshal(scalet)
		    if err != nil {
		        return ""
		    }
				return string(b)
			}
		}
	}
	return ""	
}
