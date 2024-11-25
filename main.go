package main

/*
#cgo CFLAGS: -I./csrc
#cgo LDFLAGS: -L./csrc -lmdns
#include "mdns_wrapper.h"
#include <stdlib.h>
*/
import "C"
import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/cloudflare/circl/dh/kyber/kyber768"
	"github.com/cloudflare/circl/sign/sphincs"
)

const (
	serviceName = "TestService" // Example service name for mDNS
	hostname    = "test.local"  // Example hostname for mDNS
	port        = 5353          // Port for mDNS service
)

// StartMDNSService starts the mDNS service with the given parameters using the legacy C implementation.
func StartMDNSService(serviceName, hostname string, port int) {
	cServiceName := C.CString(serviceName)
	cHostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(cServiceName))
	defer C.free(unsafe.Pointer(cHostname))

	log.Printf("Starting mDNS service: %s on %s:%d", serviceName, hostname, port)
	C.start_mdns_service(cServiceName, cHostname, C.int(port))
}

// StopMDNSService stops the mDNS service using the legacy C implementation.
func StopMDNSService() {
	log.Println("Stopping mDNS service")
	C.stop_mdns_service()
}

// GenerateKyberKeyPair generates a CRYSTALS-Kyber keypair.
func GenerateKyberKeyPair() (kyber768.PrivateKey, kyber768.PublicKey) {
	privateKey, publicKey := kyber768.GenerateKeyPair(rand.Reader)
	return privateKey, publicKey
}

// GenerateSPHINCSKeyPair generates a SPHINCS+ keypair.
func GenerateSPHINCSKeyPair() (sphincs.PrivateKey, sphincs.PublicKey) {
	privateKey, publicKey, err := sphincs.GenerateKey(rand.Reader, sphincs.ParamsSHA256_128f)
	if err != nil {
		log.Fatalf("Failed to generate SPHINCS+ keys: %v", err)
	}
	return privateKey, publicKey
}

// TestMDNSService tests the mDNS service functionality and showcases cryptographic enhancements.
func TestMDNSService() {
	// Start mDNS service
	StartMDNSService(serviceName, hostname, port)

	// Generate cryptographic keys
	kyberPriv, kyberPub := GenerateKyberKeyPair()
	sphincsPriv, sphincsPub := GenerateSPHINCSKeyPair()

	// Display keys for debugging
	fmt.Printf("CRYSTALS-Kyber Public Key: %x\n", kyberPub)
	fmt.Printf("SPHINCS+ Public Key: %x\n", sphincsPub)

	// Stop mDNS service
	StopMDNSService()
}

func main() {
	// Ensure shard directory exists (optional for this test)
	err := os.MkdirAll("./shards", 0755)
	if err != nil {
		log.Fatalf("Failed to create shard directory: %v", err)
	}

	// Test the mDNS service with cryptographic enhancements
	TestMDNSService()
}
