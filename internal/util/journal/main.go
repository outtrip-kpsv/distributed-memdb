package journal

import (
	"crypto/sha256"
	"encoding/base64"
	"hash/fnv"
	"math/rand"
	"strings"
	"team01/internal/proto/node"
	"team01/internal/util"
)

type Journal struct {
	data map[string]string // [hashUUID] hashCluster
}

func NewJournal() *Journal {
	return &Journal{data: make(map[string]string)}
}

func fnv64(text string) int64 {
	algorithm := fnv.New64a()
	_, err := algorithm.Write([]byte(text))
	if err != nil {
		return 0
	}
	return int64(algorithm.Sum64())
}

func (j *Journal) HashRequest(req string, nodes *node.KnownNodes) []string {
	uuidHash := hashUUID(req)
	clusterHash, addresses := hashCluster(nodes)
	fullHash := hashStr(uuidHash + clusterHash)

	//TODO проверить соответствие и продумать дальнейшую логику
	j.data[uuidHash] = clusterHash

	rnd := rand.New(rand.NewSource(fnv64(fullHash)))
	rnd.Shuffle(len(addresses), func(i, j int) {
		addresses[i], addresses[j] = addresses[j], addresses[i]
	})
	return addresses
}

func (j *Journal) GetClusterHash(uuid string) string {
	if val, ok := j.data[hashUUID(uuid)]; ok {
		return val
	}
	return ""
}

func hashUUID(req string) string {
	return hashStr(req)

}

func hashCluster(nodes *node.KnownNodes) (string, []string) {
	text := util.KnowNodesToString(nodes)
	return hashStr(strings.Join(text, "-")), text
}

func hashStr(txt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(txt))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
