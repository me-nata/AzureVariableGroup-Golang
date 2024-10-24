package azure

import (
	"time"
)

type Project struct {
    ID             string    `json:"id"`
    Name           string    `json:"name"`
    Description    string    `json:"description"` 
    URL            string    `json:"url"`
    State          string    `json:"state"`
    Revision       int       `json:"revision"`
    Visibility     string    `json:"visibility"`
    LastUpdateTime time.Time `json:"lastUpdateTime"`
}

type Projects []Project