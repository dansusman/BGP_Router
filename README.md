# cs3700-proj3 - J. Adisoetjahya, D. Susman

## High Level Approach

To implement this BGP Router, we started with the router_skeleton provided by the instructors. Using the given starter code allowed us to narrow in on what was required to make the router perform as expected.
Since the milestone only required implementation of update messages, dump/table messages, and data messages, we avoided collasce and revoke (at least for now).

We started using Golang initially, but ran into some issues running the simulator, so we switched to Python. This made things a lot easier on us. We deciphered the skeleton, defined our data, and began working
on the forwarding table/routes/relations fields in the Router class. Since the packets are represented using JSON, most of the information we needed was readily available and easy to parse/work with. There were
a lot of intricacies we had to be careful to consider, such as bitwise AND'ing addresses, forwarding packets to the correct neigbors, and working with multiple open sockets/ports.

Overall, the skeleton served us well and ensured we stayed on task/heading in the right direction. There are certainly things to improve upon in future iterations of this project (full submission next Friday),
but for the most part, with the help of lecture notes and Google searches, our router ended up pretty well implemented.

Storing our forwarding table as an array of dictionaries ended up being a great choice, since we could easily iterate over the entries in the table and access any and all information needed (IP Address, Netmask, etc.).

## Challenges

Some of the challenges we faced stemmed from figuring out how to network program in Python. Since both of us completed previous projects in Golang, we started out our implementation in Go. However, we quickly discovered combability issues due to the Simulation being written in Python. Furthermore, methods such as using select() and poll() were favored to manage multiple sockets, which simply did not exist in Go. Therefore, we decided to scrap our implementation in Go and switch to coding in Python. Although coding in Python proved to be much more straight forward given the project instructions, it did take some time to get over the learning curve. With careful research and assistance from Piazza responses, we were eventually able to overcome our unfamiliarity with network programming in the language.

There were parts of the starter code that were confusing to us, not just because it was written in Python but also in the logic. One example was "get_routes()" which is essentially choosing the best route to forward the packet to. This included the series of network rules that were unclear to us, but were figured out as soon as we went over the lecture slides. In general, going over the lecture slides helped us understand the starter code better.

After settling with our implementation, additional challenges arose when we began testing through the simulation. The first problem was a failure to send all packets, so we had to narrow down which types were being dropped and where we were failing in our code. After a series of fixes, we began to have issues with our DUMP response, which was resolved after playing around with the destination of the TABLE message.

### Phase 2
For this project, there were a lot of conceptual challenges that we struggled with, particularly towards the end of the implementation. All up to longest prefix matching, our problems were usually resolved by looking at the lecture notes or re-reading the instructions more carefully. However, once we were tasked with aggregation and disaggregation, we found that modifying our code was more difficult than we anticipated. Before we dove into implementing aggregation, we talked through possible cases where aggregation was needed. In addition, we discussed similarities between networks and netmasks in the context of aggregation. Even with a clear idea of what we needed to do, the actual pythonic implementation was very cumbersome. The logic regarding disaggregation gave us the most trouble, so we were unfortunately unable to pass the last test (6-3).

## Testing the Program

Testing this project was a bit of a headache, since the code wouldn't run on either of our machines. Thus, we followed the following procedure to test our code/fix small bugs: Make a change, Git Push, Git Pull on
Khoury VM, run ./sim milestone. For the most part, this is how we tested and debugged. We utilized print statements that displayed our dump tables and update messages as they were being sent. This allowed us
to compare what was being sent to various sockets with what we thought should be sent. We caught many issues using print statements, while consulting notes/online docs.

Running ./sim milestone over and over allowed us to catch what was failing/passing. We used the simulator strategically to focus on the parts of the code where things were going wrong. At first, the error messages
were confusing or nonexistent, but as we made slight improvement after slight improvement, we got more and more helpful error messages. After solving many key errors and conceptual mishaps, we passed both milestone
tests on the Khoury machines, made some small tweaks, and wrote this README.
