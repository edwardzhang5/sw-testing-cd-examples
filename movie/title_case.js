
const dontCapitalize = ['and', 'a', 'as', 'the', 'for', 'or'];

function titleCased(title) {
  // check type and throw error if not a string
  if (typeof(title) !== 'string') {
    throw new TypeError;
  }

  // set lowercase to handle multi-cased worsed
  title = title.toLowerCase();

  // capitalize characters after hyphen
  title = title.replace(/(-.)/g, (x) => {
    const newString = '-' + x[1].toUpperCase();
    return newString;
  });

  // create array of words then map over to uppercase the first letter as needed
  const words = title.split(' ');
  const titleCasedWords = words.map((word, index) => {
    if (index == 0) {

      return word[0].toUpperCase() + word.substring(1);
    } else {
      if (dontCapitalize.includes(word)) { // dont capitalize certain words
        return word;
      }
      else {
        return word[0].toUpperCase() + word.substring(1);
      }
    }
  })

  // join the array to create a single string
  return titleCasedWords.join(' ');
}

module.exports.titleCased = titleCased;
