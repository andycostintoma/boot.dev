def selection_sort(arr):
    for i in range(len(arr)):
        smallest_idx = i
        for j in range(i + 1, len(arr)):
            if arr[j] < arr[smallest_idx]:
                smallest_idx = j
        arr[i], arr[smallest_idx] = arr[smallest_idx], arr[i]
    return arr

arr = [4, 2, 7, 3, 6, 2, 5, 1]
print("Sorted array is:", selection_sort(arr))
