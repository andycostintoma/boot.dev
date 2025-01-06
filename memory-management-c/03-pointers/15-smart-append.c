#include <stdio.h>
#include <string.h>
#include <assert.h>

typedef struct {
    size_t length;
    char buffer[64];
} TextBuffer;

int smart_append(TextBuffer *dest, const char *src) {
    if (!dest || !src) return 1;
    const int max_buffer_size = 64;
    const size_t src_len = strlen(src);
    const size_t available_space = max_buffer_size - dest->length - 1; // 1 for NULL terminator
    if (src_len > available_space) {
        strncat(dest->buffer, src, available_space);
        dest->length = max_buffer_size - 1;
        return 1;
    }
    strcat(dest->buffer, src);
    dest->length += src_len;
    return 0;
}

void test_return_1_for_null_value() {
    TextBuffer dest;
    const char *src = NULL;
    int result = smart_append(&dest, src);
    assert(result == 1 && "Should return 1 for null value");
}

void test_smart_append_empty_buffer() {
    TextBuffer dest;
    strcpy(dest.buffer, "");
    dest.length = 0;
    const char *src = "Hello";
    int result = smart_append(&dest, src);
    assert(result == 0 && "Should return 0 for successful append");
    assert(strcmp(dest.buffer, "Hello") == 0 && "Buffer should contain 'Hello'");
    assert(dest.length == 5 && "Length should be 5");
}

void test_smart_append_full_buffer() {
    TextBuffer dest;
    strcpy(dest.buffer, "This is a very long string that will fill up the entire buffer.");
    dest.length = 63;
    const char *src = " Extra";
    int result = smart_append(&dest, src);
    assert(result == 1 && "Should return 1 for unsuccessful append");
    assert(
        strcmp(dest.buffer, "This is a very long string that will fill up the entire buffer.") == 0 &&
        "Buffer should remain unchanged");
    assert(dest.length == 63 && "Length should remain 63");
}

void test_smart_append_overflow() {
    TextBuffer dest;
    strcpy(dest.buffer, "This is a long string");
    dest.length = 21;
    const char *src = " that will fill the whole buffer and leave no space for some of the chars.";
    int result = smart_append(&dest, src);
    assert(result == 1 && "Should return 1 for overflow append");
    assert(
        strcmp(dest.buffer, "This is a long string that will fill the whole buffer and leave") == 0 &&
        "Buffer should be filled to capacity");
    assert(dest.length == 63 && "Length should be 63 after overflow append");
}

int main() {
    test_return_1_for_null_value();
    test_smart_append_empty_buffer();
    test_smart_append_full_buffer();
    test_smart_append_overflow();

    printf("All tests passed!\n");

    return 0;
}
