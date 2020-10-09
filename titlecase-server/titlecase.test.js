const titleCase = require('./titlecase');

describe('TitleCase returns an appropriate movie title for the text entered', () => {
	it('should capitalize the letter after the hyphen in the title', () => {
		const res = titleCase('x-men mission-critical');
		expect(res).toBe('X-Men Mission-Critical');
	});
	it('should not capitalize common prepositions or articles', () => {
		const res = titleCase('the young and the restless');
		expect(res).toBe('The Young and the Restless');
	});

	it('should not break on numbers', () => {
		expect(titleCase('24')).toBe('24');
	});

	it('should capitalize multiword titles', () => {
		const res = titleCase('the roaD');
		expect(res).toBe('The Road');
	});
	it('should properly capitalize mixed case strings', () => {
		expect(titleCase('suPERman')).toBe('Superman');
	});

	it('should capitalize a single word title', () => {
		expect(titleCase('superman')).toBe('Superman');
	});

	it('should capilize a single letter title', () => {
		const res = titleCase('v');
		expect(res).toEqual('V');
	});

	it('should only accept a string as an argument', () => {
		expect(() => titleCase(3)).toThrow();
	});

	it('should return a string', () => {
		expect(typeof titleCase('movie')).toBe('string');
	});
});

describe('Jest Test Framework works as it should', () => {
	it('runs correctly', () => {
		expect(true).toBeTruthy();
	});
});
