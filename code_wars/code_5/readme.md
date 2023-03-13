Write a function named first_non_repeating_letter that takes a string input, and returns the first character that is not repeated anywhere in the string.

For example, if given the input 'stress', the function should return 't', since the letter t only occurs once in the string, and occurs first in the string.

As an added challenge, upper- and lowercase letters are considered the same character, but the function should return the correct case for the initial letter. For example, the input 'sTreSS' should return 'T'.

If a string contains all repeating characters, it should return an empty string ("") or None -- see sample tests.

tests:

```
var _ = Describe("Basic Tests", func() {
   It("should handle simple tests", func() {
     Expect(FirstNonRepeating("a")).To(Equal("a"))
     Expect(FirstNonRepeating("stress")).To(Equal("t"))
     Expect(FirstNonRepeating("moonmen")).To(Equal("e"))
   })
    It("should handle empty strings", func() {
     Expect(FirstNonRepeating("")).To(Equal(""))
    })
    It("should handle all repeating strings", func() {
     Expect(FirstNonRepeating("abba")).To(Equal(""))
     Expect(FirstNonRepeating("aa")).To(Equal(""))
    })
    It("should handle odd characters", func() {
     Expect(FirstNonRepeating("~><#~><")).To(Equal("#"))
     Expect(FirstNonRepeating("hello world, eh?")).To(Equal("w"))
    })
    It("should handle letter cases", func() {
     Expect(FirstNonRepeating("sTreSS")).To(Equal("T"))
     Expect(FirstNonRepeating("Go hang a salami, I'm a lasagna hog!")).To(Equal(","))
    })
})
```
