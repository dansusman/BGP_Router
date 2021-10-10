package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

// Reading the given connection until a newline character
func readAll(conn net.Conn) (message []byte, err error) {
	// creating a buffer store the accumulator
	var buf bytes.Buffer
	// NewReader creates reader of default size
	scanner := bufio.NewReader(conn)
	// looping through the response to ensure the entire message is read
	for {
		tmp, e := scanner.ReadString('\n')
		if e != nil {
			if e != io.EOF {
				return make([]byte, 0), e
			}
			break
		}
		// adding the bytes read in this iteration to the buffer
		print(tmp)
	}
	return buf.Bytes(), nil
}

const DEBUG bool = false

// Message Fields
const TYPE string = "type"
const SRCE string = "src"
const DEST string = "dst"
const MESG string = "msg"
const TABL string = "table"

// Message Types
const DATA string = "data"
const DUMP string = "dump"
const UPDT string = "update"
const RVKE string = "revoke"
const NRTE string = "no route"

// Update Message Fields
const NTWK string = "network"
const NMSK string = "netmask"
const ORIG string = "origin"
const LPRF string = "localpref"
const APTH string = "ASPath"
const SORG string = "selfOrigin"

// internal route info
const CUST string = "cust"
const PEER string = "peer"
const PROV string = "prov"

type Router struct {
	// List of IP Addresses mapped to port number (index)
	routes []string
	// Copy of each update announcement in JSON
	updates []map[string]string
	// IP Address to type of relationship (Customer, peer, or provider)
	relations map[string]string
	// IP Address mapped to its open connection
	sockets map[string]net.Conn
}

type Network struct {
	network string
	netmask string
	netType string
	AS      string
}

type Msg struct {
	network    string
	netmask    string
	localpref  int
	ASPath     []int
	origin     string
	selfOrigin bool
}

type Message struct {
	mType string
	src   string
	dest  string
	msg   Msg
}

type Packet struct {
	asn      int
	networks []Network
	messages []Message
}

func main() {
	asn, networks := parseCmdArgs()
	println(asn)
	router := makeRouter(networks)
	srcif := ""
	for {
		//     socks = select.select(self.sockets.values(), [], [], 0.1)[0]
		//     for conn in socks:
		//         try:
		//             k = conn.recv(65535)
		//         except:
		//             # either died on a connection reset, or was SIGTERM's by parent
		//             return
		//         if k:
		//             for sock in self.sockets:
		//                 if self.sockets[sock] == conn:
		//                     srcif = sock
		//             msg = json.loads(k)
		//             if not self.handle_packet(srcif, msg):
		//                 self.send_error(conn, msg)
		//         else:
		//             return
		// return
		socks := router.sockets
		for ip, conn := range socks {
			message, err := readAll(conn)
			checkError(err)
			println(message)
			if message != nil {
				if socks[ip] == conn {
					srcif = ip
				}
				var packet Packet
				err := json.Unmarshal(message, &packet)
				checkError(err)
				println(packet)
				if !handle_packet(srcif, packet) {
					send_error(conn, packet.messages)
				} else {
					return
				}
			}
		}
		return
	}

}

func parseCmdArgs() (int, []string) {
	if len(os.Args) < 2 {
		panic("Please specify an ASN for your router!")
	}
	asn, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Please include a valid ASN integer!")
	}

	return asn, os.Args[2:]
}

func makeRouter(networks []string) Router {
	router := new(Router)
	//     self.routes = []
	//     self.updates = []
	//     self.relations = {}
	//     self.sockets = {}
	//     for relationship in networks:
	//         network, relation = relationship.split("-")
	//         if DEBUG:
	//             print("Starting socket for", network, relation)
	//         self.sockets[network] = socket.socket(socket.AF_UNIX, socket.SOCK_SEQPACKET)
	//         self.sockets[network].setblocking(0)
	//         self.sockets[network].connect(network)
	//         self.relations[network] = relation
	//     return

	for _, relationship := range networks {
		splitArr := strings.Split(relationship, "-")
		network, relation := splitArr[0], splitArr[1]
		if DEBUG {
			fmt.Println("Starting socket for ", network, relation)
		}
		router.sockets[network] = startSocket(network)
		// set blocking?
		router.relations[network] = relation
	}
	return *router
}

func startSocket(network string) net.Conn {
	connection, err := net.Dial("unix", network)
	checkError(err)
	return connection
}

func checkError(err error) {
	if err != nil {
		panic("Error occurred: " + err.Error())
	}
}

// Lookup all valid routes for an address
func lookup_routes(daddr string) []string {
	return nil
}

// Select the route with the shortest AS Path
func get_shortest_as_path(routes []string) string {
	return ""
}

func get_highest_preference(routes []string) string {
	return ""
}

// Select self originating routes
func get_self_origin(routes []string) {

}

// Select origin routes: IGP > EGP > UNK
func get_origin_routes(routes []string) {

}

// Don't allow Peer->Peer, Peer->Prov, or Prov->Peer forwards
func filter_relationships(srcif string, routes []string) {

}

// Select the best route for a given address
func get_route(srcif string, daddr string) {
	// peer = None
	// routes = lookup_routers(daddr)
	// # Rules go here
	// if routes:
	//     # 1. Highest Preference
	//     routes = self.get_highest_preference(routes)
	//     # 2. Self Origin
	//     routes = self.get_self_origin(routes)
	//     # 3. Shortest ASPath
	//     routes = self.get_shortest_as_path(routes)
	//     # 4. IGP > EGP > UNK
	//     routes = self.get_origin_routes(routes)
	//     # 5. Lowest IP Address
	//     # TODO
	//     # Final check: enforce peering relationships
	//     routes = self.filter_relationships(srcif, routes)
	// return self.sockets[peer] if peer else None

}

func forward(srcif string, packet Packet) {

}

func coalesce() {

}

func update(srcif string, packet Packet) {

}

func revoke(packet Packet) {

}

func dump(packet Packet) {

}

func handle_packet(srcif string, packet Packet) bool {
	return true

}

func send_error(conn net.Conn, msg []Message) {

}
