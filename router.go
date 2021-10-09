package main

import "net"

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
	routes []string
    updates []string
    relations map[string]int
    sockets []int
}

func main() {
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

	}


}

func makeRouter(networks []string) {
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
}

// Lookup all valid routes for an address
func lookup_routes(daddr string) []string {

}

// Select the route with the shortest AS Path
func get_shortest_as_path(routes []string) string {

}

func get_highest_preference(routes []string) string {

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

func forward(srcif string, packet string) {

}

func coalesce() {

}

func update(srcif string, packet string) {

}

func revoke(packet string) {

}

func dump(packet string) {

}

func handle_packet(srcif string, packet string) {

}

func send_error(conn net.Conn, msg string) {

}