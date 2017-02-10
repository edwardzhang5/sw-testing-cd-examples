let expect = require('chai').expect;
let should = require('chai').should();

let titleCased = require('../movie/title_case.js').titleCased;

// test suite
describe('Titlecase Function', () => {
  // titlecase function tests
  it('should fix multi-cased words', () => {
    titleCased('suPERman ReTurnS').should.equal('Superman Returns');
  });

  it('should ignore articles, prepositions etc.', () => {
    titleCased('young and the Restless').should.equal('Young and the Restless');
  });

  it('should accept hyphenated words', () => {
    titleCase('X-men and the mission-critical mission').should.equal('X-Men and the Mission-Critical Mission');
  });

  it('should throw a TypeError for input other than a string', () => {
     should.Throw(() => titleCase(1));
   });

  it('should capitalize a single word as a movie title', () => {
    titleCased('titanic').should.equal('Titanic');
  });

  it('should convert a single character to a title', () => {
    expect(titleCased('v')).to.equal('V');
  });

  it('should return a string', () => {
    titleCased('a string').should.be.a('string');
  });

  it('should take a title string and return a title-cased string', () => {
    expect(titleCased('dazed and confused')).to.equal('Dazed and Confused');
    expect(titleCased('the great gatsby')).to.equal('The Great Gatsby');
  });
});

describe('Mocha test framework', () => {

  it('should run the assertion', () => {

    expect(true).to.be.true;
    
  });
});
