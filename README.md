# Word Helper

A command line tool that wraps the [Datamuse API](https://www.datamuse.com/api/).

## Usage
Invoke the script like `./word_helper [flags]`

Example: `./word_helper -rhyme=dog -means=amphibian`
```
  -c string
      Number of results to return (default "100")
  -comes-after string
      The word before the desired word
  -comes-before string
      The word after the desired word
  -means string
      Word or phrase with a similar meaning to the desired word
  -related-to string
      Words commonly used in works including the desired word
  -rhyme string
      A word that the desired word rhymes with
  -slant-rhyme string
      A word that the desired word has a slant rhyme with
  -sounds-like string
      A word that the desired word should sound like
  -spelling string
      Spelling of desired word; use ? as a placeholder for 1 letter; use * as a placeholder for multiple letters
  -topics string
      Up to 5 words relating to the desired word
  -v  Show verbose API response
```
## Building

* Have the Golang environment setup on your computer
* Run `go build` inside of the project
