


#
#
#
#Without using a built-in string split method, implement a split(s, c) method, which receives a string s and a character c and splits s at each occurrence of c, returning a list of strings.


def split(s, c):
    result = []
    if not s:
        return []

    word = ""
    for ch in s:
        if ch == c:
            result.append(word)
            word = ""
        else:
            word += ch

    result.append(word)

    return result



def run_tests():
  tests = [
      # Example 1 from the book
      ("split by space", ' ', ["split", "by", "space"]),
      # Example 2 from the book
      ("beekeeper needed", 'e', ["b", "", "k", "", "p", "r n", "", "d", "d"]),
      # Example 3 from the book
      ("/home/./..//Documents/", '/',
          ["", "home", ".", "..", "", "Documents", ""]),
      # Example 4 from the book
      ("", '?', []),
      # Edge case - empty string with various delimiters
      ("", ' ', []),
      ("", '\n', []),
      ("", '', []),
      # Edge case - single character string
      ("a", 'a', ["", ""]),
      ("a", 'b', ["a"]),
      # Edge case - no splits
      ("hello", 'x', ["hello"]),
      ("hello", '?', ["hello"]),
      # Edge case - all splits
      ("aaa", 'a', ["", "", "", ""]),
      # Edge case - special characters
      ("\n\n\n", '\n', ["", "", "", ""]),
      ("tab\tseparated\ttext", '\t', ["tab", "separated", "text"]),
      # Edge case - consecutive delimiters
      ("one,,two,,,three", ',', ["one", "", "two", "", "", "three"]),
      # Edge case - delimiter at start/end
      (",start,middle,end,", ',', ["", "start", "middle", "end", ""]),
      # Edge case - mixed length strings
      ("short,medium string,very very long string", ',', [
          "short", "medium string", "very very long string"]),
      # Edge case - whitespace handling
      ("  leading space", ' ', ["", "", "leading", "space"]),
      ("trailing space  ", ' ', ["trailing", "space", "", ""]),
      # Edge case - numbers and special chars
      ("123,456,789", ',', ["123", "456", "789"]),
      ("!@#$%", '@', ["!", "#$%"]),
  ]
  for s, c, want in tests:
    got = split(s, c)
    assert got == want, f"\nsplit({s}, {c}): got: {got}, want: {want}\n"

run_tests()
