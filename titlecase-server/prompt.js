const prompt = require('prompt');
const titleCase = require('./titlecase');

prompt.start();

prompt.get('title', (err, res) => {
	console.log(titleCase(res.title));
});
