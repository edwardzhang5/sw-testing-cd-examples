const _ = require('lodash');

const dontCapitalize = ['and', 'or', 'the', 'for', 'but', 'to', 'a'];

const titlecase = (title) => {
  if (typeof title !== 'string') {
    throw new TypeError();
  }

  title = title.toLowerCase();

  title = title.replace(/(-.)/g, (x) => {
    return x.toUpperCase();
  });

  const words = title.split(' ');

  const titleWords = _.map(words, (word, index) => {
    if (index == 0) {
      return `${word[0].toUpperCase()}${word.substring(1)}`;
    } else if (dontCapitalize.includes(word)) {
      return word;
    } else {
      return `${word[0].toUpperCase()}${word.substring(1)}`;
    }
  });

  return titleWords.join(' ');
};

module.exports = titlecase;
