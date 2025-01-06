def bubble_sort(arr):
    end = len(arr)
    swapping = True

    # If no swaps were made during the pass, the array is sorted
    while swapping:
        swapping = False
        for i in range(1, end):
            if arr[i - 1] > arr[i]:
                arr[i - 1], arr[i] = arr[i], arr[i - 1]
                swapping = True
        print(f"After pass: {arr} -> {arr[end-1]} bubbled")
        end -= 1

arr = [3, 7, 4, 12, 6, 5]
print("Initial array: ", arr)
bubble_sort(arr)
print("Sorted array: ", arr)
