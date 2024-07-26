package request

type CreateCardRequest struct {
	Type       string `json:"type"`
	Number     string `json:"number"`
	Alias      string `json:"name"`
	Expiration string `json:"expiration"`
}
