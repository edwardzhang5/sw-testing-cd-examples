const jwt = require('jsonwebtoken');

const token = jwt.sign(
  {
    exp: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 30,
    team: 'team1',
    course: 'drbyronw-F20',
  },
  'super-secret'
);

console.log('Token: ', token);
