#include <stdio.h>
#include <assert.h>
#include <string.h>

typedef enum {
    HTTP_BAD_REQUEST = 400,
    HTTP_UNAUTHORIZED = 401,
    HTTP_NOT_FOUND = 404,
    HTTP_TEAPOT = 418,
    HTTP_INTERNAL_SERVER_ERROR = 500
} HttpErrorCode;

char *http_to_str(HttpErrorCode code) {
    switch (code) {
        case HTTP_BAD_REQUEST: return "400 Bad Request";
        case HTTP_UNAUTHORIZED: return "401 Unauthorized";
        case HTTP_NOT_FOUND: return "404 Not Found";
        case HTTP_TEAPOT: return "418 I AM A TEAPOT!";
        case HTTP_INTERNAL_SERVER_ERROR: return "500 Internal Server Error";
        default: return "Unknown HTTP status code";
    }
}

int main() {
    // assert examples for HttpErrorCode
    assert(HTTP_BAD_REQUEST == 400);
    assert(HTTP_UNAUTHORIZED == 401);
    assert(HTTP_NOT_FOUND == 404);
    assert(HTTP_TEAPOT == 418);
    assert(HTTP_INTERNAL_SERVER_ERROR == 500);

    // assert examples for http_to_str function
    assert(strcmp(http_to_str(HTTP_BAD_REQUEST), "400 Bad Request") == 0);
    assert(strcmp(http_to_str(HTTP_UNAUTHORIZED), "401 Unauthorized") == 0);
    assert(strcmp(http_to_str(HTTP_NOT_FOUND), "404 Not Found") == 0);
    assert(strcmp(http_to_str(HTTP_TEAPOT), "418 I AM A TEAPOT!") == 0);
    assert(strcmp(http_to_str(HTTP_INTERNAL_SERVER_ERROR), "500 Internal Server Error") == 0);
    assert(strcmp(http_to_str(600), "Unknown HTTP status code") == 0);

    // Variable declarations for the second part of the main function
    int logLevel = 1; // Example log level value
    int errorCode = 4; // Example error code value

    // Log level switch-case
    switch (logLevel) {
        case 0: // assuming LOG_DEBUG = 0
            printf("Debug logging enabled\n");
            break;
        case 1: // assuming LOG_INFO = 1
            printf("Info logging enabled\n");
            break;
        case 2: // assuming LOG_WARN = 2
            printf("Warning logging enabled\n");
            break;
        case 3: // assuming LOG_ERROR = 3
            printf("Error logging enabled\n");
            break;
        default:
            printf("Unknown log level: %d\n", logLevel);
            break;
    }

    // Error code switch-case
    switch (errorCode) {
        case 1:
        case 2:
        case 3:
            // 1, 2, and 3 are all minor errors
            printf("Minor error occurred. Please try again.\n");
            break;
        case 4:
        case 5:
            // 4 and 5 are major errors
            printf("Major error occurred. Restart required.\n");
            break;
        default:
            printf("Unknown error.\n");
            break;
    }

    printf("All asserts passed and code executed successfully!\n");

    return 0;
}
