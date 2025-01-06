#include <stdio.h>

void printMessageOne();
void printMessageTwo();
void printMessageThree();
void printStackPointerDiff();

int main() {
    printMessageOne(); // Main function calls printMessageOne
    return 0;
}

void printMessageOne() {
    const char *message = "Dark mode?\n";   // Local variable stored on stack
    printStackPointerDiff();                // Check the stack pointer difference
    printf("%s\n", message);                // Print message
    printMessageTwo();                      // Call the next function
}

void printMessageTwo() {
    const char *message = "More like...\n"; // Local variable stored on stack
    printStackPointerDiff();                // Check the stack pointer difference
    printf("%s\n", message);                // Print message
    printMessageThree();                    // Call the next function
}

void printMessageThree() {
    const char *message = "dark roast.\n";  // Local variable stored on stack
    printStackPointerDiff();                // Check the stack pointer difference
    printf("%s\n", message);                // Print message
}

void printStackPointerDiff() {
    static void *last_sp = NULL;                    // Static variable to store last stack pointer
    void *current_sp = __builtin_frame_address(0);  // Get current stack pointer
    long diff = (char*)last_sp - (char*)current_sp; // Calculate difference
    if (last_sp == NULL) {
        last_sp = current_sp;                       // First call sets the initial stack pointer
        diff = 0;
    }
    printf("---------------------------------\n");
    printf("Stack pointer offset: %ld bytes\n", diff);
    printf("---------------------------------\n");
}

/*
 * Explanation of the Stack:
 * When main() calls printMessageOne(), a new stack frame is created to store the function's return address,
 * parameters (if any), and local variables.
 * Inside printMessageOne(), the local variable 'message' is created and stored in the stack frame.
 * printStackPointerDiff() is called to print the current stack pointer offset compared to the last call.
 * The process repeats as printMessageOne() calls printMessageTwo(), creating another stack frame for printMessageTwo() and so on.
 * The static variable last_sp keeps track of the stack pointer from the previous call to calculate the difference.
 * The stack frames are deallocated when the functions return, resetting the stack pointer to its previous value.
*/