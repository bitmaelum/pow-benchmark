package work

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"math"
	"math/big"
	"strconv"
	"sync"
)

var running int
var mutex = &sync.Mutex{}

/*
 * Proof of work consists of SHA256 hashing data with an additional counter
 *
 *      counter = 0
 *      do
 *          counter += 1
 *          hash = SHA256(data + counter)
 *      until X number of left bits of hash is 0[^1]
 *
 * The idea is that is takes time to calculate a SHA256 from data. Because
 * the SHA256 is distributed evenly, there is no way to know how many bits
 * on the left are 0. By increasing the counter with 1 every time and hashing
 * again, we get many different hashes. As soon as we found a hash with
 * X bits 0 on the left, we use the counter value as the proof for the work
 * that has been done.
 *
 * ^1 Even though we are checking against bits to be 0, the actual check is if
 *    the hash is lower than the target hash.
 */

// ProofOfWork represents a proof-of-work which either can be completed or not
type ProofOfWork struct {
	Bits  int    `json:"bits"`
	Data  string `json:"data"`
	Proof uint64 `json:"proof,omitempty"`
}

// New generates a new ProofOfWork structure.
func New(bits int, data string, proof uint64) *ProofOfWork {
	pow := &ProofOfWork{
		Bits:  bits,
		Data:  data,
		Proof: proof,
	}

	return pow
}

// GenerateWorkData generates random work
func GenerateWorkData() (string, error) {
	data := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, data)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

// HasDoneWork returns true if this instance already has done proof-of-work
func (pow *ProofOfWork) HasDoneWork() bool {
	return pow.Proof > 0
}

// Work actually does the proof-of-work
func (pow *ProofOfWork) Work(cores int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	found := make(chan uint64)
	for i:=0; i<cores; i++ {
		// fmt.Printf("Starting core %d\n", i)
		mutex.Lock()
		running++
		mutex.Unlock()
		go pow.workOnCore(ctx, uint64(i), uint64(cores), found)
	}
	pow.Proof = <- found
	// fmt.Printf("Found proof: %d\n", pow.Proof)
}

func (pow *ProofOfWork) workOnCore(ctx context.Context, start, step uint64, found chan uint64) {
	var hashInt big.Int

	defer func() {
		mutex.Lock()
		running--
		mutex.Unlock()
	}()

	// Hash must be less than this
	target := big.NewInt(1)
	target = target.Lsh(target, uint(256-pow.Bits))

	var counter uint64 = start
	for counter < math.MaxInt64 {
		// Break when context is cancelled (other go thread has found our proof)
		select {
		case <-ctx.Done():
			// fmt.Printf("done with go routine (%d %d)\n", start, step)
			return
		default:
			// fmt.Printf("Hashing %d (core %d) (%d)\n", counter, start, running)

			// 1st round of SHA256
			hash := sha256.Sum256(bytes.Join([][]byte{
				[]byte(pow.Data),
				intToHex(counter),
			}, []byte{}))

			// 2nd round of SHA256
			hash = sha256.Sum256(hash[:])
			hashInt.SetBytes(hash[:])

			// Is it less than our target, then we have done our work
			if hashInt.Cmp(target) == -1 {
				found <- counter
				return
			}
		}

		// Higher, so we must do more work. Increase counter and try again
		counter += step
	}
}

// IsValid returns true when the given work can be validated against the proof
func (pow *ProofOfWork) IsValid() bool {
	var hashInt big.Int

	// 1st round
	hash := sha256.Sum256(bytes.Join([][]byte{
		[]byte(pow.Data),
		intToHex(pow.Proof),
	}, []byte{}))

	// 2nd round of SHA256
	hash = sha256.Sum256(hash[:])
	hashInt.SetBytes(hash[:])

	target := big.NewInt(1)
	target = target.Lsh(target, uint(256-pow.Bits))

	return hashInt.Cmp(target) == -1
}

// convert a large number to hexadecimal bytes
func intToHex(n uint64) []byte {
	return []byte(strconv.FormatUint(n, 16))
}

