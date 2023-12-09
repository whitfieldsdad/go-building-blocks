package bb

import "github.com/denisbrodbeck/machineid"

var (
	AppId = "d5be1754-035e-4242-9bbe-debc86154f8b"
)

func GetHostId() (string, error) {
	hostId, err := machineid.ProtectedID(AppId)
	if err != nil {
		return "", err
	}
	m := map[string]interface{}{
		"host_id": hostId,
	}
	return NewUUID5FromMap(AppId, m)
}
