def power_set(input_set):

    if not input_set:
        return [[]]

    final = []
    first = input_set[0]
    rest = input_set[1:]
    subsets = power_set(rest)

    for subset in subsets:
        final.append([first] + subset)
        final.append(subset)

    return final


print(power_set([1, 2, 3]))
