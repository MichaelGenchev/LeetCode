


def palindrome(word):

    l, r = 0, len(word) - 1

    while l < r:
        if word[l] != word[r]:
            return False
        l += 1
        r -= 1
    return True




print(palindrome("level"))

print(palindrome("naan"))
print(palindrome("hello"))
