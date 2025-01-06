def quick_sort(arr, low, high):
    if low < high:
        pivot_index = partition(arr, low, high)
        quick_sort(arr, low, pivot_index - 1)
        quick_sort(arr, pivot_index + 1, high)

def partition(arr, low, high):
    pivot_index = (low + high) // 2
    pivot = arr[pivot_index]

    # put the pivot at the end
    arr[high], arr[pivot_index] = arr[pivot_index], arr[high]

    i = low
    for j in range(low, high):
        if arr[j] < pivot:
            arr[i], arr[j] = arr[j], arr[i]
            i += 1

    # put the pivot at the right position
    arr[i], arr[high] = arr[high], arr[i]
    return i


arr = [4, 2, 7, 3, 6, 2, 5, 1]
n = len(arr)
quick_sort(arr, 0, n - 1)
print("Sorted array is:", arr)
