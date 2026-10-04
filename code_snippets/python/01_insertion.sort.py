def insertion_sort(list):
    for i in range(len(list)):
        key = list[i]
        j = i -1
        while j>=0 and list[j] > key:
            list[j+1] = list[j]
            j = j - 1
            print(f" status for i: {i}, j: {j}, key : {key}, list: {list}")
        list[j+1] = key
    return list

if __name__ == "__main__":
    unsorted = [4,5,9,1,34,5,13,76,24,9,22]
    sorted = insertion_sort(unsorted)
    print(f" for unsorted {unsorted} \n the sorted list is {sorted}")
