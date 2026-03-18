



def join(arr, s):

    res = ""

    for i, el in enumerate(arr):
        res += el
        if i != len(arr) - 1:
            res += s
    return res



def run_tests():
  tests = [
      # Example 1 from the book
      (["join", "by", "space"], " ", "join by space"),
      # Example 2 from the book
      (["b", "", "k", "", "p", "r n", "", "d", "d!!"],
          "ee", "beeeekeeeepeer neeeedeed!!"),
      # Edge case - empty arrays
      ([], "x", ""),
      ([], "", ""),
      ([], "long separator", ""),
      # Edge case - single element arrays
      (["a"], "x", "a"),
      ([""], "x", ""),
      (["multiple words"], "x", "multiple words"),
      # two element arrays
      (["a", "b"], "", "ab"),
      (["a", "b"], " ", "a b"),
      (["", ""], ",", ","),
      # Edge case - empty strings in array
      (["", "", ""], ",", ",,"),
      (["hello", "", "world"], " ", "hello  world"),
      # special characters
      (["\n", "\t"], ",", "\n,\t"),
      (["tab", "separated"], "\t", "tab\tseparated"),
      # long separators
      (["short", "strings"], "very long separator",
       "shortvery long separatorstrings"),
      # mixed content
      (["123", "abc", "!@#", "   "], "|", "123|abc|!@#|   "),
      # whitespace handling
      (["  leading", "trailing  ", "  both  "],
          "|", "  leading|trailing  |  both  "),
      # numbers and special chars
      (["123", "456"], "-", "123-456"),
      (["!@#", "$%^"], "&", "!@#&$%^"),
  ]
  for arr, s, want in tests:
    got = join(arr, s)
    assert got == want, f"\njoin({arr}, {s}): got: {got}, want: {want}\n"

run_tests()
