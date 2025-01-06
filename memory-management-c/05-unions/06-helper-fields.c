#include <stdio.h>
#include <stdint.h>
#include <assert.h>
#include <string.h>

// Define the union
typedef union PacketHeader {
    struct {
        uint16_t src_port;
        uint16_t dest_port;
        uint32_t seq_num;
    } tcp_header;
    uint8_t raw[8];
} packet_header_t;

void test_packet_header_size() {
    assert(sizeof(packet_header_t) == 8 && "PacketHeader union should be 8 bytes");
}

void test_tcp_header_fields() {
    packet_header_t header;
    header.tcp_header.src_port = 0x1234;
    header.tcp_header.dest_port = 0x5678;
    header.tcp_header.seq_num = 0x9ABCDEF0;

    assert(header.tcp_header.src_port == 0x1234 && "src_port should be 0x1234");
    assert(header.tcp_header.dest_port == 0x5678 && "dest_port should be 0x5678");
    assert(header.tcp_header.seq_num == 0x9ABCDEF0 && "seq_num should be 0x9ABCDEF0");
}

void test_field_to_raw_consistency() {
    packet_header_t header = {0};
    header.tcp_header.src_port = 0x1234;
    header.tcp_header.dest_port = 0x5678;
    header.tcp_header.seq_num = 0x9ABCDEF0;

    assert(header.raw[0] == 0x34 && "[0] should be 0x34");
    assert(header.raw[1] == 0x12 && "[1] should be 0x12");
    assert(header.raw[2] == 0x78 && "[2] should be 0x78");
    assert(header.raw[3] == 0x56 && "[3] should be 0x56");
    assert(header.raw[4] == 0xF0 && "[4] should be 0xF0");
    assert(header.raw[5] == 0xDE && "[5] should be 0xDE");
    assert(header.raw[6] == 0xBC && "[6] should be 0xBC");
    assert(header.raw[7] == 0x9A && "[7] should be 0x9A");
}

int main() {
    // Run the tests
    test_packet_header_size();
    test_tcp_header_fields();
    test_field_to_raw_consistency();

    printf("All tests passed.\n");

    return 0;
}