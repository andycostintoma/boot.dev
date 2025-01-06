#include <assert.h>
#include <stdio.h>

// Function to update the file data
void update_file(int filedata[200], int new_filetype, int new_num_lines) {
    filedata[1] = new_num_lines;   // Update the number of lines
    filedata[2] = new_filetype;    // Update the file type
    filedata[199] = 0;             // Set the last element to 0
}

// Arrays in C are passed by reference, meaning the function receives a pointer
// to the original array. Thus, modifications within the function affect the original array.
int main() {
    {
        int filedata[200] = {0};  // Initialize array with zeros
        update_file(filedata, 1, 100);   // Call the function to update the array
        // Assertions to check the updates
        assert(filedata[1] == 100);   // Number of lines should be updated to 100
        assert(filedata[2] == 1);     // File type should be updated to 1
        assert(filedata[199] == 0);   // Last element should be set to 0
    }

    printf("All assertions passed.\n");
    return 0;
}